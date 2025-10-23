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
func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

func CountAliveneighbours(x uint, y uint, grid Grid) uint {
	count := 0
	s := uint(grid.size - 1)
	if x == 0 && y == 0 {
		count += b2i(grid.data[x][y+1])   // right
		count += b2i(grid.data[x+1][y])   // below
		count += b2i(grid.data[x+1][y+1]) // diagonal
		return uint(count)
	}
	if x == 0 && y == s {
		count += b2i(grid.data[x][y-1])
		count += b2i(grid.data[x+1][y-1])
		count += b2i(grid.data[x+1][y])
		return uint(count)
	}
	if x == s && y == 0 {
		count += b2i(grid.data[x-1][y])
		count += b2i(grid.data[x-1][y+1])
		count += b2i(grid.data[x][y+1])
		return uint(count)
	}
	if x == s && y == s {
		count += b2i(grid.data[x-1][y])
		count += b2i(grid.data[x-1][y-1])
		count += b2i(grid.data[x][y-1])
		return uint(count)
	}
	if x == 0 {
		count += b2i(grid.data[x][y-1])
		count += b2i(grid.data[x+1][y-1])
		count += b2i(grid.data[x+1][y])
		count += b2i(grid.data[x+1][y+1])
		count += b2i(grid.data[x][y+1])

		return uint(count)
	}
	if x == s {
		count += b2i(grid.data[x][y-1])
		count += b2i(grid.data[x-1][y-1])
		count += b2i(grid.data[x-1][y])
		count += b2i(grid.data[x-1][y+1])
		count += b2i(grid.data[x][y+1])

		return uint(count)
	}
	if y == 0 {
		count += b2i(grid.data[x-1][y])
		count += b2i(grid.data[x-1][y+1])
		count += b2i(grid.data[x][y+1])
		count += b2i(grid.data[x+1][y])
		count += b2i(grid.data[x+1][y+1])

		return uint(count)
	}

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
