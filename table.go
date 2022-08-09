package onigiri

type table interface {
	Delete(e Entity)
}

// Table is a map of entities to indices into a dense slice of components of type T.
// It allows fast insertion and removal, but the order of components are not guaranteed.
type Table[T any] struct {
	entities   map[Entity]int
	components []T
}

// NewTable creates a new map with a capacity of (size) components.
func NewTable[T any](size int) *Table[T] {
	return &Table[T]{
		entities:   make(map[Entity]int, size),
		components: make([]T, size),
	}
}

// Get returns a reference to the component indexed by the entity.
// If the entity is not in the table, nil is returned.
func (t *Table[T]) Get(e Entity) *T {
	if idx, ok := t.entities[e]; ok {
		return &t.components[idx]
	} else {
		return nil
	}
}

// Set inserts a entity-component pair.
func (t *Table[T]) Set(e Entity, c T) {
	// length is always last index + 1
	t.entities[e] = len(t.components)
	t.components = append(t.components, c)
}

// Delete removes the component by entity, swapping the index of the last element in place.
func (t *Table[T]) Delete(e Entity) {
	idx := t.entities[e]
	last := len(t.components) - 1

	t.components[idx] = t.components[last]
	t.components = t.components[:len(t.components)-1]
	delete(t.entities, e)
}
