package game

import (
	"fmt"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ScreenWidth  = 900
	ScreenHeight = 800
)

var (
	cam rl.Camera2D

	playerSprite rl.Texture2D
	playerSrc    rl.Rectangle
	playerDest   rl.Rectangle

	frameCount int
)

type Game struct {
	Player Player

	PlayArea1 rl.Vector2
	PlayArea2 rl.Vector2
	PlayArea3 rl.Vector2
	PlayArea4 rl.Vector2

	Slime []Slime

	WindowShouldClose bool
}

func NewGame() (g Game) {
	g.Init()
	return
}

func (g *Game) Init() {
	g.Player = Player{rl.NewVector2(ScreenWidth/2, ScreenHeight/2), 1.8}
	playerSrc = rl.NewRectangle(0, 0, 34, 34)
	playerDest = rl.NewRectangle(ScreenWidth/2, ScreenHeight/2, 100, 100)
	g.CreatePlayArea()
	frameCount = 0

	cam = rl.NewCamera2D(g.Player.Position, g.Player.Position, 0, 1.1)
}

func (g *Game) CreatePlayArea() {
	g.PlayArea1 = rl.NewVector2(g.Player.Position.X-(ScreenWidth/2), g.Player.Position.Y-(ScreenHeight/2))
	g.PlayArea2 = rl.NewVector2(g.Player.Position.X+(ScreenWidth/2), g.Player.Position.Y-(ScreenHeight/2))
	g.PlayArea3 = rl.NewVector2(g.Player.Position.X-(ScreenWidth/2), g.Player.Position.Y+(ScreenHeight/2))
	g.PlayArea4 = rl.NewVector2(g.Player.Position.X+(ScreenWidth/2), g.Player.Position.Y+(ScreenHeight/2))
}

func (g *Game) TargetPlayer() {
	for i := 0; i < len(g.Slime); i++ {
		g.Slime[i].Position.X += (g.Player.Position.X - g.Slime[i].Position.X) / 250
		g.Slime[i].Position.Y += (g.Player.Position.Y - g.Slime[i].Position.Y) / 250
	}
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
		g.Slime = append(g.Slime, Slime{rl.NewVector2(float32(SpawnX), float32(SpawnY)), 500})
	}
}

func (g *Game) DespawnEnemy() {
	for i := 0; i < len(g.Slime); i++ {
		if (g.Slime[i].Position.X < g.PlayArea1.X || g.Slime[i].Position.X > g.PlayArea2.X) || (g.Slime[i].Position.Y < g.PlayArea1.Y || g.Slime[i].Position.Y > g.PlayArea3.Y) {
			//crazy delete element magic
			g.Slime = append(g.Slime[:i], g.Slime[i+1:]...)
		}

	}
}

func (g *Game) Quit() {
	//unload assets
}

func (g *Game) Update() {
	if rl.WindowShouldClose() {
		g.WindowShouldClose = true
	}

	//spawn enemy on an interval
	g.SpawnEnemy()

	//update play area whenever player moves
	//effects spawn region for enemies
	g.CreatePlayArea()

	//move enemies towards the player
	g.TargetPlayer()

	//Despawn Enemy
	g.DespawnEnemy()
	//update camera to follow player
	cam.Target = g.Player.Position

	//debug
	if frameCount%60 == 0 {
		fmt.Println(len(g.Slime))
		fmt.Println(g.PlayArea1, g.PlayArea2, g.PlayArea3, g.PlayArea4)
	}
	//update frame count
	frameCount++
}

func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Blue)
	rl.BeginMode2D(cam)

	//render da slimes (they're just circles rn :3)
	for i := 0; i < len(g.Slime); i++ {
		rl.DrawCircleV(g.Slime[i].Position, 10, rl.Brown)
	}

	rl.DrawCircleV(g.Player.Position, 10, rl.White)
	rl.EndDrawing()
}

func (g *Game) Input() {
	if rl.IsKeyDown(rl.KeyW) {
		g.Player.Position.Y -= g.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyS) {
		g.Player.Position.Y += g.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyA) {
		g.Player.Position.X -= g.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyD) {
		g.Player.Position.X += g.Player.Speed
	}
}
