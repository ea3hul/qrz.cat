App.controller('PropagacioCtrl', function($scope,$http){

	$scope.solar = {};

	var getSolar = function(){

		var url = "http://localhost:8088/api/solar/activitat";

			$http.get(url)
			.success(function(data){
				$scope.solar = data;

			});
	}

	getSolar();
  
});