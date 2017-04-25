angular.module('starbucks').component('order', {
  templateUrl: 'components/order/order.template.html',

  controller: ['$state', 'OrdersService',
    function OrderController($state, OrdersService) {
      var self = this;

      this.item = {
        name: '',
        qty: 0,
        milk: '',
        size: ''
      };

      // TODO: Form validation
      this.addToBag = function() {
        OrdersService.addItem(self.item);
        $state.go('home');
      };
    }
  ]
});