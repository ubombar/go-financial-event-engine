/*
Copyright © 2024 Ufuk Bombar ufukbombar@gmail.com
*/
package main

import (
	"fmt"
	"time"

	"github.com/ubombar/go-financial-event-engine/pkg/gofee"
)

func NewConstantEventCallback(valueDifference int) func(s gofee.State) gofee.State {
	return func(s gofee.State) gofee.State {
		state := s.DeepCopy()
		state.Accounts("default").SetValue(state.Accounts("default").Value() + valueDifference)
		return state
	}
}

func main() {
	// cmd.Execute()
	genesisEvent := gofee.NewGenesisEvent(time.Now())

	event1 := gofee.NewConstantEvent(time.Now(), NewConstantEventCallback(+10_00))
	event2 := gofee.NewConstantEvent(time.Now(), NewConstantEventCallback(-10_00))
	event3 := gofee.NewConstantEvent(time.Now(), NewConstantEventCallback(-10_00))

	uuidGenesis := genesisEvent.UUID()
	uuidEvent1 := event1.UUID()
	uuidEvent2 := event2.UUID()
	uuidEvent3 := event3.UUID()

	genesisEvent.AddFork(event1)
	genesisEvent.AddFork(event2)
	genesisEvent.AddFork(event3)

	// Get the state recorder
	stateRecorder := gofee.NewDefaultRecorder()

	// Create the initial state, start with 100$
	initialState := gofee.NewConcreeteState()
	initialState.Accounts("default").SetValue(100_00)

	genesisEvent.Propagate(initialState, stateRecorder)

	fmt.Printf("event g %v\n", stateRecorder.Get(uuidGenesis))
	fmt.Printf("event 1 %v\n", stateRecorder.Get(uuidEvent1))
	fmt.Printf("event 2 %v\n", stateRecorder.Get(uuidEvent2))
	fmt.Printf("event 3 %v\n", stateRecorder.Get(uuidEvent3))
}
