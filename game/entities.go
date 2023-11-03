package game

import rl "github.com/gen2brain/raylib-go/raylib"

type Slime struct {
	Texture  rl.Texture2D
	Position rl.Vector2

	Speed float32
	HP    int
}

func NewSlime(texture rl.Texture2D, position rl.Vector2, speed float32, hp int) Slime {
	return Slime{
		Texture:  texture,
		Position: position,
		Speed:    speed,
		HP:       hp,
	}
}

func (s *Slime) DrawSlime(Player Player) {
	if s.Position.X < Player.Position.X {
		width := s.Texture.Width
		height := s.Texture.Height

		rl.DrawTextureRec(s.Texture, rl.NewRectangle(float32(width), 0, -float32(width), float32(height)), s.Position, rl.White)
	} else {
		rl.DrawTextureEx(s.Texture, s.Position, 0, 1, rl.White)
	}
}
