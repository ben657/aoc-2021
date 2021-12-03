package main

import (
	"ben657/aoc21_util"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	lines := aoc21_util.ReadInput()

	var highBitCounts [12]uint

	for _, line := range lines {
		for i, char := range line {
			if char == '1' {
				highBitCounts[i]++
			}
		}
	}

	var result uint

	for i := 0; i < 12; i++ {
		if highBitCounts[i] > uint(len(lines)/2) {
			result = result | (1 << uint(11-i))
		}
	}

	lResult := uint(1) ^ result

	fmt.Println("High: ", result)
	fmt.Println("Low: ", lResult)
	fmt.Println("Answer: ", result*lResult)

	fmt.Println("Time: ", time.Since(start))
}
