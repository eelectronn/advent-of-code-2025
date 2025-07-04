package solver

import (
	"bytes"
	"fmt"
	"strconv"
)

type day02 struct{}

type triplet struct {
	red   int
	green int
	blue  int
}

func (t triplet) isLessOrEqual(o triplet) bool {
	return t.red <= o.red && t.green <= o.green && t.blue <= o.blue
}

func parseInput(input []byte) [][]triplet {
	lines := bytes.Split(input, []byte("\n"))
	games := make([][]triplet, 0, len(lines))
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		_, drawsStr, found := bytes.Cut(line, []byte(":"))
		if !found {
			continue
		}
		draws := bytes.Split(drawsStr, []byte(";"))
		gameDraws := make([]triplet, 0, len(draws))
		for _, draw := range draws {
			values := getDrawValue(draw)
			gameDraws = append(gameDraws, triplet{
				red:   values["red"],
				green: values["green"],
				blue:  values["blue"],
			})
		}
		games = append(games, gameDraws)
	}

	return games

}

func getDrawValue(draw []byte) map[string]int {
	colorsStr := bytes.Split(draw, []byte(","))
	colorTriple := make(map[string]int)
	for _, colorStr := range colorsStr {
		colorStr = bytes.TrimSpace(colorStr)
		numBytes, color, found := bytes.Cut(colorStr, []byte(" "))
		if found {
			num, err := strconv.Atoi(string(numBytes))
			if err != nil {
				panic("parse int error")
			}
			colorTriple[string(color)] = num
		}
	}
	return colorTriple
}

func (d *day02) SolvePartOne(input []byte) string {
	parsed := parseInput(input)

	all := triplet{
		red:   12,
		green: 13,
		blue:  14,
	}

	answer := 0
outer:
	for i, game := range parsed {
		for _, draw := range game {
			if !draw.isLessOrEqual(all) {
				continue outer
			}
		}
		answer += i + 1
	}

	return fmt.Sprintf("%v", answer)
}

func (d *day02) SolvePartTwo(input []byte) string {
	parsed := parseInput(input)

	answer := 0

	for _, game := range parsed {
		min := triplet{}
		for _, draw := range game {
			min.red = maxInt(min.red, draw.red)
			min.green = maxInt(min.green, draw.green)
			min.blue = maxInt(min.blue, draw.blue)
		}
		answer += min.red * min.green * min.blue
	}

	return fmt.Sprintf("%v", answer)
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
