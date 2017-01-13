---
layout: post
title: Custom Tasks In Grunt
date: 2017-01-13 13:45
categories:
- javascript
- grunt
---

I've been using [Grunt](https://github.com/gruntjs/grunt) for a while now. It's been such a great experience that I haven't touched any other build tool. There hasn't been a need for me to explore Yeoman or Webpack. My `Gruntfile` structure has stayed pretty consistent for the projects I've used it in. The first task that I always install is [grunt-contrib-watch](https://github.com/gruntjs/grunt-contrib-watch) so I can run things like sass and coffeescript compilation when files change.

Up until this past week, that's all I've done with grunt and I've been re-running tests manually. In one of my projects, I have a few scripts configured in my package.json to create my distribution files and run tests. I wanted to do those two things again when files changed so I was curious to see if I could attach them as part of the grunt watch configuration I had. After a few minutes of reading the grunt docs on the website and seeing how the grunt-contrib-sass task was implemented, I set up a couple new tasks in my Gruntfile with less than 10 lines of code and no new dependencies.

Here's my task for running tests:

``` coffeescript
{spawnSync} = require 'child_process'

grunt.registerTask 'test', () ->
  {status} = spawnSync 'npm', ['test'], stdio: 'inherit'
  status is 0
```

As you can see it's pretty simple. The `spawnSync` function is from Node's [child_process](https://nodejs.org/api/child_process.html) so there's nothing new to install though it needs to be `required`. I won't go into the details of the `spawnSync` method signature but all this is doing is calling my npm script in a shell process and then returning whether the job finished successfully (i.e. the status code is 0). It's important to return false if the task fails so that if other tasks are chained after this one, then they won't be invoked if the current one fails. In my setup, the task to browserify my app follows this and I don't want that to happen if tests are failing.

With a new task registered it can be invoked by the watch task simply by putting the task name in the watch configuration:

``` coffeescript
watch:
  js:
    files: ['src/**/**.js', '__tests__/**/**.js']
    tasks: ['test', 'browserify', 'dist:dev']
```

Now the tests will run first when javascript files change and if they pass, then the app will be browserified and the 'dist:dev' task moves files to a different directory.

This was a small win for me and I'm happy I can set up simple tasks like this with a few lines of code and no extra dependencies in my package.json. As long as I don't need hot reloading of my code, I can stay away from webpack.
