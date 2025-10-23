package main

import "fmt"

type Grid struct {
	size int
	data [][]bool
}

func NewGrid(size uint, live ...int) Grid {
	grid := Grid{
		size: int(size),
		data: make([][]bool, size),
	}

	for i := range grid.data {
		grid.data[i] = make([]bool, size)
	}
	for i := 0; i < len(live); i = i + 2 {
		x := live[i]
		y := live[i+1]
		grid.data[x][y] = true
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
