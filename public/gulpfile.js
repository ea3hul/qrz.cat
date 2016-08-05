var gulp = require('gulp');
var $ = require('gulp-load-plugins')();
var BOWER_DIR = "bower_components/";

gulp.task('app.css', function(){
  return gulp.src("css/*.css")
  .pipe($.minifyCss())
  .pipe($.concat("qrz.cat.min.css"))
  .pipe(gulp.dest("dist/css"));
});

gulp.task('app.js', function(){
  return gulp.src('js/**/*.js')
  .pipe($.plumber())
  .pipe($.order([
    "js/app.js"
    ]))
  .pipe($.sourcemaps.init())
  .pipe($.ngAnnotate())
  .pipe($.uglify())
  .pipe($.concat("qrz.cat.min.js"))
  .pipe($.sourcemaps.write())
  .pipe(gulp.dest("dist/js"))
  .pipe($.livereload());
});

gulp.task('icons', function() {
  return gulp.src([
   BOWER_DIR + '/font-awesome/fonts/**.*',
   BOWER_DIR + '/bootstrap/fonts/**.*',
   BOWER_DIR + '/weather-icons/font/**.*'
   ]) 
  .pipe(gulp.dest('dist/fonts'));
});

gulp.task('vendor.css', function(){
  return gulp.src([
    BOWER_DIR + "bootstrap-toggle/css/bootstrap-toggle.min.css",
    BOWER_DIR + "bootstrap/dist/css/bootstrap.css",
    BOWER_DIR + "font-awesome/css/font-awesome.css",
    BOWER_DIR + "angular-bootstrap-lightbox/dist/angular-bootstrap-lightbox.css",
    BOWER_DIR + "leaflet/dist/leaflet.css"
    ])
  .pipe($.minifyCss())
  .pipe($.concat("vendor.min.css"))
  .pipe(gulp.dest("dist/css"));
});

gulp.task("vendor.js", function() {
  return gulp.src([
    BOWER_DIR + "jquery/dist/jquery.min.js",
    BOWER_DIR + "bootstrap/dist/js/bootstrap.min.js",
    BOWER_DIR + "bootstrap-toggle/js/bootstrap-toggle.min.js",
    BOWER_DIR + "angular/angular.min.js",
    BOWER_DIR + "angular-route/angular-route.min.js",
    BOWER_DIR + "angular-resource/angular-resource.min.js",
    BOWER_DIR + "angular-websocket/dist/angular-websocket.min.js",
    BOWER_DIR + "leaflet/dist/leaflet.js",
    BOWER_DIR + "socket.io-client/socket.io.js",
    BOWER_DIR + "angular-bootstrap/ui-bootstrap-tpls.min.js",
    BOWER_DIR + "suncalc/suncalc.js",
    BOWER_DIR + "angular-bootstrap-lightbox/dist/angular-bootstrap-lightbox.js",
    BOWER_DIR + "d3/d3.js",
    BOWER_DIR + "ngRadialGauge/src/ng-radial-gauge-dir.js",
    BOWER_DIR + "moment/moment.js",
    BOWER_DIR + "moment/locale/ca.js"
    ])
  .pipe($.plumber())
  .pipe($.uglify())
  .pipe($.concat("vendor.min.js"))
  .pipe(gulp.dest("dist/js"));
});

gulp.task("default", [
  "app.js",
  "app.css",
  "vendor.js",
  "vendor.css",
  "icons"
  ]);

gulp.task("watch", function() {
  $.livereload.listen();
  gulp.watch("js/*.js", ["app.js"]);
  gulp.watch("js/**/*.js", ["app.js"]);
  gulp.watch("css/*.css", ["app.css"]);
});