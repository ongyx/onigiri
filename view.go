package onigiri

// View is a cached query that can be used by systems to search for entities with specific components.
type View struct {
	world *World
	sig   Signature
}

// NewView creates a new view with the signatures of the components.
func NewView(w *World, components ...ComponentKind) *View {
	v := &View{world: w}
	for _, c := range components {
		v.sig.Set(c.kind())
	}
	return v
}

// Each calls the function for each entity that matches the view's signature.
// Entities are not guaranteed to be in order; if you want to sort them before iteration use Filter.
func (v *View) Each(f func(e Entity)) {
	for e, sig := range v.world.entities {
		if sig.Contains(v.sig) {
			f(e)
		}
	}
}

// Filter appends all entities that match the view's signature to a buffer and returns it.
// The buffer may be nil, and can be reused across calls to Filter to reduce allocation.
func (v *View) Filter(buf []Entity) []Entity {
	// NOTE(ongyx): this avoids alloc if the buffer already has an underlying array
	buf = buf[:0]

	for e, sig := range v.world.entities {
		if sig.Contains(v.sig) {
			buf = append(buf, e)
		}
	}

	return buf
}
