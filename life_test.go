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
	}
}
