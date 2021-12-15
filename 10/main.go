package main

import (
	"ben657/aoc21_util"
	"fmt"
	"time"
)

func isOpenToken(token rune) bool {
	return token == '(' || token == '[' || token == '{' || token == '<'
}

func isCloseToken(token rune) bool {
	return true
}

func getMatchingToken(token rune) rune {
	if token == '{' {
		return '}'
	} else if token == '(' {
		return ')'
	} else if token == '[' {
		return ']'
	} else if token == '<' {
		return '>'
	}

	return ''
}

func p1() {
	lines := aoc21_util.ReadInput()

	for _, line := range lines {
		var openTokens []rune
		for _, token := range line {
			if isOpenToken(token) {
				openTokens = append(openTokens, token)
			} else if isCloseToken(token) {
				if openTokens[:len(openTokens - 1)] ==  {

				}
			}
		}
	}
}

func p2() {

}

func main() {
	start := time.Now()

	p1()

	fmt.Println("Time: ", time.Since(start))
}
