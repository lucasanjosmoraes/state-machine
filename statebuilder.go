package main

type Builder struct {
	CurrentState CurrentState
}

func New(state CurrentState) StateBuilder {
	return Builder{
		CurrentState: state,
	}
}

func (b Builder) NewStateMachine() State {
	stateMachine := &StateMachine{}

	aState := &AState{
		StateMachine: stateMachine,
	}
	bState := &BState{
		StateMachine: stateMachine,
	}
	cState := &CState{
		StateMachine: stateMachine,
	}
	dState := &DState{}

	stateMachine.A = aState
	stateMachine.B = bState
	stateMachine.C = cState
	stateMachine.D = dState

	stateMachine.loadState(b.CurrentState)

	return stateMachine
}
