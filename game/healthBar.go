package game

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
	}
}

func (b *HealthBar) Draw() {
	for _, item := range b.bar {
		item.Draw()
	}
}

func NewHealthBar(screenHeight int32) *HealthBar {
	var b []*HealthTile
	for i := 10; i < 300; i+=30{
		b = append(b, NewHealthTile(int32(i), screenHeight-40))
	}
	return &HealthBar{
		bar: b,
		FilledTo: 10,
	}	
}