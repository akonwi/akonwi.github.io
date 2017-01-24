---
layout: post
title: Inline Styles In React.js
date: 2017-01-22 15:20
categories:
- javascript
- react
- css
---

As part of my refactoring and getting my pet React project up to date with the latest cool stuff, I took a foray into inline styles for my components.
As I was moving my styles from [SASS](http://sass-lang.com/) into plain javascript, there were a few things I found I didn't like.

The first was that with inline styles on the components as javascript objects, it isn't possible to combine styles by passing multiple style objects. This isn't really a big deal and it makes sense because that's how the DOM works. The style attribute of DOM nodes takes a string containing a css property declarations and React flattens an object into a string appropriate for the style attribute. Ideally, what I would like to do is provide an array of style objects like so:

{% highlight  javascript %}
const styles = {
  button: {
    color: 'black',
    fontSize: '12px'
  },
  small: {
    fontSize: '8px'
  },
  red: {
    color: 'red'
  }
}

<div style={[styles.button, styles.small, styles.red]}>...</div>
{% endhighlight %}

Since that's not possible, I just created a function to combine multiple styles.

{% highlight  javascript %}
function button(styles) {
  return styles.reduce((final, style) => Object.assign(final, styles[style]), styles.button)
}
// ...then
<div style={button(['small', 'red'])}>...</div>
{% endhighlight %}

That was sufficient for what I wanted but later when I came across pseudo-selectors like `:hover` I realized I needed the help of a library to do it inline. I had previously heard of [Radium](https://github.com/FormidableLabs/radium) and [Aphrodite](https://github.com/Khan/aphrodite) so I compared the two. Radium seemed great and most like what I wanted but I was turned off by the API. Radium decorates React components so the `@Radium` annotation is necessary above the ES6 class or components created with `React.createClass` need to be exported as Radium components like `Radium(MyComponent)` because Radium is actually a higher-order component that wraps other components. For me that makes it a little to involved in my code. I want to declare my styles and reference them without other components involved. Thankfully, Aphrodite provides the ability to create a simple stylesheet and then at run-time, it generates css styles and classnames for them so application of styles moves back to the `className` attribute of components rather than `style`. Aphrodite's tag line ("It's inline styles, but they work!") on Github is misleading because they aren't inline styles but rather a stylesheet created through js. I wonder if there are other libraries that might provide the ability to style pseudo-selectors without generating css.

With Aphrodite, I realized that tests that were written to find nodes based on their css classes or assert that they had particular classes based on their state, were no longer valid because the class names were dynamically generated at the time of rendering the components. The solution to this is to put the style declaration in a separate module which can be referenced from both the component and the test. This becomes kind of ugly though and seems to defeat the purpose of components being self contained.

Another thing I don't really like about doing this is that if you have styles declared that are for the children of a selector i.e. this in SASS:

{% highlight css %}
.title-bar
  *
    vertical-align: middle
    display: inline-block
{% endhighlight %}

now has to be applied individually to those child elements. If the children are all the same component and can be dynamically generated, then it's not much of a fuss. Unfortunately, this wasn't the case for me because all the child elements were different. I think the solution to this is probably to refactor the styling to be based on flex-box because I'm probably doing something wrong.

Overall, I'm pretty happy with this refactoring because I removed the need for SASS in the project and that's one less step in my build process and I can keep a small css file with some base styling. I need to do more research though to find a library that doesn't create a stylesheet but also accommodates for pseudo-selectors.
