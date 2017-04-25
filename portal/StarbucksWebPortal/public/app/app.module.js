var starbucks = angular.module('starbucks', ['ui.router']);

starbucks.config(function($stateProvider, $urlRouterProvider, $locationProvider) {
  $locationProvider.html5Mode({
    enabled: true,
    requireBase: false
  });

  $stateProvider.state('home', {
    url: '/',
    component: 'home'
  });

  $stateProvider.state('order', {
    url: '/order',
    component: 'order'
  });

  $stateProvider.state('bag', {
    url: '/bag',
    component: 'bag'
  });

  $stateProvider.state('pay', {
    url: '/pay',
    component: 'pay'
  });

  $urlRouterProvider.otherwise('/');
});

starbucks.run(function() {});