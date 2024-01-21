package gofee

import (
	"time"

	"github.com/google/uuid"
)

type EventType int

const (
	// Constant event
	EVENT_TYPE_CONST EventType = iota
	EVENT_TYPE_VARIABLE
)

type Event interface {
	// Gets the event uuid
	UUID() uuid.UUID

	// This is used for getting the events in front.
	// It is an array because the event tree might be forked.
	Next() []Event

	// Gets the event type
	Type() EventType

	// Adds the given event to the event tree
	AddFork(Event) bool

	// Accumulates the changed states into map object
	Propagate(State, Recorder) bool
}

// CONSTANT EVENT
// ConstantEvent is a fixed type of event. The volume and the date
// will be constant. When added to the event tree it will only
// append and will not fork the tree.
type constantEvent struct {
	Event

	uuid                uuid.UUID
	nextPointer         []Event
	datetime            time.Time
	stateChangeCallback func(State) State
}

func NewConstantEvent(datetime time.Time, stateChangeCallback func(State) State) Event {
	return &constantEvent{
		uuid:                uuid.New(),
		nextPointer:         make([]Event, 0),
		datetime:            datetime,
		stateChangeCallback: stateChangeCallback,
	}
}

func NewGenesisEvent(datetime time.Time) Event {
	return NewConstantEvent(datetime, func(s State) State {
		return s.DeepCopy()
	})
}

func (e *constantEvent) UUID() uuid.UUID {
	return e.uuid
}

func (e *constantEvent) Next() []Event {
	return e.nextPointer
}

func (e *constantEvent) Type() EventType {
	return EVENT_TYPE_CONST
}

// Adds the given event to the event tree. If there is no next and date is
// smaller cannot add the event.
func (e *constantEvent) AddFork(event Event) bool {
	switch event.Type() {
	case EVENT_TYPE_CONST:
		// If the given event is before our current event, we cannot add.
		// Note that the events are assumed to not hapen at the same time.
		if e.datetime.Compare(event.(*constantEvent).datetime) != -1 {
			return false
		}

		// If the given event has no next, add it to the next list
		// Else the there is some then add the event to each of them since const
		// event's propagate to all other forks.
		if len(e.nextPointer) == 0 {
			e.nextPointer = append(e.nextPointer, event)
			return true
		} else {
			// We need to have a rollback mechanism if not all the events succesfully
			// adds the event to the event tree. But for now I will assume they do
			for _, nextEvent := range e.nextPointer {
				nextEvent.AddFork(event)
			}
			return true
		}
	}
	return false
}

func (e *constantEvent) Propagate(given State, recorder Recorder) bool {
	// Calculate the state chance
	current := e.stateChangeCallback(given)

	// Save the state with even't uuid
	recorder.Save(e.uuid, current)

	// Same thing applies here, the function return might be false and we might
	// need to do a rollback. Haven't implemented yet.
	for _, nextEvent := range e.nextPointer {
		nextEvent.Propagate(current, recorder)
	}

	return true
}
