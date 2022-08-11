package onigiri

// Query returns the component by generic type in the world.
// The type must have already been registered with Register.
func Query[T any](w *World) *Component[T] {
	return newComponent[T](w, w.tables[typeof[T]()])
}

// Register allocates a table with capacity for (size) components in the world.
// Subsequent calls to Register will panic if it was already called once,
// or there are too many component types registered.
func Register[T any](w *World, size int) *Component[T] {
	if w.tableID >= 64 {
		panic("component: too many registered")
	}

	w.tableID++

	e := entry{
		table: NewTable[T](size),
		id:    w.tableID,
	}
	w.tables[typeof[T]()] = e

	return newComponent[T](w, e)
}
