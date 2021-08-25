package main

import "log"

type DState struct{}

func (s *DState) ToB() {
	log.Printf("doing nothing (state D)")
}

func (s *DState) ToC() {}

func (s *DState) ToD() {}
