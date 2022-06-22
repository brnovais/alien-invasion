package game

import (
	"bytes"
	"testing"
)

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
