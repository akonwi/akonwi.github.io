---
layout: post
title: CSS Precedence
date: 2015-03-24 16:30
published: true
categories:
- css
---

Today, I had a phone interview for a web developer role and one of the questions I was asked was "How does CSS precedence work?" Unfortunately,
I didn't know so decided to look into it afterwards.

So it turns out that precedence is determined by specificity of the selector. Precedence can be written in the form of `(a,b,c,d)`.

I'll explain what each of those letters represents.

- a: Inline styles
- b: Id selector
- c: Class, Pseudo class, Attribute Selector
- d: Element, Pseudo Element selector

In terms of specificity, we can sort of quantify this in the `(a,b,c,d)` notation where each letter is a numeric value representing the count
of that corresponding component in the selector. For instance, a rule `div.centered { height: 100px; }` has a specificity of (0,0,1,1) because there are one of each Element and class selectors. Now `div.centered` takes precedence over something like `input { color: red; }`, which has a specificity of (0,0,0,1). Both of these rules would be trumped by something more precise like `div.tilted a#main-link { color: black; }` because its specificity is (0,1,1,2).

For clarity:
{% highlight css %}
input { color: red; } //(0,0,0,1)
div.centered { height: 100px; } //(0,0,1,1)
div.tilted a#main-link { color: black } //(0,1,1,1)
{% endhighlight %}

In terms of precedence: `div.tilted a#main-link` > `div.centered` > `input`, and in conclusion, precedence depends on specificity of selectors with inline styles being the ultimate priority and id's following after.
