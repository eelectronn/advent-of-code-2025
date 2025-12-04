package main

import (
	"fmt"
	"main/solver"
)

func main() {
	day := 4

	sample := false
	// sample := true
	// part := 1
	part := 2

	fmt.Println(solver.Solve(day, part, sample))
}
