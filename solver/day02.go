package solver

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type day02 struct{}

func (d *day02) parseInput(input []byte) [][]string {
	lines := bytes.Split(input, []byte(","))
	var ranges [][]string
	for _, line := range lines {
		low, high, _ := bytes.Cut(line, []byte("-"))
		ranges = append(ranges, []string{string(low), string(high)})
	}
	return ranges
}

func (d *day02) LowestEventDigitNumberGreaterThan(n string) string {
	if len(n)%2 == 0 {
		// already even length
		return n
	}
	return "1" + strings.Repeat("0", len(n))
}

func (d *day02) HighestEventDigitNumberLessThan(n string) string {
	if len(n)%2 == 0 {
		// already even length
		return n
	}
	return strings.Repeat("9", len(n)-1)
}

func (d *day02) SolvePartOne(input []byte) string {
	ranges := d.parseInput(input)
	fmt.Println(ranges)
	answer := 0

	for _, r := range ranges {
		low := d.LowestEventDigitNumberGreaterThan(r[0])
		high := d.HighestEventDigitNumberLessThan(r[1])
		lowInt, _ := strconv.Atoi(low)
		highInt, _ := strconv.Atoi(high)
		if lowInt > highInt {
			continue
		}
		fmt.Println("Range:", r, "->", low, high)
		startRangeLow := low[:len(low)/2]
		startRangeHigh := high[:len(high)/2]
		endRangeLow := low[len(low)/2:]
		endRangeHigh := high[len(high)/2:]

		startRangeLowInt, _ := strconv.Atoi(startRangeLow)
		startRangeHighInt, _ := strconv.Atoi(startRangeHigh)
		endRangeLowInt, _ := strconv.Atoi(endRangeLow)
		endRangeHighInt, _ := strconv.Atoi(endRangeHigh)

		fmt.Println("  Start range ints:", startRangeLowInt, startRangeHighInt)
		fmt.Println("  End range ints:", endRangeLowInt, endRangeHighInt)

		fmt.Println("  Up range:", startRangeLow, startRangeHigh)
		fmt.Println("  Down range:", endRangeLow, endRangeHigh)

		first := strings.Repeat(startRangeLow, 2)
		firstInt, _ := strconv.Atoi(first)
		if firstInt >= lowInt && firstInt <= highInt {
			fmt.Println(first)
			answer += firstInt
		}

		for i := startRangeLowInt + 1; i < startRangeHighInt; i++ {
			num := strings.Repeat(fmt.Sprintf("%v", i), 2)
			numInt, _ := strconv.Atoi(num)
			fmt.Println(num)
			answer += numInt
		}

		last := strings.Repeat(startRangeHigh, 2)
		lastInt, _ := strconv.Atoi(last)
		if lastInt >= lowInt && lastInt <= highInt && firstInt != lastInt {
			fmt.Println(last)
			answer += lastInt
		}

		fmt.Println()
	}

	return fmt.Sprintf("%v", answer)
}

func (d *day02) SolvePartTwo(input []byte) string {
	ranges := d.parseInput(input)
	fmt.Println(ranges)

	answer := 0

	for _, r := range ranges {
		low := r[0]
		high := r[1]
		lowInt, _ := strconv.Atoi(low)
		highInt, _ := strconv.Atoi(high)
		for i := lowInt; i <= highInt; i++ {
			numStr := fmt.Sprintf("%v", i)

			usedNumbers := make(map[string]struct{})
			for l := 1; l < len(numStr); l += 1 {
				if len(numStr)%l != 0 {
					continue
				}
				pattern := numStr[:l]
				repeated := strings.Repeat(pattern, len(numStr)/l)
				if repeated == numStr {
					if _, exists := usedNumbers[repeated]; !exists {
						usedNumbers[repeated] = struct{}{}
						fmt.Println("Found repeating pattern:", numStr, "->", pattern)
						answer += i
					}
				}
			}
		}
	}

	return fmt.Sprintf("%v", answer)
}
