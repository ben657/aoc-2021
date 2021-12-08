package main

import (
	"ben657/aoc21_util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func iterate(fish *[]uint64) {
	for i, f := range *fish {
		if f == 0 {
			*fish = append(*fish, 8)
			(*fish)[i] = 6
		} else {
			(*fish)[i] = f - 1
		}
	}
}

func p1() {
	lines := aoc21_util.ReadInput()

	numStrs := strings.Split(lines[0], ",")
	var fish []uint64
	for _, numStr := range numStrs {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}
		fish = append(fish, uint64(num))
	}

	for i := 0; i < 256; i++ {
		iterate(&fish)
	}

	fmt.Println("Part 1: ", len(fish))
}

func p2() {

}

func main() {
	start := time.Now()

	p1()

	fmt.Println("Time: ", time.Since(start))
}
