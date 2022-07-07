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

func DrawCheckerboard(size int) {
	var isGray bool
	for y := -100; y < 1100; y += size {
		for x := -300; x < 3000; x += size {
			if isGray {
				rl.DrawRectangle(int32(x),int32(y),int32(size), int32(size),dataio.Palette["verylightgray"])
			} else {
				rl.DrawRectangle(int32(x),int32(y),int32(size), int32(size),dataio.Palette["white"])
			}
			isGray = !isGray
		} 
		isGray = !isGray
	}
}

func Game(conf Config) {
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
		player.Update(conf.FPS)
		rl.BeginDrawing()
			rl.ClearBackground(rl.RayWhite)
			camera.Target = rl.Vector2{X: player.Rect.X + player.Rect.Width/2, Y: player.Rect.Y + player.Rect.Height/2}
			rl.BeginMode2D(camera)
				DrawCheckerboard(15)
				player.Draw()
			rl.EndMode2D()
			rl.DrawText(fmt.Sprintf("version %s, build %d", conf.Version, conf.Build), 20, conf.ScreenHeight - 30, 20, rl.Gold)
			rl.DrawText(fmt.Sprintf("X: %0.1f Y: %0.1f", player.Rect.X, player.Rect.Y), 20, 20, 15, rl.Gray)
			rl.DrawText(fmt.Sprintf("X Velocity: %0.2f Y Velocity: %0.2f", player.Velocity.X, player.Velocity.Y), 20, 35, 15, rl.Gray)
		rl.EndDrawing()
	}
}