package day24

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type Point struct {
	x, y, z float64
}

type Eq struct {
	a, b, c float64
}

type Stone struct {
	start Point
	end   Point
	vel   Point
	eq    Eq
	name  int
}

var hailstone = color.New(color.FgRed, color.Bold).SprintFunc()
var multiple = color.New(color.FgGreen, color.Bold).SprintFunc()
var startColor = color.New(color.FgCyan, color.Bold).SprintFunc()
var endColor = color.New(color.FgBlue, color.Bold).SprintFunc()
var collisionColor = color.New(color.FgHiYellow, color.Bold, color.BlinkRapid).SprintFunc()

var space = color.New(color.Bold, color.FgBlack).SprintFunc()
var emptySpace = color.New(color.Bold, color.FgBlack).SprintFunc()
var axis = color.New(color.FgWhite).SprintFunc()

func ReadInput(input []string, min float64, max float64) int {
	initialStones := map[string]Stone{}
	maxX, maxY, maxZ := max, max, max
	minX, minY, minZ := min, min, min

	for rowIndex, r := range input {
		p := strings.Split(strings.Split(r, " @")[0], ", ")
		v := strings.Split(strings.Split(r, "@ ")[1], ", ")

		posAsInt := []float64{}
		velAsInt := []float64{}
		for _, s := range p {
			n, _ := strconv.Atoi(strings.Trim(s, " "))
			posAsInt = append(posAsInt, float64(n))
		}

		for _, e := range v {
			n, _ := strconv.Atoi(strings.Trim(e, " "))
			velAsInt = append(velAsInt, float64(n))
		}

		possibleSteps := float64(0)
		tx, ty, tz := float64(0), float64(0), float64(0)

		if velAsInt[0] > 0 {
			tx = (maxX - posAsInt[0]) / velAsInt[0]
		} else if velAsInt[0] < 0 {
			tx = (minX - posAsInt[0]) / velAsInt[0]
		}

		if velAsInt[1] > 0 {
			ty = (maxY - posAsInt[1]) / velAsInt[1]
		} else if velAsInt[1] < 0 {
			ty = (minY - posAsInt[1]) / velAsInt[1]
		}

		if velAsInt[2] > 0 {
			tz = (maxZ - posAsInt[2]) / velAsInt[2]
		} else if velAsInt[2] < 0 {
			tz = (minZ - posAsInt[2]) / velAsInt[2]
		}

		possibleSteps = math.Min(tx, math.Min(ty, tz))

		initialStones[fmt.Sprintf("%+v", rowIndex)] = Stone{
			name: rowIndex,
			start: Point{
				x: posAsInt[0],
				y: posAsInt[1],
				z: posAsInt[2],
			},
			end: Point{
				x: posAsInt[0] + (possibleSteps * velAsInt[0]),
				y: posAsInt[1] + (possibleSteps * velAsInt[1]),
				z: posAsInt[2] + (possibleSteps * velAsInt[2]),
			},
			vel: Point{
				x: velAsInt[0],
				y: velAsInt[1],
				z: velAsInt[2],
			},
			eq: Eq{
				a: velAsInt[1],
				b: -velAsInt[0],
				c: velAsInt[1]*posAsInt[0] - velAsInt[0]*posAsInt[1],
			},
		}
	}

	collisions := 0
	for _, stone := range initialStones {

		for _, otherStone := range initialStones {
			if otherStone.name == stone.name {
				continue
			}

			eqCheck, point := doEquationsIntersect(stone.eq, otherStone.eq, minX, maxX, minY, maxY)
			// ensure point is not smaller than the origin of a stone // exclude "past"
			if eqCheck && ((point.x-stone.start.x)*stone.vel.x >= 0 && (point.y-stone.start.y)*stone.vel.y >= 0 &&
				(point.x-otherStone.start.x)*otherStone.vel.x >= 0 && (point.y-otherStone.start.y)*otherStone.vel.y >= 0) {
				fmt.Print("Line Equotation based check Check: ", eqCheck, point, "\n")
				collisions++
			}
		}
	}

	return collisions / 2 // dividie by two since every collision is recorded twice
}

// this checks intersections based on line equation instead of segments, seems to be more accurate
// i guess there is something off wiht how i set the "end" based on which direction i can
// add the least amount of segments to until maxX/maxY/minY/minY are reached
// Found issue -> when setting possible steps i was not respecting floating point values
// https://en.wikipedia.org/wiki/Line%E2%80%93line_intersection#Given_two_line_equations
func doEquationsIntersect(eq Eq, eq2 Eq, minX, maxX, minY, maxY float64) (bool, Point) {
	if eq.a*eq2.b == eq.b*eq2.a {
		return false, Point{0, 0, 0}
	}

	x := (eq.c*eq2.b - eq2.c*eq.b) / (eq.a*eq2.b - eq2.a*eq.b)
	y := (eq2.c*eq.a - eq.c*eq2.a) / (eq.a*eq2.b - eq2.a*eq.b)

	if x >= minX && x <= maxX && y >= minY && y <= maxY {
		return true, Point{x, y, 0}
	}

	return false, Point{0, 0, 0}
}

func crossProduct(p1, p2, p3 Point) float64 {
	return (p2.x-p1.x)*(p3.y-p1.y) - (p2.y-p1.y)*(p3.x-p1.x)
}

// two points given per line / no segments
// still the most useful for painting purposes. Guess it's up to being not all too precise
// https://en.wikipedia.org/wiki/Line%E2%80%93line_intersection#Given_two_points_on_each_line
func doSegmentsIntersect(p1, p2, p3, p4 Point, tol, minX, maxX, minY, maxY, minZ, maxZ float64) (bool, Point) {
	intersects := crossProduct(p1, p3, p4)*crossProduct(p2, p3, p4) <= tol &&
		crossProduct(p3, p1, p2)*crossProduct(p4, p1, p2) <= tol

	if intersects {
		denominator := (p1.x-p2.x)*(p3.y-p4.y) - (p1.y-p2.y)*(p3.x-p4.x)
		if denominator == 0 {
			return true, Point{x: 0, y: 0}
		}

		t := (((p1.x - p3.x) * (p3.y - p4.y)) - ((p1.y - p3.y) * (p3.x - p4.x))) / denominator
		u := (((p1.x - p3.x) * (p1.y - p2.y)) - ((p1.y - p3.y) * (p1.x - p2.x))) / denominator

		px, py := (p1.x + t*(p2.x-p1.x)), (p1.y + t*(p2.y-p1.y))

		if t >= 0 && t <= 1 && u >= 0 && u <= 1 {
			// if px >= minX && px <= maxX && py >= minY && py <= maxY {
			return true, Point{x: px, y: py}
			// }
		}
	}

	return false, Point{}
}

func printMap(stones map[string]Stone, minX, maxX, minY, maxY, minZ, maxZ float64, view string, currentZ float64, collisionMap map[string]Point) {
	fmt.Printf("--------------- \n")
	fmt.Printf("Drawing Min|Max x:%+v|%+v,y:%+v|%+v,z:%+v|%+v,%+v at level: %+v \n", minX, maxX, minY, maxY, minZ, maxZ, view, currentZ)
	fmt.Printf("--------------- \n")

	//
	padding := float64(25)
	resolution := float64(1)

	for i := (maxY + padding) / resolution; i >= (minY-padding)/resolution; i-- {
		fmt.Print("\n")
		for j := (minX - padding) / resolution; j <= (maxX+padding)/resolution; j++ {
			foundStones := []string{}
			isStart := false
			isEnd := false
			foundCollision := false

			if j == (minX-padding)/resolution {
				fmt.Print(axis(fmt.Sprintf("%3v", i)))
			}

			for _, stone := range stones {

				if ok, _ := doSegmentsIntersect(
					Point{float64(j * resolution), float64((i) * resolution), float64(0)},
					Point{float64(j * resolution), float64((i) * resolution), float64(0)},
					stone.start,
					stone.end,
					1,
					minX,
					maxX,
					minY,
					maxY,
					minZ,
					maxZ,
				); ok {
					foundStones = append(foundStones, strconv.Itoa(stone.name))
				}
				if stone.start.x == j && stone.start.y == i {
					isStart = true
				} else if stone.end.x == j && stone.end.y == i {
					isEnd = true
				}
			}

			for _, collision := range collisionMap {
				if math.Ceil(collision.x) == float64(j*resolution) && math.Ceil(collision.y) == float64(i*resolution) {
					foundCollision = true
				}
			}

			if len(foundStones) <= 0 {
				if j < minX || i < minY || j > maxX || i > maxY {
					fmt.Print(emptySpace(fmt.Sprintf("%4s", "*")))
				} else {
					fmt.Print(space(fmt.Sprintf("%4s", "*")))
				}
			} else {
				if len(foundStones) > 1 {
					if foundCollision {
						fmt.Print(collisionColor(fmt.Sprintf("%4s", strings.Join(foundStones, ","))))
					} else if isStart {
						fmt.Print(startColor(fmt.Sprintf("%4s", strings.Join(foundStones, ","))))
					} else if isEnd {
						fmt.Print(endColor(fmt.Sprintf("%4s", strings.Join(foundStones, ","))))
					} else {
						fmt.Print(multiple(fmt.Sprintf("%4s", strings.Join(foundStones, ","))))
					}
				} else {
					if foundCollision {
						fmt.Print(collisionColor(fmt.Sprintf("%4s", strings.Join(foundStones, ","))))
					} else if isStart {
						fmt.Print(startColor(fmt.Sprintf("%4s", foundStones[0])))
					} else if isEnd {
						fmt.Print(endColor(fmt.Sprintf("%4s", foundStones[0])))
					} else {
						fmt.Print(hailstone(fmt.Sprintf("%4s", foundStones[0])))
					}
				}
			}
		}
	}

	fmt.Print("\n\n   ")
	for i := (minX - padding) / resolution; i <= (maxX+padding)/resolution; i++ {
		fmt.Print(axis(fmt.Sprintf("%4d", int(i))))
	}
	fmt.Print("\n")
}
