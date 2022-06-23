package game

import (
	"bytes"
	"testing"
)

// TestWorldLoad tests if the world is correctly loaded.
func TestWorldLoad(t *testing.T) {
	var buffer bytes.Buffer

	buffer.WriteString("Foo north=Bar west=Baz south=Qu-ux\n")
	buffer.WriteString("Bar south=Foo west=Bee\n")

	world := World{}
	world.Initialize()
	world.load(&buffer)

	if len(world.cities) != 5 {
		t.Error("Failed to load all nodes")
	}

	if len(world.GetCityByName("Foo").roads) != 3 {
		t.Error("Failed to load all edges to Foo")
	}

	if len(world.GetCityByName("Bar").roads) != 2 {
		t.Error("Failed to load all edges to Bar")
	}
}

// TestWorldPrint tests if the world is correctly printed.
func TestWorldPrint(t *testing.T) {
	var buffer bytes.Buffer

	buffer.WriteString("Foo north=Bar west=Baz south=Qu-ux\n")
	buffer.WriteString("Bar south=Foo west=Bee\n")

	bufString := buffer.String()

	world := World{}
	world.Initialize()
	world.load(&buffer)

	if world.Print() != bufString {
		t.Error("Failed to load all nodes")
	}
}

// TestWorldAddCity tests if the city is added as a node to the graph.
func TestWorldAddCity(t *testing.T) {
	world := World{}
	world.Initialize()
	world.AddCity("Foo")

	if world.GetCityByName("Foo") == nil {
		t.Error("Failed to add city as node")
	}
}

// TestWorldConnect tests if we can add edges to the graph.
func TestWorldConnect(t *testing.T) {
	world := World{}
	world.Initialize()

	cityFoo := world.AddCity("Foo")
	cityBar := world.AddCity("Bar")

	world.Connect("south", cityFoo, cityBar)

	if len(world.GetCityByName("Foo").roads) != 1 {
		t.Error("Failed to connect node edges")
	}

	if len(world.GetCityByName("Bar").roads) != 1 {
		t.Error("Failed to connect node edges")
	}
}
