---
layout: post
title: You Don't Need lodash
date: 2017-03-30 12:00
categories:
- javascript
- lodash
- dependencies
---

The javascript ecosystem today, is full of *stuff*. By stuff, I mean a ton of libraries, tools, and frameworks. A lot of these provide value
and I think because it is so easy to create them and publish them, we often become too reliant on the libraries to grab the low hanging fruit for us.
A genre of libraries that has become incredibly prevalent in javascript development is the general purpose utility library, e.g. underscore/lodash, and zepto, jQuery.
jQuery was (and sort of still is) amazing. It soothed a lot of pain we felt related to DOM manipulation and it had a pretty impressive api that has inspired many
javascript libraries since then.
Underscore did some amazing things for us in terms of allowing us to write more functional code and handle collections with some fancy tricks.
Then lodash came along to sweeten the deal by being more performant and providing some extra functionality as well as (arguably) better usability or ergonomics.
Today, in the year of 2017, I propose we start the lodash diet because we don't need to rely on these libraries as much and I'll tell you why.

## Ecmascript + Browsers Grow
The Ecmascript specifications and standards that drive the javascript language features and adoption in browsers are always changing and maturing.
This means the usual go-to functions like `filter`, `forEach`, `map`, and `find` are being adopted as core features of the javascript language and browsers.
A lot of folks would be surprised to know that things like sorting an array have been around for a very long time (since Chrome/Firefox 1.0 and IE 5.5).
Plenty of us have to support some version of IE and I get it. I just want to say there are a ton of things we reach for lodash to accomplish that even IE9 has already implemented.
And for the record, `Array.prototype.find`, is one of the few things all Internet Explore browsers don't have but thankfully, there are lightweight polyfills that can be used for IE.
I use IE9 as a base line because I figure most of the user facing traffic we build for today in 2017 is likely not coming from a browser running anything less.

## Why Does this matter?
The lodash diet allows us to be less reliant on utility libraries and gives us the opportunity to ship less code.
The browser has limitations and the fact that we build web applications means we have to consider latency as a factor when we force our users
to download megabytes of javascript and assets on a page.
Regardless of whether we're building single page applications or traditional server driven applications,
downloading and parsing assets from the server still happens and we should be doing all that we can to reduce how much *needs* to be transmitted and executed.

Beyond just our obligation to our end users, we should also consider the pain for ourselves. Dependency pruning is an activity we as developers should be actively engaged in.
NPM has made it stupidly easy to bring in dependencies without a thought and that can lead to projects that heavily rely on code external to the team or product,
which in turn means more complexity for the whole application.
Though open source libraries are great, let's not forget that they are usually created by people who are bored or solving
an *immediate* personal need and then these projects are maintained by developers in their spare time beyond that.
*NOTE:* I'm not saying lodash was created by any bored people and it will eventually be abandoned. But let's not forget about left-pad-gate.

## Immediate Steps To Start The diet
The lodash project is pretty modular and all the functions are exported as individual packages on npm.
If you really need help from a lodash function, I suggest reaching for exactly what it is that you want and not the entire project.
Importing 10-16 individual lodash functions might end up being as costly as the entire project because most of the functions use some
common helper code (a.k.a feature sets) and each package will have its own copy of that shared code, meaning there will be duplicated code.
At that point, the diet has turned into over-consumption of little snacks.
Thankfully, there's a strategy for this too.

Since the lodash project is aware of these problems, they have webpack and babel plugins to support reducing how much lodash is consumed.
Using the [babel-plugin-lodash](https://www.npmjs.com/package/babel-plugin-lodash),
we can include all of lodash as an NPM dependency and then at build time, cherry-pick exactly which functions we want.
To take it a step further, those common chunks of code I mentioned earlier can sometimes be replaced with noops or smaller alternatives if the [lodash-webpack-plugin](https://www.npmjs.com/package/lodash-webpack-plugin) is also used.
