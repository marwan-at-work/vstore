package vstore

import (
	"github.com/gopherjs/vecty"
)

// StoreComponent must be satisfied by components to connect to a store.
// Also, vecty.Component is embedded to preserve component behavior.
type StoreComponent interface {
	vecty.Component
	Connect(store Store)
}

// storeComponent embeds a component to add store pubsub logic.
// Also, StoreComponent is embedded to preserve component behavior.
type storeComponent struct {
	StoreComponent
	store *store
}

// Connect links the store to a given component and returns the component.
func (s *store) Connect(comp StoreComponent) StoreComponent {
	comp.Connect(s)
	return &storeComponent{StoreComponent: comp, store: s}
}

// Mount subscribes the component to the store and optionally calls mount.
func (c *storeComponent) Mount() {
	c.store.sub(c)
	if mounter, ok := c.StoreComponent.(vecty.Mounter); ok {
		mounter.Mount()
	}
}

// Unmount unsubscribes the component to the store and optionally calls unmount.
func (c *storeComponent) Unmount() {
	c.store.unsub(c)
	if unmounter, ok := c.StoreComponent.(vecty.Unmounter); ok {
		unmounter.Unmount()
	}
}
