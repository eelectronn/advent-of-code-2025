package solver

import (
	"bytes"
	"fmt"
	"math"
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
			iDiff := int(math.Abs(float64(p1.i-p2.i))) + 1
			jDiff := int(math.Abs(float64(p1.j-p2.j))) + 1
			if iDiff*jDiff > largestArea {
				largestArea = iDiff * jDiff
			}
		}
	}

	return fmt.Sprintf("%v", largestArea)
}

func (d *day09) printGrid(boundryPoints map[point]struct{}, iMax, jMax int) {
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
	fmt.Println(iMax)
	fmt.Println(jMax)
	pointsLoop := append(points, points[0])

	boundryPoints := make(map[point]struct{})
	iWithCorners := make(map[int]struct{}, 0)
	jWithCorners := make(map[int]struct{}, 0)
	for i := 1; i < len(pointsLoop); i++ {
		p := pointsLoop[i-1]
		q := pointsLoop[i]

		iWithCorners[q.i] = struct{}{}
		jWithCorners[q.j] = struct{}{}

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

	// d.printGrid(boundryPoints, iMax, jMax)

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

	directionsOnPoints := make(map[point]map[dir]struct{})
	for _, l := range lineSegments {
		lDir := l.direction()
		if _, exist := directionsOnPoints[l.from]; !exist {
			directionsOnPoints[l.from] = make(map[dir]struct{})
		}
		if _, exist := directionsOnPoints[l.to]; !exist {
			directionsOnPoints[l.to] = make(map[dir]struct{})
		}
		directionsOnPoints[l.from][lDir] = struct{}{}
		directionsOnPoints[l.to][lDir] = struct{}{}
		p := l.from
		q := l.to
		if p.i == q.i {
			l := p.j
			h := q.j
			if l > h {
				l, h = h, l
			}
			for b := l; b <= h; b++ {
				if _, exist := directionsOnPoints[point{p.i, b}]; !exist {
					directionsOnPoints[point{p.i, b}] = make(map[dir]struct{})
				}
				directionsOnPoints[point{p.i, b}][lDir] = struct{}{}
			}
		} else {
			l := p.i
			h := q.i
			if l > h {
				l, h = h, l
			}
			for b := l; b <= h; b++ {
				if _, exist := directionsOnPoints[point{b, p.j}]; !exist {
					directionsOnPoints[point{b, p.j}] = make(map[dir]struct{})
				}
				directionsOnPoints[point{b, p.j}][lDir] = struct{}{}
			}
		}
	}

	fmt.Println("here")
	iIncludedPoints := make(map[int]map[int]struct{})
	for i := range iWithCorners {
		inside := false
		included := make(map[int]struct{}, 0)
		for j := 0; j < jMax; j++ {
			if _, exist := boundryPoints[point{i, j}]; exist || inside {
				included[j] = struct{}{}
			}
			if pointDirs, exist := directionsOnPoints[point{i, j}]; exist {
				_, upExist := pointDirs[dir{-1, 0}]
				_, downExist := pointDirs[dir{1, 0}]
				if upExist {
					inside = true
				} else if downExist {
					inside = false
				}
			}
		}
		iIncludedPoints[i] = included
	}

	// fmt.Println(iInludedPoints)

	fmt.Println("here")
	jIncludedPoints := make(map[int]map[int]struct{})
	for j := range jWithCorners {
		inside := false
		included := make(map[int]struct{}, 0)
		for i := 0; i < iMax; i++ {
			if _, exist := boundryPoints[point{i, j}]; exist || inside {
				included[i] = struct{}{}
			}
			if pointDirs, exist := directionsOnPoints[point{i, j}]; exist {
				_, rightExist := pointDirs[dir{0, 1}]
				_, leftExist := pointDirs[dir{0, -1}]
				if rightExist {
					inside = true
				} else if leftExist {
					inside = false
				}
			}
		}
		jIncludedPoints[j] = included
	}

	largestArea := -1
	for x := range points {
	nextRect:
		for y := x + 1; y < len(points); y++ {
			p1 := points[x]
			p2 := points[y]
			topLeft := point{i: min(p1.i, p2.i), j: min(p1.j, p2.j)}
			topRight := point{i: min(p1.i, p2.i), j: max(p1.j, p2.j)}
			bottomLeft := point{i: max(p1.i, p2.i), j: min(p1.j, p2.j)}
			bottomRight := point{i: max(p1.i, p2.i), j: max(p1.j, p2.j)}

			top := directedLineSegment{from: topLeft, to: topRight}
			right := directedLineSegment{from: topRight, to: bottomRight}
			bottom := directedLineSegment{from: bottomLeft, to: bottomRight}
			left := directedLineSegment{from: topLeft, to: bottomLeft}

			validPts := iIncludedPoints[top.from.i]
			for k := top.from.j; k <= top.to.j; k++ {
				if _, exist := validPts[k]; !exist {
					continue nextRect
				}
			}

			validPts = iIncludedPoints[bottom.from.i]
			for k := bottom.from.j; k <= bottom.to.j; k++ {
				if _, exist := validPts[k]; !exist {
					continue nextRect
				}
			}

			validPts = jIncludedPoints[right.from.j]
			for k := right.from.i; k <= right.to.i; k++ {
				if _, exist := validPts[k]; !exist {
					continue nextRect
				}
			}

			validPts = jIncludedPoints[left.from.j]
			for k := left.from.i; k <= left.to.i; k++ {
				if _, exist := validPts[k]; !exist {
					continue nextRect
				}
			}
			iDiff := int(math.Abs(float64(p1.i-p2.i))) + 1
			jDiff := int(math.Abs(float64(p1.j-p2.j))) + 1
			if iDiff*jDiff > largestArea {
				largestArea = iDiff * jDiff
			}
		}
	}

	fmt.Println(largestArea)

	// fmt.Println(jIncludedPoints)

	return "answer"
}
