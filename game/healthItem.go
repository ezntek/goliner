package game

import (
	"liner/dataio"

	"github.com/gen2brain/raylib-go/raylib" 
)

type HealthTile struct {
	BorderTile rl.Rectangle
	HealthItem rl.Rectangle
	IsHidden bool
}

func (t *HealthTile) Draw() {
	rl.DrawRectangleRec(t.BorderTile, dataio.Palette["lightgray"])
	if !t.IsHidden {
		rl.DrawRectangleRec(t.HealthItem, dataio.Palette["red"])
	}
}

func NewHealthTile(posx int32, posy int32) *HealthTile {
	return &HealthTile{
		BorderTile: rl.NewRectangle(float32(posx), float32(posy), 25,25),
		HealthItem: rl.NewRectangle(float32(posx)+5, float32(posy)+5, 15,15),
		IsHidden: true,
	}
}