package main

import (
	"ben657/aoc21_util"
	"ben657/aoc21_util/vec"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := aoc21_util.ReadInput()

	pos := vec.Vec{0, 0}

	for _, line := range lines {
		parts := strings.Split(line, " ")
		direction := parts[0]
		distance, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}

		switch direction {
		case "forward":
			pos.Add(vec.Vec{distance, 0})
		case "up":
			pos.Add(vec.Vec{0, distance})
		case "down":
			pos.Add(vec.Vec{0, -distance})
		}
	}

	fmt.Println(pos.ToStr())
	fmt.Println(pos.X * -pos.Y)
}
