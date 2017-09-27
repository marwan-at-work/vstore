package main

import (
	"github.com/cathalgarvey/fmtless"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/marwan-at-work/vstore"
	"github.com/marwan-at-work/vstore/example/store"
)

func main() {
	vecty.RenderBody(&body{
		store: vstore.New(&store.State{}),
	})
}

type body struct {
	vecty.Core
	store vstore.Store
}

func (b *body) Render() *vecty.HTML {
	return elem.Body(
		b.store.Connect(NewMainComp()),
	)
}

type mainComp struct {
	vecty.Core
	store vstore.Store
}

// NewMainComp is the mainComp constructor
func NewMainComp() vstore.StoreComponent {
	return &mainComp{}
}

func (comp *mainComp) Connect(store vstore.Store) {
	comp.store = store
}

func (m *mainComp) Render() *vecty.HTML {
	state := m.store.State().(*store.State)
	return elem.Div(
		elem.Div(
			vecty.Text(fmt.Sprint(state.Number)),
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
	m.store.Dispatch(store.Decrement{})
}

func (m *mainComp) onInc(e *vecty.Event) {
	m.store.Dispatch(store.Increment{})
}
