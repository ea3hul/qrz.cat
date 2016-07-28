App.controller('CercarCtrl', function($scope,$location){

	$scope.cerca = function(){
 		$location.path("/db/" + $scope.indicatiu);
	}
  
});