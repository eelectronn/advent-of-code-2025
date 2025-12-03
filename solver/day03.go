package solver

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
)

type day03 struct{}

func (d *day03) parseInput(input []byte) [][]int {
	// implement parsing logic
	parsed := [][]int{}
	lines := bytes.Split(input, []byte("\n"))
	for _, line := range lines {
		row := []int{}
		for _, char := range line {
			i, _ := strconv.Atoi(string(char))
			row = append(row, i)

		}
		parsed = append(parsed, row)
	}

	return parsed
}

func (d *day03) SolvePartOne(input []byte) string {
	parsedInput := d.parseInput(input)
	fmt.Println(parsedInput)
	answer := 0
	for i := 0; i < len(parsedInput); i++ {
		max := 0
		for j := 0; j < len(parsedInput[i]); j++ {
			for k := j + 1; k < len(parsedInput[i]); k++ {
				num := 10*parsedInput[i][j] + parsedInput[i][k]
				if num > max {
					max = num
				}
			}
		}
		fmt.Println(max)
		answer += max
	}
	return fmt.Sprintf("%d", answer)
}

func (d *day03) SolvePartTwo(input []byte) string {
	parsedInput := d.parseInput(input)
	answer := 0
	for i := 0; i < len(parsedInput); i++ {
		if len(parsedInput[i]) == 0 {
			continue
		}
		max := 0
		leftMostIndex := 0
		for j := 12; j > 0; j-- {
			maxDigit := -1
			newLeftMostIndex := leftMostIndex
			for k := len(parsedInput[i]) - j; k >= leftMostIndex; k-- {
				if parsedInput[i][k] >= maxDigit {
					maxDigit = parsedInput[i][k]
					newLeftMostIndex = k + 1
				}
			}
			leftMostIndex = newLeftMostIndex
			max += maxDigit * int(math.Pow(10, float64(j-1)))
		}
		fmt.Println(max)
		answer += max
	}
	return fmt.Sprintf("%d", answer)
}
