
var path = require('path');
var express = require('express');
var router = express.Router();

var order = require('./../models/Order');

 //Get all Orders for a customer
router.get('/orders/:customerName', function(req, res) {
    console.log(req.params.customerName);
    order.find({ customerName : req.params.customerName }, function(err, orders) {
        if (err)
            res.send(err);
        res.json({'orders' : orders}); // return all orders
    });
});

 //Place an Order
router.post('/order/place', function(req, res) {
    var newOrder = new order(req.body);
    console.log(newOrder);
    newOrder.save(function(err) {
        if(err)
            message = {"error": true, "message" : "Error Processing the Order"};
        else
            message = {"error": false, "message" : "Order Placed Successfully"};
        res.json(message);
    });
});

// Edit an Order
router.put('/order/:orderId', function(req, res){
    console.log(req.params.orderId);
    order.findOneAndUpdate({ _id : req.params.orderId }, req.body, function(err, order){
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

module.exports = router;
