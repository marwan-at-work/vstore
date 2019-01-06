package main

import (
	"github.com/cathalgarvey/fmtless"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"marwan.io/vstore"
	number "marwan.io/vstore/example/store"
)

func main() {
	vecty.RenderBody(&body{
		store: vstore.New(&number.State{}),
	})
}

type body struct {
	vecty.Core
	store vstore.Store
}

func (b *body) Render() vecty.ComponentOrHTML {
	return elem.Body(
		b.store.Connect(NewMainComp()),
	)
}

type mainComp struct {
	vecty.Core
	Dispatch func(action interface{})
	Number   int
}

// NewMainComp is the mainComp constructor
func NewMainComp() vstore.StoreComponent {
	return &mainComp{}
}

func (m *mainComp) Connect(store vstore.Store) {
	m.Dispatch = store.Dispatch

	state := store.State().(*number.State)
	m.Number = state.Number
}

func (m *mainComp) Render() vecty.ComponentOrHTML {
	return elem.Div(
		elem.Div(
			vecty.Text(fmt.Sprint(m.Number)),
		),
		elem.Button(
			vecty.Text("inc"),
			vecty.Markup(
				event.Click(m.onInc),
			),
		),
		elem.Button(
			vecty.Text("dec"),
			vecty.Markup(
				event.Click(m.onDec),
			),
		),
	)
}

func (m *mainComp) onDec(e *vecty.Event) {
	m.Dispatch(number.Decrement{})
}

func (m *mainComp) onInc(e *vecty.Event) {
	m.Dispatch(number.Increment{})
}
