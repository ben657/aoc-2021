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

func p2iterate(fish *[9]uint) {
	oldFish := *fish
	for i := len(fish) - 2; i > 0; i-- {
		fish[i] = oldFish[i+1]
	}

	fish[len(fish)-1] = fish[0]
	fish[len(fish)-3] += fish[0]
	fish[0] = oldFish[1]
}

func p2() {
	lines := aoc21_util.ReadInput()

	numStrs := strings.Split(lines[0], ",")
	var fish [9]uint
	for _, numStr := range numStrs {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}
		fish[num]++
	}

	for i := 0; i < 256; i++ {
		fmt.Println(i)
		p2iterate(&fish)
	}

	count := uint(0)
	for _, f := range fish {
		fmt.Println(f)
		count += f
	}
	fmt.Println("Part 2: ", count)
}

func main() {
	start := time.Now()

	p2()

	fmt.Println("Time: ", time.Since(start))
}
