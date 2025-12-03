package solver

import (
	"fmt"
	"os"
)

type Solver interface {
	SolvePartOne(input []byte) string
	SolvePartTwo(input []byte) string
}

var solvers = map[int]Solver{
	1: &day01{},
	2: &day02{},
	3: &day03{},
}

func readFile(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Solve(day, part int, sample bool) string {
	solver, exists := solvers[day]
	if !exists {
		return fmt.Sprintf("No solver found for day %d, part %d", day, part)
	}
	filePath := fmt.Sprintf("input/%02d.txt", day)
	if sample {
		filePath = fmt.Sprintf("input/sample/%02d.txt", day)
	}
	input, err := readFile(filePath)
	if err != nil {
		return fmt.Sprintf("Error reading input for day %d: %v", day, err)
	}
	if part == 1 {
		return solver.SolvePartOne(input)
	}
	return solver.SolvePartTwo(input)

}
