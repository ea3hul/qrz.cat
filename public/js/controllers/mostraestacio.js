App.controller('MostraEstacioCtrl', function($scope,$routeParams,$http,$window,$sce,Lightbox){

	var mapTargets = [null, null, null];

	var indicatiu = $routeParams.indicatiu;

	$scope.images = [];

	$scope.session = {};
	$scope.callsign = {};
	$scope.dxcc = {};
	$scope.astro = {};
	$scope.poblacio = {};

	var url = "http://www.qrz.cat:8088/api/qrz/";
	var urlent = "http://www.qrz.cat:8088/api/dxcc/";
	var urlpb = "http://www.qrz.cat:8088/api/geo/poblacio/"

	$http({
		method: "GET",
		url: url + indicatiu +"/bio"
	}).then(function successCallback(response) {

		var htmldata = response.data.split("Â").join("").split("Ã³").join("ó").split("Ã").join("í")
		.split("í±").join("ñ").split("í¡").join("á").split("í©").join("é").split("í ").join("à")
		.split("í¨").join("è");

		$scope.htmlString = $sce.trustAsHtml(htmldata);
		
		
	}, function errorCallback(response) {
		console.log(response);
	});

	$http({
		method: "GET",
		url: url + indicatiu
	}).then(function successCallback(response) {
		$scope.session = response.data.QRZDatabase.Session;
		if ($scope.session.Error == undefined){
			$scope.callsign = response.data.QRZDatabase.Callsign;
			$http({
				method: "GET",
				url: urlent + $scope.callsign.dxcc
			}).then(function successCallback(response) {
				$scope.session = response.data.QRZDatabase.Session;
				if ($scope.session.Error == undefined){
					$scope.dxcc = response.data.QRZDatabase.DXCC;

					$scope.images = [
					{
						'url': $scope.callsign.image,
						'thumbUrl': $scope.callsign.image,
						'caption': $scope.callsign.call
					}];

					if ($scope.dxcc.name == "Spain"){

						$http({
							method:"GET",
							url: urlpb +  $scope.callsign.lat +"/"+ $scope.callsign.lon
						}).then(function successCallback(response) {

							$scope.poblacio = response.data;

							
							switch ($scope.poblacio.comunitat){
								case "Catalunya":
								$scope.dxcc.cc = "CT";
								break;
								case "Andalucía":
								$scope.dxcc.cc = "_AN";
								break;
								case "Valencia":
								$scope.dxcc.cc = "_VC";
								break;
								case "Madrid":
								$scope.dxcc.cc = "_M";
								break;
								case "Murcia":
								$scope.dxcc.cc = "_MU";
								break;
								case "País Vasco":
								$scope.dxcc.cc = "_PV";
								break;
								case "Aragón":
								$scope.dxcc.cc = "_AR";
								break;
								case "Navarra":
								$scope.dxcc.cc = "_NA";
								break;
								case "Cantabria":
								$scope.dxcc.cc = "_S";
								break;
								case "Galicia":
								$scope.dxcc.cc = "_GA";
								break;
								case "La Rioja":
								$scope.dxcc.cc = "_LO";
								break;
								case "Asturias":
								$scope.dxcc.cc = "_AS";
								break;
								case "Castilla León":
								$scope.dxcc.cc = "_CL";
								break;
								case "Castilla La Mancha":
								$scope.dxcc.cc = "_CM";
								break;
								case "Extremadura":
								$scope.dxcc.cc = "_EX";
								break;	

								default:
								break;

							}

						}, function errorCallback (response){
							console.log(response);
						});
					}

					if ($scope.dxcc.name == "Balearic Islands"){
						$scope.dxcc.cc = "_IB";
					}

					if ($scope.dxcc.name == "Ceuta and Melilla"){
						$scope.dxcc.cc = "_CL";
					}

					if ($scope.dxcc.name == "Canary Islands"){
						$scope.dxcc.cc = "_CN";
					}

					$scope.astro = SunCalc.getTimes(Date.now(), $scope.callsign.lat, $scope.callsign.lon);

				}else{
					$window.alert($scope.session.Error);
				}

			}, function errorCallback(response) {
				console.log(response);
			});
		}else{
			$window.alert($scope.session.Error);
		}
		
	}, function errorCallback(response) {
		console.log(response);
	});

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

		var lat = $scope.callsign.lat;
		var lon = $scope.callsign.lon;

		var coord = [lat,lon];	
		setMapTarget(coord,2);
	}

	$scope.openLightboxModal = function (index) {
		Lightbox.openModal($scope.images, index);
	}


});