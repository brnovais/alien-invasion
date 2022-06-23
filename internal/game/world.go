package game

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// World is the main graph represeting the non-existent world of X.
type World struct {
	// Cities is a private adjacency list of cities representing a graph
	// like structure, where city is the node and roads the edges.
	cities []*City

	// Indexed values to search for cities faster based on its name.
	index map[string]int
}

// Initialize the world and all required data structures.
func (w *World) Initialize() {
	// Allocate memory to store the cities.
	w.cities = make([]*City, 0)

	// Allocate memory to store the city name index.
	w.index = make(map[string]int)
}

func (w *World) Read(filePath string) {
	// Try to open the file with data about nodes (city) and edges (roads).
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Load the world map from file to memory.
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

		// If the city is not in the index, it means the city needs
		// to be added to the graph before using it.
		cityIndex, exist := w.index[cityName]
		if !exist {
			// Add a new city to the graph.
			cityIndex = w.AddCity(cityName)
		}

		// If the city was added from a file, prioritize it to be prited.
		w.GetCity(cityIndex).Print(true)

		// Loop through the roads in/out of the city.
		for i := 1; i < len(s); i++ {
			// The directions are separated from their respective
			// cities with an equals (=) sign.
			dir := strings.Split(s[i], "=")

			// Followed by 1-4 directions (north, south, east, or west)
			roadName := dir[0]
			destName := dir[1]

			// If the city is not in the index, it means the city needs
			// to be added to the graph before using it.
			destIndex, exist := w.index[destName]
			if !exist {
				// Add a new city to the graph.
				destIndex = w.AddCity(destName)

				// And create an edge between these cities.
				w.Connect(roadName, cityIndex, destIndex)
			} else {
				//w.GetCity(cityIndex).GetRoad(destIndex).SetName(roadName)
				for _, r := range w.cities[cityIndex].roads {
					if r.dest == w.cities[destIndex] {
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

	for _, v := range w.cities {
		isolated := len(v.roads) == 0

		// Verify if this is a city added directly to the map or it's just an edge.
		if !v.destroyed && (v.print || isolated) {
			sb.WriteString(v.name)

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

// AddCity is responsible to add a new node to the graph and create a index based on its name.
func (w *World) AddCity(cityName string) int {
	// Use the length as the last index.
	cityIndex := len(w.cities)

	// Append the city, making the list larger and a valid cityIndex.
	w.cities = append(w.cities, &City{
		name:  cityName,
		roads: make([]*Road, 0),
		print: false,
	})

	// Set the city in our index.
	w.index[cityName] = cityIndex

	// Return the newly created index.
	return cityIndex
}

// GetCity is responsible to get the city based on its identifier.
func (w *World) GetCity(index int) *City {
	return w.cities[index]
}

// GetCityByName is responsible to get the city based on its name. It uses an index for fast lookup.
func (w *World) GetCityByName(name string) *City {
	return w.cities[w.index[name]]
}

// Connect two cities creating an edge between nodes.
func (w *World) Connect(roadName string, city1, city2 int) {
	// Append the road to the first city.
	w.cities[city1].roads = append(w.cities[city1].roads, &Road{
		name: roadName,
		dest: w.cities[city2],
	})

	// Append the road to the second city.
	w.cities[city2].roads = append(w.cities[city2].roads, &Road{
		name: "",
		dest: w.cities[city1],
	})
}

// Size returns the amount of cities (nodes) we have on our world.
func (w *World) Size() int {
	return len(w.cities)
}
