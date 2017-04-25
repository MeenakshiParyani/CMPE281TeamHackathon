angular.module('starbucks').component('bag', {
  templateUrl: 'components/bag/bag.template.html',

  controller: ['$state', 'OrdersService',
    function BagController($state, OrdersService) {
      var self = this;

      this.items = OrdersService.getPendingItems();

      this.orderMode = '';
      this.customerName = '';

      this.placeOrder = function() {
        // TODO: Form validation

        OrdersService.placeOrder({
          mode: self.orderMode,
          customerName: self.customerName
        }).then(function() {
          $state.go('pay');
        });
      };
    }
  ]
});