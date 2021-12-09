package main

import (
	"ben657/aoc21_util"
	"fmt"
	"math"
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

func sumBelow(n int) int {
	sum := n
	for i := 1; i < n; i++ {
		sum += i
	}
	return sum
}

func p1() {
	lines := aoc21_util.ReadInput()

	numStrs := strings.Split(lines[0], ",")
	var nums []int
	min := math.MaxInt
	max := 0
	for _, numStr := range numStrs {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)

		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	minT := math.MaxInt
	for i := min; i <= max; i++ {
		t := 0
		for _, num := range nums {
			t += sumBelow(abs(i - num))
		}
		if t < minT {
			minT = t
		}
	}

	fmt.Println("Part 1: ", minT)
}

func p2() {

}

func main() {
	start := time.Now()

	p1()

	fmt.Println("Time: ", time.Since(start))
}
