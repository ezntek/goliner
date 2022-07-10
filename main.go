package main

import (
	"fmt"
	"liner/game"
)

func main() {
	conf := game.Config{
		Build: 8,
		FPS: 60,
		ScreenWidth: 1152,
		ScreenHeight: 648,
		Debug: true,
		Version: "alpha 0.2.1",
	}
	fmt.Printf("build: %d\n", conf.Build)
	game.Game(conf)
}
