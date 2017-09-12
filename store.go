package vstore

// Store is the main redux-like store. Use it to access the State properties
// through the ComponentConstructor.
type Store struct {
	r    Reducer
	mws  []Middleware
	subs []*SubFunc
}

// SubFunc is a subscription callback.
type SubFunc func(s *Store)

// Middleware is a simple middleware that gets called before any Reduce calls.
type Middleware func(action interface{})

// CreateStore returns a store that you can make global dispatches to in order
// to update the passed global state.
func CreateStore(r Reducer, mws ...Middleware) *Store {
	s := Store{r: r, mws: mws}

	return &s
}

// Dispatch takes an action and passes it to all middlewares and the Reduce function.
// It also calls any subscription functions if exist.
func (s *Store) Dispatch(action interface{}) {
	for _, m := range s.mws {
		m(action)
	}

	s.r.Reduce(action)

	for _, sub := range s.subs {
		(*sub)(s)
	}
}

// Subscribe subscribes a callback to any store updates.
func (s *Store) Subscribe(callback SubFunc) func() {
	s.subs = append(s.subs, &callback)

	return func() {
		i := -1
		for idx, sub := range s.subs {
			if sub == &callback {
				i = idx
				break
			}
		}

		if i != -1 {
			s.subs = append(s.subs[0:i], s.subs[i:]...)
		}
	}
}

// GetState returns global state.
func (s *Store) GetState() interface{} {
	return s.r
}
