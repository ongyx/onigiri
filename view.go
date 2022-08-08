package onigiri

const viewSize = 64

// View is a cached query that can be used by systems to search for entities with specific components.
type View struct {
	entities []Entity
	sig      Signature
}

// NewView creates a new view with the signatures of the components.
func NewView(components ...ComponentKind) *View {
	v := &View{entities: make([]Entity, 0, viewSize)}
	for _, c := range components {
		v.sig.Set(c.kind())
	}
	return v
}

// Query searches for all entities that matches the view's signature.
// Entities are not guaranteed to be in order.
func (v *View) Query(w *World) {
	// NOTE(ongyx): reslicing to zero avoids having to alloc again
	v.entities = v.entities[:0]

	// TODO(ongyx): maybe use a 'dirty' world flag or callbacks to indicate a entity has changed signature?
	for e, sig := range w.entities {
		if sig.Contains(v.sig) {
			v.entities = append(v.entities, e)
		}
	}
}

// Each calls the function on each entity queried.
func (v *View) Each(f func(e Entity)) {
	for _, e := range v.entities {
		f(e)
	}
}
