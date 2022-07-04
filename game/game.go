package game

import (
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
	var state string = "game"
	rl.InitWindow(conf.ScreenWidth, conf.ScreenHeight, "wow goliner")
	rl.SetTargetFPS(conf.FPS)

	var camera rl.Camera2D
	var player = Player{
		Rect: rl.Rectangle{X: 20, Y:20, Width:50, Height:100},
		Velocity: rl.Vector2{X:0,Y:0},
		Color: dataio.Palette["lightblue"],
	}
	
	camera.Offset = rl.Vector2{X: float32(conf.ScreenWidth)/2, Y: float32(conf.ScreenHeight)/2}
	camera.Zoom = 1.1
	camera.Rotation = 0.0

	for !rl.WindowShouldClose() {
		if state == "game" {
			player.Update()
			rl.BeginDrawing()
				camera.Target = rl.Vector2{X: player.Rect.X + player.Rect.Width/2, Y: player.Rect.Y + player.Rect.Height/2}
				rl.BeginMode2D(camera)
					rl.DrawRectangle(30,200,100,150,dataio.Palette["brickred"])
					player.Draw()
				rl.EndMode2D()
			rl.EndDrawing()
		}
	}
}