angular.module('starbucks').controller('orderCtrl', [ '$rootScope', '$scope', '$http', '$state', function($rootScope, $scope, $http, $state) {

	
	$scope.addToCart = function () {

		console.log("Bev type :" + $scope.orderItem.bevName) ;
		$rootScope.order.bagItems.push($scope.orderItem)
		resetOrderItem () ;
		$state.go ( 'home' ) ;

	}


	function resetOrderItem (){

		$scope.orderItem.bevName = "" ;
		$scope.orderItem.qty = 0 ;
		$scope.orderItem.milkType = "" ;
		$scope.orderItem.cupSize	 = "" ;


	}

} ]);
