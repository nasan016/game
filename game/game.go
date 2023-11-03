package game

import (
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
	playerSprite      rl.Texture2D
	slimeSprite       rl.Texture2D

	shotDelay      int
	shotsRemaining int
	background     Background
)

type Game struct {
	Player  Player
	Slimes  []Slime
	Bullets []Bullet

	PlayArea1 rl.Vector2
	PlayArea2 rl.Vector2
	PlayArea3 rl.Vector2
	PlayArea4 rl.Vector2

	WindowShouldClose bool
}

func NewGame() (g Game) {
	g.Init()
	return
}

func (g *Game) Init() {
	backgroundSprite := rl.LoadTexture("res/background.png")
	screenSize := rl.NewVector2(float32(ScreenWidth), float32(ScreenHeight))
	background = InitBackground(backgroundSprite, screenSize, rl.NewVector2(34, 34))

	playerSprite = rl.LoadTexture("res/player-sprites.png")
	slimeSprite = rl.LoadTexture("res/slime.png")

	g.Player = InitPlayer(playerSprite)

	g.CreatePlayArea()

	cam = rl.NewCamera2D(rl.NewVector2(float32(ScreenWidth/2)-34, float32(ScreenHeight/2)-34), rl.NewVector2(float32(g.Player.PlayerDest.X-(g.Player.PlayerDest.Width/2)), float32(g.Player.PlayerDest.Y-g.Player.PlayerDest.Height/2)), 0, 2.0)

	shotDelay = 20
	shotsRemaining = 0
}

func (g *Game) Unload() {
	rl.UnloadTexture(playerSprite)
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

	g.Player.UpdateFrame()
	g.Player.ResetInvincibility()
	g.Player.IsReloading()
	//debug
	if frameCount%60 == 0 {

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
	background.Draw(g.Player.Position)
	//draw all the slimes in the list
	for _, slime := range g.Slimes {
		slime.DrawSlime(g.Player)
	}

	g.Player.DrawPlayer()
	g.DrawBullets()

	rl.EndMode2D()

	//UI Stuff
	DrawHPBar(*g)
	DrawEXPBar(*g)
	DrawReloading(*g)
	rl.EndDrawing()
}

func (g *Game) Input() {

	g.Player.Move()
	g.BulletInput()
}
