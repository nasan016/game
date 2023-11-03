package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Bullet struct {
	Position rl.Vector2
	Velocity rl.Vector2
	Damage   int
	Active   bool
}

func InitBullet(position rl.Vector2, velocity rl.Vector2, damage int) Bullet {
	return Bullet{
		Position: position,
		Velocity: velocity,
		Damage:   damage,
		Active:   true,
	}
}

func (g *Game) UpdateBullets() {
	for i := 0; i < len(g.Bullets); i++ {
		if g.Bullets[i].Active {
			g.Bullets[i].Position.X += g.Bullets[i].Velocity.X
			g.Bullets[i].Position.Y += g.Bullets[i].Velocity.Y
		}
		if g.Bullets[i].Position.X < g.PlayArea1.X || g.Bullets[i].Position.X > g.PlayArea2.X {
			g.Bullets[i].Active = false
		}
	}
}

func (g *Game) DrawBullets() {
	for _, bullet := range g.Bullets {
		if bullet.Active {
			rl.DrawCircleV(bullet.Position, 8, rl.Black)
		}
	}
}

func (g *Game) BulletInput() {
	if !g.Player.Reloading {
		if rl.IsKeyDown(rl.KeyRight) && shotsRemaining == 0 {

			g.Bullets = append(g.Bullets, InitBullet(rl.NewVector2(g.Player.PlayerDest.X+17, g.Player.PlayerDest.Y+17), rl.NewVector2(5, 0), 1))
			g.Bullets = append(g.Bullets, InitBullet(rl.NewVector2(g.Player.PlayerDest.X+17, g.Player.PlayerDest.Y+17), rl.NewVector2(5, 0), 1))
			g.Bullets = append(g.Bullets, InitBullet(rl.NewVector2(g.Player.PlayerDest.X+17, g.Player.PlayerDest.Y+17), rl.NewVector2(5, 0), 1))
			shotsRemaining = shotDelay
			g.Player.BulletsShot++
		}
		if rl.IsKeyDown(rl.KeyLeft) && shotsRemaining == 0 {

			g.Bullets = append(g.Bullets, InitBullet(rl.NewVector2(g.Player.PlayerDest.X+17, g.Player.PlayerDest.Y+17), rl.NewVector2(5, 0), 1))
			g.Bullets = append(g.Bullets, InitBullet(rl.NewVector2(g.Player.PlayerDest.X+17, g.Player.PlayerDest.Y+17), rl.NewVector2(5, 0), 1))
			g.Bullets = append(g.Bullets, InitBullet(rl.NewVector2(g.Player.PlayerDest.X+17, g.Player.PlayerDest.Y+17), rl.NewVector2(5, 0), 1))
			shotsRemaining = shotDelay
			g.Player.BulletsShot++
		}
		if rl.IsKeyDown(rl.KeyUp) && shotsRemaining == 0 {
			g.Bullets = append(g.Bullets, InitBullet(rl.NewVector2(g.Player.PlayerDest.X+17, g.Player.PlayerDest.Y+17), rl.NewVector2(5, 0), 1))
			g.Bullets = append(g.Bullets, InitBullet(rl.NewVector2(g.Player.PlayerDest.X+17, g.Player.PlayerDest.Y+17), rl.NewVector2(5, 0), 1))
			g.Bullets = append(g.Bullets, InitBullet(rl.NewVector2(g.Player.PlayerDest.X+17, g.Player.PlayerDest.Y+17), rl.NewVector2(5, 0), 1))
			shotsRemaining = shotDelay
			g.Player.BulletsShot++
		}

		if rl.IsKeyDown(rl.KeyDown) && shotsRemaining == 0 {
			g.Bullets = append(g.Bullets, InitBullet(rl.NewVector2(g.Player.PlayerDest.X+17, g.Player.PlayerDest.Y+17), rl.NewVector2(5, 0), 1))
			g.Bullets = append(g.Bullets, InitBullet(rl.NewVector2(g.Player.PlayerDest.X+17, g.Player.PlayerDest.Y+17), rl.NewVector2(5, 0), 1))
			g.Bullets = append(g.Bullets, InitBullet(rl.NewVector2(g.Player.PlayerDest.X+17, g.Player.PlayerDest.Y+17), rl.NewVector2(5, 0), 1))
			shotsRemaining = shotDelay
			g.Player.BulletsShot++
		}
		print("YO")
	}
}
