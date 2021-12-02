package main

import (
	"ben657/aoc21_util"
	"ben657/aoc21_util/vec"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	lines := aoc21_util.ReadInput()

	pos := vec.Vec{X: 0, Y: 0}

	for _, line := range lines {
		parts := strings.Split(line, " ")
		direction := parts[0]
		distance, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}

		switch direction {
		case "forward":
			pos.Add(vec.Vec{X: distance, Y: 0})
		case "up":
			pos.Add(vec.Vec{X: 0, Y: distance})
		case "down":
			pos.Add(vec.Vec{X: 0, Y: -distance})
		}
	}

	fmt.Println(pos.ToStr())
	fmt.Println(pos.X * -pos.Y)

	fmt.Println("Took: ", time.Since(start))
}
