angular.module('starbucks').controller('myBagCtrl', [ '$rootScope', '$scope', '$http', '$state', function($rootScope, $scope, $http, $state) {



	$scope.placeOrder = function(){
		console.log("Order Placed");
		$rootScope.customerName = $scope.customerName;
		console.log($rootScope.customerName);
		console.log($scope.inRadio);

		$rootScope.order = {
				"customerName" : $rootScope.customerName,
				"items" : $rootScope.bagItems,
				"location" : $scope.inRadio
		};

		$http({
			method : "POST",
			url : '/api/order',
			data : $rootScope.order
		}).success(function(response) {
			console.log('post response ' + response);
			$rootScope.orderId = response.orderId;
		}).error(function(err) {
			console.log(err);
		});
	}

}]);