package onigiri

import "reflect"

// ComponentKind represents a specific kind of component in a world.
type ComponentKind interface {
	kind() uint64
}

// Component is a unique identifer for a component type.
type Component[T any] struct {
	w  *World
	id uint64
	t  *Table[T]
}

func NewComponent[T any](w *World) *Component[T] {
	if w.tableID >= 64 {
		panic("component: too many registered")
	}

	rt := typeof[T]()
	if _, ok := w.tables[rt]; ok {
		panic("component: type already registered")
	}

	id := w.tableID
	w.tableID++

	t := NewTable[T](256)
	w.tables[rt] = t

	return &Component[T]{w, id, t}
}

func (c *Component[T]) Get(e Entity) *T {
	return c.t.Get(e)
}

func (c *Component[T]) Set(e Entity, com T) {
	c.t.Set(e, com)
	c.w.entities[e].Set(c.id)
}

func (c *Component[T]) Delete(e Entity) {
	c.t.Delete(e)
	c.w.entities[e].Clear(c.id)
}

func (c *Component[T]) kind() uint64 {
	return c.id
}

func typeof[T any]() reflect.Type {
	var v T
	return reflect.TypeOf(v)
}
