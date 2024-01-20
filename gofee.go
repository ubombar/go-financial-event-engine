/*
Copyright © 2024 Ufuk Bombar ufukbombar@gmail.com
*/
package main

import (
	"fmt"
	"time"

	"github.com/ubombar/go-financial-event-engine/pkg/gofee"
)

func NewGenesisEventCallback(initialValue int) func(s gofee.State) gofee.State {
	return func(s gofee.State) gofee.State {
		state := gofee.NewConcreeteState()
		state.Accounts("default").SetValue(initialValue)
		return state
	}
}

func NewConstantEventCallback(valueDifference int) func(s gofee.State) gofee.State {
	return func(s gofee.State) gofee.State {
		state := gofee.NewConcreeteState()
		state.Accounts("default").SetValue(state.Accounts("default").Value() + valueDifference)
		return state
	}
}

func main() {
	// cmd.Execute()

	rootEvent := gofee.NewConstantEvent(time.Now(), NewGenesisEventCallback(100_00))

	event1 := gofee.NewConstantEvent(time.Now(), NewConstantEventCallback(+100_00))
	event2 := gofee.NewConstantEvent(time.Now(), NewConstantEventCallback(-22_00))
	event3 := gofee.NewConstantEvent(time.Now(), NewConstantEventCallback(+12_00))

	rootEvent.AddFork(event1)
	rootEvent.AddFork(event2)
	rootEvent.AddFork(event3)

	fmt.Printf("rootEvent: %v\n", rootEvent)
}
