angular.module('starbucks').controller('myBagCtrl', [ '$rootScope', '$scope', '$http', '$state', function($rootScope, $scope, $http, $state) {
	console.log($rootScope.bagItems);
	$scope.placeOrder = function(){
		console.log("Order Placed");

	// 	$http({
	// 		method : "POST",
	// 		url : 'http://localhost:8080/order',
	// 		data : {
	// 			"email_id" : $scope.email_id,
	// 			"passwd" : $scope.passwd
	// 		}
	// 	}).success(function(data) {
	// 		//checking the response data for statusCode
	// 		if (data.statusCode == 401) {
	// 			$scope.invalid_login = false;
	// 			$scope.validlogin = true;

	// 		} else {

	// 			$.sessionTimeout({
	// 				logoutUrl : '/logout',
	// 				warnAfter : 5 * 60 * 1000,
	// 				redirAfter : 30 * 60 * 1000,

	// 				onRedir : function() {
	// 					console.log("onRedir :Your session has expired");

	// 				}
	// 			});

	// 			$rootScope.totalCostOfCart = data.totalCostOfCart;
	// 			$rootScope.cartItems = JSON.parse(data.cartItems);

	// 			$rootScope.user_dtls = JSON.parse(data.userDtls);
	// 			console.log("In client : userDtls : Dtls as string :" + $rootScope.user_dtls);
	// 			redirectToHomePage();
	// 			$state.transitionTo('ebayHome');
	// 		}
	// 		//Making a get call to the '/redirectToHomepage' API
	// 		//window.location.assign("/homepage"); 
	// 	}).error(function(error) {
	// 		$scope.validlogin = true;
	// 		$scope.invalid_login = true;
	// 	});
	}

}]);