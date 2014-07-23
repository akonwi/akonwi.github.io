module.exports = (grunt) ->
  grunt.initConfig
    sass:
      dist:
        options:
          style: 'compressed'
          noCache: true
        files:
          'css/styles.css': '_sass/styles.sass'
          'css/syntax.css': '_sass/syntax.sass'

  grunt.loadNpmTasks 'grunt-contrib-sass'
  grunt.registerTask 'default', ['sass']
