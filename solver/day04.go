package solver

import (
	"bytes"
	"fmt"
)

type day04 struct{}

type coord struct {
	i, j int
}

func (d *day04) parseInput(input []byte) [][]byte {
	parsed := bytes.Split(input, []byte("\n"))
	return parsed[:len(parsed)-1]
}

func (d *day04) findRemovableCoords(grid [][]byte) []coord {
	removable := []coord{}
	for i, line := range grid {
		for j, c := range line {
			if c != byte('@') {
				continue
			}
			numberOfNeighbors := 0
			for di := -1; di <= 1; di++ {

				for dj := -1; dj <= 1; dj++ {
					if di == 0 && dj == 0 {
						continue
					}
					ni, nj := i+di, j+dj
					if ni >= 0 && ni < len(grid) && nj >= 0 && nj < len(line) && grid[ni][nj] == byte('@') {
						numberOfNeighbors++
					}
				}
			}
			if numberOfNeighbors < 4 {
				removable = append(removable, coord{i, j})
			}
		}
	}
	return removable
}

func (d *day04) SolvePartOne(input []byte) string {
	return fmt.Sprintf("%d", len(d.findRemovableCoords(d.parseInput(input))))
}

func (d *day04) SolvePartTwo(input []byte) string {
	grid := d.parseInput(input)
	answer := 0
	for {
		removable := d.findRemovableCoords(grid)
		if len(removable) == 0 {
			break
		}
		answer += len(removable)
		for _, c := range removable {
			grid[c.i][c.j] = byte('.')
		}
	}
	return fmt.Sprintf("%d", answer)
}
