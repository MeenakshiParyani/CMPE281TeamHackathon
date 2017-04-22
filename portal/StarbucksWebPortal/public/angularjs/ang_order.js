angular.module('starbucks').controller('orderCtrl', [ '$rootScope', '$scope', '$http', '$state', function($rootScope, $scope, $http, $state) {

	
	$scope.addToCart = function () {

		console.log("Bev type :" + $scope.orderItem.bevName) ;

		var orderItemBag = {
			name : $scope.orderItem.bevName ,
			qty : $scope.orderItem.qty,
			milk : $scope.orderItem.milkType,
			size : $scope.orderItem.cupSize 
		} ;

		$rootScope.bagItems.push(orderItemBag);
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
