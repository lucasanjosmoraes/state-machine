package main

import "log"

type CState struct {
	StateMachine *StateMachine
}

func (s *CState) ToB() {
	// retrieve previous data from external sources if needed or retrieve data from all previous states in the next steps
	log.Printf("doing nothing")
}

func (s *CState) ToC() {
	// retrieve previous data from external sources if needed
	log.Printf("doing nothing")
}

func (s *CState) ToD() {
	log.Printf("doing something and going to state D")

	s.StateMachine.setState(s.StateMachine.D)
}
