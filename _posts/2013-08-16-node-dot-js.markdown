---
layout: post
title: Node.js
date: 2013-08-15 14:51
categories:
- javascript
- node.js
- web
---

So I started using Node.js a few days ago to work on a side project(details to come later). I could have chosen Rails since I'm more familiar with that but Node is the ideal language for this. This project is supposed to be a quick, real-time, but simple web app. Node is all about non-blocking I/O, meaning that as a command or method is running, the whole app doesn't stop and wait for that command to finish but instead the program continues. Rails is totally blocking I/O and you can see it while the database is querying because the page takes some time to load.

Node is very much asynchronous in that everything is happening almost at once. With that comes the 'callback' function. Callbacks are the node equivalent to blocks in Ruby. When you pass a callback, which is just a function, to a method call in node, that method will run and when it is done, your callback will be called. It's like saying "Go do this thing, and I don't care when you'll be done so do this[callback] as well when you're done." Great right? Yea but this asynchronocity has tripped me up in a couple of places. Most notably while setting up authentication. The user model has a class method `User.exists(username)` which checks if the given username is already in use when someone attempts signing up and then simply returns true or false.

The full method looks like this:
{% highlight coffeescript %}
exists: (username) ->
  this.findOne username, (err, user) ->
    if user
      return true
    else
      return false
{% endhighlight %}

And then in the 'controller,' I called the method expecting it to return that boolean and create the user if `User.exists(username)` was false.
{% highlight coffeescript %}
if User.exists username
  # don't create it
else
  # User.create...
{% endhighlight %}

Problems
---------------------
1. Using `if User.exists username` WILL be `undefined`, which is false-y and not true so the code block continues with creation.
2. Using `if User.exists username is true` will throw an error of comparison of an undefined type or something of the sort.
3. Even though my `exists` method was trying to return true, only the callback from finding the user returned true.
4. So trying to return the return of the callback still failed becaused due to asynchronocity, it didn't exist yet and we're back to problem 1.

I figured these out by printing to the console from inside the `User.exists()` and from inside the controller calling it. In the console, the controller printed the user didn't exist and continued and then a couple of lines later, the exists  method call would print that the user did in fact exist! My controller was saying to the method "I don't care when you're done, I'm going to continue doing work." Welcome to Node I guess.

Solution
--------------------
So this is where callbacks come in to play. Once I added a callback to the parameters of the exists method like so:
{% highlight coffeescript %}
exists: (username, callback) ->
  this.findOne username: username, (err, user) ->
    if user
      callback true
    else
      callback false
{% endhighlight %}

And the controller:
{% highlight coffeescript %}
User.exists username, (exists) ->
  if exists
    # Don't create it
  else
    # User.create...
{% endhighlight %}

If you followed what the exists method was doing earlier, it calls the `findOne()` method which will find the first thing in the database matching given criteria and it also takes a callback. This means, my method kept going while trying to find the user in the database. By taking a callback in my method, I transferred what the controller was doing to happen when the database query was finished. A callback in a callback in a call. Wow...I love programming and I'm loving javascript even more after a few days of Node.js.
