angular.module('starbucks').controller('orderCtrl', [ '$rootScope', '$scope', '$http', '$state', function($rootScope, $scope, $http, $state) {

	
	$scope.addToCart = function () {

		console.log("Bev type :" + $scope.orderItem.bevName) ;

		var orderItemBag = {
			bevName : $scope.orderItem.bevName ,
			qty : $scope.orderItem.qty,milkType : $scope.orderItem.milkType,
			cupSize : $scope.orderItem.cupSize } ;
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
