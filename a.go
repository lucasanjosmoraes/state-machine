package main

import "log"

type AState struct {
	StateMachine *StateMachine
}

func (s *AState) ToB() {
	log.Printf("doing something and going to state B")

	s.StateMachine.setState(s.StateMachine.B)
}

func (s *AState) ToC() {
	panic("invalid method call")
}

func (s *AState) ToD() {
	panic("invalid method call")
}
