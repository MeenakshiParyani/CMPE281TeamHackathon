angular.module('starbucks').controller('myBagCtrl', [ '$rootScope', '$scope', '$http', '$state', function($rootScope, $scope, $http, $state) {



	$scope.placeOrder = function(){
		console.log("Order Placed");
		$rootScope.order = {
				"customerName" : $rootScope.customerName,
				"items" : $rootScope.bagItems,
				"location" : $scope.inRadio
		};
		$rootScope.order.customerName = $scope.customerName;
		$http({
			method : "POST",
			url : '/api/order',
			data : $rootScope.order
		}).success(function(response) {
			console.log('post response ' + response);
			$rootScope.order.orderId = response.orderId;
			$state.go('payment');
		}).error(function(err) {
			console.log(err);
		});
	}

}]);