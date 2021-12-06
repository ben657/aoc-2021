package main

import (
	"ben657/aoc21_util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Cell struct {
	val    int
	marked bool
}

func main() {
	start := time.Now()

	lines := aoc21_util.ReadInput()

	inputStrings := strings.Split(lines[0], ",")
	var inputs []int
	for _, inputString := range inputStrings {
		input, err := strconv.Atoi(inputString)
		if err != nil {
			panic(err)
		}
		inputs = append(inputs, input)
	}

	var boards [][5][5]Cell
	for i := 2; i < len(lines); i += 6 {
		var board [5][5]Cell
		for j := 0; j < 5; j++ {
			line := lines[i+j]
			currNumStr := ""
			numIndex := 0

			for k, char := range line {
				if char == ' ' {
					if len(currNumStr) > 0 {
						num, err := strconv.Atoi(currNumStr)
						if err != nil {
							panic(err)
						}
						board[j][numIndex].val = num
						numIndex++
					}
					currNumStr = ""
				} else {
					currNumStr += string(char)
					if k == len(line)-1 {
						num, err := strconv.Atoi(currNumStr)
						if err != nil {
							panic(err)
						}
						board[j][numIndex].val = num
					}
				}
			}
		}
		boards = append(boards, board)
	}

	result := runInputs(boards, inputs)

	fmt.Println(result)
	fmt.Println("Time: ", time.Since(start))
}

func runInputs(boards [][5][5]Cell, inputs []int) int {
	for _, num := range inputs {
		for b, board := range boards {
			for r, row := range board {
				for c, cell := range row {
					if cell.val == num {
						boards[b][r][c].marked = true
						result, sum := findBingo(boards[b], r, c)
						if result {
							return sum
						}
					}
				}
			}
		}
	}

	return -1
}

func findBingo(board [5][5]Cell, row int, col int) (bool, int) {
	win := true
	for r := 0; r < 5; r++ {
		if !board[r][col].marked {
			win = false
			break
		}
	}

	if win {
		sum := 0
		for r := 0; r < 5; r++ {
			sum += board[r][col].val
		}
		return true, sum
	}

	win = true
	for c := 0; c < 5; c++ {
		if !board[row][c].marked {
			win = false
			break
		}
	}

	if win {
		sum := 0
		for c := 0; c < 5; c++ {
			sum += board[row][c].val
		}
		return true, sum
	}

	return false, -1
}
