package game

// City is a node in our graph. Its the main object
// in the non-existent world of X.
type City struct {
	// Name of this node, used as a key to identify the city.
	name string

	// Edges of this node, connecting nodes.
	roads []*Road

	// Cached value indicating if we need to print this city.
	print bool

	// Indexed value indicating the amount of aliens in the city.
	aliens int

	// Cached value indicating if this city was destroyed.
	destroyed bool
}

func (c *City) Print(print bool) {
	c.print = print
}

func (c *City) IsDestroyed() bool {
	return c.destroyed
}

func (c *City) Destroy() {
	c.destroyed = true
}

func (c *City) DelRoad(dest *City) {
	l := len(c.roads)

	for i := 0; i < l; i++ {
		if c.roads[i].dest == dest {
			l--
			c.roads[i], c.roads[l] = c.roads[l], c.roads[i]
		}
	}

	c.roads = c.roads[:l]
}

// Road is an edge in our graph, creating connections between cities.
type Road struct {
	// Name of the road, used right now to store the direction.
	// The main utility for this field is to explore the idead
	// of multiple paths going to the same city.
	name string

	// The destination node (city) of this road.
	dest *City
}
