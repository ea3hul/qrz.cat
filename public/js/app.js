/* App Module */
var App = angular.module('App', [
	"ngRoute", 
	"ngResource",
	"ui.bootstrap",
	"bootstrapLightbox"
]);

/* Rutes */
App.config(['$routeProvider', '$httpProvider', function ($routeProvider, $httpProvider) {

	$routeProvider.
	when('/db/:indicatiu', {
		templateUrl: 'views/mostraestacio.html',
		controller: 'MostraEstacioCtrl'
	}).
	otherwise({
		redirectTo: '/'
	});
}]);

