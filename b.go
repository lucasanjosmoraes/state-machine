package main

import "log"

type BState struct {
	StateMachine *StateMachine
}

func (s *BState) ToB() {
	// retrieve previous data from external sources if needed
	log.Printf("doing nothing")
}

func (s *BState) ToC() {
	log.Printf("doing something and going to state C")

	s.StateMachine.setState(s.StateMachine.C)
}

func (s *BState) ToD() {
	panic("invalid method call")
}
