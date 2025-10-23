package main

import (
	"testing"
)

func TestNewGrid(t *testing.T) {
	grid := NewGrid(3, 0, 0, 0, 1)
	if grid.size != 3 {
		t.Errorf("Expected 3 but got %d\n", grid.size)
	}
	if len(grid.data) != 3 {
		t.Errorf("Expected 3 but got %d\n", len(grid.data))
	}
	for i := 0; i < int(grid.size); i++ {
		if len(grid.data[i]) != 3 {
			t.Errorf("Expected 3 but got %d\n", len(grid.data[i]))
		}
		for j := 0; j < int(grid.size); j++ {
			Expected := false
			if i == 0 && j == 0 || i == 0 && j == 1 {
				Expected = true
			}
			if grid.data[i][j] != Expected {
				t.Error("expected", Expected, "but got", grid.data[i][j])
			}

		}
	}
}

func TestDisplayGrid(t *testing.T) {
	grid := NewGrid(3)
	grid.data[0][1] = true
	grid.data[2][2] = true
	expected := ".#.\n...\n..#\n"
	result := DisplayGrid(grid)
	if result != expected {
		t.Error("Expected", expected, "but got", result)
	}
}

func TestNeighbourTopleftCorner(t *testing.T) {
	grid := NewGrid(4, 0, 0, 0, 1)
	actual := CountAliveneighbours(0, 0, grid)
	Expected := uint(1)
	if actual != Expected {
		t.Error("Expected", Expected, "bt got", actual)
	}
	grid = NewGrid(4, 0, 0, 0, 1, 1, 1)
	actual = CountAliveneighbours(0, 0, grid)
	Expected = uint(2)
	if actual != Expected {
		t.Error("Expected", Expected, "bt got", actual)
	}
	grid = NewGrid(4, 0, 0, 0, 1, 1, 0, 1, 1)
	actual = CountAliveneighbours(0, 0, grid)
	Expected = uint(3)
	if actual != Expected {
		t.Error("Expected", Expected, "bt got", actual)
	}
}
func TestNeighbourToprightCorner(t *testing.T) {
	grid := NewGrid(4, 0, 3, 0, 2)
	actual := CountAliveneighbours(0, 3, grid)
	Expected := uint(1)
	if actual != Expected {
		t.Error("Expected", Expected, "bt got", actual)
	}
	grid = NewGrid(4, 0, 3, 0, 2, 1, 2)
	actual = CountAliveneighbours(0, 3, grid)
	Expected = uint(2)
	if actual != Expected {
		t.Error("Expected", Expected, "bt got", actual)
	}
	grid = NewGrid(4, 0, 3, 0, 2, 1, 2, 1, 3)
	actual = CountAliveneighbours(0, 3, grid)
	Expected = uint(3)
	if actual != Expected {
		t.Error("Expected", Expected, "bt got", actual)
	}
}
func TestNeighbourBottomleftCorner(t *testing.T) {
	grid := NewGrid(4, 3, 0, 2, 0)
	actual := CountAliveneighbours(3, 0, grid)
	Expected := uint(1)
	if actual != Expected {
		t.Error("Expected", Expected, "bt got", actual)
	}
	grid = NewGrid(4, 3, 0, 2, 0, 2, 1)
	actual = CountAliveneighbours(3, 0, grid)
	Expected = uint(2)
	if actual != Expected {
		t.Error("Expected", Expected, "bt got", actual)
	}
	grid = NewGrid(4, 3, 0, 2, 0, 2, 1, 3, 1)
	actual = CountAliveneighbours(3, 0, grid)
	Expected = uint(3)
	if actual != Expected {
		t.Error("Expected", Expected, "bt got", actual)
	}
}
func TestNeighbourBottomrightCorner(t *testing.T) {
	grid := NewGrid(4, 3, 3, 3, 2)
	actual := CountAliveneighbours(3, 3, grid)
	Expected := uint(1)
	if actual != Expected {
		t.Error("Expected", Expected, "bt got", actual)
	}
	grid = NewGrid(4, 3, 3, 3, 2, 2, 2)
	actual = CountAliveneighbours(3, 3, grid)
	Expected = uint(2)
	if actual != Expected {
		t.Error("Expected", Expected, "bt got", actual)
	}
	grid = NewGrid(4, 3, 3, 3, 2, 2, 2, 2, 3)
	actual = CountAliveneighbours(3, 3, grid)
	Expected = uint(3)
	if actual != Expected {
		t.Error("Expected", Expected, "bt got", actual)
	}
}
func TestNeighbourTopedge(t *testing.T) {
	grid := NewGrid(4, 0, 0)
	actual := CountAliveneighbours(0, 1, grid)
	Expected := uint(1)
	if actual != Expected {
		t.Error("Expected", Expected, "bt got", actual)
	}
	grid = NewGrid(4, 0, 0, 1, 0)
	actual = CountAliveneighbours(0, 1, grid)
	Expected = uint(2)
	if actual != Expected {
		t.Error("Expected", Expected, "bt got", actual)
	}
	grid = NewGrid(4, 0, 0, 1, 0, 1, 1)
	actual = CountAliveneighbours(0, 1, grid)
	Expected = uint(3)
	if actual != Expected {
		t.Error("Expected", Expected, "bt got", actual)
	}
	grid = NewGrid(4, 0, 0, 1, 0, 1, 1, 1, 2)
	actual = CountAliveneighbours(0, 1, grid)
	Expected = uint(4)
	if actual != Expected {
		t.Error("Expected", Expected, "bt got", actual)
	}
	grid = NewGrid(4, 0, 0, 1, 0, 1, 1, 1, 2, 0, 2, 0, 1)
	actual = CountAliveneighbours(0, 1, grid)
	Expected = uint(5)
	if actual != Expected {
		t.Error("Expected", Expected, "bt got", actual)
	}
}
func TestNeighbourBottomedge(t *testing.T) {
	grid := NewGrid(4, 3, 0)
	actual := CountAliveneighbours(3, 1, grid)
	Expected := uint(1)
	if actual != Expected {
		t.Error("Expected", Expected, "bt got", actual)
	}
	grid = NewGrid(4, 3, 0, 2, 0)
	actual = CountAliveneighbours(3, 1, grid)
	Expected = uint(2)
	if actual != Expected {
		t.Error("Expected", Expected, "bt got", actual)
	}
	grid = NewGrid(4, 3, 0, 2, 0, 2, 1)
	actual = CountAliveneighbours(3, 1, grid)
	Expected = uint(3)
	if actual != Expected {
		t.Error("Expected", Expected, "bt got", actual)
	}
	grid = NewGrid(4, 3, 0, 2, 0, 2, 1, 2, 2)
	actual = CountAliveneighbours(3, 1, grid)
	Expected = uint(4)
	if actual != Expected {
		t.Error("Expected", Expected, "bt got", actual)
	}
	grid = NewGrid(4, 3, 0, 2, 0, 2, 1, 2, 2, 3, 2)
	actual = CountAliveneighbours(3, 1, grid)
	Expected = uint(5)
	if actual != Expected {
		t.Error("Expected", Expected, "bt got", actual)
	}
}
