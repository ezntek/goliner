package main

import (
	"fmt"
	"liner/game"
)

func main() {
	conf := game.Config{
		Build: 7,
		FPS: 60,
		ScreenWidth: 1024,
		ScreenHeight: 576,
		Debug: true,
		Version: "alpha 0.2",
	}
	fmt.Printf("build: %d\n", conf.Build)
	game.Game(conf)
}