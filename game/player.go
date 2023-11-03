package game

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	PlayerDest   rl.Rectangle
	PlayerSrc    rl.Rectangle
	Position     rl.Vector2
	PlayerSprite rl.Texture2D
	Speed        float32
	HP           int
	LVL          int
}

type Bullets struct {
	Position rl.Vector2
	Velocity rl.Vector2
	Damage   int
	Active   bool
}
