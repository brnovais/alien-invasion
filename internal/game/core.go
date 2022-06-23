package game

import (
	"fmt"
	"math/rand"

	"github.com/brnovais/alien-invasion/internal/config"
)

// Game represents the configuration and data used to simulate the alien invasion.
type Game struct {
	// Configuration for this game instance.
	config config.Config

	// Non-existent world of X.
	world World

	// Mad aliens trying to invade our world.
	aliens []*Alien

	// Cached value of aliens still alive used to optimize execution stop.
	alive int
}

// Configure the game based on the configuration reader.
func (g *Game) Configure(reader config.Reader) {
	// Read the whole configuration and store it.
	g.config = reader.Read()
}

// Initialize the game and all required data structures.
func (g *Game) Initialize(seed int64) {
	// Initialize the world data.
	g.world.Initialize()

	// Read the world map from a file.
	g.world.Read(g.config.MapFile)

	// Prepare the seed to use random numbers.
	rand.Seed(seed)

	// Unleashes the aliens in our world.
	g.unleasheAliens()
}

// Run the alien invasion simulation.
func (g *Game) Run() {
	// The program should run until all the aliens have been
	// destroyed, or each alien has moved at least 10,000 times.
	for i := 0; i < g.config.Iterations; i++ {
		g.fight()

		// Check if we have enough aliens to fight.
		if g.alive <= 1 {
			break
		}

		g.move()
	}
}

// Print the results of the simulation.
func (g *Game) Print() {
	// Print the world using the same format as the input file.
	fmt.Println(g.world.Print())
}

// You should create a program that reads in the world map,
// creates N aliens, and unleashes them.
func (g *Game) unleasheAliens() {
	worldSize := g.world.Size()

	// Cached value used to optmize execution stop.
	g.alive = g.config.Aliens
	// Allocate memory to store all aliens.
	g.aliens = make([]*Alien, g.config.Aliens)

	for i := 0; i < g.config.Aliens; i++ {
		// Find a random city to deploy this alien!
		rnd := rand.Intn(worldSize)
		city := g.world.GetCity(rnd)

		city.aliens++

		// Create a new alien in this city.
		g.aliens[i] = &Alien{
			id:   i,
			city: city,
		}
	}
}

// Arena where the aliens fight to the death.
func (g *Game) fight() {
	for _, city := range g.world.cities {
		// Aliens meet each other and explode everything fighting.
		if !city.IsDestroyed() && city.aliens > 1 {

			fmt.Print(city.name, " has been destroyed by")

			for _, alien := range g.aliens {
				if alien.city == city {
					city.aliens--

					fmt.Print(" alien ", alien.id)
					if city.aliens > 0 {
						fmt.Print(" and")
					}

					// Kill the alien!
					alien.city = nil
				}
			}

			fmt.Println("!")

			// We keep this city in the graph for now, but remove all of its roads.
			// This will make it impossible to any alien move to the city.
			for _, roads := range city.roads {
				roads.dest.DelRoad(city)
			}

			// Destroy the city.
			city.Destroy()
			// And reduce the amount o aliens alive.
			g.alive -= city.aliens
		}
	}
}

// Mad aliens moving around the world.
func (g *Game) move() int {
	count := 0

	for _, alien := range g.aliens {
		// If the alien doesn't have a city, it's dead.
		if alien.city != nil {
			l := len(alien.city.roads)

			if l != 0 {
				alien.city.aliens--

				if l == 1 {
					alien.city = alien.city.roads[0].dest
				} else {
					// If the city has many roads, choose one randomly.
					rnd := rand.Intn(l)
					alien.city = alien.city.roads[rnd].dest
				}

				alien.city.aliens++
			}

			count++
		}
	}

	return count
}
