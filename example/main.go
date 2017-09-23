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
		store: vstore.CreateStore(&store.State{}),
	})
}

type body struct {
	vecty.Core
	store *vstore.Store
}

func (b *body) Render() *vecty.HTML {
	return elem.Body(
		b.store.NewComponent(NewMainComp),
	)
}

type mainComp struct {
	vecty.Core
	Number   int `vecty:"prop"`
	Dispatch func(action interface{})
}

// NewMainComp is the mainComp constructor
func NewMainComp(s *vstore.Store) vecty.Component {
	state := s.GetState().(*store.State)

	return &mainComp{
		Number:   state.Number,
		Dispatch: s.Dispatch,
	}
}

func (m *mainComp) Render() *vecty.HTML {
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
	m.Dispatch(store.Decrement{})
}

func (m *mainComp) onInc(e *vecty.Event) {
	m.Dispatch(store.Increment{})
}
