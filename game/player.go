package game

import (
	"fmt"
	"image/color"
	"math"
	"github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Rect rl.Rectangle
	Color color.RGBA
	Velocity rl.Vector2
	health int32
	gravityForce float32
	counter int32
	stopIncrementingCounter bool
	frames int32
}

func (player *Player) Draw() {
	rl.DrawRectangleRec(player.Rect, player.Color)
}

func (player *Player) controls(frames int32, fps int32, floorHeight float32) {
	var speed float32 = 0.5
	if rl.IsKeyPressed(rl.KeyZ) {
		if player.health >= 1 {
			player.health--
		}
	}
	if rl.IsKeyPressed(rl.KeyX) {
		if player.health < 12 { 
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
		player.counter++
		fmt.Println(player.counter)
		if player.counter <= 20 {
			player.Velocity.Y -= 1.7
		}
		if player.counter >= 20 && player.Rect.Y > floorHeight - 5{
			player.gravityForce = 1.5
			player.stopIncrementingCounter = true
		}	
	}	
	if rl.IsKeyReleased(rl.KeySpace) {
		player.Velocity.Y = 3.5
		//tickTimer = 0
		player.gravityForce = 1
		player.counter = 0
		player.stopIncrementingCounter = false
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
	if player.health > 12{
		player.health = 12
	}
	
	player.gravity(player.frames, fps, 1000)
	player.controls(player.frames, fps, 1000)
	player.Rect.X += player.Velocity.X
	player.Rect.Y += player.Velocity.Y
}

func NewPlayer(rect rl.Rectangle, color color.RGBA) *Player {
	return &Player{
		Rect: rect,
		Color: color,
		Velocity: rl.Vector2{X:0,Y:0},
		health: 12,
		gravityForce: 1,
		stopIncrementingCounter: false,
		counter: 0,
	}
}
