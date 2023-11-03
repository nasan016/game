package game

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	PlayerDest       rl.Rectangle
	PlayerSrc        rl.Rectangle
	Position         rl.Vector2
	PlayerSprite     rl.Texture2D
	Speed            float32
	HP               int
	LVL              int
	frameCount       int
	playerMoving     bool
	playerDirection  int
	shotDelay        int
	shotsRemaining   int
	FrameCount       int
	playerFrame      int
	IsInvincible     bool
	InvincibilityEnd time.Time
	BulletsShot      int
	ReloadTimer      time.Time
	Reloading        bool
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
		FrameCount:      0,
		playerFrame:     0,
		BulletsShot:     0,
		Reloading:       false,
	}
}

/*
Logic for player getting hit is split into two parts
GetHit() and ResetInvincibility()
*/

func (p *Player) IsReloading() {
	if p.BulletsShot == 2 {
		p.BulletsShot = 0
		p.Reloading = true
		p.ReloadTimer = time.Now().Add(1500 * time.Millisecond)
	}

	if time.Now().After(p.ReloadTimer) {
		p.Reloading = false
	}
}

func (p *Player) GetHit() {
	if !p.IsInvincible {
		p.HP -= 10
		p.IsInvincible = true
		p.InvincibilityEnd = time.Now().Add(1 * time.Second)
	}
}

func (p *Player) ResetInvincibility() {
	if p.IsInvincible && time.Now().After(p.InvincibilityEnd) {
		p.IsInvincible = false
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
	if p.IsInvincible {
		if p.FrameCount%30 < 5 {

		} else {
			rl.DrawTexturePro(p.PlayerSprite, p.PlayerSrc, p.PlayerDest, rl.NewVector2(p.PlayerDest.Width-34, p.PlayerDest.Height-34), 0, rl.DarkBlue)
		}
	} else {
		rl.DrawTexturePro(p.PlayerSprite, p.PlayerSrc, p.PlayerDest, rl.NewVector2(p.PlayerDest.Width-34, p.PlayerDest.Height-34), 0, rl.White)
	}
}
