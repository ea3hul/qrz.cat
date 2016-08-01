App.controller('DashBoardCtrl', function($scope,$http){

	$scope.solar = {};


	var getSolar = function(){

		var url = "http://www.qrz.cat:8088/api/solar/activitat";

		var verd = "#8DCA2F";
		var groc = "#FDC702";
		var taronja = "#FF7700";
		var vermell = "#C50200";
		var blau = "#0000FF";

		$http.get(url)
		.success(function(data){
			$scope.solar = data.solar.solardata;

			$scope.kvalue = $scope.solar.kindex;

			$scope.kupperLimit = 9;
			$scope.klowerLimit = 0;
			$scope.kunit = "";
			$scope.kprecision = 2;

			$scope.kranges = [
			{
				min: 0,
				max: 2.99,
				color: verd
			},
			{
				min: 3,
				max: 4.99,
				color: verd
			},
			{
				min: 5,
				max: 5.99,
				color: groc
			},
			{
				min: 6,
				max: 6.99,
				color: groc
			},
			{
				min: 7,
				max: 7.99,
				color: vermell
			},

			{
				min: 8,
				max: 8.99,
				color: vermell
			},

			{
				min: 9,
				max: 9,
				color: vermell
			}

			];

			$scope.avalue = $scope.solar.aindex;

			$scope.aupperLimit = 400;
			$scope.alowerLimit = 0;
			$scope.aunit = "";
			$scope.aprecision = 2;

			$scope.aranges = [
			{
				min: 0,
				max: 400,
				color: blau
			}];

			$scope.fluxvalue = $scope.solar.solarflux;

			$scope.fluxupperLimit = 300;
			$scope.fluxlowerLimit = 62.5;
			$scope.fluxunit = "";
			$scope.fluxprecision = 2;

			$scope.fluxranges = [
			{
				min: 62.5,
				max: 79.99,
				color: vermell
			},
			{
				min: 80,
				max: 104.99,
				color: groc
			},
			{
				min: 105,
				max: 300,
				color: verd
			}
			];

			$scope.spotsvalue = $scope.solar.sunspots;

			$scope.spotsupperLimit = 250;
			$scope.spotslowerLimit = 0;
			$scope.spotsunit = "";
			$scope.spotsprecision = 2;

			$scope.spotsranges = [
			{
				min: 0,
				max: 250,
				color: blau
			}];


		});
	}

	getSolar();

});