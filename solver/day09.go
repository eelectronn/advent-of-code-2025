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

func (d *day09) SolvePartTwo(input []byte) string {
	points, iMax, jMax := d.parseInput(input)

	fmt.Println(iMax, jMax)

	points = append(points, points[0])

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

	fmt.Println("lines done")

	// for i := 0; i < iMax; i++ {
	// for j := 0; j < jMax; j++ {
	// c := "."
	// if _, exist := boundryPoints[point{i, j}]; exist {
	// c = "#"
	// }
	// fmt.Print(c)
	// }
	// fmt.Println()
	// }

	queue := NewQueue[point]()
	queue.Enqueue(point{0, 0})
	outside := make(map[point]struct{})

	for queue.Len() > 0 {
		curr, _ := queue.Dequeue()
		if _, seen := outside[curr]; seen {
			continue
		}
		outside[curr] = struct{}{}
		for _, iOffset := range []int{-1, 0, 1} {
			for _, jOffset := range []int{-1, 0, 1} {
				if iOffset == 0 && jOffset == 0 {
					continue
				}

				n := point{curr.i + iOffset, curr.j + jOffset}
				if 0 <= n.i && n.i < iMax && 0 <= n.j && n.j < jMax {
					if _, boundryPoint := boundryPoints[n]; !boundryPoint {
						queue.Enqueue(n)
					}
				}

			}
		}
	}

	// for i := 0; i < iMax; i++ {
	// for j := 0; j < jMax; j++ {
	// c := "."
	// if _, exist := boundryPoints[point{i, j}]; exist {
	// c = "#"
	// } else if _, exist := outside[point{i, j}]; exist {
	// c = "O"
	// }
	// fmt.Print(c)
	// }
	// fmt.Println()
	// }

	usable := make(map[point]struct{})
	for i := 0; i < iMax; i++ {
		for j := 0; j < jMax; j++ {
			p := point{i, j}
			if _, exist := boundryPoints[p]; exist {
				usable[p] = struct{}{}
			}
			if _, exist := outside[p]; !exist {
				usable[p] = struct{}{}
			}
		}
	}

	// for i := 0; i < iMax; i++ {
	// for j := 0; j < jMax; j++ {
	// c := "."
	// if _, exist := usable[point{i, j}]; exist {
	// c = "#"
	// }
	// fmt.Print(c)
	// }
	// fmt.Println()
	// }

	return "answer"
}
