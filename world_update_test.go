package onigiri

import (
	"fmt"
	"testing"
)

func TestWorldUpdate(t *testing.T) {
	w := UpdateScene()

	for i := 0; i < 5; i++ {
		w.Update()
	}
}

type Text string

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

func UpdateScene() *World {
	w := NewWorld(1)

	text := Register[Text](w, 1)

	e := w.Spawn()
	text.Set(e, Text("Hello World!"))

	w.Register(&PrintSystem{})

	return w
}
