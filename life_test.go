package main

import "testing"

func TestNewGrid(t *testing.T) {
	grid := NewGrid(3)
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
			if grid.data[i][j] {
				t.Errorf("Expected false bt got true")
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
