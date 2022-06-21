package main

import (
	"github.com/brnovais/alien-invasion/alinv"
)

func main() {
	game := alinv.NewGame()

	// Configure the game using command-line arguments.
	game.Configure(
		alinv.NewConfigReader(alinv.CommandLineConfigReader),
	)

	// Initialize the game and all required data structures.
	game.Initialize()
	// Run the alien invasion simulation.
	game.Run()
	// Print the results of the simulation.
	game.Print()
}
