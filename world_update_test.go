package onigiri

import (
	"fmt"
	"testing"
)

type Text string

type PrintSystem struct {
	text *Component[Text]
	view *View
}

func (ps *PrintSystem) Init(w *World) {
	ps.text = NewComponent[Text](w)
	ps.view = NewView(w, ps.text)
}

func (ps *PrintSystem) Update(w *World) error {
	ps.view.Each(func(e Entity) {
		fmt.Println(e, *ps.text.Get(e))
	})

	return nil
}

func TestWorld(t *testing.T) {
	// create world
	w := NewWorld(16)

	// register components
	tc := NewComponent[Text](w)
	tc.Register(16)

	// add components to entities
	e := w.Spawn()
	tc.Set(e, Text("Hello World!"))

	// add systems
	w.Register(&PrintSystem{})

	for i := 0; i < 5; i++ {
		w.Update()
	}
}
