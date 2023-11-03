package game

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawEXPBar(g Game) {
	rl.DrawRectangle(0, 0, int32(ScreenWidth), 40, rl.Black)
	rl.DrawRectangle(5, 5, int32(ScreenWidth)-10, 40-10, rl.Black)
	rl.DrawRectangle(5, 5, int32(ScreenWidth*.5)-10, 40-10, rl.Yellow)
	rl.DrawRectangle(5, 5, int32(100), 40-10, rl.Pink)
	rl.DrawRectangle(105, 5, int32(20), 40-10, rl.LightGray)
	rl.DrawText(fmt.Sprintf("LVL: %02d", g.Player.LVL), 10, 10, 20, rl.Black)
}

func DrawHPBar(g Game) {
	rl.DrawText(fmt.Sprintf("HP: %02d / 100", g.Player.HP), 5, 40, 20, rl.Black)
}

func DrawReloading(g Game) {
	if g.Player.Reloading {
		rl.DrawText("RELOADING.", ScreenWidth/2-55, ScreenHeight/2-54, 20, rl.Black)
	}
}
