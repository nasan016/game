package game

import rl "github.com/gen2brain/raylib-go/raylib"

type Slime struct {
	Texture rl.Texture2D

	Position rl.Vector2
	Speed    float32
	HP       int
}
