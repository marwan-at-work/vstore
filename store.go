package vstore

// Store is the main redux-like store. Use it to access the State properties
// through the ComponentConstructor.
type Store struct {
	r    Reducer
	mws  []Middleware
	subs map[*SubFunc]bool
}

// SubFunc is a subscription callback.
type SubFunc func(s *Store)

// Middleware is a simple middleware that gets called before any Reduce calls.
type Middleware func(action interface{})

// CreateStore returns a store that you can make global dispatches to in order
// to update the passed global state.
func CreateStore(r Reducer, mws ...Middleware) *Store {
	s := Store{r: r, mws: mws, subs: map[*SubFunc]bool{}}

	return &s
}

// Dispatch takes an action and passes it to all middlewares and the Reduce function.
// It also calls any subscription functions if exist.
func (s *Store) Dispatch(action interface{}) {
	for _, m := range s.mws {
		m(action)
	}

	s.r.Reduce(action)

	for sub := range s.subs {
		(*sub)(s)
	}
}

// Subscribe subscribes a callback to any store updates.
func (s *Store) Subscribe(callback SubFunc) func() {
	s.subs[&callback] = true

	return func() {
		delete(s.subs, &callback)
	}
}

// GetState returns global state.
func (s *Store) GetState() interface{} {
	return s.r
}
