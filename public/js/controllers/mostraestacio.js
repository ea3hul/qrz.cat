App.controller('MostraEstacioCtrl', function($scope,$routeParams,$http,$window,$sce,Lightbox,REST_API){

	var mapTargets = [null, null, null];

	var indicatiu = $routeParams.indicatiu;

	$scope.alerts = [];

    $scope.closeAlert = function(index) {
        $scope.alerts.splice(index, 1);
    };

	$scope.images = [];

	$scope.session = {};
	$scope.callsign = {};
	$scope.dxcc = {};
	$scope.astro = {};
	$scope.poblacio = {};

	var urlbase = REST_API.HostName + ":" + REST_API.Port;

	var url = urlbase + "/api/qrz/";
	var urlent = urlbase + "/api/dxcc/";
	var urlpb = urlbase + "/api/geo/poblacio/";


	$http({
		method: "GET",
		url: url + indicatiu +"/bio"
	}).then(function successCallback(response) {

		var htmldata = response.data.split("Â").join("").split("Ã³").join("ó").split("Ã").join("í")
		.split("í±").join("ñ").split("í¡").join("á").split("í©").join("é").split("í ").join("à")
		.split("í¨").join("è").split("â€“").join("-").split("í“").join("Ó");

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

					if ($scope.dxcc.name == "Spain" || $scope.dxcc.name == "Balearic Islands"  || $scope.dxcc.name == "Ceuta and Melilla"  || $scope.dxcc.name == "Canary Islands"){

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
								case "Canarias":
								$scope.dxcc.cc = "_CN";
								break;
								case "Ceuta y Melilla":
								$scope.dxcc.cc = "_CE";
								break;
								case "Islas Baleares":
								$scope.dxcc.cc = "IB";
								default:
								break;

							}
							if ($scope.callsign.call.indexOf("3") !== -1){
								$scope.dxcc.cc = "CT";
							}

							if ($scope.callsign.call.indexOf("7") !== -1){
								$scope.dxcc.cc = "_AN";
							}
						}, function errorCallback (response){
							console.log(response);
						});
					}

					

					$scope.astro = SunCalc.getTimes(Date.now(), $scope.callsign.lat, $scope.callsign.lon);

				}else{
					$scope.alerts.push({ type: "danger", msg: $scope.session.Error});
				}

			}, function errorCallback(response) {
				console.log(response);
			});
		}else{
			$scope.alerts.push({ type: "danger", msg: $scope.session.Error});
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