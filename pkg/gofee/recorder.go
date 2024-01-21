package gofee

import "github.com/google/uuid"

// Used for recording the state change diagram of the events
type Recorder interface {
	Save(uuid.UUID, State)

	Get(uuid.UUID) State
}

type defaultRecorder struct {
	Recorder

	// Maps the event uuid to the state. State represents
	// the version after the state change callback.
	states map[uuid.UUID]State
}

// Make sure DefaultRecorder implements Recorder interface
var _ Recorder = (*defaultRecorder)(nil)

func NewDefaultRecorder() Recorder {
	return &defaultRecorder{
		states: make(map[uuid.UUID]State),
	}
}

func (r *defaultRecorder) Save(uid uuid.UUID, state State) {
	r.states[uid] = state
}

func (r *defaultRecorder) Get(uid uuid.UUID) State {
	return r.states[uid]
}
