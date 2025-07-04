package solver

import (
	"fmt"
	"slices"
	"strings"
)

type day01 struct{}

func (d *day01) SolvePartOne(input string) string {
	lines := strings.Split(input, "\n")
	answer := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		lineBytes := []byte(line)
		firstDigit := findFirstDigit(lineBytes)
		slices.Reverse(lineBytes)
		lastDigit := findFirstDigit(lineBytes)
		if firstDigit == -1 || lastDigit == -1 {
			continue
		}
		answer += firstDigit*10 + lastDigit

	}
	return fmt.Sprintf("%d", answer)
}

func (d *day01) SolvePartTwo(input string) string {
	return "Not implemented"
}

func findFirstDigit(line []byte) int {
	for _, char := range line {
		if char >= '0' && char <= '9' {
			return int(char - '0')
		}
	}
	return -1
}
