package game

import (
	"testing"

	"github.com/brnovais/alien-invasion/internal/config"
)

// TestCoreUnleasheAliens tests if the aliens are correctly deployed.
func TestCoreUnleasheAliens(t *testing.T) {
	type test struct {
		alien int
		city  string
	}

	tests := []test{
		{0, "Bee"},
		{1, "Qu-ux"},
		{2, "Baz"},
		{3, "Bee"},
		{4, "Qu-ux"},
	}

	game := Game{}

	game.Configure(TestReader{})
	game.Initialize(10)

	for _, tc := range tests {
		cityName := game.aliens[tc.alien].city.name

		if cityName != tc.city {
			t.Fatalf("Failed to unleash alien %v: %v, got: %v", tc.alien, tc.city, cityName)
		}
	}
}

// TestCoreFightDestroyCity tests if the city is destroyed after a fight.
func TestCoreFightDestroyCity(t *testing.T) {
	type test struct {
		city      string
		destroyed bool
	}

	tests := []test{
		{"Foo", false},
		{"Bar", false},
		{"Baz", false},
		{"Qu-ux", true},
		{"Bee", true},
	}

	game := Game{}

	game.Configure(TestReader{})
	game.Initialize(10)

	game.fight()

	for _, tc := range tests {
		destroyed := game.world.GetCityByName(tc.city).destroyed

		if destroyed != tc.destroyed {
			t.Fatalf("Failed to destroy city %v: %v, got: %v", tc.city, tc.destroyed, destroyed)
		}
	}
}

// TestCoreFightKillAliens tests if the alien is killed after a fight.
func TestCoreFightKillAliens(t *testing.T) {
	type test struct {
		alien  int
		killed bool
	}

	tests := []test{
		{0, true},
		{1, true},
		{2, false},
		{3, true},
		{4, true},
	}

	game := Game{}

	game.Configure(TestReader{})
	game.Initialize(10)

	game.fight()

	for _, tc := range tests {
		killed := game.aliens[tc.alien].city == nil

		if killed != tc.killed {
			t.Fatalf("Failed to kill alien %v: %v, got: %v", tc.alien, tc.killed, killed)
		}
	}
}

// TestCoreMove tests if aliens can move between cities.
func TestCoreMove(t *testing.T) {
	type test struct {
		alien int
		city  string
	}

	tests := []test{
		{0, "Bar"},
		{1, "Foo"},
		{2, "Foo"},
		{3, "Bar"},
		{4, "Foo"},
	}

	game := Game{}

	game.Configure(TestReader{})
	game.Initialize(10)

	game.move()

	for _, tc := range tests {
		city := game.aliens[tc.alien].city

		if city.name != tc.city {
			t.Fatalf("Failed to move alien %v: %v, got: %v", tc.alien, tc.city, city)
		}
	}
}

// TestReader is an auxiliary struct to help with tests.
type TestReader struct {
}

// Read a simple configuration from memory.
func (r TestReader) Read() config.Config {
	return config.Config{
		Aliens:     5,
		Iterations: 1,
		MapFile:    "../../testdata/data.txt",
	}
}
