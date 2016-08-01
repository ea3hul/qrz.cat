/* App Module */
var App = angular.module('App', [
	"ngRoute", 
	"ngResource",
	"ui.bootstrap",
	"bootstrapLightbox",
	"ngRadialGauge"
]);

/* Rutes */
App.config(['$routeProvider', '$httpProvider', function ($routeProvider, $httpProvider) {

	$routeProvider.
	when('/db/:indicatiu', {
		templateUrl: 'views/mostraestacio.html',
		controller: 'MostraEstacioCtrl'
	}).
	when('/solar',{
		templateUrl:'views/propagacio.html',
		controller: 'PropagacioCtrl'
	}).
	when('/', {
		templateUrl: 'views/dashboard.html',
		controller: 'DashBoardCtrl'
	}).
		when('/quant', {
		templateUrl: 'views/quant.html'
	}).
	when('/repetidors', {
		templateUrl: 'views/repetidors.html'
	}).
	otherwise({
		redirectTo: '/'
	});
}]);

