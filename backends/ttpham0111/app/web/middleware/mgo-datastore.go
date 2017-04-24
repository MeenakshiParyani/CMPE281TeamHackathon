package middleware

import (
	"crypto/tls"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net"
	"strings"
	"ttpham0111/app/models"
)

type MgoDatastore struct {
	s *mgo.Session
	c *mgo.Collection
}

func NewMgoDatastore(dbUrl string) *MgoDatastore {
	var s *mgo.Session
	var err error
	useSSL := false

	sslFlag := "ssl=true"
	if strings.Contains(dbUrl, sslFlag) {
		dbUrl = strings.Replace(dbUrl, sslFlag, "", 1)
		useSSL = true
	}

	tlsConfig := &tls.Config{}
	dialInfo, err := mgo.ParseURL(dbUrl)
	if err != nil {
		panic(err)
	}

	if useSSL {
		dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
			conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
			return conn, err
		}
	}

	s, err = mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	}

	c := s.DB("").C("orders")
	return &MgoDatastore{s, c}
}

func (d *MgoDatastore) Close() {
	d.s.Close()
}

func (d *MgoDatastore) GetOrders() (error, []models.Order) {
	orders := make([]models.Order, 0)
	err := d.c.Find(nil).All(&orders)
	return err, orders
}

func (d *MgoDatastore) GetOrder(orderID string) (error, *models.Order) {
	if !bson.IsObjectIdHex(orderID) {
		return NoResultFound{}, nil
	}

	var order models.Order
	err := d.c.FindId(bson.ObjectIdHex(orderID)).One(&order)
	if err == mgo.ErrNotFound {
		return NoResultFound{}, nil
	}
	return err, &order
}

func (d *MgoDatastore) CreateOrder(order *models.Order) (error, *models.Order) {
	err := d.c.Insert(order)
	return err, order
}

func (d *MgoDatastore) UpdateOrder(orderID string, order *models.Order) (error, *models.Order) {
	if !bson.IsObjectIdHex(orderID) {
		return NoResultFound{}, nil
	}

	err := d.c.UpdateId(bson.ObjectIdHex(orderID), order)
	if err == mgo.ErrNotFound {
		return NoResultFound{}, nil
	}
	return err, order
}

func (d *MgoDatastore) DeleteOrder(orderID string) error {
	if !bson.IsObjectIdHex(orderID) {
		return NoResultFound{}
	}

	err := d.c.RemoveId(bson.ObjectIdHex(orderID))
	return err
}
