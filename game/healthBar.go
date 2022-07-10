package game

import (
	"github.com/gen2brain/raylib-go/raylib"
	"liner/dataio"
)

type HealthBar struct {
	bar []*HealthTile
	FilledTo int
}

func (b *HealthBar) Update(playerHP int) {
	for i, item := range b.bar {
		if i <= playerHP {
			item.IsHidden = false
		} else {
			item.IsHidden =  true
		}
		if i <= 11 && i >= 9 {
			item.Color = dataio.Palette["pink"]
		} else if i <= 8 && i >= 6 {
			item.Color = dataio.Palette["red"]
		} else if i <= 5 && i >= 3 {
			item.Color = dataio.Palette["darkred"]
		} else {
			item.Color = dataio.Palette["brickred"]
		}
	}
}

func (b *HealthBar) Draw(screenHeight int32) {
	rl.DrawRectangle(10, screenHeight - 80, 320, 75, rl.Gray)
	rl.DrawText("Player Health", 20, screenHeight - 70, 20, rl.DarkGray)
	for _, item := range b.bar {
		item.Draw()
	}
}

func NewHealthBar(screenHeight int32) *HealthBar {
	var b []*HealthTile
	for i := 20; i < 360; i+=25{
		b = append(b, NewHealthTile(int32(i), screenHeight-40))
	}
	return &HealthBar{
		bar: b,
		FilledTo: 12,
	}	
}
