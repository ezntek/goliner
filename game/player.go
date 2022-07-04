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

func (player Player) Draw() {
	rl.DrawRectangleRec(player.Rect, player.Color)
}

func (player Player) controls() {
	
}

func (player Player) Update() {
	player.controls()
	if rl.IsKeyDown(rl.KeyLeft) {
		player.Velocity.X -= 0.5
	}
	if rl.IsKeyDown(rl.KeyRight) {
		player.Velocity.Y += 0.5
	}
	if rl.IsKeyReleased(rl.KeyLeft) || rl.IsKeyReleased(rl.KeyRight) {
		player.Velocity.X = 0
	}

	player.Rect.X += player.Velocity.X
	player.Rect.Y += player.Velocity.Y
}