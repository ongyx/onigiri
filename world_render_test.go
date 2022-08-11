package onigiri

import (
	"errors"
	"image"
	"image/color"
	"math"
	"math/rand"
	"runtime"
	"testing"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	spriteNum  = 8
	spriteSize = 16

	maxDistance = 4

	radian = math.Pi / 180

	screenSize = 256
)

// TODO(ongyx): Is there a better way to signal early termination?
var errDone = errors.New("done")

func init() {
	rand.Seed(time.Now().Unix())
}

func TestWorldRender(t *testing.T) {
	// NOTE(ongyx): OpenGL panics on VMs with a `no context` error, so lock to the main thread for now.
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	g := &game{RenderScene()}
	if err := ebiten.RunGame(g); err != errDone {
		t.Error(err)
	}
}

type game struct {
	*World
}

func (g *game) Layout(w, h int) (sw, sh int) {
	return screenSize, screenSize
}

type Position image.Point

type Velocity struct {
	dx, dy int
}

type Transform struct {
	rotation int
}

type Sprite *ebiten.Image

type MovementSystem struct {
	view *View
}

func (m *MovementSystem) Init(w *World) {
	m.view = NewView(w, Query[Position](w), Query[Velocity](w), Query[Transform](w))
}

func (m *MovementSystem) Update(w *World) error {
	pos := Query[Position](w)
	vlc := Query[Velocity](w)
	tf := Query[Transform](w)

	m.view.Each(func(e Entity) {
		p := pos.Get(e)
		v := vlc.Get(e)
		t := tf.Get(e)

		p.X += v.dx
		p.Y += v.dy

		if p.X >= screenSize {
			p.X %= screenSize
		}

		if p.Y >= screenSize {
			p.Y %= screenSize
		}

		t.rotation++
	})

	return nil
}

type RenderSystem struct {
	view *View
}

func (r *RenderSystem) Init(w *World) {
	r.view = NewView(w, Query[Position](w), Query[Transform](w), Query[Sprite](w))
}

func (r *RenderSystem) Update(w *World) error {
	return nil
}

func (r *RenderSystem) Render(w *World, img *ebiten.Image) {
	pos := Query[Position](w)
	tf := Query[Transform](w)
	sprite := Query[Sprite](w)

	img.Fill(color.Black)

	r.view.Each(func(e Entity) {
		p := pos.Get(e)
		t := tf.Get(e)
		s := sprite.Get(e)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Rotate(float64(t.rotation) * radian)
		op.GeoM.Translate(float64(p.X), float64(p.Y))

		img.DrawImage(*s, op)
	})
}

type TimerSystem struct {
	duration time.Duration
	timer    time.Timer
}

func (t *TimerSystem) Init(w *World) {
	t.timer = *time.NewTimer(t.duration)
}

func (t *TimerSystem) Update(w *World) error {
	select {
	case <-t.timer.C:
		return errDone
	default:
		return nil
	}
}

func RenderScene() *World {
	w := NewWorld(spriteNum)

	pos := Register[Position](w, spriteNum)
	vlc := Register[Velocity](w, spriteNum)
	tf := Register[Transform](w, spriteNum)
	sprite := Register[Sprite](w, spriteNum)

	img := ebiten.NewImage(spriteSize, spriteSize)
	img.Fill(color.White)

	for i := 0; i < 7; i++ {
		e := w.Spawn()

		pos.Set(e, Position{rand.Intn(screenSize), rand.Intn(screenSize)})

		vlc.Set(e, Velocity{rand.Intn(maxDistance) + 1, rand.Intn(maxDistance) + 1})

		tf.Set(e, Transform{rand.Intn(360)})

		sprite.Set(e, img)
	}

	w.Register(&MovementSystem{}, &RenderSystem{}, &TimerSystem{duration: 5 * time.Second})

	return w
}
