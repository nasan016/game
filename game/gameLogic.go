package game

import (
	"fmt"
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// -------------------------------------------------------------------
func (g *Game) CreatePlayArea() {
	g.PlayArea1 = rl.NewVector2(g.Player.Position.X-(ScreenWidth/2), g.Player.Position.Y-(ScreenHeight/2))
	g.PlayArea2 = rl.NewVector2(g.Player.Position.X+(ScreenWidth/2), g.Player.Position.Y-(ScreenHeight/2))
	g.PlayArea3 = rl.NewVector2(g.Player.Position.X-(ScreenWidth/2), g.Player.Position.Y+(ScreenHeight/2))
	g.PlayArea4 = rl.NewVector2(g.Player.Position.X+(ScreenWidth/2), g.Player.Position.Y+(ScreenHeight/2))
}

func (g *Game) SpawnEnemy() {
	randInt := rand.Intn(4)

	minX := int64(g.PlayArea1.X)
	maxX := int64(g.PlayArea2.X)

	minY := int64(g.PlayArea1.Y)
	maxY := int64(g.PlayArea3.Y)

	var SpawnX int64
	var SpawnY int64

	switch randInt {
	case 0:
		//spawn at bottom edge
		SpawnX = rand.Int63n(int64(maxX-minX+1)) + minX
		SpawnY = maxY
	case 1:
		//spawn at right edge
		SpawnX = maxX
		SpawnY = rand.Int63n(int64(maxY-minY+1)) + minY
	case 2:
		//spawn at left edge
		SpawnX = minX
		SpawnY = rand.Int63n(int64(maxY-minY+1)) + minY
	default:
		//spawn at top edge
		SpawnX = rand.Int63n(int64(maxX-minX+1)) + minX
		SpawnY = minY
	}

	if frameCount%20 == 0 {
		g.Slimes = append(g.Slimes, NewSlime(slimeSprite, rl.NewVector2(float32(SpawnX), float32(SpawnY)), 2, 50))
	}
}

func (g *Game) DespawnEnemy() {
	for i := 0; i < len(g.Slimes); i++ {
		if (g.Slimes[i].Position.X < g.PlayArea1.X || g.Slimes[i].Position.X > g.PlayArea2.X) || (g.Slimes[i].Position.Y < g.PlayArea1.Y || g.Slimes[i].Position.Y > g.PlayArea3.Y) {
			//crazy delete element magic
			g.Slimes = append(g.Slimes[:i], g.Slimes[i+1:]...)
			fmt.Println("OUT OF BOUNDS")
		}

	}
}

func (g *Game) CheckCollision() {
	playerRect := rl.NewRectangle(g.Player.Position.X+4, g.Player.Position.Y+4, 28, 28)

	for i := range g.Slimes {
		slimeRect := rl.NewRectangle(g.Slimes[i].Position.X, g.Slimes[i].Position.Y, 34, 34)
		if rl.CheckCollisionRecs(playerRect, slimeRect) {
			g.Player.GetHit()
		}
	}
}

func (g *Game) TargetPlayer() {
	var separationRadius float32 = 35
	for i := range g.Slimes {
		direction := rl.Vector2Subtract(g.Player.Position, g.Slimes[i].Position)
		direction = rl.Vector2Normalize(direction)

		g.Slimes[i].Position.X += direction.X * g.Slimes[i].Speed
		g.Slimes[i].Position.Y += direction.Y * g.Slimes[i].Speed
		for j := range g.Slimes {
			if i != j {
				dx := g.Slimes[i].Position.X - g.Slimes[j].Position.X
				dy := g.Slimes[i].Position.Y - g.Slimes[j].Position.Y
				distance := rl.Vector2Length(rl.NewVector2(dx, dy))

				if distance < separationRadius {
					separationAmount := float32((separationRadius - distance) / 2.0)
					angle := math.Atan2(float64(dy), float64(dx))
					separationVector := rl.NewVector2(separationAmount*float32(math.Cos(angle)), separationAmount*float32(math.Sin(angle)))
					g.Slimes[i].Position = rl.Vector2Add(g.Slimes[i].Position, separationVector)
				}
			}
		}
	}
}

//------------------------------------------------------------
