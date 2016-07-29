App.controller('DashBoardCtrl', function($scope,$http){

	$scope.solar = {};

	var getSolar = function(){

		var url = "https://www.dxcluster.co.uk/index.php?/api/solar";

			$http.get(url)
			.success(function(data){
				console.log(data);
				$scope.solar = data;

			});
	}

	getSolar();

});