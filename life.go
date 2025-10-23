package main

import "fmt"

type Grid struct {
	size int
	data [][]bool
}

func DisplayGrid(g Grid) string {
	return "[][]"
}
func CountAliveneighbours(x, y uint, size uint) uint {
	return 0
}
func RunGeneration(g Grid) Grid {
	return g

}

func NewGrid(size uint) Grid {
	grid := Grid{
		size: int(size),
		data: make([][]bool, size),
	}

	for i := range grid.data {
		grid.data[i] = make([]bool, size)
	}

	return grid
}

func main() {
	grid := NewGrid(3)
	fmt.Println(grid)
	// for {
	// 	grid = RunGeneration(grid)
	// 	fmt.Println(DisplayGrid(grid))
	// }
}
