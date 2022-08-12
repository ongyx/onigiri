# DEPRECIATED

Onigiri has been merged into [Bento] under the `ecs/` folder.

# Onigiri

Experimental ECS (Entity Component System) framework for [Ebitengine].

**DISCLAIMER**: Onigiri is still in the early stages of development! The API is subject to breaking changes.

## Rationale

When I got around to creating a demo game with Bento, I realised that it was difficult to separate scripting a game object from the concrete implementation.
With this in mind, I decided to try an ECS approach with generics from scratch.

As with Bento, Ebitengine provides the foundation for efficient updating/drawing.
Onigiri aims to provide niceties on top of that found in other game engine systems.

In future, some parts of Bento such as animations, fonts, and drawing utilities such as `Vec` may be moved into their own `bento` subdirectory.
I recognize they are still useful to do common tasks when creating a game.

## Credits

Hajime Hoshi for creating Ebiten.

## License

Onigiri is licensed under the MIT License.

[Bento]: https://github.com/ongyx/bento
[Ebitengine]: https://github.com/hajimehoshi/ebiten
