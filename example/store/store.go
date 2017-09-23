package store

// State is the UI's global store
type State struct {
	Number int
}

// Reduce updates the state's number based on
// the appropriate action.
func (s *State) Reduce(action interface{}) {
	switch action.(type) {
	case Increment:
		s.Number++
	case Decrement:
		s.Number--
	}
}
