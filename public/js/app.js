/* App Module */
var App = angular.module('App', [
	"ngRoute", 
	"ngResource",
	"ui.bootstrap",
	"bootstrapLightbox",
	"ngWebSocket",
	"ngRadialGauge"
]).constant("REST_API", {
        "HostName": "http://www.qrz.cat",
        "WsHostname": "ws://www.qrz.cat",
        "Port": 8088
})

App.filter('condicions', function () {
  return function (condicio) {
  	var res = condicio;
	switch(condicio){
		case "Poor":
			res = "Dolentes";
			break;
		case "Good":
			res = "Bones";
			break;
		case "Fair":
			res = "Amb posibilitats";
			break;
		case "day":
			res = "Durant el dia";
			break;
		case "night":
			res = "Durant la nit";
			break;
		case "Band Closed":
			res = "Banda tancada";
			break;
		default:
		break;
	}
    return res;
  };
});

/* Rutes */
App.config(['$routeProvider', '$httpProvider', function ($routeProvider, $httpProvider) {

	$routeProvider.
	when('/db/:indicatiu/:sufix', {
		templateUrl: 'views/mostraestacio.html',
		controller: 'MostraEstacioCtrl'
	}).
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
	when('/aprs', {
		templateUrl: 'views/aprs.html',
		controller: 'AprsCtrl'
	}).
	when('/entitats',{
		templateUrl: "views/entitats.html",
		controller: "EntitatsCtrl"
	}).
	otherwise({
		redirectTo: '/'
	});
}]);


