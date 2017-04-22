
var path = require('path');
var express = require('express');
var router = express.Router();

var order = require('./../models/Order');

 //Get all Orders for a customer
router.get('/orders', function(req, res) {
    //console.log(req.params.customerName);
    // order.find({ customerName : req.params.customerName }, function(err, orders) {
       order.find(function(err, orders) {
        if (err)
            res.send(err);
        res.json({'orders' : orders}); // return all orders
    });
});

 //Place an Order
router.post('/order', function(req, res) {
    var newOrder = new order(req.body);
    console.log(newOrder);
    newOrder.save(function(err, order) {
        if(err){
            message = {"error": true, "message" : "Error Processing the Order"};
            console.log(err);
        }
        else
            message = {"error": false, "message" : "Order Placed Successfully", "orderId" : order._id};
        res.json(message);
    });
});

// Edit an Order
router.put('/order/:orderId', function(req, res){
    console.log(req.params.orderId);
    order.findOneAndUpdate({ _id : req.params.orderId }, req.body, function(err, order){
        console.log('edited order' + order);
        if(err)
            message =  { "error" : true, "message": err};
        else
            message =  { "error" : false, "message": "Order Edited Successfully" };
        res.json(message);
    });
});

// Cancel an Order
router.delete('/order/:orderId', function(req, res){
    console.log(req.params.orderId);
    order.remove({ _id: req.params.orderId }, function(err, order){
        if(err)
            message =  { "error" : true, "message": err};
        else
            message =  { "error" : false, "message": "Order Cancelled Successfully" };
        res.json(message);
    });

});

// Pay for an Order
router.put('/order/:orderId/pay', function(req, res){
    console.log(req.params.orderId);
    order.findOneAndUpdate({ _id: req.params.orderId }, {"status" : "Paid"}, function(err, order){
        if(err)
            message =  { "error" : true, "message": err};
        else
            message =  { "error" : false, "message": "Order Paid Successfully", "status" : "Paid" };
        res.json(message);
    });

});

module.exports = router;
