package main

import "fmt"

type Grid struct {
	size int
	data [][]bool
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
func DisplayGrid(g Grid) string {
	result := ""
	for _, row := range g.data {
		for _, coloum := range row {
			if coloum {
				result += "#"
			} else {
				result += "."
			}
		}
		result += "\n"
	}
	return result
}

func CountAliveneighbours(x, y uint, size uint) uint {
	return 0
}
func RunGeneration(g Grid) Grid {
	return g

}

func main() {
	grid := NewGrid(3)
	fmt.Println(grid)
	grid.data[0][1] = true
	grid.data[2][2] = true
	// for {
	// 	grid = RunGeneration(grid)
	fmt.Println(DisplayGrid(grid))
	// }
}
