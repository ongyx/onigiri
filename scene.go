package onigiri

// Scene creates and initializes a world by spawning entities, registering components, and systems.
type Scene interface {
	Setup() *World
}
