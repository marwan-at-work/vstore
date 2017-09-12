package vstore

// Reducer takes an action and should
// modify the state based on that action.
type Reducer interface {
	Reduce(action interface{})
}

// CombineReducers combines reducer functions into a single one.
// Use when creating a global redux state
func CombineReducers(reducers ...Reducer) Reducer {
	return &reduceAggergator{reducers}
}

// reduceAggregator is a minimal type
// that can call all the reducers that are given to it.
type reduceAggergator struct {
	reducers []Reducer
}

// Reduce calls all reduce funcs.
func (r *reduceAggergator) Reduce(action interface{}) {
	for _, rd := range r.reducers {
		rd.Reduce(action)
	}
}
