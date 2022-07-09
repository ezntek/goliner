package game

import (
	"image/color"
	"math"
	"fmt"
	"github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Rect rl.Rectangle
	Color color.RGBA
	Velocity rl.Vector2
	health int32
}

func (player *Player) Draw() {
	rl.DrawRectangleRec(player.Rect, player.Color)
}

func (player *Player) controls(frames int32, fps int32, floorHeight float32) {
	var tickTimer int32
	var stopIncrementingTickTimer bool = false
	var speed float32 = 0.5
	if rl.IsKeyPressed(rl.KeyZ) {
		if player.health >= 1 {
			player.health--
		}
	}
	if rl.IsKeyPressed(rl.KeyX) {
		if player.health < 10 { 
			player.health++
		}
	}
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

	if rl.IsKeyDown(rl.KeySpace) {
		if frames % 6 == 0 && !stopIncrementingTickTimer { 
			tickTimer++
			fmt.Println("add")
		}
		if tickTimer <= 4 && !stopIncrementingTickTimer {
			player.Velocity.Y -= 1.7
		}
		if tickTimer > 5 && player.Rect.Y > floorHeight - 5{
			fmt.Println("stop")
			player.Velocity.Y = 3
		}
	}
	if rl.IsKeyReleased(rl.KeySpace) {
		player.Velocity.Y = 0
		tickTimer = 0
		stopIncrementingTickTimer = false
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
	if player.health > 10{
		player.health = 10
	}
	var ticks, frames int32
	frames++
	if ticks % 6 == 0 {
		ticks++
	}
	player.gravity(frames, fps, 1000)
	player.controls(frames, fps, 1000)
	player.Rect.X += player.Velocity.X
	player.Rect.Y += player.Velocity.Y
}

func NewPlayer(rect rl.Rectangle, color color.RGBA) *Player {
	return &Player{
		Rect: rect,
		Color: color,
		Velocity: rl.Vector2{X:0,Y:0},
		health: 5,
	}
}