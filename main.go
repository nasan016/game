package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	screenWidth = 1000
	screenHeight = 480
)

var (
	running = true
	bkgColor = rl.NewColor(147, 211, 196, 255)
	posx int32 = 100
	posy int32 = 100
)

func drawScene() {
	rl.DrawRectangle(posx, posy, 60, 60, rl.White)
}

func input() {
	if rl.IsKeyDown(rl.KeyW) { posy -= 1 }
	if rl.IsKeyDown(rl.KeyS) { posy += 1 }
	if rl.IsKeyDown(rl.KeyA) { posx -= 1 }
	if rl.IsKeyDown(rl.KeyD) { posx += 1 }
}

func update() {
	running = !rl.WindowShouldClose()
}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(bkgColor)
	drawScene()

	rl.EndDrawing()
}

func init() {
	rl.InitWindow(screenWidth, screenHeight, "prototype")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)
}

func quit() {
	rl.CloseWindow()
}

func main() {

	for running {
		input()
		update()
		render()
	}

	quit()
}