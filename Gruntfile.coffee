module.exports = (grunt) ->
  grunt.initConfig
    sass:
      dist:
        options:
          style: 'compressed'
          noCache: true
        files:
          'css/styles.css': '_sass/styles.sass'
          # 'css/syntax.css': '_sass/syntax.sass'
    watch:
      source:
        files: '**/**.sass'
        tasks: 'sass'

  grunt.loadNpmTasks 'grunt-contrib-sass'
  grunt.loadNpmTasks 'grunt-contrib-watch'
  grunt.registerTask 'default', ['watch']
