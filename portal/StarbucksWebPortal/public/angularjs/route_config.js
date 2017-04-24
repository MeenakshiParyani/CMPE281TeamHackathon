var starbucks = angular.module('starbucks', [ 'ui.router' ]);
starbucks.config(function($stateProvider, $urlRouterProvider, $locationProvider) {
	$locationProvider.html5Mode({
  		enabled: true,
 		requireBase: false
});
	$stateProvider.state('home', {
		url : '/',
		views : {
			'header' : {
				templateUrl : 'templates/header.html',
				controller : 'headerCtrl'
			},
			'content' : {
				templateUrl : 'templates/welcome_page.html',
				controller : 'welcomeCtrl'
			}
		}
		
	});
	$stateProvider.state('order', {
		url : '/',
		views : {
			'header' : {
				templateUrl : 'templates/header.html',
				controller : 'headerCtrl'
			},
			'content' : {
				templateUrl : 'templates/order.html',
				controller : 'orderCtrl'
			}
		}
		
	});

	$urlRouterProvider.otherwise('/');
});

starbucks.run([ '$state','$rootScope', function($state,$rootScope) {
	$state.transitionTo('home');
	$rootScope.store_location  = "Location" ;

} ]);