angular.module('starbucks').component('orders', {
  templateUrl: 'component/orders/orders.component.html',

  controller: ['OrdersService',
    function PayController(OrdersService) {
      this.confirm = function() {
        // TODO: Form validation

        OrdersService.pay().then(function() {
          $state.go('home');
        });
      };
    }
  ]
});