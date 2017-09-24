package vstore

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

// ComponentConstructor is a function that connects a Component's props
// to the Store's State. This will get called everytime a Store updates,
// therefore you can just use it instantiate your Component struct and
// give it the updated property from Store.GetState().
type ComponentConstructor func(s *Store) vecty.Component

// NewComponent returns a vecty Component that will rerender everytime
// the store is updated.
func (s *Store) NewComponent(cc ComponentConstructor) vecty.Component {
	c := &comp{cc: cc, s: s}

	return c
}

type comp struct {
	vecty.Core
	cc    ComponentConstructor
	s     *Store
	unsub func()
	test  interface{}
}

func (c *comp) Render() *vecty.HTML {
	return elem.Div(
		c.cc(c.s),
	)
}

func (c *comp) Mount() {
	c.unsub = c.s.Subscribe(c.callback)
}

func (c *comp) Unmount() {
	c.unsub()
}

func (c *comp) callback(s *Store) {
	c.s = s

	vecty.Rerender(c)
}
