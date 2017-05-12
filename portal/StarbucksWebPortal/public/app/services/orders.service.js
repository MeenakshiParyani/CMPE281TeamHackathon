angular.module('starbucks').service('OrdersService', function($http) {
  var service = {
    orderId: '',
    currentLocation: '',
    pendingItems: [],
    numPendingItems: 0,

    setLocation: function(location) {
      service.currentLocation = location;
    },

    getCurrentLocation: function() {
      return service.currentLocation;
    },

    getPendingItems: function() {
      return service.pendingItems;
    },

    addItem: function(item) {
      service.pendingItems.push(item);
      service.numPendingItems++;
    },

    removeItem: function(i) {
      service.pendingItems.splice(i, 1);
      service.numPendingItems--;
    },

    placeOrder: function(mode, customerName) {
      return $http.post('/api/order', {
        location: mode,
        customerName: customerName,
        items: service.pendingItems
      }).then(function(res) {
        service.orderId = res.orderId;
        service.pendingItems = []; 
      });
    },

    removeOrder: function(orderId) {
      console.log("Order ID to be removed : " + orderId );
      return $http.delete('/api/order/' + orderId);
    },

    pay: function() {
      return $http.post('/api/order/' + service.orderId + '/pay').then(function() {
        service.orderId = '';
      });
    },

    getAllOrders: function(callback) {
     $http.get('/api/orders').then(function(res) {
        callback(res.data.orders || res.data);
      });
    }   
  };

  return service;
});