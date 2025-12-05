package solver

import (
	"bytes"
	"fmt"
	"slices"
	"strconv"
)

type day05 struct{}

type rng struct {
	low  int
	high int
}

func (d *day05) parseInput(input []byte) ([]rng, []int) {
	database, queries, _ := bytes.Cut(input, []byte("\n\n"))
	var parsedDatabase []rng
	for _, line := range bytes.Split(database, []byte("\n")) {
		if len(line) == 0 {
			continue
		}
		low, high, _ := bytes.Cut(line, []byte("-"))
		lowInt, _ := strconv.Atoi(string(low))
		highInt, _ := strconv.Atoi(string(high))
		parsedDatabase = append(parsedDatabase, rng{low: lowInt, high: highInt})
	}

	var parsedQueries []int
	for _, line := range bytes.Split(queries, []byte("\n")) {
		if len(line) == 0 {
			continue
		}
		queryInt, _ := strconv.Atoi(string(line))
		parsedQueries = append(parsedQueries, queryInt)
	}

	return parsedDatabase, parsedQueries
}

func (d *day05) SolvePartOne(input []byte) string {
	database, queries := d.parseInput(input)
	answer := 0

	for _, query := range queries {
		for _, record := range database {
			if query >= record.low && query <= record.high {
				answer++
				break
			}
		}
	}

	return strconv.Itoa(answer)
}

func (d *day05) SolvePartTwo(input []byte) string {
	database, _ := d.parseInput(input)

	slices.SortFunc(database, func(a, b rng) int {
		if a.low != b.low {
			return a.low - b.low
		}
		return a.high - b.high
	})

	disjointRanges := []rng{database[0]}
	for _, record := range database[1:] {
		last := &disjointRanges[len(disjointRanges)-1]
		if record.low <= last.high {
			if record.high > last.high {
				last.high = record.high
			}
		} else {
			disjointRanges = append(disjointRanges, record)
		}
	}

	answer := 0
	for _, record := range disjointRanges {
		answer += record.high - record.low + 1
	}

	return fmt.Sprintf("%v", answer)
}
