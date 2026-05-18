---
layout: post
title: Put Some Backbone Into It
date: 2013-10-14 00:02
categories:
- node.js
- backbone.js
- web
---
So my current main project is an app I'm building with Node.js. I chose Node because it's a fullstack system that uses the same language(javascript) on both server and client. Node is built to be fast and lightweight. So anyway, I want this web app to be fast and lightweight as well and that's why I'm making it a SPA(Single Page App) with the help of Backbone.

The Backbone documentation is very simple. I learned next to nothing the first 5 times I read through it. There was a bounty of tutorials online though but only a handful of helpful ones. They all demonstrated the 'Todo List' app which was good for demonstrating the basics but nothing gave any info about using Backbone with a RESTful backend. Somehow though, I found an online book called [Developing Backbone.js Applications](http://addyosmani.github.io/backbone-fundamentals/) by Addy Osmani. Parts of the book are a bit outdated in terms of how Backbone works now but the key concepts are still the same. I'm going to outline key concepts I've learned(some of these may pertain only to me but this blog is mainly for me anyway).

## Models are the 'backbone' of your app #
  Models in Backbone can be anything but they are only representations of actual objects. For the most part, the whole app is built around models. So in Backbone, a model is an object with some values in it. These values can be data for html elements or data from a server-side database. When a change is made to your model, that change stays on the client. In order to persist these changes in a db, the Model#sync method needs to be called and Backbone actually does it for you once you call save() on a model. Now you're probably wondering "If the models are so important, what's their role in the view of the app?" Keep reading.

## Views cover and visually represent models #
  A view is essentially an HTML element with its own content. A view can be a div with tons of other views in it or just an object that fills and renders other views. Each view can have it's own model that it represents and/or a whole collection of models(see below for collections). Views are usually made up of templates that are very similar to ERB templates in Rails and when you render a view, you compile it and pass in the variables that you want to use in the template and construct your template to portray that data as you please. But not every view has to have data from a model, so they can just be static templates. Like I said, they are just HTML elements you put into the DOM whenever and wherever you want.

  Another treat we get from views is that each one can handle its own events. Events being anything from a click to a contextmenu event. I don't even know what a contextmenu event is but I can't wait to find out. So each view can have events attached to it and these events are similary to jquery/javascript events in that you attach a callback that takes an 'event' parameter.

## Collections are literally collections #
  Very simple. It's very common that a class that represents a model will have methods like `Item.find(id)`. Well in Backbone, the collection kind of houses those class methods and is a bucket of all the models of that class. The collection is an array I'm pretty sure that just holds all the models as objects and has methods like `find`, `get`, `add`, `remove` and tons more. If you are familiar with Rails, SQL, or any database, then you can think of a collection as a SQL table but the trick is that you have to manually fill in the collection before you use it. This means you have to load in the models from database into the collection or manually add them individually. The only automatic loading of data Backbone does is when you make changes on the client. Nothing persists to the server unless you explicitly want it to.

That's a quick jist of Backbone that I wish I was aware of before using it. But maybe I'm just dumb and everyone else can understand that from reading their documentation. Anyway, I've hooked up my Backbone app to the Node.js backend with mongoDB and everything is working pretty smoothly. There are still plenty of things I still need to refine or probably re-do but I think I'll do another post to document some of my techniques and design of my app but I'm open to better ways of doing things.
