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

func (d *day07) printGrid(grid [][]rune) {
	for _, line := range grid {
		fmt.Println(string(line))
	}
	fmt.Println("--------------------------------------------------------------------------")
}

func (d *day07) solve(input []byte) (int, int) {
	grid := d.parseInput(input)
	iMax := len(grid)
	jMax := len(grid[0])
	start := d.findStart(grid)

	queue := NewQueue[index]()
	seen := make(map[index]struct{})
	queue.Enqueue(start)

	answer := 0

	for queue.Len() > 0 {
		curr, ok := queue.Dequeue()
		if !ok {
			break
		}
		if _, exist := seen[curr]; exist {
			continue
		}
		seen[curr] = struct{}{}

		next := index{curr.i + 1, curr.j}
		if !(0 <= next.i && next.i < iMax && 0 <= next.j && next.j < jMax) {
			continue
		}

		switch grid[next.i][next.j] {
		case '.':
			{
				queue.Enqueue(next)
				grid[next.i][next.j] = '|'
			}
		case '^':
			{
				answer += 1
				for _, splitOffset := range []int{-1, 1} {
					split := index{next.i, next.j + splitOffset}
					if 0 <= split.j && split.j < jMax {
						queue.Enqueue(split)
						grid[split.i][split.j] = '|'
					}
				}
			}
		}
		// d.printGrid(grid)
	}

	return answer, 0
}

func (d *day07) SolvePartOne(input []byte) string {
	answer, _ := d.solve(input)
	return fmt.Sprintf("%v", answer)
}

func (d *day07) SolvePartTwo(input []byte) string {
	_, answer := d.solve(input)
	return fmt.Sprintf("%v", answer)
}
