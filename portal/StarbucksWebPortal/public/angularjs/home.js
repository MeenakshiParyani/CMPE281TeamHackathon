var home = angular.module("home", [ "ngRoute" ]);
home.config(function($routeProvider) {
	// $locationProvider.html5Mode(true)
	$routeProvider.when("/some_url", {
		templateUrl : "templates/some_template.html",
		controller : "someCtrl"
	}).otherwise({
		templateUrl : "templates/welcome_page.html",
		controller : "welcomePageCtrl"
	});
});

home.controller("someCtrl", function($scope) {

});

home.controller("welcomePageCtrl", function($scope) {

});
