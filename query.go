package onigiri

import "reflect"

type Query[T any] struct {
	t reflect.Type
	w *World
	a Archetype[T]
}

func NewQuery[T any](w *World) *Query[T] {
	q := &Query[T]{t: typeof[T](), w: w}

	if a, ok := w.archetypes[q.t]; ok {
		q.a = a.(Archetype[T])
	}

	return q
}

func (q *Query[T]) Init(size int) {
	if q.a == nil {
		q.Use(NewTable[T](size))
	}
}

func (q *Query[T]) Use(a Archetype[T]) {
	q.a = a
	q.w.archetypes[q.t] = q.a
}

func (q *Query[T]) Get(e Entity) *T {
	return q.a.Get(e)
}

func (q *Query[T]) Insert(e Entity, c T) {
	q.a.Insert(e, c)
}
