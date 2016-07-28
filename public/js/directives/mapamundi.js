App.directive("worldmap", function($window, $timeout) {
    return {
        restrict: "E",
        scope: {
            maptarget: '='
        },
        link: function(scope, ele, attrs) {
            var offline = L.tileLayer("/images/mapa/{z}/{x}/{y}.png", {
                maxZoom: 4
            });

            

            L.Icon.Default.imagePath = '/images';

            var map = L.map(ele[0], {
                center: [41.37, 1.81],
                zoom: 2,
                layers: offline
            });

            var marker = L.marker([41.37, 1.81]).addTo(map);

            scope.$watch("maptarget", function(newValue, oldValue) {
                if (!newValue || newValue == oldValue) return;

                var pos = L.latLng(newValue);
                if (!pos) return;

                map.panTo(pos);
                marker.setLatLng(pos);

                $timeout(function() {
                    map.invalidateSize(true);
                }, 200);
            });
      }}
});