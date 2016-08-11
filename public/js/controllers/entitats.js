App.controller('EntitatsCtrl', function($scope,$http,REST_API){

	$scope.entitat = {};
	$scope.entitats = [];
	var mapTargets = [null, null, null];
	$scope.dxcc = {
		id: "3"
	}

	var getEntitats = function(){

		var urlbase = REST_API.HostName + ":" + REST_API.Port;

		var url = urlbase + "/api/dxcc";

		$http.get(url)
		.success(function(data){
			
			$scope.entitats = (data.QRZDatabase.DXCC);
			$scope.entitat = $scope.entitats[3];

		});
	}

	getEntitats();

	entitatChanged = function (id) {
		
		angular.forEach($scope.entitats, function(entitat, key) {
			angular.forEach(entitat, function (valor, clau) {

				if ( clau == "dxcc" && valor == id){
					$scope.entitat = entitat;
					$scope.mostrarMapa();
				}
			});
		});
	}

	function setMapTarget(coord, priority) {
		mapTargets[priority] = coord;
		for (var i = 0; i < mapTargets.length; i++) {
			if (mapTargets[i]) {
				updateMapTarget(mapTargets[i]);
				return;
			}
		}
		updateMapTarget(null);
	}

	function updateMapTarget(coord) {
		$scope.maptarget = coord;
	}

	$scope.mostrarMapa = function(){

		var lat = $scope.entitat.lat;
		var lon = $scope.entitat.lon;

		var coord = [lat,lon];	
		setMapTarget(coord,2);
	}


	$scope.$watch("dxcc.id", entitatChanged);

});