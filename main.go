package main

import "log"

func main() {
	log.Printf("------ Starting from state A ------")
	builderFromA := New(A)
	smFromA := builderFromA.NewStateMachine()
	smFromA.ToB()
	smFromA.ToC()
	smFromA.ToD()

	log.Printf("------ Starting from state B ------")
	builderFromB := New(B)
	smFromB := builderFromB.NewStateMachine()
	smFromB.ToB()
	smFromB.ToC()
	smFromB.ToD()

	log.Printf("------ Starting from state C ------")
	builderFromC := New(C)
	smFromC := builderFromC.NewStateMachine()
	smFromC.ToB()
	smFromC.ToC()
	smFromC.ToD()

	log.Printf("------ Starting from state D ------")
	builderFromD := New(D)
	smFromD := builderFromD.NewStateMachine()
	smFromD.ToB()
	smFromD.ToC()
	smFromD.ToD()
}
