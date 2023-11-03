package game

import rl "github.com/gen2brain/raylib-go/raylib"

type Background struct {
	Texture    rl.Texture2D
	ScreenSize rl.Vector2
	TileSize   rl.Vector2
	CamView    rl.Rectangle
}

func InitBackground(texture rl.Texture2D, screenSize rl.Vector2, tileSize rl.Vector2) Background {
	return Background{
		Texture:    texture,
		ScreenSize: screenSize,
		TileSize:   tileSize,
		CamView:    rl.NewRectangle(0, 0, float32(ScreenWidth), float32(ScreenHeight)),
	}
}

func (b *Background) Draw(playerPosition rl.Vector2) {

	bgX := int32(playerPosition.X) - ScreenWidth/2
	bgY := int32(playerPosition.Y) - ScreenHeight/2

	rl.DrawTexture(b.Texture, bgX, bgY, rl.White)

}
