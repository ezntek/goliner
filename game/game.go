package game

import (
	"fmt"
	"liner/dataio"

	"github.com/gen2brain/raylib-go/raylib"
)

type Config struct {
	Version string
	Build int32
	FPS int32
	ScreenWidth int32
	ScreenHeight int32
	Debug bool
}

func Game(conf Config) {
	var state string = "title"
	rl.InitWindow(conf.ScreenWidth, conf.ScreenHeight, "wow goliner")
	rl.SetTargetFPS(conf.FPS)

	for !rl.WindowShouldClose() {
		if state == "title" {
			if rl.IsKeyPressed(rl.KeyEnter) { state = "game" }
			rl.BeginDrawing()
				rl.ClearBackground(rl.RayWhite)
				rl.DrawText("goLiner alpha 0.1", 20, 20, 60, dataio.Palette["navy"])
				if conf.Debug {
					rl.DrawText(fmt.Sprintf("build %d", conf.Build), 20, conf.ScreenHeight - 60, 30, dataio.Palette["navy"])
				}
			rl.EndDrawing()
		} else if state == "game" { 
			rl.BeginDrawing()
				rl.ClearBackground(rl.RayWhite)
				rl.DrawText("it works!", 20,20,20,rl.Red)
			rl.EndDrawing()
		}
	}
}