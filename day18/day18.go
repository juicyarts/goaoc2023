package day18

import (
	"math"
	"strconv"
	"strings"
)

type DirCfg struct {
	d int
	a int
}

var dirCfg = map[string]DirCfg{
	"0": {
		d: 0,
		a: 1,
	},
	"2": {
		d: 0,
		a: -1,
	},
	"1": {
		d: 1,
		a: 1,
	},
	"3": {
		d: 1,
		a: -1,
	},
}

func Travel(input []string) (int, int) {
	points := [][]int{}
	x, y := 0, 0
	for i := 0; i < len(input); i++ {
		inst := strings.Split(input[i], " ")
		_, _, color := inst[0], inst[1], inst[2]
		dir := color[7:8]
		a, _ := strconv.ParseInt(color[2:7], 16, 64)
		dirConfig := dirCfg[dir]

		for j := 0; j < int(a); j++ {
			if dirConfig.d == 0 {
				if dirConfig.a > 0 {
					x++
				}
				if dirConfig.a < 0 {
					x--
				}
			}

			if dirConfig.d == 1 {
				if dirConfig.a < 0 {
					y--
				}
				if dirConfig.a > 0 {
					y++
				}
			}
			points = append(points, []int{x, y})
		}
	}

	area := 0

	// shoelace https://en.wikipedia.org/wiki/Shoelace_formula
	for i := 0; i < len(points); i++ {
		if i > 1 {
			area += points[i][0] * (points[i-1][1] - points[(i+1)%len(points)][1])
		}
	}

	area = int(math.Abs(float64(area))) / 2

	// picks theorem https://en.wikipedia.org/wiki/Pick%27s_theorem
	area = area - len(points)/2 + 1

	return len(points), area + len(points)
}
