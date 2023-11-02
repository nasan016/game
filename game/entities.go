package game

import rl "github.com/gen2brain/raylib-go/raylib"

type Slime struct {
	Position rl.Vector2
	Speed    float32

	Texture rl.Texture2D
}
