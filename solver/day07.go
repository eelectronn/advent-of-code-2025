package solver

import (
	"bytes"
	"fmt"
)

type index struct {
	i, j int
}

type day07 struct{}

func (d *day07) parseInput(input []byte) [][]rune {
	lines := bytes.Split(input, []byte("\n"))
	lineStr := make([][]rune, 0, len(lines))
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		lineStr = append(lineStr, []rune(string(line)))
	}
	return lineStr
}

func (d *day07) findStart(grid [][]rune) index {
	for i, line := range grid {
		for j, c := range line {
			if c == 'S' {
				return index{i, j}
			}
		}
	}
	return index{-1, -1}
}

func (d *day07) solve(input []byte) (int, int) {
	grid := d.parseInput(input)
	iMax := len(grid)
	jMax := len(grid[0])
	curr := d.findStart(grid)

	numberOfPaths := make([][]int, len(grid))
	for i := range numberOfPaths {
		numberOfPaths[i] = make([]int, len(grid[0]))
	}
	numberOfPaths[curr.i][curr.j] = 1
	grid[curr.i][curr.j] = '|'

	answerPartOne := 0
	for i := curr.i; i < iMax-1; i++ {
		nextI := i + 1
		for j := 0; j < jMax; j++ {
			paths := numberOfPaths[i][j]
			if grid[i][j] != '|' {
				continue
			}
			switch grid[nextI][j] {
			case '^':
				answerPartOne += 1
				for _, splitOffset := range []int{-1, 1} {
					split := index{nextI, j + splitOffset}
					if 0 <= split.j && split.j < jMax {
						numberOfPaths[split.i][split.j] += paths
						grid[split.i][split.j] = '|'
					}
				}
			default:
				grid[nextI][j] = '|'
				numberOfPaths[nextI][j] += paths
			}
		}
	}

	answerPartTwo := 0
	for _, n := range numberOfPaths[len(numberOfPaths)-1] {
		answerPartTwo += n
	}

	return answerPartOne, answerPartTwo
}

func (d *day07) SolvePartOne(input []byte) string {
	answer, _ := d.solve(input)
	return fmt.Sprintf("%v", answer)
}

func (d *day07) SolvePartTwo(input []byte) string {
	_, answer := d.solve(input)
	return fmt.Sprintf("%v", answer)
}
