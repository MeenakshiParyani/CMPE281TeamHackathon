angular.module('starbucks').component('header', {
  templateUrl: 'components/header/header.template.html',

  controller: ['$scope', '$state', 'OrdersService',
    function HeaderController($scope, $state, OrdersService) {
      var self = this;

      $scope.OrdersService = OrdersService;
      
      this.numPendingItems = OrdersService.numPendingItems;
      $scope.$watch('OrdersService.numPendingItems', function() {
        self.numPendingItems = OrdersService.numPendingItems;
      });

      this.location = OrdersService.getCurrentLocation();

      this.setLocation = function(location) {
        OrdersService.setLocation(location);
        self.location = OrdersService.getCurrentLocation();
      };

      this.setState = function(state) {
        $state.go(state);
      };
    }
  ]
});