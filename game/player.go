package game

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	PlayerDest      rl.Rectangle
	PlayerSrc       rl.Rectangle
	Position        rl.Vector2
	PlayerSprite    rl.Texture2D
	Speed           float32
	HP              int
	LVL             int
	frameCount      int
	playerMoving    bool
	playerDirection int
	shotDelay       int
	shotsRemaining  int
	Bullets         []Bullets
	FrameCount      int
	playerFrame     int
}

func InitPlayer(playerSprite rl.Texture2D) Player {
	return Player{
		PlayerDest:      rl.NewRectangle(ScreenHeight/2, ScreenWidth/2, 34, 34),
		PlayerSrc:       rl.NewRectangle(0, 0, 34, 34),
		Position:        rl.NewVector2(ScreenHeight/2, ScreenWidth/2),
		PlayerSprite:    playerSprite,
		Speed:           2,
		HP:              100,
		LVL:             0,
		frameCount:      0,
		playerMoving:    false,
		playerDirection: 0,
		shotDelay:       20,
		shotsRemaining:  0,
		Bullets:         []Bullets{},
		FrameCount:      0,
		playerFrame:     0,
	}
}

func (p *Player) Move() {
	p.playerMoving = false

	if rl.IsKeyDown(rl.KeyW) {
		p.PlayerDest.Y -= p.Speed
		p.playerMoving = true
		p.playerDirection = 1
	}
	if rl.IsKeyDown(rl.KeyS) {
		p.PlayerDest.Y += p.Speed
		p.playerMoving = true
		p.playerDirection = 0
	}
	if rl.IsKeyDown(rl.KeyA) {
		p.PlayerDest.X -= p.Speed
		p.playerMoving = true
		p.playerDirection = 2
	}
	if rl.IsKeyDown(rl.KeyD) {
		p.PlayerDest.X += p.Speed
		p.playerMoving = true
		p.playerDirection = 3
	}

	if rl.IsKeyDown(rl.KeyD) && rl.IsKeyDown(rl.KeyA) {
		if rl.IsKeyDown(rl.KeyW) {
			p.playerDirection = 1
		} else {
			p.playerDirection = 0
		}
	}

	p.Position.X = p.PlayerDest.X
	p.Position.Y = p.PlayerDest.Y
}

func (p *Player) UpdateFrame() {
	if p.FrameCount%15 == 1 {
		if p.playerMoving {
			p.playerFrame++
		} else {
			p.playerFrame = 0
		}
	}

	if p.playerFrame > 2 {
		p.playerFrame = 1
	}

	p.PlayerSrc.X = p.PlayerSrc.Width * float32(p.playerFrame)
	p.PlayerSrc.Y = p.PlayerSrc.Height * float32(p.playerDirection)
	p.FrameCount++

	if p.FrameCount == 60 {
		p.FrameCount = 0
	}
}

func (p *Player) DrawPlayer() {
	rl.DrawTexturePro(p.PlayerSprite, p.PlayerSrc, p.PlayerDest, rl.NewVector2(p.PlayerDest.Width-34, p.PlayerDest.Height-34), 0, rl.White)
}

type Bullets struct {
	Position rl.Vector2
	Velocity rl.Vector2
	Damage   int
	Active   bool
}
