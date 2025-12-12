package solver

import (
	"bytes"
	"fmt"
	"math/bits"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/draffensperger/golp"
)

type day10 struct{}

type machine struct {
	lights  []int
	buttons [][]int
	joltage []int
}

func (d *day10) parseInput(input []byte) []machine {
	lines := bytes.Split(input, []byte("\n"))
	parsedInput := make([]machine, 0, len(lines))
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		re := regexp.MustCompile(`^\[(?P<square>[^\]]*)\](?P<rounds>(?:\s*\([^)]*\))*)\s*\{(?P<curly>[^}]*)\}$`)

		match := re.FindStringSubmatch(string(line))
		if match == nil {
			panic("no match")
		}

		groups := make(map[string]string)
		for i, name := range re.SubexpNames() {
			if i > 0 && name != "" {
				groups[name] = match[i]
			}
		}

		lights := make([]int, 0, len(groups["square"]))
		for _, c := range groups["square"] {
			if c == '#' {
				lights = append(lights, 1)
			} else {
				lights = append(lights, 0)
			}
		}

		joltageRaw := strings.Split(groups["curly"], ",")
		joltage := make([]int, 0, len(joltageRaw))
		for _, jRaw := range joltageRaw {
			j, _ := strconv.Atoi(jRaw)
			joltage = append(joltage, j)
		}

		reRound := regexp.MustCompile(`\(([^)]*)\)`)
		roundMatches := reRound.FindAllStringSubmatch(groups["rounds"], -1)

		buttons := make([][]int, 0, len(roundMatches))
		for _, rm := range roundMatches {
			button := make([]int, 0)
			buttonNumsRaw := strings.Split(rm[1], ",")
			for _, numRaw := range buttonNumsRaw {
				num, _ := strconv.Atoi(numRaw)
				button = append(button, num)
			}
			buttons = append(buttons, button)
		}
		parsedInput = append(parsedInput, machine{
			lights:  lights,
			buttons: buttons,
			joltage: joltage,
		})
	}
	return parsedInput
}

func (d *day10) SolvePartOne(input []byte) string {
	machines := d.parseInput(input)
	finalAnswer := 0
	for _, m := range machines {
		limit := 1 << len(m.buttons)
		choices := make([]int, 0, limit)
		for i := 0; i < limit; i++ {
			choices = append(choices, i)
		}
		slices.SortFunc(choices, func(a, b int) int {
			a1 := bits.OnesCount(uint(a))
			b1 := bits.OnesCount(uint(b))
			if a1 != b1 {
				return a1 - b1
			}
			return a - b
		})

		answer := -1
		for _, c := range choices {
			lightState := make([]int, len(m.lights))
			for i := 0; i < len(m.buttons); i++ {
				selectBit := 1 << i
				if c&selectBit == selectBit {
					for _, l := range m.buttons[i] {
						lightState[l] = (lightState[l] + 1) % 2
					}
				}
			}
			same := true
			for i := 0; i < len(m.lights); i++ {
				if m.lights[i] != lightState[i] {
					same = false
					break
				}
			}
			if same {
				answer = bits.OnesCount(uint(c))
				break
			}

		}
		finalAnswer += answer
	}
	return fmt.Sprintf("%v", finalAnswer)
}

func (d *day10) SolvePartTwo(input []byte) string {
	machines := d.parseInput(input)
	answer := 0.0
	for _, m := range machines {
		lightToButtons := make([][]float64, len(m.joltage))
		for i := 0; i < len(m.buttons); i++ {
			for _, l := range m.buttons[i] {
				if len(lightToButtons[l]) == 0 {
					lightToButtons[l] = make([]float64, len(m.buttons))
				}
				lightToButtons[l][i] = 1.0
			}
		}
		lp := golp.NewLP(0, len(m.buttons))
		for i := range len(lightToButtons) {
			lp.AddConstraint(lightToButtons[i], golp.EQ, float64(m.joltage[i]))
		}
		for i := range len(m.buttons) {
			lp.SetInt(i, true)
		}
		obj := make([]float64, len(m.buttons))
		for i := range len(m.buttons) {
			obj[i] = 1
		}
		lp.SetObjFn(obj)
		lp.Solve()
		answer += lp.Objective()
	}

	return fmt.Sprintf("%v", answer)
}
