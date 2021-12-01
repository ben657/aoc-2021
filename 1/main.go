package main

import (
	"ben657/aoc21_util"
	"fmt"
	"strconv"
)

func main() {
	lines := aoc21_util.ReadInput()
	var nums []int

	for i := 0; i < len(lines)-2; i++ {
		total := 0
		for j := 0; j < 3; j++ {
			num, err := strconv.Atoi(lines[i+j])
			if err != nil {
				break
			}
			total += num
		}
		nums = append(nums, total)
	}

	increases := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			increases++
		}
	}

	fmt.Println("Total increases: ", increases)
}
