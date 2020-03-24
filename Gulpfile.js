require('es6-promise').polyfill();
const spawn = require('child_process').spawn;

/* ****************************************************************************************************
 *  Variables                                                                                          *
 **************************************************************************************************** */

const confGlobal = require('./gulp/config/gulp-global.json');
const confPlugins = require('./gulp/config/gulp-plugins.json');
const exec = require('child_process').exec;
const gulp = require('gulp');
const plugins = require("gulp-load-plugins")({
    pattern: ['gulp-*', 'gulp.*'],
    replaceString: /\bgulp[\-.]/
});
const runSequence = require('gulp4-run-sequence');
const gulpif = require('gulp-if');
const del = require('del');

const imagemin = require('gulp-imagemin');

const uglify_js = require('uglify-es');
const composer = require('gulp-uglify/composer');
const pump = require('pump');

const minify = composer(uglify_js, console);

/**
 * This solution to execute and watch a shell function from Gulp is
 * adapted from https://stackoverflow.com/a/10232330/3232832
 *
 * e.g. callSpawn( 'ping', [ '-c 5', 'google.com' ], cb );
 */
function callSpawn(command, arguments, cb) {
    const call = spawn(command, arguments);

    // stdout
    call.stdout.on('data', function (data) {
        // like console.log, but without the trailing newline
        process.stdout.write(data.toString());
    });

    // stderr
    call.stderr.on('data', function (data) {
        console.log(data.toString());

        cb(data.toString());
    });

    call.on('exit', function (code) {
        console.log('child process exited with code ' + code.toString());

        cb();
    });
}


/* ****************************************************************************************************
 *                                                                                                     *
 *  MAIN TASKS                                                                                         *
 *                                                                                                     *
 **************************************************************************************************** */

gulp.task('dev', async function () {
    confGlobal.isDevelop = true;
    runSequence(['js', 'css', 'copy:assets'], ['watch', 'hugo:server']);
});

gulp.task('dev:nowatch', async function () {
    confGlobal.isDevelop = true;
    runSequence(['js', 'css', 'copy:assets'], 'hugo:server:nowatch');
});

gulp.task('prod', async function () {
    confGlobal.isDevelop = false;
    runSequence('clean', ['js', 'css', 'copy:assets:minify'], 'css:clean', 'hugo:build');
});

/* ****************************************************************************************************
 *                                                                                                     *
 *  SUBTASKS                                                                                           *
 *                                                                                                     *
 **************************************************************************************************** */

gulp.task('js', function () {

    var sourcemaps = require('gulp-sourcemaps');

    return gulp.src([
            'node_modules/jquery/dist/jquery.js',
            './assets/js/**/*.js'
        ])
        .pipe(plugins.plumber({
            handleError: function (err) {
                console.log(err);
                this.emit('end');
            }
        }))
        .pipe(sourcemaps.init())
        .pipe(plugins.jshint(confPlugins.jshintOptions))
        .pipe(plugins.jshint.reporter('jshint-stylish'))
        .pipe(plugins.concat('app.js'))
        .pipe(gulpif(!confGlobal.isDevelop, minify({
            mangle: true
        })))
        .pipe(gulpif(!confGlobal.isDevelop, plugins.stripDebug()))
        .pipe(gulpif(confGlobal.enableGZIP, plugins.gzip(confPlugins.gzipOptions)))
        .pipe(sourcemaps.write('./static/js/', {
            includeContent: false
        }))
        .pipe(gulp.dest('./static/js/'));
});

gulp.task('css', function () {
    var autoprefixer = require('autoprefixer');
    var cssgrace = require('cssgrace');
    var pseudoelements = require('postcss-pseudoelements');
    var cssnano = require('cssnano');

    var processors = [
        autoprefixer(),
        //cssgrace,
        pseudoelements
    ];

    if (!confGlobal.isDevelop) {
        processors = [
            autoprefixer(),
            //cssgrace,
            pseudoelements,
            cssnano
        ];
    }

    return gulp.src('./assets/css/*.css')
        .pipe(plugins.plumber({
            handleError: function (err) {
                console.log(err);
                this.emit('end');
            }
        }))
        // .pipe(plugins.scssLint(confPlugins.scssLint))
        .pipe(plugins.sass())
        .pipe(plugins.postcss(processors))
        .pipe(gulpif(confGlobal.enableGZIP, plugins.gzip(confPlugins.gzipOptions)))
        .pipe(gulp.dest('./static/css/'));
});

gulp.task('css:clean', function () {
    console.log('Removing unused css styles...');
    return gulp.src('./public/css/*.css')
        .pipe(gulpif(!confGlobal.isDevelop, plugins.uncss({
            html: './public/**/*.html'
        })))
        .pipe(gulp.dest('./public/css/'));
});

gulp.task('watch', function () {
    gulp.watch('./assets/css/**/*.css', gulp.parallel('css'));
    gulp.watch('./assets/js/**/*.js', gulp.parallel('js'));
    gulp.watch('./assets/images/**/*.+(jpg|jpeg|gif|png)', gulp.parallel('copy:assets'));
});


/* ****************************************************************************************************
 *                                                                                                     *
 *  HELPERS                                                                                            *
 *                                                                                                     *
 **************************************************************************************************** */

gulp.task('copy:assets', function () {
    return gulp.src(['./assets/images/**/*.+(jpg|jpeg|gif|png)'])
        .pipe(gulp.dest('./static/images/'));
});

gulp.task('copy:assets:minify', function () {
    return gulp.src(['./assets/images/**/*.+(jpg|jpeg|gif|png)'])
        .pipe(gulp.dest('./static/images/'));
});

gulp.task('clean', function () {
    console.log('Deleting public folder...');
    return del('./public/');
});

gulp.task('hugo:server', function (cb) {
    exec("hugo serve --watch --enableGitInfo --forceSyncStatic --gc --cleanDestinationDir --ignoreCache --templateMetrics --templateMetricsHints --buildFuture --buildDrafts --disableFastRender", function (err, stdout, stderr) {
        console.log(stdout);
        console.log(stderr);
        cb(err);
    })
});

gulp.task('hugo:server:nowatch', function (cb) {
    exec('hugo serve --watch=false', function (err, stdout, stderr) {
        console.log(stdout);
        console.log(stderr);
        cb(err);
    })
});

gulp.task('hugo:build', function (cb) {
    exec('hugo', function (err, stdout, stderr) {
        console.log(stdout);
        console.log(stderr);
        cb(err);
    })
});

// runSequence(['js','css','copy:assets'], 'watch', 'hugo:server');
gulp.task('default', async function () {
    runSequence('clean', ['js', 'css', 'copy:assets'], ['watch', 'hugo:server'])
});