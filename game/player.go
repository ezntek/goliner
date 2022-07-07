package game

import (
	"image/color"
	"github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Rect rl.Rectangle
	Color color.RGBA
	Velocity rl.Vector2
}

func (player *Player) Draw() {
	rl.DrawRectangleRec(player.Rect, player.Color)
}

func (player *Player) controls(frames int32, fps int32) {
	if rl.IsKeyDown(rl.KeyLeft) {
		player.Velocity.X -= 0.5
	}
	if rl.IsKeyDown(rl.KeyRight) {
		player.Velocity.X += 0.5
	} 
	if rl.IsKeyReleased(rl.KeyLeft) || rl.IsKeyReleased(rl.KeyRight) {
		var count int
		if player.Velocity.X > 0.3 {
			if frames % (fps/10) == 0 && count < 6 {
				count++
				if count < 5 {
					player.Velocity.X -= player.Velocity.X/2
				} else {
					count = 0
				}
			}
		}
	}
}

func (player *Player) Update(fps int32) {
	var ticks, frames int32
	frames++
	if ticks % (fps/10) == 0 {
		ticks++
	}
	player.controls(frames, fps)
	player.Rect.X += player.Velocity.X
	player.Rect.Y += player.Velocity.Y
}