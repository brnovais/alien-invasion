package game

import (
	"testing"
)

// TestCityDelRoad tests if the edge is removed from the graph.
func TestCityDelRoad(t *testing.T) {
	funCity := &City{}
	boringCity := &City{}

	funCity.roads = []*Road{
		{"north", boringCity},
	}

	boringCity.roads = []*Road{
		{"south", funCity},
	}

	funCity.DelRoad(boringCity)

	if len(funCity.roads) != 0 {
		t.Error("Failed to delete the road")
	}

	if len(boringCity.roads) != 1 {
		t.Error("Failed to delete the road")
	}
}
