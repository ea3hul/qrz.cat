App.controller('PropagacioCtrl', function($scope,$http,REST_API){

	$scope.solardata = {};
	
	var getSolar = function(){

		var urlbase = REST_API.HostName + ":" + REST_API.Port;

		var url = urlbase + "/api/solar/activitat";

		$http.get(url)
		.success(function(data){
			
			$scope.solardata = data.solar.solardata;
		});
	}

	getSolar();

});