package main

import (
	"github.com/brnovais/alien-invasion/alinv"
)

func main() {
	game := alinv.NewGame()

	game.Configure()
	game.Initialize()
	game.Run()
}
