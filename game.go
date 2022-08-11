package onigiri

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// GameOptions provides configuration for how the game is run.
type GameOptions struct {
	// Size sets the logical screen size.
	Size image.Point
	// Scale sets the logical to physical pixel scale.
	Scale float64
}

// DefaultGameOptions returns the default game options.
// By default, the scale is set to the device scale factor as defined by Ebiten.
func DefaultGameOptions() *GameOptions {
	return &GameOptions{
		Scale: 1,
	}
}

// Game is a wrapper around a World that implements the ebiten.Game interface.
type Game struct {
	scene   Scene
	world   *World
	options *GameOptions
}

func NewGame(initial Scene, options *GameOptions) *Game {
	if options == nil {
		options = DefaultGameOptions()
	}

	g := &Game{options: options}
	g.SetScene(initial)

	return g
}

func (g *Game) SetScene(scene Scene) {
	g.scene = scene
	g.world = scene.Setup()
}

func (g *Game) Update() error {
	return g.world.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.world.Draw(screen)
}

func (g *Game) Layout(lw, lh int) (pw, ph int) {
	size := g.options.Size

	if size != (image.Point{}) {
		lw = size.X
		lh = size.Y
	}

	scale := g.options.Scale
	pw = int(float64(lw) * scale)
	ph = int(float64(lh) * scale)

	return
}
