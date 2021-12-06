package main

import (
	"ben657/aoc21_util"
	"fmt"
	"time"
)

const bits = 12

func p1() {
	lines := aoc21_util.ReadInput()

	var highBitCounts [bits]uint

	for _, line := range lines {
		for i, char := range line {
			if char == '1' {
				highBitCounts[i]++
			}
		}
	}

	var result uint

	for i := 0; i < bits; i++ {
		if highBitCounts[i] > uint(len(lines)/2) {
			result = result | (1 << uint(11-i))
		}
	}

	lResult := uint(1) ^ result

	fmt.Println("High: ", result)
	fmt.Println("Low: ", lResult)
	fmt.Println("Answer: ", result*lResult)
}

func bitStrToInt(bitStr string) uint {
	var result uint

	for i, char := range bitStr {
		if char == '1' {
			result = result | (1 << uint(bits-1-i))
		}
	}

	return result
}

func p2() {
	lines := aoc21_util.ReadInput()

	var oxygenLines []string
	oxygenLines = append(oxygenLines, lines[:]...)

	var co2Lines []string
	co2Lines = append(co2Lines, lines[:]...)

	for i := 0; i < bits; i++ {
		if len(oxygenLines) > 1 {
			highBits := 0
			lowBits := 0
			for _, line := range oxygenLines {
				if line[i] == '1' {
					highBits++
				} else {
					lowBits++
				}
			}

			oTarget := byte('1')
			if lowBits > highBits {
				oTarget = byte('0')
			}

			for j := 0; j < len(oxygenLines); j++ {
				if oxygenLines[j][i] != oTarget {
					// Replace it with the last element, slice off the last element
					oxygenLines = append(oxygenLines[:j], oxygenLines[j+1:]...)
					j--
				}
			}
		}

		if len(co2Lines) > 1 {
			highBits := 0
			lowBits := 0
			for _, line := range co2Lines {
				if line[i] == '1' {
					highBits++
				} else {
					lowBits++
				}
			}

			co2Target := byte('0')
			if highBits < lowBits {
				co2Target = byte('1')
			}

			for j := 0; j < len(co2Lines); j++ {
				if co2Lines[j][i] != co2Target {
					// Replace it with the last element, slice off the last element
					co2Lines = append(co2Lines[:j], co2Lines[j+1:]...)
					j--
				}
			}
		}
	}

	oxygen := bitStrToInt(oxygenLines[0])
	co2 := bitStrToInt(co2Lines[0])

	fmt.Println("Oxygen: ", oxygen)
	fmt.Println("CO2: ", co2)
	fmt.Println("Answer: ", oxygen*co2)
}

func main() {
	start := time.Now()

	p2()

	fmt.Println("Time: ", time.Since(start))
}
