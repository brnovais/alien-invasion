package game

import (
	"fmt"

	"github.com/brnovais/alien-invasion/internal/config"
)

type Game struct {
	// Configuration for this game instance.
	config config.Config

	// Non-existent world of X.
	world World
}

func (g *Game) Configure(reader config.ConfigReader) {
	// Read the whole configuration and store it.
	g.config = reader.Read()
}

// Initialize the game and all required data structures.
func (g *Game) Initialize() {
	// Initialize the world data.
	g.world.Initialize()

	// Read the world map from a file.
	g.world.Read(g.config.MapFile)
}

// Run the alien invasion simulation.
func (g *Game) Run() {
}

// Print the results of the simulation.
func (g *Game) Print() {
	// Print the world using the same format as the input file.
	fmt.Println(g.world.Print())
}
