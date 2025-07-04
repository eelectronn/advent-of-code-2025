package solver

import (
	"bytes"
	"fmt"
	"slices"
)

type day01 struct{}

func (d *day01) SolvePartOne(input []byte) string {
	lines := bytes.Split(input, []byte("\n"))
	answer := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		firstDigit := findFirstDigit(line)
		slices.Reverse(line)
		lastDigit := findFirstDigit(line)
		if firstDigit == -1 || lastDigit == -1 {
			continue
		}
		answer += firstDigit*10 + lastDigit

	}
	return fmt.Sprintf("%d", answer)
}

func (d *day01) SolvePartTwo(input []byte) string {
	replaceMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	reversedReplaceMap := map[string]string{
		"eno":   "1",
		"owt":   "2",
		"eerht": "3",
		"ruof":  "4",
		"evif":  "5",
		"xis":   "6",
		"neves": "7",
		"thgie": "8",
		"enin":  "9",
	}
	answer := 0
	lines := bytes.Split(input, []byte("\n"))
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		firstDigit := findFirstDigitWithWords(line, replaceMap)
		slices.Reverse(line)
		lastDigit := findFirstDigitWithWords(line, reversedReplaceMap)
		if firstDigit == -1 || lastDigit == -1 {
			continue
		}
		answer += firstDigit*10 + lastDigit
	}
	return fmt.Sprintf("%d", answer)
}

func findFirstDigitWithWords(line []byte, replaceMap map[string]string) int {
	updatedLine := make([]byte, 0, len(line))
outside:
	for i := 0; i < len(line); i++ {
		for word, digit := range replaceMap {
			if bytes.HasPrefix(line[i:], []byte(word)) {
				updatedLine = append(updatedLine, []byte(digit)...)
				i += len(word)
				continue outside
			}
		}
		updatedLine = append(updatedLine, line[i])
	}
	return findFirstDigit(updatedLine)
}

func findFirstDigit(line []byte) int {
	for _, char := range line {
		if char >= '0' && char <= '9' {
			return int(char - '0')
		}
	}
	return -1
}
