App.controller('AprsCtrl', function($scope,wsData){

	$scope.aprs = [];

	$scope.aprs = $scope.aprs.push(wsData.methods);

	console.log($scope.aprs);

});