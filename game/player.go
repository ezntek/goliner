package game

import (
	"image/color"
	"math"
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

func (player *Player) controls(frames int32, fps int32, speed float32) {
	if rl.IsKeyDown(rl.KeyLeft) {
		player.Velocity.X -= speed
	} else if !rl.IsKeyDown(rl.KeyLeft) && player.Velocity.X < 0 {
		player.Velocity.X += speed * 3
		if math.Abs(float64(player.Velocity.X)) == 0.5 {
			player.Velocity.X = 0
		}
	}
	if rl.IsKeyDown(rl.KeyRight) {
		player.Velocity.X += speed
	} else if !rl.IsKeyDown(rl.KeyRight) && player.Velocity.X > 0 {
		player.Velocity.X -= speed * 3
		if math.Abs(float64(player.Velocity.X)) <= 1 {
			player.Velocity.X = 0
		}
	}
}

func (player  *Player) gravity(frames int32, fps int32, floorHeight float32) {
    if player.Rect.Y < floorHeight {
		player.Velocity.Y++
	} else if player.Rect.Y > floorHeight {
		player.Rect.Y = floorHeight
	}
	if player.Velocity.Y > 0 && player.Rect.Y == floorHeight {
		player.Velocity.Y = 0
	}
}
func (player *Player) Update(fps int32) {
	var ticks, frames int32
	frames++
	if ticks % (fps/10) == 0 {
		ticks++
	}
	player.gravity(frames, fps, 1000)
	player.controls(frames, fps, 0.5)
	player.Rect.X += player.Velocity.X
	player.Rect.Y += player.Velocity.Y
}