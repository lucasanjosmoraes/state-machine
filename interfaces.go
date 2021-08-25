package main

type CurrentState int

const (
	A CurrentState = iota + 1
	B
	C
	D
)

type State interface {
	ToB()
	ToC()
	ToD()
}

type StateBuilder interface {
	NewStateMachine() State
}
