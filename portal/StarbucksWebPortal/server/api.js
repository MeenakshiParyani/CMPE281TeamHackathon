var express = require('express');
var request = require('request');

var router = express.Router();

router.get('/', function(req, res) {
  res.status(200).json({status: 'OK'});
});


router.get('/order/:orderID', function(req, res) {
  request(process.env.GATEWAY_URL + '/' + req.query.location + '/order/' + req.params.orderID, function(err, r, body) {
    if (err) { throw err; }
    res.append('Content-Type', 'application/json').status(r.statusCode).send(body);
  });
});

router.post('/order', function(req, res) {
  var options = {
    url: process.env.GATEWAY_URL + '/' + req.query.location + '/order',
    method: 'POST',
    json: req.body
  };

  request(options, function(err, r, body) {
    if (err) { throw err; }
    res.append('Content-Type', 'application/json').status(r.statusCode).send(body);
  });
});

router.post('/order/:orderID/pay', function(req, res) {
  var options = {
    url: process.env.GATEWAY_URL + '/' + req.query.location + '/order/' + req.params.orderID + '/pay',
    method: 'POST',
    json: req.body
  };

  request(options, function(err, r, body) {
    if (err) { throw err; }
    res.append('Content-Type', 'application/json').status(r.statusCode).send(body);
  });
});

router.put('/order/:orderID', function(req, res) {
  var options = {
    url: process.env.GATEWAY_URL + '/' + req.query.location + '/order/' + req.params.orderID,
    method: 'PUT',
    json: req.body
  };

  request(options, function(err, r, body) {
    if (err) { throw err; }
    res.append('Content-Type', 'application/json').status(r.statusCode).send(body);
  });
});

router.delete('/order/:orderID', function(req, res) {
  request.delete(process.env.GATEWAY_URL + '/' + req.query.location + '/order/' + req.params.orderID, function(err, r, body) {
    if (err) { throw err; }
    res.append('Content-Type', 'application/json').status(r.statusCode).send(body);
  });
});

router.get('/orders', function(req, res) {
  console.log(process.env.GATEWAY_URL + '/' + req.query.location + '/orders');
  request.get(process.env.GATEWAY_URL + '/' + req.query.location + '/orders', function(err, r, body) {
    if (err) { throw err; }
    res.append('Content-Type', 'application/json').status(r.statusCode).send(body);
  });
});

module.exports = router;