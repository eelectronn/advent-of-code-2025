package solver

import (
	"fmt"
	"strconv"
	"strings"
)

type day12 struct{}

type testCase struct {
	w, l          int
	numberOfGifts []int
}

func (d *day12) parseInput(input []byte) ([]int, []testCase) {
	inputS := string(input)
	lines := strings.Split(inputS, "\n\n")
	giftShapesRaw := lines[:6]
	giftNumberOfBlocks := make([]int, 6)
	for i := range giftShapesRaw {
		count := 0
		for _, c := range giftShapesRaw[i] {
			if c == '#' {
				count += 1
			}
		}
		giftNumberOfBlocks[i] = count
	}

	testCases := make([]testCase, 0)
	lines = strings.Split(lines[6], "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		dim, numsRaw, _ := strings.Cut(line, ": ")
		dimWRaw, dimLRaw, _ := strings.Cut(dim, "x")
		dimW, _ := strconv.Atoi(dimWRaw)
		dimL, _ := strconv.Atoi(dimLRaw)
		nums := make([]int, 0, 6)
		for _, numRaw := range strings.Split(numsRaw, " ") {
			num, _ := strconv.Atoi(numRaw)
			nums = append(nums, num)
		}
		testCases = append(testCases, testCase{
			w:             dimW,
			l:             dimL,
			numberOfGifts: nums,
		})
	}
	return giftNumberOfBlocks, testCases
}

func (d *day12) SolvePartOne(input []byte) string {
	_, testCases := d.parseInput(input)
	answer := 0
	for _, test := range testCases {
		wNormal := test.w - (test.w % 3)
		lNormal := test.l - (test.l % 3)

		numberOfGifts := 0
		for _, n := range test.numberOfGifts {
			numberOfGifts += n
		}
		if wNormal*lNormal >= 9*numberOfGifts {
			answer += 1
		}
	}
	return fmt.Sprintf("%v", answer)
}

func (d *day12) SolvePartTwo(input []byte) string {
	return "answer"
}
