package solver

import (
	"bytes"
	"fmt"
	"strings"
)

type day11 struct{}

func (d *day11) parseInput(input []byte) map[string][]string {
	graph := make(map[string][]string)
	lines := bytes.Split(input, []byte("\n"))
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		ls := string(line)
		node, adj, _ := strings.Cut(ls, ": ")
		graph[node] = strings.Split(adj, " ")
	}
	return graph
}

func (d *day11) waysOutPartOne(start string, g map[string][]string, memo map[string]int) int {
	if start == "out" {
		return 1
	}
	if m, exist := memo[start]; exist {
		return m
	}
	ways := 0
	for _, n := range g[start] {
		ways += d.waysOutPartOne(n, g, memo)
	}
	memo[start] = ways
	return ways
}

func (d *day11) waysOutPartTwo(start string, dacExist, fftExist bool, g map[string][]string, memo map[string]int) int {
	if start == "out" {
		if dacExist && fftExist {
			return 1
		}
		return 0
	}
	k := fmt.Sprintf("%v|%v|%v", start, dacExist, fftExist)
	if m, exist := memo[k]; exist {
		return m
	}
	ways := 0
	for _, n := range g[start] {
		dac := dacExist || n == "dac"
		fft := fftExist || n == "fft"
		ways += d.waysOutPartTwo(n, dac, fft, g, memo)
	}
	memo[k] = ways
	return ways
}

func (d *day11) SolvePartOne(input []byte) string {
	g := d.parseInput(input)
	memo := make(map[string]int)
	return fmt.Sprintf("%v", d.waysOutPartOne("you", g, memo))
}

func (d *day11) SolvePartTwo(input []byte) string {
	g := d.parseInput(input)
	memo := make(map[string]int)
	return fmt.Sprintf("%v", d.waysOutPartTwo("svr", false, false, g, memo))
}
