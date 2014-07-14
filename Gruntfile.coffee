module.exports = (grunt) ->
  grunt.initConfig
    sass:
      dist:
        options:
          style: 'compressed'
          noCache: true
        files:
          'css/styles.css': '_sass/styles.scss'
          'css/syntax.css': '_sass/syntax.scss'

  grunt.loadNpmTasks 'grunt-contrib-sass'
  grunt.registerTask 'default', ['sass']
