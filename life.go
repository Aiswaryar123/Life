package main

import (
	"fmt"
	"time"
)

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
		count += b2i(grid.data[x][y+1])
		count += b2i(grid.data[x+1][y])
		count += b2i(grid.data[x+1][y+1])
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
	if y == s {
		count += b2i(grid.data[x-1][y])
		count += b2i(grid.data[x-1][y-1])
		count += b2i(grid.data[x][y-1])
		count += b2i(grid.data[x+1][y])
		count += b2i(grid.data[x+1][y-1])

		return uint(count)
	}
	count += b2i(grid.data[x-1][y-1])
	count += b2i(grid.data[x-1][y])
	count += b2i(grid.data[x-1][y+1])
	count += b2i(grid.data[x][y+1])
	count += b2i(grid.data[x+1][y+1])
	count += b2i(grid.data[x+1][y])
	count += b2i(grid.data[x+1][y-1])
	count += b2i(grid.data[x][y-1])

	return uint(count)

}
func RunGeneration(g Grid) Grid {
	newGrid := NewGrid(uint(g.size))

	for i := 0; i < g.size; i++ {
		for j := 0; j < g.size; j++ {
			cell := g.data[i][j]
			aliveNeighbours := CountAliveneighbours(uint(i), uint(j), g)

			if cell {

				if aliveNeighbours < 2 || aliveNeighbours > 3 {

					newGrid.data[i][j] = false
				} else {

					newGrid.data[i][j] = true
				}
			} else {

				if aliveNeighbours == 3 {

					newGrid.data[i][j] = true
				} else {

					newGrid.data[i][j] = false
				}
			}
		}
	}

	return newGrid
}

func main() {
	fmt.Println("Welcome to Game of Life!")
	grid := NewGrid(8, 0, 0, 1, 1, 2, 0, 2, 1, 2, 2)

	fmt.Print(DisplayGrid(grid))

	for {
		fmt.Print("\033[H\033[2J\033[3J")
		fmt.Print(DisplayGrid(grid))
		time.Sleep(1 * time.Second)
		grid = RunGeneration(grid)
	}
}
