package game

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	PlayerDest   rl.Rectangle
	PlayerSrc    rl.Rectangle
	Position     rl.Vector2
	PlayerSprite rl.Texture2D

	Speed float32
}
