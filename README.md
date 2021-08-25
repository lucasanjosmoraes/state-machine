# State Machine

Example of how to build a State Machine in Go.

`State` is where we define all the methods that we will use from our State Machine. `AState`, `BState`, `CState` and `DState`
are implementations of `State`. `StateMachine` also implements `State` because we're going to call it instead of calling
the states directly.

We load all states into the state machine using the `Builder` struct. It instantiates all states with a state machine
reference to allow the states to move the current state to the next one.

The current state is an attribute of `Builder`, but it can be retrieved from an external source and then passed to the `loadState`
method.