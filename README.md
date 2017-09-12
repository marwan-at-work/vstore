# VStore

A redux-like implementation for [Vecty](https://www.github.com/gopherjs/vecty)

# Example:

Reducer: 

```go
// reducer.go

type State struct {
    Name string
    Age  int
}

// Reduce calls all combined reducers
func (s *State) Reduce(action interface{}) {
	switch a := action.(type) {
        case actions.ChangeName:
            s.Name = a.NewName
        default: // unrecognized action.
    }
}

// NewState returns the initial state. 
func NewState() *State {
    return &State{Name: "person", Age: 20}
}
```

On page load: 

```go
// main.go

func main() {
	store := vstore.CreateStore(reducer.NewState())
	vecty.RenderBody(&body.Body{
		Store: store,
	})
}
```

Body Component: 

```go
// body.go

type Body struct {
    vecty.Core
    Name    string `vecty:"prop"`
    Age     int `vecty:"prop"`
}

// NewBody is a mandatory constructor to update the props on dispatch.
func NewBody(s *vstore.Store) vecty.Component {
    state := s.GetState().(*reducer.State)

    return &Body{
        Name: state.Name,
        Age: state.Age,
    }
}

func (b *Body) Render() *vecty.HTML {
    return elem.Body(
        elem.Div(
            vecty.Text(b.Name),
        ),
    )
}
```

## Status 

WIP.