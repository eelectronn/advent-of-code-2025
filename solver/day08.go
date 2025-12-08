package solver

import (
	"bytes"
	"fmt"
	"math"
	"slices"
	"strconv"
)

type loc struct {
	x, y, z int
}

type day08 struct{}

func (d *day08) parseInput(input []byte) []loc {
	lines := bytes.Split(input, []byte("\n"))
	locs := make([]loc, 0)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		vals := bytes.Split(line, []byte(","))
		valInts := make([]int, 0, 3)
		for _, val := range vals {
			valInt, _ := strconv.Atoi(string(val))
			valInts = append(valInts, valInt)
		}
		locs = append(locs, loc{x: valInts[0], y: valInts[1], z: valInts[2]})
	}
	return locs
}

func (d *day08) getDistancesAndMap(locs []loc) ([]float64, map[float64][]loc) {
	distances := make([]float64, 0)
	pairs := make(map[float64][]loc)
	for i := 0; i < len(locs); i++ {
		for j := i + 1; j < len(locs); j++ {
			d := math.Sqrt(
				math.Pow(float64(locs[i].x-locs[j].x), 2) +
					math.Pow(float64(locs[i].y-locs[j].y), 2) +
					math.Pow(float64(locs[i].z-locs[j].z), 2),
			)
			distances = append(distances, d)
			pairs[d] = []loc{locs[i], locs[j]}
		}
	}

	slices.Sort(distances)
	return distances, pairs
}

func (d *day08) SolvePartOne(input []byte) string {
	locs := d.parseInput(input)
	disjointSet := NewDisjointSet[loc]()

	for _, l := range locs {
		disjointSet.Add(l)
	}

	distances, pairs := d.getDistancesAndMap(locs)

	for _, d := range distances[:10] {
		pair := pairs[d]
		disjointSet.Join(pair[0], pair[1])
	}

	clusterCount := make(map[loc]int)

	for _, loc := range locs {
		if p, exist := disjointSet.Parent(loc); exist {
			if _, ok := clusterCount[p]; !ok {
				clusterCount[p] = 0
			}
			clusterCount[p] += 1
		}
	}

	clusterCountList := make([]int, 0, len(clusterCount))
	for _, v := range clusterCount {
		clusterCountList = append(clusterCountList, v)
	}

	slices.Sort(clusterCountList)
	slices.Reverse(clusterCountList)

	answer := clusterCountList[0] * clusterCountList[1] * clusterCountList[2]

	return fmt.Sprintf("%v", answer)
}

func (d *day08) SolvePartTwo(input []byte) string {
	locs := d.parseInput(input)
	disjointSet := NewDisjointSet[loc]()

	for _, l := range locs {
		disjointSet.Add(l)
	}

	distances, pairs := d.getDistancesAndMap(locs)

	for _, d := range distances {
		pair := pairs[d]
		disjointSet.Join(pair[0], pair[1])
		if disjointSet.NumberOfSets() == 1 {
			return fmt.Sprintf("%v", pair[0].x*pair[1].x)
		}
	}
	return "not-found"
}
