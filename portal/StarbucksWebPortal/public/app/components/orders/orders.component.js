angular.module('starbucks').component('orders', {
  templateUrl: 'components/orders/orders.component.html',

  controller: ['OrdersService',
    function OrdersController(OrdersService) {
      var self = this;

      this.orders = OrdersService.getAllOrders();

      this.removeOrder = function(orderId) {
        OrdersService.removeOrder(orderId).then(function() {
          self.orders.forEach(function(order, i) {
            if (order.id === orderId) self.orders.splice(i, 1);
          });
        });
      };
    }
  ]
});
