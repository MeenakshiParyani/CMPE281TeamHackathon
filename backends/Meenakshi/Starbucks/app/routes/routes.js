
var path = require('path');
var express = require('express');
var router = express.Router();

var order = require('./../models/Order');

 //Get all Orders
router.get('/api/orders', function(req, res) {
    order.find(function(err, orders) {
        if (err)
            res.send(err);

        res.json({'orders' :orders}); // return all orders
    });
});

 //Place an Order
router.get('/api/order/place', function(req, res) {
            
    Order.find(function(err, orders) {

        if (err)
            res.send(err);

        res.json(orders); // return all orders
    });
});


module.exports = router;
