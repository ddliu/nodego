var gulp = require('gulp');
var clean = require('gulp-clean');
var uglify = require('gulp-uglify');
var remoteSrc = require('gulp-remote-src');
var bindata = require('gulp-gobin');

var FILES = [
    "src/node.js",
    "lib/_debugger.js",
    "lib/_linklist.js",
    "lib/_stream_duplex.js",
    "lib/_stream_passthrough.js",
    "lib/_stream_readable.js",
    "lib/_stream_transform.js",
    "lib/_stream_writable.js",
    "lib/assert.js",
    "lib/buffer.js",
    "lib/child_process.js",
    "lib/cluster.js",
    "lib/console.js",
    "lib/constants.js",
    "lib/crypto.js",
    "lib/dgram.js",
    "lib/dns.js",
    "lib/domain.js",
    "lib/events.js",
    "lib/freelist.js",
    "lib/fs.js",
    "lib/http.js",
    "lib/https.js",
    "lib/module.js",
    "lib/net.js",
    "lib/os.js",
    "lib/path.js",
    "lib/punycode.js",
    "lib/querystring.js",
    "lib/readline.js",
    "lib/repl.js",
    "lib/stream.js",
    "lib/string_decoder.js",
    "lib/sys.js",
    "lib/timers.js",
    "lib/tls.js",
    "lib/tty.js",
    "lib/url.js",
    "lib/util.js",
    "lib/vm.js",
    "lib/zlib.js"
];

var URL = "https://raw.githubusercontent.com/joyent/node/v0.11.13/"

gulp.task('bindata', function() {
    return remoteSrc(FILES, {base: URL})
        .pipe(uglify({
            compress: false // src/node.js might get lost after compression
        }))
        .pipe(bindata('bindata.go', {package: 'nodego'}))
        .pipe(gulp.dest('../'))
});

gulp.task('default', ['bindata']);