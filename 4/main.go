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

func contains(slice []int, val int) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

func runInputs(boards [][5][5]Cell, inputs []int) int {
	var completedBoards []int
	for _, num := range inputs {
		for b, board := range boards {
			for r, row := range board {
				for c, cell := range row {
					if cell.val == num {
						boards[b][r][c].marked = true
						result, sum := findBingo(boards[b], r, c)
						if result {
							if !contains(completedBoards, b) {
								completedBoards = append(completedBoards, b)
							}
							if len(completedBoards) == len(boards) {
								return sum * num
							}
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
		return true, sumUnmarked(board)
	}

	win = true
	for c := 0; c < 5; c++ {
		if !board[row][c].marked {
			win = false
			break
		}
	}

	if win {
		return true, sumUnmarked(board)
	}

	return false, -1
}

func sumUnmarked(board [5][5]Cell) int {
	sum := 0
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			if !board[r][c].marked {
				sum += board[r][c].val
			}
		}
	}
	return sum
}
