package game

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// Non-existent world of X.
type World struct {
	// Cities is a private map of city names and its data.
	// Its a graph like structure similar to an adjacency list.
	cities map[string]*City
}

// Initialize the world and all required data structures.
func (w *World) Initialize() {
	// Allocate memory to store the cities.
	w.cities = make(map[string]*City)
}

func (w *World) Read(filePath string) {
	// Try to open the file with data about nodes (city) and edges (roads).
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w.load(file)
}

// Read the world map from a buffer, with one city per line.
// The city name is first, followed by 1-4 directions (north, south,
// east, or west). Each one represents a road to another city that
// lies in that direction.
//
// For example:
//
//	Foo north=Bar west=Baz south=Qu-ux
//	Bar south=Foo west=Bee
func (w *World) load(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		t := scanner.Text()
		s := strings.Split(t, " ")

		// The city name is first.
		cityName := s[0]

		_, exist := w.cities[cityName]
		if !exist {
			w.cities[cityName] = &City{
				name:  cityName,
				roads: make([]*Road, 0),
				print: true,
			}
		} else {
			w.cities[cityName].print = true
		}

		// Loop through the roads in/out of the city.
		for i := 1; i < len(s); i++ {
			// The directions are separated from their respective
			// cities with an equals (=) sign.
			dir := strings.Split(s[i], "=")

			// Followed by 1-4 directions (north, south, east, or west)
			roadName := dir[0]
			destName := dir[1]

			_, exist := w.cities[destName]
			if !exist {
				dest := &City{
					name:  destName,
					roads: make([]*Road, 1),
				}

				dest.roads[0] = &Road{
					name: "",
					dest: w.cities[cityName],
				}

				w.cities[destName] = dest

				w.cities[cityName].roads = append(w.cities[cityName].roads, &Road{
					name: roadName,
					dest: dest,
				})
			} else {
				for _, r := range w.cities[cityName].roads {
					if r.dest == w.cities[destName] {
						r.name = roadName
						break
					}
				}
			}
		}
	}
}

// Print the world using the same format as the input file.
func (w *World) Print() string {
	// Use a string buffer to optmize string processing.
	var sb strings.Builder

	for k, v := range w.cities {
		// Verify if this is a city added directly to the map or it's just an edge.
		if v.print {

			sb.WriteString(k)
			for _, r := range v.roads {
				if r.name != "" {
					sb.WriteString(" ")
					sb.WriteString(r.name)
					sb.WriteString("=")
					sb.WriteString(r.dest.name)
				}
			}
			sb.WriteString("\n")
		}
	}

	return sb.String()
}

// City is a node in our graph. Its the main object
// in the non-existent world of X.
type City struct {
	// Name of this node, used as a key to identify the city.
	name string

	// Edges of this node, connecting nodes.
	roads []*Road

	// Cached value indicating if we need to print this city.
	print bool
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
