package solver

import (
	"bytes"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

type day09 struct{}

type point struct {
	i, j int
}

type dir struct {
	i, j int
}

type interval struct {
	l, h int
}

type directedLineSegment struct {
	from, to point
}

func (dls *directedLineSegment) direction() dir {
	return dir{
		i: sign(dls.to.i - dls.from.i),
		j: sign(dls.to.j - dls.from.j),
	}
}

func sign(n int) int {
	if n < 0 {
		return -1
	} else if n == 0 {
		return 0
	}
	return 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (d *day09) parseInput(input []byte) ([]point, int, int) {
	lines := bytes.Split(input, []byte("\n"))
	var points []point
	iMax, jMax := -1, -1
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		xS, yS, _ := strings.Cut(string(line), ",")
		x, _ := strconv.Atoi(xS)
		y, _ := strconv.Atoi(yS)
		points = append(points, point{i: y, j: x})
		if x > jMax {
			jMax = x
		}
		if y > iMax {
			iMax = y
		}
	}
	return points, iMax + 2, jMax + 2
}

func (d *day09) SolvePartOne(input []byte) string {
	points, _, _ := d.parseInput(input)

	largestArea := -1
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			p1 := points[i]
			p2 := points[j]
			iDiff := int(math.Abs(float64(p1.i - p2.i + 1)))
			jDiff := int(math.Abs(float64(p1.j - p2.j + 1)))
			if iDiff*jDiff > largestArea {
				largestArea = iDiff * jDiff
			}
		}
	}

	return fmt.Sprintf("%v", largestArea)
}

func (d *day09) printGrid(points []point, iMax, jMax int) {
	boundryPoints := make(map[point]struct{})
	for i := 1; i < len(points); i++ {
		p := points[i-1]
		q := points[i]

		boundryPoints[p] = struct{}{}
		if p.i == q.i {
			l := p.j
			h := q.j
			if l > h {
				l, h = h, l
			}
			for b := l; b <= h; b++ {
				boundryPoints[point{p.i, b}] = struct{}{}
			}
		} else {
			l := p.i
			h := q.i
			if l > h {
				l, h = h, l
			}
			for b := l; b <= h; b++ {
				boundryPoints[point{b, p.j}] = struct{}{}
			}
		}
	}

	for i := 0; i < iMax; i++ {
		for j := 0; j < jMax; j++ {
			c := "."
			if _, exist := boundryPoints[point{i, j}]; exist {
				c = "#"
			}
			fmt.Print(c)
		}
		fmt.Println()
	}
}

func (d *day09) SolvePartTwo(input []byte) string {
	points, iMax, jMax := d.parseInput(input)

	pointsLoop := append(points, points[0])

	d.printGrid(pointsLoop, iMax, jMax)

	var lineSegments []directedLineSegment
	for i := 1; i < len(pointsLoop); i++ {
		from := pointsLoop[i-1]
		to := pointsLoop[i]

		lineSegments = append(lineSegments, directedLineSegment{from: from, to: to})
	}

	lsLoop := append(append([]directedLineSegment{}, lineSegments...), lineSegments[0])

	// Turn counter
	rightTurns := 0
	for p := 1; p < len(lsLoop); p++ {
		prev := lsLoop[p-1]
		curr := lsLoop[p]

		prevDir := prev.direction()
		currDir := curr.direction()

		crossProduct := prevDir.i*currDir.j - prevDir.j*currDir.i
		if crossProduct < 0 {
			rightTurns += 1
		} else {
			rightTurns -= 1
		}
	}
	fmt.Println(rightTurns)

	verticalLines := make(map[int][]directedLineSegment)
	horizaontalLines := make(map[int][]directedLineSegment)

	for _, l := range lineSegments {
		switch l.direction() {
		case dir{1, 0}, dir{-1, 0}:
			{
				if _, exist := verticalLines[l.from.j]; !exist {
					verticalLines[l.from.j] = make([]directedLineSegment, 0)
				}
				verticalLines[l.from.j] = append(verticalLines[l.from.j], l)
			}
		case dir{0, 1}, dir{0, -1}:
			{
				if _, exist := horizaontalLines[l.from.i]; !exist {
					horizaontalLines[l.from.i] = make([]directedLineSegment, 0)
				}
				horizaontalLines[l.from.i] = append(horizaontalLines[l.from.i], l)
			}
		}
	}

	verticalPoints := mapKeys(verticalLines)
	horizontalPoints := mapKeys(horizaontalLines)

	slices.Sort(verticalPoints)
	slices.Sort(horizontalPoints)

	

	fmt.Printf("%v\n", verticalLines)
	fmt.Printf("%v\n", horizaontalLines)

	return "answer"
}

func mapKeys[T comparable, U any](m map[T]U) []T {
	l := make([]T, 0, len(m))
	for k := range m {
		l = append(l, k)
	}
	return l
}
