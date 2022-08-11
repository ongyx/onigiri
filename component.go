package onigiri

import "reflect"

// ComponentKind represents a specific kind of component in a world.
type ComponentKind interface {
	kind() uint8
}

// Component is a unique identifer for a component type.
type Component[T any] struct {
	world *World
	table *Table[T]
	id    uint8
}

func newComponent[T any](w *World, e entry) *Component[T] {
	return &Component[T]{
		world: w,
		table: e.table.(*Table[T]),
		id:    e.id,
	}
}

// Get gets the component associated with the entity.
func (c *Component[T]) Get(e Entity) *T {
	return c.table.Get(e)
}

// Set adds the component to the entity.
func (c *Component[T]) Set(e Entity, com T) {
	c.table.Set(e, com)
	c.world.entities[e].Set(c.id)
}

// Delete removes the component from the entity.
func (c *Component[T]) Delete(e Entity) {
	c.table.Delete(e)
	c.world.entities[e].Clear(c.id)
}

// kind returns the kind of component (internal world ID)
func (c *Component[T]) kind() uint8 {
	return c.id
}

// typeof returns the reflected type from a generic type.
func typeof[T any]() reflect.Type {
	var v T
	return reflect.TypeOf(v)
}
