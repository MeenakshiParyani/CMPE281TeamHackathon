angular.module('starbucks').component('orders', {
  templateUrl: 'components/orders/orders.component.html',

  controller: ['OrdersService',
    function OrdersController(OrdersService) {
      this.orders = OrdersService.getAllOrders();
    }
  ]
});
