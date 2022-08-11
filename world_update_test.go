package onigiri

import (
	"fmt"
	"testing"
)

func TestWorldUpdate(t *testing.T) {
	g := NewGame(&UpdateScene{}, nil)

	for i := 0; i < 5; i++ {
		g.Update()
	}
}

// components
type Text string

// systems
type PrintSystem struct {
	view *View
}

func (ps *PrintSystem) Init(w *World) {
	text := Query[Text](w)

	ps.view = NewView(w, text)
}

func (ps *PrintSystem) Update(w *World) error {
	text := Query[Text](w)

	ps.view.Each(func(e Entity) {
		fmt.Println(e, *text.Get(e))
	})

	return nil
}

// scenes
type UpdateScene struct{}

func (us *UpdateScene) Setup() *World {
	// create world
	w := NewWorld(16)

	// register components
	text := Register[Text](w, 16)

	// add components to entities
	e := w.Spawn()
	text.Set(e, Text("Hello World!"))

	// add systems
	w.Register(&PrintSystem{})

	return w
}
