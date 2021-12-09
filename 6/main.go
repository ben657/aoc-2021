package main

import (
	"ben657/aoc21_util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func p1iterate(fish *[]uint64) {
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
		p1iterate(&fish)
	}

	fmt.Println("Part 1: ", len(fish))
}

func p2iterate(fish *[8]uint) {
	oldFish := *fish
	for i := 6; i > 0; i-- {
		fish[i] = oldFish[i-1]
	}

	fish[7] = 0
	for i := uint(0); i < fish[0]; i++ {
		fish[7]++
		fish[4]++
	}
}

func p2() {
	lines := aoc21_util.ReadInput()

	numStrs := strings.Split(lines[0], ",")
	var fish [8]uint
	for _, numStr := range numStrs {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}
		fish[num]++
	}

	for i := 0; i < 256; i++ {
		p2iterate(&fish)
	}

	fmt.Println("Part 1: ", len(fish))
}

func main() {
	start := time.Now()

	p2()

	fmt.Println("Time: ", time.Since(start))
}
