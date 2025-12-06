package solver

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type day06 struct{}

func (d *day06) parseInputPartOne(input []byte) ([][]int, []string) {
	lines := bytes.Split(input, []byte("\n"))
	if len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}
	operations := lines[len(lines)-1]
	lines = lines[:len(lines)-1]

	nums := make([][]int, 0, len(lines))
	for _, line := range lines {
		fields := strings.Fields(string(line))
		row := make([]int, 0, len(fields))
		for _, val := range fields {
			valInt, _ := strconv.Atoi(val)
			row = append(row, valInt)
		}
		nums = append(nums, row)
	}

	return nums, strings.Fields(string(operations))
}

func (d *day06) SolvePartOne(input []byte) string {
	nums, ops := d.parseInputPartOne(input)

	answer := 0
	for i := range nums[0] {
		result := 0
		if ops[i] == "*" {
			result = 1
		}
		for j := range nums {
			if ops[i] == "+" {
				result += nums[j][i]
			} else {
				result *= nums[j][i]
			}
		}
		answer += result
	}

	return fmt.Sprintf("%v", answer)
}

func (d *day06) SolvePartTwo(input []byte) string {
	lines := bytes.Split(input, []byte("\n"))
	if len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}
	operations := strings.Fields(string(lines[len(lines)-1]))
	// fmt.Println(operations)
	lines = lines[:len(lines)-1]

	verticalStr := make([]string, 0, len(lines[0]))
	for i := range lines[0] {
		s := make([]byte, 0, len(lines))
		for j := range lines {
			s = append(s, lines[j][i])
		}
		verticalStr = append(verticalStr, string(s))
	}

	allQuestions := make([][]int, 0)
	question := make([]int, 0)
	for _, val := range verticalStr {
		valTrim := strings.TrimSpace(val)
		if len(valTrim) == 0 {
			allQuestions = append(allQuestions, question)
			question = make([]int, 0)
		} else {
			valInt, _ := strconv.Atoi(valTrim)
			question = append(question, valInt)
		}
	}
	if len(question) > 0 {
		allQuestions = append(allQuestions, question)
	}

	answer := 0
	for i := range operations {
		nums := allQuestions[i]
		op := operations[i]
		result := 0
		if op == "*" {
			result = 1
		}
		for _, n := range nums {
			switch op {
			case "+":
				result += n
			case "*":
				result *= n
			}
		}
		answer += result
	}

	return fmt.Sprintf("%v", answer)
}
