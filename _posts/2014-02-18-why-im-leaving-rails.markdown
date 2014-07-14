---
layout: post
title: "Why I'm leaving rails"
date: 2014-02-18 01:15
comments: true
categories:
- rails
---

I gave up Ruby on Rails pretty quickly after I put my first app, intrnhuntr.com, into production. But recently, I've been reading about other developer's complaints about the Rails framework and the Rails way. One of the arguments being brought up is that the framework is "too big" and [not for beginners](http://rob.yurkowski.net/post/17610425880/rails-is-definitely-not-for-beginners). Here is pretty good blog post that identifies how [Ruby on Rails is too complicated](http://www.woohooitsbacon.com/rails-needs-change/).


Those posts aren't the main reason I'm leaving Rails but they are relevant and inspired this post. The main reason I'm leaving Rails is because it doesn't make me feel like a productive programmer. I've been using Node.js and Coffeescript lately and everything is smooth. There is no scaffolding and somehow having a full-blown app generated that I need to muck around and change. Sure, you can remove what you don't need/want but why muck around when I can build what I need from scratch and keep it lightweight and efficient? My latest project has been completely all server side. It started out as a desktop app, built with [node-webkit](http://github.com/rogerwang/node-webkit) I could use to study French vocab, and then I decided to host it on Github pages as a SPA using localstorage for a database. All I needed was a Backbone.js, and a localstorage adapter. Now I've realized that a web app really just needs to be the html and the javascript. Servers should feed the app but it shouldn't __be__ the app. Rails is the app and it's monolithic. Node supports the app, performs 2x as well as Rails, and probably has fewer (coffeescript) lines than Rails.

Despite my departure from the Rails community, I am thankful for what I've learned from web development in Rails. My MVC knowledge came from Rails and it's helped me pick up on front-end frameworks like Backbone.js and Ember.js. Both of which are designed by large contributors to the Rails community. Rails is complex, it probably will need to slim down in order to compete in the future but it has good points and principles from which we can learn from as developers.
