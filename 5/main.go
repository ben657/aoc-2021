package main

import (
	"ben657/aoc21_util"
	"ben657/aoc21_util/vec"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func abs(i int) int {
	if i < 0 {
		i = -i
	}
	return i
}

func min(i1 int, i2 int) int {
	if i1 < i2 {
		return i1
	}
	return i2
}

func max(i1 int, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}

func p1() {
	lines := aoc21_util.ReadInput()

	var pairs [][2]vec.Vec
	var maxX, maxY int
	for _, line := range lines {
		parts := strings.Split(line, " ")

		p1Str := parts[0]
		p1 := parsePosStr(p1Str)
		if p1.X > maxX {
			maxX = p1.X
		}
		if p1.Y > maxY {
			maxY = p1.Y
		}

		p2Str := parts[2]
		p2 := parsePosStr(p2Str)
		if p2.X > maxX {
			maxX = p2.X
		}
		if p2.Y > maxY {
			maxY = p2.Y
		}

		if p1.X == p2.X || p1.Y == p2.Y {
			pairs = append(pairs, [2]vec.Vec{p1, p2})
		}
	}

	var grid [][]int
	for x := 0; x <= maxX+1; x++ {
		grid = append(grid, make([]int, maxY+1))
	}

	for _, pair := range pairs {
		p1 := pair[0]
		p2 := pair[1]
		if abs(p1.X-p2.X) > abs(p1.Y-p2.Y) {
			for x := min(p1.X, p2.X); x <= max(p1.X, p2.X); x++ {
				grid[p1.Y][x] += 1
			}
		} else {
			for y := min(p1.Y, p2.Y); y <= max(p1.Y, p2.Y); y++ {
				grid[y][p1.X] += 1
			}
		}
	}

	var count int
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if grid[x][y] > 1 {
				count++
			}
		}
	}

	fmt.Println("Count: ", count)
}

func p2() {
	lines := aoc21_util.ReadInput()

	var pairs [][2]vec.Vec
	var maxX, maxY int
	for _, line := range lines {
		parts := strings.Split(line, " ")

		p1Str := parts[0]
		p1 := parsePosStr(p1Str)
		if p1.X > maxX {
			maxX = p1.X
		}
		if p1.Y > maxY {
			maxY = p1.Y
		}

		p2Str := parts[2]
		p2 := parsePosStr(p2Str)
		if p2.X > maxX {
			maxX = p2.X
		}
		if p2.Y > maxY {
			maxY = p2.Y
		}

		pairs = append(pairs, [2]vec.Vec{p1, p2})
	}

	var grid [][]int
	for x := 0; x <= maxX+1; x++ {
		grid = append(grid, make([]int, maxY+1))
	}

	for _, pair := range pairs {
		p1 := pair[0]
		p2 := pair[1]

		xDir := 0
		if p1.X > p2.X {
			xDir = -1
		} else if p1.X < p2.X {
			xDir = 1
		}

		yDir := 0
		if p1.Y > p2.Y {
			yDir = -1
		} else if p1.Y < p2.Y {
			yDir = 1
		}

		p := p1

		l := max(abs(p1.X-p2.X), abs(p1.Y-p2.Y))
		for i := 0; i <= l; i++ {
			grid[p.Y][p.X] += 1
			p.X += xDir
			p.Y += yDir
		}
	}

	var count int
	for y := 0; y <= maxY; y++ {
		// fmt.Println(grid[y])
		for x := 0; x <= maxX; x++ {
			if grid[x][y] > 1 {
				count++
			}
		}
	}

	fmt.Println("Count: ", count)
}

func parsePosStr(posStr string) vec.Vec {
	parts := strings.Split(posStr, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return vec.Vec{X: x, Y: y}
}

func main() {
	start := time.Now()

	p2()

	fmt.Println("Time: ", time.Since(start))
}
