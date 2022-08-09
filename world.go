package onigiri

import (
	"reflect"
	"sync/atomic"

	"github.com/hajimehoshi/ebiten/v2"
)

const maxTables = 64

type entry struct {
	table
	id uint8
}

// World is the container of all data related to entities, components and systems.
type World struct {
	tables  map[reflect.Type]entry
	tableID uint8

	entities map[Entity]*Signature
	entityID uint64

	systems   []System
	renderers []Renderer

	init bool
}

// NewWorld creates a new world with an inital capacity of (size) entities.
func NewWorld(size int) *World {
	return &World{
		tables:   make(map[reflect.Type]entry, maxTables),
		entities: make(map[Entity]*Signature, size),
	}
}

// Spawn creates a new entity and adds it to the world.
// Entity IDs are guaranteed to be unique to seperate worlds, and not collide with any despawned entities.
func (w *World) Spawn() Entity {
	// TODO(ongyx): recycle despawned entities?
	e := Entity{id: atomic.AddUint64(&w.entityID, 1)}

	var s Signature
	w.entities[e] = &s

	return e
}

// Despawn destroys the entity, removing it from the world.
func (w *World) Despawn(e Entity) {
	for _, t := range w.tables {
		t.Delete(e)
	}

	delete(w.entities, e)
}

// Register adds the systems to the world.
// A System may optionally implement Renderer, whose Render method is called when drawing to the screen.
func (w *World) Register(systems ...System) {
	for _, s := range systems {
		w.systems = append(w.systems, s)
		if r, ok := s.(Renderer); ok {
			w.renderers = append(w.renderers, r)
		}
	}
}

// Update updates the state of the world's systems in the order they were registered.
// This implements the ebiten.Game interface.
func (w *World) Update() error {
	for _, s := range w.systems {
		if !w.init {
			s.Init(w)
		}

		if err := s.Update(w); err != nil {
			return err
		}
	}

	// all systems initalised by this point
	w.init = true

	return nil
}

// Draw renders all renderers to the screen in the order they were registered.
// This implements the ebiten.Game interface.
func (w *World) Draw(screen *ebiten.Image) {
	for _, r := range w.renderers {
		r.Render(w, screen)
	}
}
