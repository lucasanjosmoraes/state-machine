package main

type StateMachine struct {
	A State
	B State
	C State
	D State

	CurrentState State
}

func (sm *StateMachine) ToB() {
	sm.CurrentState.ToB()
}

func (sm *StateMachine) ToC() {
	sm.CurrentState.ToC()
}

func (sm *StateMachine) ToD() {
	sm.CurrentState.ToD()
}

func (sm *StateMachine) loadState(s CurrentState) {
	switch s {
	case A:
		sm.setState(sm.A)
	case B:
		sm.setState(sm.B)
	case C:
		sm.setState(sm.C)
	case D:
		sm.setState(sm.D)
	}
}

func (sm *StateMachine) setState(s State) {
	sm.CurrentState = s
}
