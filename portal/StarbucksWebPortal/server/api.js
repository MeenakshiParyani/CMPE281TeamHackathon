var express = require('express');
var request = require('request');

var router = express.Router();

router.get('/', function(req, res) {
  res.status(200).json({status: 'OK'});
});


router.get(':location/order/:orderID', function(req, res) {
  request(process.env.GATEWAY_URL + '/' + req.params.location + '/order/' + req.params.orderID, function(err, r, body) {
    if (err) { throw err; }
    res.append('Content-Type', 'application/json').status(r.statusCode).send(body);
  });
});

router.post(':location/order', function(req, res) {
  var options = {
    url: process.env.GATEWAY_URL + '/' + req.params.location + '/order',
    method: 'POST',
    json: req.body
  };

  request(options, function(err, r, body) {
    if (err) { throw err; }
    res.append('Content-Type', 'application/json').status(r.statusCode).send(body);
  });
});

router.post(':location/order/:orderID/pay', function(req, res) {
  var options = {
    url: process.env.GATEWAY_URL + '/' + req.params.location + '/order/' + req.params.orderID + '/pay',
    method: 'POST',
    json: req.body
  };

  request(options, function(err, r, body) {
    if (err) { throw err; }
    res.append('Content-Type', 'application/json').status(r.statusCode).send(body);
  });
});

router.put(':location/order/:orderID', function(req, res) {
  var options = {
    url: process.env.GATEWAY_URL + '/' + req.params.location + '/order/' + req.params.orderID,
    method: 'PUT',
    json: req.body
  };

  request(options, function(err, r, body) {
    if (err) { throw err; }
    res.append('Content-Type', 'application/json').status(r.statusCode).send(body);
  });
});

router.delete(':location/order/:orderID', function(req, res) {
  request.delete(process.env.GATEWAY_URL + '/' + req.params.location + '/order/' + req.params.orderID, function(err, r, body) {
    if (err) { throw err; }
    res.append('Content-Type', 'application/json').status(r.statusCode).send(body);
  });
});

router.get(':location/orders', function(req, res) {
  request.get(process.env.GATEWAY_URL + '/' + req.params.location + '/orders', function(err, r, body) {
    if (err) { throw err; }
    res.append('Content-Type', 'application/json').status(r.statusCode).send(body);
  });
});

module.exports = router;