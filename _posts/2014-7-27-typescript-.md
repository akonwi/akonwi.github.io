---
layout: post
title: Typescript
date: 2014-07-27 20:58
categories:
- javascript
---

This summer, I've had to write a lot of Typescript for my front-end development at my internship.
Picking it up was not too difficult because the main difference between it and vanilla javascript,
is that you *can* declare a type for your variables and functions. The type declarations are completely optional.
But if you don't do the typings then what's the point right? Typescript was created by a Microsoft developer for the
purpose of making it easier for C# developers to read and write javascript by allowing them to see variable types.

Here are some pros and cons I've noticed with Typescript:

### Pros
- Typing (obviously)
- Intuitive inheritance and classes (for the most part)
- Comes with moduling support for AMD and Commonjs
- Somewhat follows ECMA6 syntax

### Cons
- Typing
- Typing can create verbose/cluttered code
- `.d.ts` files, which I'll discuss later

So why is 'typing' both a pro and a con? Well, obviously it's good because it can make the code more legible
and sometimes help us be more organized. The negative to that is that this is java__SCRIPT__.
It was designed to be lightweight and duck typed because it is kind of a scripting language especially now that
we have node.js. The typing in javascript is inherent to me in my opinion. But again, that's my opinion.

My other large issue with Typescript is that in order to use a third party library not in core javascript, like
JQuery for example, you need a definition file i.e. `jquery.d.ts` in addition to the external library's file. This definition file is like a header file
in C. It declares the interface of the library that you can use in Typescript. Having a definition file makes sense
because this is Typescript but it adds another layer of difficulty in getting to work quickly I think.

Personally, I won't keep writing Typescript after this. I've written more Coffeescript than anything else though
probably and I love it. But after writing so much Typescript, which is almost vanilla javascript, I realized there
 is not a huge amount of benefit in coffeescript aside from syntactic sugar (which I really like).
 But the arguments about coffeescript being a barrier to development for plain js developers and
 that it isn't actually javascript are crap. If you remove the parens and semi-colons from javascript,
 you still have objects, functions, and prototyping. When ECMA6 comes, I won't have an issue picking
 it up because it's no different from coffeescript or vanilla javascript.
