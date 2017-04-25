angular.module('starbucks').component('pay', {
  templateUrl: 'component/pay/pay.template.html',

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