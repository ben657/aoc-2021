package main

import (
	"ben657/aoc21_util"
	"fmt"
	"strings"
	"time"
)

var patterns = [10]string{"abcefg", "cf", "acdeg", "acdfg", "bcdf", "abdfg", "abdefg", "acf", "abcdefg", "abcdfg"}

type entry struct {
	signals []string
	digits  []string
	decoded int
}

func findDigit(oPatterns [10]int, occurrences int) int {
	for i, pattern := range oPatterns {
		if pattern == occurrences {
			return i
		}
	}

	return -1
}

func pow(a, b int) int {
	if b == 0 {
		return 1
	}

	return a * pow(a, b-1)
}

func p1() {
	occurrences := make(map[rune]int)
	for _, s := range patterns {
		for _, r := range s {
			occurrences[r]++
		}
	}

	var oPatterns [10]int
	for i, pattern := range patterns {
		sum := 0
		for _, r := range pattern {
			sum += occurrences[r]
		}
		oPatterns[i] = sum
	}

	fmt.Println(oPatterns)

	lines := aoc21_util.ReadInput()
	fmt.Println(int('a'))

	// var digitCounts [10]int
	var entries []*entry
	for _, line := range lines {
		parts := strings.Split(line, " ")
		e := entry{parts[:10], parts[11:], 0}
		entries = append(entries, &e)

		occurrences := make(map[rune]int)
		for _, s := range e.signals {
			for _, r := range s {
				occurrences[r]++
			}
		}

		var result int
		for i, s := range e.digits {
			sum := 0
			for _, r := range s {
				sum += occurrences[r]
			}
			result += findDigit(oPatterns, sum) * pow(10, 3-i)
		}

		e.decoded = result
	}

	total := 0
	for _, e := range entries {
		total += (*e).decoded
	}

	fmt.Println("Total: ", total)
}

func p2() {

}

func main() {
	start := time.Now()

	p1()

	fmt.Println("Time: ", time.Since(start))
}
