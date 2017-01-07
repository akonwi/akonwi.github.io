---
layout: post
title: Introducing Jest Into A React Project
date: 2017-01-10 12:40
categories:
- javascript
- react
---

Recently, I decided to catch up on what's going on in the front-end javascript world because it's been a solid 14 months since I worked on a modern javascript project that wasn't the Atom editor. I have a chrome extension I've been using as a pet project for me to experiment with React and test my opinions. I'm writing this as I've just started porting my app from React v0.14.7 to v15.4.2. My build tool for this project is Grunt with browserify and babelify to use ES6 and jsx. There were no tests and I want to get my feet wet testing React components so after some research it seems Jest is a great way to go.

Thankfully, getting up and running with jest in my existing project wasn't too much of a hassle. I followed Jest's [getting started guide](http://facebook.github.io/jest/docs/tutorial-react.html) and did `yarn add -D jest babel-jest babel-preset-es2015 babel-preset-react enzyme`. I substituted`react-test-renderer` for [enzyme](https://github.com/airbnb/enzyme) which handles rendering in tests and has useful matchers for asserting on components. Jest requires having a `.babelrc` file, which I didn't have before. The simplest configs you need to get up and running looks like this:

``` javascript
//.babelrc
{
  "presets": ["es2015", "react"]
}
```

With that, you just need to run `jest` from your project's root directory and all the tests in the `__tests__` directory will be run. I had to rename my test directory from 'spec' to `__tests__` when the jest output said no tests were found after I first ran it.

The first component I wanted to test was very simple:

``` javascript
// Thing.jsx
export default function({url, title}) {
  const openLink = (e) => {...}

  return (
    <li>
      <a href='#' onClick={openLink}>{ title }</a>
    </li>
  )
}
```

This is the entire file. Notice how there're no `import` statements at the top. At the very least I should be importing React right? No because up until now I've just included a distribution of React and React-dom in the project and included them in the head of my main html file. Before you throw your hands up in disgust, remember this is an extension for Google Chrome, so everything is bundled up and run inside of your browser so it's okay to do things the 'old way' because the assets are already there and there are no extra requests made. Anyway, to make this component testable with Jest and independent of a global environment where React is declared I needed to `yarn add react react-dom` to download react into my node environment and simply add `import React from 'react'` to the top of the file. Now we can write a test like this:

``` javascript
// __tests__/Thing-test.js
import React from 'react'
import {shallow} from 'enzyme'
import Thing from 'path/to/Thing'

describe("Thing", () => {
  const url = 'foo://bar.com'
  const title = 'Foobar'

  it("creates a list item with a link", () => {
    const thing = shallow(<Thing title={title} url={url}/>)
    const link = thing.find('a')

    expect(thing.is('li')).toBe(true)
    expect(link.text()).toEqual(title)
  })
})
```

Running `jest` will now execute those tests. Awesome! My component is tested now. It's time to rebuild my project and try it out in chrome so I run `grunt browserify` and I get some weird error that looks like this:

``` shell
>> ReferenceError: [BABEL] /Users/angoh/stuff/project/src/another-file.jsx: Unknown option: /Users/angoh/stuff/project/.babelrc.presets while parsing file: /Users/angoh/stuff/project/src/another-file.jsx
Warning: Error running grunt-browserify. Use --force to continue.
```

It seems browserify doesn't understand my `.babelrc` options because I've just created one and I didn't use one before. Well I just need to `yarn upgrade browserify babelify` to get the latest because it has been a while since I created this project and the babel ecosystem has surely changed a lot. After that, I could rebuild the extension and try it out in Chrome.
