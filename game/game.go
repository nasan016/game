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

	backgroundOffsetX int32
	backgroundOffsetY int32

	slimeSprite rl.Texture2D

	shotDelay      int
	shotsRemaining int
)

type Game struct {
	Player Player

	PlayArea1 rl.Vector2
	PlayArea2 rl.Vector2
	PlayArea3 rl.Vector2
	PlayArea4 rl.Vector2

	Slimes []Slime

	Bullets           []Bullets
	WindowShouldClose bool
}

func NewGame() (g Game) {
	g.Init()
	return
}

func (g *Game) Init() {
	backgroundTexture = rl.LoadTexture("res/background.png")
	tileWidth = backgroundTexture.Width
	tileHeight = backgroundTexture.Height

	playerSprite = rl.LoadTexture("res/player-sprites.png")
	slimeSprite = rl.LoadTexture("res/slime.png")

	g.Player = Player{
		rl.NewRectangle(ScreenHeight/2, ScreenWidth/2, 34, 34),
		rl.NewRectangle(0, 0, 34, 34),
		rl.NewVector2(g.Player.PlayerDest.X, g.Player.PlayerDest.Y),
		playerSprite,
		1.5,
		100,
		0,
	}

	g.CreatePlayArea()
	frameCount = 0

	cam = rl.NewCamera2D(rl.NewVector2(float32(ScreenWidth/2)-34, float32(ScreenHeight/2)-34), rl.NewVector2(float32(g.Player.PlayerDest.X-(g.Player.PlayerDest.Width/2)), float32(g.Player.PlayerDest.Y-g.Player.PlayerDest.Height/2)), 0, 2.0)

	shotDelay = 20
	shotsRemaining = 0
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

func (g *Game) CheckCollision() {
	playerRect := rl.NewRectangle(g.Player.Position.X+4, g.Player.Position.Y+4, 28, 28)

	for i := range g.Slimes {
		slimeRect := rl.NewRectangle(g.Slimes[i].Position.X, g.Slimes[i].Position.Y, 34, 24)
		if rl.CheckCollisionRecs(playerRect, slimeRect) {
			fmt.Println("hit")
			g.Player.HP -= 10
		}
	}
}

func (g *Game) TargetPlayer() {
	var seperationRadius float32 = 35
	for i := range g.Slimes {
		g.Slimes[i].Position.X += (g.Player.Position.X - g.Slimes[i].Position.X) / 300
		g.Slimes[i].Position.Y += (g.Player.Position.Y - g.Slimes[i].Position.Y) / 300

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
		g.Slimes = append(g.Slimes, Slime{slimeSprite, rl.NewVector2(float32(SpawnX), float32(SpawnY)), 500, 50})
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

func (g *Game) Quit() {
	//unload assets
}

func (g *Game) Update() {
	if rl.WindowShouldClose() {
		g.WindowShouldClose = true
	}

	g.Player.Position.X = g.Player.PlayerDest.X
	g.Player.Position.Y = g.Player.PlayerDest.Y

	backgroundOffsetX = int32(g.Player.Position.X) % int32(tileWidth)
	backgroundOffsetY = int32(g.Player.Position.Y) % int32(tileHeight)

	//spawn enemy on an interval
	g.SpawnEnemy()

	//update play area whenever player moves
	//effects spawn region for enemies
	g.CreatePlayArea()

	//move enemies towards the player
	g.TargetPlayer()

	//lose health
	g.CheckCollision()
	//Despawn Enemy
	g.DespawnEnemy()

	g.UpdateBullets()
	//update camera to follow player
	cam.Target = g.Player.Position

	if shotsRemaining > 0 {
		shotsRemaining--
	}
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
		fmt.Println(g.Player.Position)
		fmt.Println(g.Bullets)
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

	rl.BeginMode2D(cam)
	//TODO fix dis background
	for y := backgroundOffsetY - backgroundTexture.Height; y < int32(rl.GetScreenHeight()); y += backgroundTexture.Height {
		for x := backgroundOffsetX - backgroundTexture.Width; x < int32(rl.GetScreenWidth()); x += backgroundTexture.Width {
			rl.DrawTexture(backgroundTexture, x, y, rl.White)
		}
	}

	//draw slimes
	for i := 0; i < len(g.Slimes); i++ {
		if g.Slimes[i].Position.X < g.Player.Position.X {
			width := slimeSprite.Width
			height := slimeSprite.Height

			rl.DrawTextureRec(g.Slimes[i].Texture, rl.NewRectangle(float32(width), 0, -float32(width), float32(height)), g.Slimes[i].Position, rl.White)
		} else {
			rl.DrawTextureEx(g.Slimes[i].Texture, g.Slimes[i].Position, 0, 1, rl.White)
		}
	}

	//draw player sprite
	rl.DrawTexturePro(g.Player.PlayerSprite, g.Player.PlayerSrc, g.Player.PlayerDest, rl.NewVector2(g.Player.PlayerDest.Width-34, g.Player.PlayerDest.Height-34), 0, rl.White)
	for _, bullet := range g.Bullets {
		if bullet.Active {
			rl.DrawCircleV(bullet.Position, 5, rl.Black)
		}
	}
	rl.EndMode2D()

	//draw the hp
	rl.DrawText(fmt.Sprintf("HP: %02d / 100", g.Player.HP), 5, 40, 20, rl.Black)

	//draw xp bar on top
	rl.DrawRectangle(0, 0, int32(ScreenWidth), 40, rl.Black)
	rl.DrawRectangle(5, 5, int32(ScreenWidth)-10, 40-10, rl.Black)
	rl.DrawRectangle(5, 5, int32(ScreenWidth*.5)-10, 40-10, rl.Yellow)
	rl.DrawRectangle(5, 5, int32(100), 40-10, rl.Pink)
	rl.DrawRectangle(105, 5, int32(20), 40-10, rl.LightGray)
	rl.DrawText(fmt.Sprintf("LVL: %02d", g.Player.LVL), 10, 10, 20, rl.Black)

	rl.EndDrawing()
}

func (g *Game) Input() {
	if rl.IsKeyDown(rl.KeyW) {
		g.Player.PlayerDest.Y -= g.Player.Speed
		playerMoving = true
		playerDirection = 1
	}
	if rl.IsKeyDown(rl.KeyS) {
		g.Player.PlayerDest.Y += g.Player.Speed
		playerMoving = true
		playerDirection = 0
	}
	if rl.IsKeyDown(rl.KeyA) {
		g.Player.PlayerDest.X -= g.Player.Speed
		playerMoving = true
		playerDirection = 2
	}
	if rl.IsKeyDown(rl.KeyD) {
		g.Player.PlayerDest.X += g.Player.Speed
		playerMoving = true
		playerDirection = 3
	}

	if rl.IsKeyDown(rl.KeyD) && rl.IsKeyDown(rl.KeyA) {
		if rl.IsKeyDown(rl.KeyW) {
			playerDirection = 1
		} else {
			playerDirection = 0
		}
	}

	g.Player.Position.X = g.Player.PlayerDest.X
	g.Player.Position.Y = g.Player.PlayerDest.Y

	if rl.IsKeyDown(rl.KeyRight) && shotsRemaining == 0 {
		x := rand.Intn(2)
		y := rand.Int()

		if y%2 == 0 {
			x = x * -1
		}
		g.Bullets = append(g.Bullets, Bullets{rl.NewVector2(g.Player.PlayerDest.X+17, g.Player.PlayerDest.Y+17), rl.NewVector2(5, float32(x)), 10, true})
		shotsRemaining = shotDelay
	}
	if rl.IsKeyDown(rl.KeyLeft) && shotsRemaining == 0 {
		x := rand.Intn(2)
		y := rand.Int()

		if y%2 == 0 {
			x = x * -1
		}
		g.Bullets = append(g.Bullets, Bullets{rl.NewVector2(g.Player.PlayerDest.X+17, g.Player.PlayerDest.Y+17), rl.NewVector2(-5, float32(x)), 10, true})
		shotsRemaining = shotDelay
	}
}
