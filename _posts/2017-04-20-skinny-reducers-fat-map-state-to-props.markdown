---
layout: post
title: Skinny Reducers + Fat MapStateToProps
date: 2017-04-20 18:36
categories:
- javascript
- redux.js
- react.js
---

Redux is a really neat library for managing application state to make it deterministic.
In a Redux application, there is a store that contains the state tree and that state changes when actions are dispatched.
The state transitions are handled by functions called reducers, which take current application state and the action to apply,
then return the new application state.
This is simple enough as a concept and when it is applied in a React app, it can become a little more complex when the state now influences a UI.
I have seen React + Redux applications become brittle and provide a poor UX because of mistakes in Redux architecture.
This is not intended to be a lesson on how to use react-redux and how 'connected' components work so I assume the readers are aware of the technologies here.

### Remember MVC?
Some folks might remember when MVC was the coolest and most beautiful pattern for web applications both server and client side.
We had the **M**odel, **V**iew, and **C**ontrollers that each had a responsibility.
The model was the data representation, the controllers mapped requests to a model and provided a view to represent the data.
The mantra for the best practice in this world was "Skinny controllers and fat models,"
meaning have the model contain its business logic and let the controller simply match the data to a view with no concern of business logic.

MVC works pretty well in traditional web applications and unfortunately, it falls short in a Redux world because
a Redux and React application is essentially an __MV__ pattern and React, which (at least originally) claimed to be just a view layer,
insists that the view is a function of data i.e. `V = f(M)`.
I think even without controllers, the "skinny \_\_\_  and fat  \_\_\_" advice can still be applied to applications using React and Redux.

### How?
Without controllers, it is very tempting and easy to put a lot of logic of how the View should be rendered inside of the application state, which in this case is the Model.
Doing this yields code like this:

{% highlight javascript %}
// in the reducer
export default function(state = defaultState, action) {
  if (action.type === FOO) {
    const someData = doABunchOfLogicWithCurrentStateAndNewAction(state, action.payload)
    const moreData = doSomeMoreStuffWithSomeData(someData)
    const moreData.thing = createThing(state, moreData)
    return { ...state, moreData }
  } else if (/*...*/) {/*...*/}
}

// in the view
export default function(props) {
  return <View {...props.moreData}/>
}
{% endhighlight %}

The likelihood of a lot of the logic happening here being about *how* the UI should look is high.
Although it is completely okay to have complex application state changes which require some helper functions,
the danger is in the situations where the code is doing more than that.
Another problem with a situation like this is that the action payload may contain a ton of information as well that may not even end up in state.

This scenario can be described as "Skinny Views and fat reducers",
which to me doesn't seem great because now the views aren't much of a function of state because they are simply templates.

### What's the alternative?
To reverse this situation, we can take advantage of the react-redux api and leverage the `mapStateToProps` function.
This function is where business logic should lie because as the name suggests, it is quite literally a mapper of application state to view relevant 'stuff'.
Using 'mapStateToProps', we can keep our reducers simple and our application state tree clear of UI litter and more about the application.

{% highlight javascript %}
// in the reducer
export default function(state = defaultState, action) {
  if (action.type === FOO)
    return {
      ...state,
      ...action.payload
    }
  } else if (/*...*/) {/*...*/}
}

// in the view
import {connect} from 'react-redux'
const View = function(props) {
  return <View {...props}/>
}

const mapStateToProps = state => {
  return {
    prop1: getViewRelevantProp1(state),
    prop2: moreComputationOfWhatViewShouldHave(state)
  }
}
export default connect(mapStateToProps)(View)
{% endhighlight %}

With this pattern, the complexity of what the UI may look like can be kept out of application state.
This allows actions that change state to be more atomic and straightforward.
Discerning what is UI logic and what is *actually* application state is not always easy and it's more of a grey area in a lot of cases.
