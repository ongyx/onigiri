package onigiri

import "github.com/hajimehoshi/ebiten/v2"

// System is an opaque state that updates entities.
type System interface {
	Init(w *World)
	Update(w *World) error
}

// Renderer is an opaque state that renders entities to the screen.
type Renderer interface {
	Render(w *World, i *ebiten.Image)
}
