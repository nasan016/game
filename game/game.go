package game

import (
	"fmt"
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ScreenWidth  = 1600
	ScreenHeight = 900
)

var (
	cam        rl.Camera2D
	frameCount int

	backgroundTexture rl.Texture2D
	tileWidth         int32
	tileHeight        int32

	playerSprite                                  rl.Texture2D
	playerFrame                                   int
	playerMoving                                  bool
	playerDirection                               int
	playerUp, playerDown, playerRight, playerLeft bool
	playerFrameCount                              int

	slimeSprite rl.Texture2D
)

type Game struct {
	Player Player

	PlayArea1 rl.Vector2
	PlayArea2 rl.Vector2
	PlayArea3 rl.Vector2
	PlayArea4 rl.Vector2

	Slimes []Slime

	WindowShouldClose bool
}

func NewGame() (g Game) {
	g.Init()
	return
}

func (g *Game) Init() {
	backgroundTexture = rl.LoadTexture("res/background.jpg")
	tileWidth = backgroundTexture.Width
	tileHeight = backgroundTexture.Height

	playerSprite = rl.LoadTexture("res/player-sprites.png")
	slimeSprite = rl.LoadTexture("res/slime.png")

	g.Player = Player{
		rl.NewRectangle(ScreenHeight/2, ScreenWidth/2, 34, 34),
		rl.NewRectangle(0, 0, 34, 34),
		rl.NewVector2(g.Player.PlayerDest.X, g.Player.PlayerDest.Y),
		playerSprite,
		2.0,
	}

	g.CreatePlayArea()
	frameCount = 0

	cam = rl.NewCamera2D(rl.NewVector2(float32(ScreenWidth/2), float32(ScreenHeight/2)), g.Player.Position, 0, 2.0)
}

func (g *Game) Unload() {
	rl.UnloadTexture(playerSprite)
}

func (g *Game) CreatePlayArea() {
	g.PlayArea1 = rl.NewVector2(g.Player.Position.X-(ScreenWidth/2), g.Player.Position.Y-(ScreenHeight/2))
	g.PlayArea2 = rl.NewVector2(g.Player.Position.X+(ScreenWidth/2), g.Player.Position.Y-(ScreenHeight/2))
	g.PlayArea3 = rl.NewVector2(g.Player.Position.X-(ScreenWidth/2), g.Player.Position.Y+(ScreenHeight/2))
	g.PlayArea4 = rl.NewVector2(g.Player.Position.X+(ScreenWidth/2), g.Player.Position.Y+(ScreenHeight/2))
}

func (g *Game) TargetPlayer() {
	var seperationRadius float32 = 30
	for i := range g.Slimes {
		g.Slimes[i].Position.X += (g.Player.Position.X - g.Slimes[i].Position.X) / 250
		g.Slimes[i].Position.Y += (g.Player.Position.Y - g.Slimes[i].Position.Y) / 250

		for j := range g.Slimes {
			if i != j {
				dx := g.Slimes[i].Position.X - g.Slimes[j].Position.X
				dy := g.Slimes[i].Position.Y - g.Slimes[j].Position.Y
				distance := rl.Vector2Length(rl.NewVector2(dx, dy))

				if distance < seperationRadius {
					seperationAmount := float64((seperationRadius - distance) / 2.0)
					angle := math.Atan2(float64(dy), float64(dx))
					seperationVector := rl.NewVector2(float32(seperationAmount*math.Cos(angle)), float32(seperationAmount*math.Sin(angle)))
					g.Slimes[i].Position = rl.Vector2Add(g.Slimes[i].Position, seperationVector)
				}
			}
		}
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
		g.Slimes = append(g.Slimes, Slime{rl.NewVector2(float32(SpawnX), float32(SpawnY)), 500, slimeSprite})
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

	if playerMoving {
		if frameCount%15 == 1 {
			playerFrame++
		}
	}

	if playerFrame > 2 {
		playerFrame = 1
	}

	g.Player.PlayerSrc.X = g.Player.PlayerSrc.Width * float32(playerFrame)
	g.Player.PlayerSrc.Y = g.Player.PlayerSrc.Height * float32(playerDirection)
	//debug
	if frameCount%60 == 0 {
		fmt.Println(len(g.Slimes))
		fmt.Println(g.PlayArea1, g.PlayArea2, g.PlayArea3, g.PlayArea4)
	}
	//update frame count
	frameCount++
	if frameCount == 60 {
		frameCount = 0
	}
}

func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	backgroundOffsetX := int(g.Player.Position.X) % int(tileWidth)
	backgroundOffsetY := int(g.Player.Position.Y) % int(tileHeight)

	for y := -backgroundOffsetY; y < rl.GetScreenHeight(); y += int(tileHeight) {
		for x := -backgroundOffsetX; x < rl.GetScreenWidth(); x += int(tileWidth) {
			rl.DrawTexture(backgroundTexture, int32(x), int32(y), rl.White)
		}
	}
	rl.BeginMode2D(cam)

	//render da Slimess (they're just circles rn :3)
	for i := 0; i < len(g.Slimes); i++ {
		rl.DrawTexture(g.Slimes[i].Texture, int32(g.Slimes[i].Position.X), int32(g.Slimes[i].Position.Y), rl.White)
	}
	rl.DrawTexturePro(g.Player.PlayerSprite, g.Player.PlayerSrc, g.Player.PlayerDest, rl.NewVector2(g.Player.PlayerDest.Width/2, g.Player.PlayerDest.Height/2), 0, rl.White)
	rl.EndDrawing()
}

func (g *Game) Input() {
	if rl.IsKeyDown(rl.KeyW) {
		g.Player.PlayerDest.Y -= g.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyS) {
		g.Player.PlayerDest.Y += g.Player.Speed
		playerMoving = true
		playerDirection = 0
	}
	if rl.IsKeyDown(rl.KeyA) {
		g.Player.PlayerDest.X -= g.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyD) {
		g.Player.PlayerDest.X += g.Player.Speed
	}

	g.Player.Position.X = g.Player.PlayerDest.X
	g.Player.Position.Y = g.Player.PlayerDest.Y
}
