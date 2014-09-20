module.exports = function (grunt) {
    grunt.loadNpmTasks('grunt-contrib-requirejs');
    grunt.loadNpmTasks('grunt-react');
    grunt.loadNpmTasks('grunt-contrib-copy');
    grunt.loadNpmTasks('grunt-contrib-cssmin');
    grunt.loadNpmTasks('grunt-contrib-clean');
    grunt.loadNpmTasks('grunt-contrib-watch');

    var requireConf = require('./requireConfig.js');

    grunt.initConfig({
        requirejs: {
            dev : {
                options: requireConf
            }
        },
        react: {
            buildJSXDeploy: {
                files: [
                    {
                        expand: true,
                        src: ['./js/**/*.js'],
                        dest: './temp/',
                        ext: '.js'
                    }
                ]
            },
            buildJSXDev: {
                files: [
                    {
                        expand: true,
                        src: ['./js/**/*.js'],
                        dest: './build/',
                        ext: '.js'
                    }
                ]
            }
        },
        copy: {
            copyLibs: {
                files: [
                    {
                        src: ['./bower_components/jquery/dist/jquery.min.js', './bower_components/react/react.min.js', './bower_components/moment/min/moment.min.js', './bower_components/history.js/scripts/bundled/html5/native.history.js'],
                        dest: './temp/js/'
                    }
                ]
            },
            copyLibsDev: {
                files: [
                    {
                        src: ['./bower_components/jquery/dist/jquery.min.js', './bower_components/react/react.js', './bower_components/moment/min/moment.min.js', './bower_components/history.js/scripts/bundled/html5/native.history.js'],
                        dest: './build/js/'
                    }
                ]
            },
            copyToBuild: {
                files: [
                    {
                        src: ['./index.html', './node_modules/requirejs/require.js', './imgs/**/*', './app.yaml'],
                        dest: './build/'
                    }
                ]
            }
        },
        cssmin: {
        	combine: {
				files: {
					'./build/css/site.css': ['./css/**/*.css']
				}
			},
            minify: {
                expand: true,
                cwd: './css/',
                src: ['*.css', '!*.min.css'],
                dest: 'build/css/',
                ext: '.min.css'
            }
        },
        clean: {
            temp: {
                src: ["./temp"]
            },
            build: {
                src: ["./build"]
            }
        },
        watch: {
            css: {
                files: ['./index.html','./css/**/*.css','./js/**/*.js'],
                tasks: ['dev'],
                options: {
                    spawn: false,
                }
            }
        }
    });

    grunt.registerTask('default', ['clean:build', 'copy:copyLibs', 'react:buildJSXDeploy', 'requirejs:dev', 'cssmin:combine',  'copy:copyToBuild', 'clean:temp']);
    grunt.registerTask('deploy', ['clean:build', 'copy:copyLibs', 'react:buildJSXDeploy', 'requirejs:dev', 'cssmin:combine',  'copy:copyToBuild', 'clean:temp']);
    grunt.registerTask('dev', ['clean:build', 'react:buildJSXDev', 'cssmin:combine', 'copy:copyToBuild', 'copy:copyLibsDev']);
    grunt.registerTask('dev-watch', ['watch:css']);
}