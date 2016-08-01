App.controller('PropagacioCtrl', function($scope,$http){

	$scope.solardata = {};
	
	var getSolar = function(){

		var url = "http://www.qrz.cat:8088/api/solar/activitat";

		$http.get(url)
		.success(function(data){
			
			$scope.solardata = data.solar.solardata;
		});
	}

	getSolar();

});