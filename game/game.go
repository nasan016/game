package game

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	ScreenWidth  = 900
	ScreenHeight = 800
)

type Game struct {
	Player Player

	WindowShouldClose bool
}

func NewGame() (g Game) {
	g.Init()
	return
}

func (g *Game) Init() {
	g.Player = Player{rl.NewVector2(float32(200), float32(200)), 2}
}

func (g *Game) Update() {
	if rl.WindowShouldClose() {
		g.WindowShouldClose = true
	}
}

func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Blue)
	rl.EndDrawing()
}

func (g *Game) Input() {

}
