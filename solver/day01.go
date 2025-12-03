package solver

import (
	"bytes"
	"fmt"
	"strconv"
)

type day01 struct{}

func (d *day01) SolvePartOne(input []byte) string {
	lines := bytes.Split(input, []byte("\n"))

	val := 50
	answer := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		dir := string(line[0])
		mag, _ := strconv.Atoi(string(line[1:]))
		switch dir {
		case "R":
			val = (val + mag) % 100
		case "L":
			val = (val - mag) % 100
		}
		if val == 0 {
			answer += 1
		}
	}

	return fmt.Sprintf("%v", answer)
}

func (d *day01) SolvePartTwo(input []byte) string {
	lines := bytes.Split(input, []byte("\n"))

	val := 50
	answer := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		dir := string(line[0])
		mag, _ := strconv.Atoi(string(line[1:]))
		switch dir {
		case "R":
			{
				answer += mag / 100
				adjustedMag := mag % 100
				rawNewVal := val + adjustedMag
				if rawNewVal > 100 {
					answer += 1
				}
				val = rawNewVal % 100

			}
		case "L":
			{
				answer += mag / 100
				adjustedMag := mag % 100
				rawNewVal := val - adjustedMag
				if rawNewVal < 0 && val != 0 {
					answer += 1
				}
				val = (rawNewVal + 100) % 100
			}
		}
		if val == 0 {
			answer += 1
		}
	}

	return fmt.Sprintf("%v", answer)
}
