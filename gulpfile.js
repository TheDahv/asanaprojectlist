var gulp = require('gulp');

var concat = require('gulp-concat');
var sourcemaps = require('gulp-sourcemaps');
var sass = require('gulp-sass');

gulp.task('scripts', function () {
  return gulp.src([
      '_assets/js/lib/*.js',
      '_assets/js/*.js'
    ])
    .pipe(sourcemaps.init())
    .pipe(concat('application.js'))
    .pipe(sourcemaps.write())
    .pipe(gulp.dest('webapp/js'));
});

gulp.task('sass', function () {
  return gulp.src('./_assets/sass/**/*.scss')
    .pipe(sass({
      'source-comments': 'map'
    }))
    .pipe(gulp.dest('./webapp/css'));
});

gulp.task('watch', function () {
  gulp.watch(['_assets/sass/**/*.scss'], ['sass']);
  gulp.watch(['_assets/js/**/*.scss'], ['scripts']);
});
