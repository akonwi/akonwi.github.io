---
layout: post
title: Closures in Javascript
date: 2015-03-24 16:30
published: true
categories:
- css
---

One of the reasons, I started this __blog__ was so I could jot down things I'm learning. Recently, I've been trying to write more plain Javascript rather than Coffeescript in order to better understand the intricacies and fundamentals of the language. A common question about javascript is always "what is a closure?" and that's what I'm exploring right now.

According to Kyle Simpson in his book series [*You Don't Know JS*](https://github.com/getify/you-dont-know-js), the formal definition of a closure is "when a function can remember and access its lexical scope even when it's invoked outside its lexical scope." This is pretty straightforward because it occurs all over javascript in the use of callbacks and other scenarios. Though closure is just about everywhere in code, it is not always directly observable. For instance a function `a` that contains another function `b` can call `b()` and that is an example of closure but not the most exemplary. A more obvious example is if our function looked like this:

{% highlight javascript %}
function a() {
  var number = 2
  var b = function () {
    return number
  }

  return b
}

var c = a()
c() // 2
{% endhighlight %}

This time, function `a` returned another function whose scope is inside of `a` and though it was called out of scope, it still had its original scope. And that's all closure is, plain and simple.
