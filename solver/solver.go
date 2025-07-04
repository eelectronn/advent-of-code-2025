package solver

import (
	"fmt"
	"os"
)

type Solver interface {
	SolvePartOne(input string) string
	SolvePartTwo(input string) string
}

var solvers = map[int]Solver{
	1: &day01{},
}

func readFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func Solve(day, part int, sample bool) string {
	solver, exists := solvers[day]
	if !exists {
		return fmt.Sprintf("No solver found for day %d, part %d", day, part)
	}
	filePath := fmt.Sprintf("input/%02d.txt", day)
	if sample {
		filePath = fmt.Sprintf("input/%02d_sample.txt", day)
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
