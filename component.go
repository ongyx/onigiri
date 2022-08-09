package onigiri

import "reflect"

// ComponentKind represents a specific kind of component in a world.
type ComponentKind interface {
	kind() uint8
}

// Component is a unique identifer for a component type.
type Component[T any] struct {
	rtype reflect.Type

	world *World
	table *Table[T]
	id    uint8
}

// NewComponent creates a new component associated with a world.
func NewComponent[T any](w *World) *Component[T] {
	c := &Component[T]{rtype: typeof[T](), world: w}

	if e, ok := w.tables[c.rtype]; ok {
		c.table = e.table.(*Table[T])
		c.id = e.id
	}

	return c
}

// Register allocates a table with capacity for (size) components in the world.
// Subsequent calls to Register will panic if it was already called once,
// or there are too many component types registered.
func (c *Component[T]) Register(size int) {
	if c.table != nil {
		panic("component: type already registered")
	}

	if c.world.tableID >= 64 {
		panic("component: too many registered")
	}

	c.world.tableID++
	c.id = c.world.tableID

	c.table = NewTable[T](size)
	c.world.tables[c.rtype] = entry{c.table, c.id}
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
