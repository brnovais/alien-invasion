package game

// Alien is the main invasor in our game. The alien is mad,
// and keeps moving around between cities, terrifying people
// and fighting its comrades.
type Alien struct {
	// id is a simple identifier of the alien.
	// It could be a string name, but it's not required.
	id int

	// City represents the city the alien is currently in.
	// If the alien doesn't have a city, it's dead.
	city *City
}
