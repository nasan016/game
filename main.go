package main

import (
	"game-engine/game"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(game.ScreenWidth, game.ScreenHeight, "PrimeFire")

	rl.SetTargetFPS(60)

	game := game.NewGame()

	for !rl.WindowShouldClose() {
		game.Update()
		game.Draw()
		game.Input()
	}

	rl.CloseWindow()

}
