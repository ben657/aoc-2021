package aoc21_util

import (
	"fmt"
	"os"
	"strings"
)

func ReadInput() []string {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println(err)
		return nil
	}

	inputString := string(data)
	lines := strings.Split(inputString, "\r\n")

	return lines
}
