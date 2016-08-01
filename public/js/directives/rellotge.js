App.directive("rellotge", function($interval,dateFilter) {
    function link($scope, element, attrs) {
        var update = function() {
            element.text(dateFilter(new Date(), "HH:mm:ss", "UTC"));
        };

        var timer = $interval(update, 1000);

        element.bind("$destroy", function() {
            $interval.cancel(timer);
        });

        update();
    }

    return ({ link: link, scope: false });
});

App.directive("rellotgelocal", function($interval,dateFilter) {
    function link($scope, element, attrs) {
        var update = function() {
            element.text(dateFilter(new Date(), "HH:mm:ss"));
        };

        var timer = $interval(update, 1000);

        element.bind("$destroy", function() {
            $interval.cancel(timer);
        });

        update();
    }

    return ({ link: link, scope: false });
});
