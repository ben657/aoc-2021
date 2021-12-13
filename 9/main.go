package main

import (
	"ben657/aoc21_util"
	"ben657/aoc21_util/vec"
	"fmt"
	"sort"
	"strconv"
	"time"
)

func getSurrounding(x, y int, grid [][]int) []int {
	h := len(grid)
	w := len(grid[0])

	surrounding := make([]int, 0)

	if x > 0 {
		surrounding = append(surrounding, grid[y][x-1])
	}
	if x < w-1 {
		surrounding = append(surrounding, grid[y][x+1])
	}
	if y > 0 {
		surrounding = append(surrounding, grid[y-1][x])
	}
	if y < h-1 {
		surrounding = append(surrounding, grid[y+1][x])
	}

	return surrounding
}

func p1() {
	lines := aoc21_util.ReadInput()

	h := len(lines)
	w := len(lines[0])

	grid := make([][]int, h)
	for i := 0; i < h; i++ {
		grid[i] = make([]int, w)
	}

	for y, line := range lines {
		for x, c := range line {
			height, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}
			grid[y][x] = height
		}
	}

	var lowPoints []vec.Vec
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			height := grid[y][x]
			surrounding := getSurrounding(x, y, grid)
			lowest := true
			for _, s := range surrounding {
				if height >= s {
					lowest = false
					break
				}
			}
			if lowest {
				lowPoints = append(lowPoints, vec.Vec{X: x, Y: y})
			}
		}
	}

	risk := 0
	for _, p := range lowPoints {
		risk += grid[p.Y][p.X] + 1
	}
}

type Cell struct {
	pos   vec.Vec
	h     int
	basin int
}

func p2() {
	lines := aoc21_util.ReadInput()

	h := len(lines)
	w := len(lines[0])

	grid := make([][]Cell, h)
	for i := 0; i < h; i++ {
		grid[i] = make([]Cell, w)
	}

	for y, line := range lines {
		for x, c := range line {
			height, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}
			grid[y][x] = Cell{pos: vec.Vec{X: x, Y: y}, h: height, basin: -1}
		}
	}

	currentBasin := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell.basin == -1 && cell.h < 9 {
				flood(cell.pos.X, cell.pos.Y, currentBasin, &grid)
				currentBasin++
			}
		}
	}

	basinSizes := make([]int, currentBasin)
	for _, row := range grid {
		for _, cell := range row {
			if cell.basin != -1 {
				basinSizes[cell.basin]++
			}
		}
	}
	sort.Ints(basinSizes)

	total := 1
	for _, size := range basinSizes[len(basinSizes)-3:] {
		total *= size
	}

	// for _, row := range grid {
	// 	for _, cell := range row {
	// 		color := "\033[0m"
	// 		if cell.basin != -1 {
	// 			color = "\033[32m"
	// 		}
	// 		fmt.Print(color, cell.basin, "\033[0m")
	// 	}
	// 	fmt.Print("\n")
	// }

	fmt.Println("Total: ", total)
}

func getSurroundingCells(x, y int, grid [][]Cell) []Cell {
	h := len(grid)
	w := len(grid[0])

	surrounding := make([]Cell, 0)

	if x > 0 {
		surrounding = append(surrounding, grid[y][x-1])
	}
	if x < w-1 {
		surrounding = append(surrounding, grid[y][x+1])
	}
	if y > 0 {
		surrounding = append(surrounding, grid[y-1][x])
	}
	if y < h-1 {
		surrounding = append(surrounding, grid[y+1][x])
	}

	return surrounding
}

func flood(x, y int, basin int, pGrid *[][]Cell) {
	grid := *pGrid
	h := len(grid)
	w := len(grid[0])

	if x < 0 || x >= w || y < 0 || y >= h {
		return
	}

	cell := grid[y][x]
	if cell.basin != -1 || cell.h == 9 {
		return
	}

	grid[y][x].basin = basin

	surrounding := getSurroundingCells(x, y, grid)
	for _, s := range surrounding {
		flood(s.pos.X, s.pos.Y, basin, pGrid)
	}
}

func main() {
	start := time.Now()

	p2()

	fmt.Println("Time: ", time.Since(start))
}
