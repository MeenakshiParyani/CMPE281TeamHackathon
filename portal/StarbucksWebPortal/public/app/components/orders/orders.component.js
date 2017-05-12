angular.module('starbucks').component('orders', {
  templateUrl: 'components/orders/orders.component.html',

  controller: ['$scope','OrdersService',
    function OrdersController($scope, OrdersService) {
      var self = this;
      
      OrdersService.getAllOrders(function(orders) {
        $scope.orders = orders;
      });


      
      this.removeOrder = function(orderId) {
        OrdersService.removeOrder(orderId).then(function() {
          console.log("Delelted order ...Now removing from array")
          
          $scope.orders.forEach(function(order, i) {
            if ((order._id || order.id) === orderId){ 
            $scope.orders.splice(i, 1);
            console.log($scope.orders);
          }
          });
        });
      };
    }
  ]
});
