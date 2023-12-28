package day21

import (
	"fmt"
	"math"

	"github.com/fatih/color"
)

var primary = color.New(color.FgRed).SprintFunc()
var secondary = color.New(color.FgYellow).SprintFunc()
var tertiary = color.New(color.FgCyan).SprintFunc()
var white = color.New(color.FgWhite).SprintFunc()

var directions = []int{
	0, 1, 2, 3,
}

type Node struct {
	next      []int
	stepsLeft int
}

func Travel(input []string, steps int) int {
	start := []int{0, 0}
	finalPlots := 0

	for lI, l := range input {
		for cI, c := range l {
			if c == 'S' {
				start = []int{lI, cI}
			}
		}
	}

	finalPlots = Walk(input, start, steps)

	return finalPlots
}

// Part 1
func getPossibleCount(m []string, start []int, stepsLeft int) (int, map[string]int, map[string]int) {
	visited := map[string]int{fmt.Sprintf("%v", start): 1}
	counter := map[string]int{}
	queue := []Node{
		{
			next:      start,
			stepsLeft: stepsLeft,
		},
	}

	for i := 0; i < len(queue); i++ {
		node := queue[i]
		nodeKey := fmt.Sprintf("%v", node.next)

		if node.stepsLeft%2 == 0 {
			counter[nodeKey] = 1
		}

		if node.stepsLeft == 0 {
			continue
		}

		var directionToNext = map[int][]int{
			0: {node.next[0] - 1, node.next[1]},
			1: {node.next[0], node.next[1] + 1},
			2: {node.next[0] + 1, node.next[1]},
			3: {node.next[0], node.next[1] - 1},
		}

		for dirIndex := 0; dirIndex < len(directions); dirIndex++ {
			next := directionToNext[dirIndex]
			nextKey := fmt.Sprintf("%v", next)
			if _, ok := visited[nextKey]; ok || next[0] < 0 || next[1] < 0 || next[0] >= len(m) || next[1] >= len(m[0]) || m[next[0]][next[1]] == '#' {
				continue
			}

			visited[nextKey] = 1
			queue = append(queue, Node{next: next, stepsLeft: node.stepsLeft - 1})
		}
	}

	return len(counter), visited, counter
}

func Walk(m []string, start []int, stepsLeft int) int {

	if len(m) != len(m[0]) {
		panic("Can't do this, the input has no quadratic form")
	}

	size := len(m)
	rowWidth := (stepsLeft / size) - 1

	fmt.Print("\n")

	tileMap := map[string]int{
		"evenTile": 0,
		"oddTile":  0,
		"rth":      0,
		"rtq":      0,
		"rbh":      0,
		"rbq":      0,
		"lth":      0,
		"ltq":      0,
		"lbh":      0,
		"lbq":      0,
		"le":       0,
		"re":       0,
		"te":       0,
		"be":       0,
	}

	tileResults := map[string]int{}

	for itemKey := range tileMap {
		fmt.Print("\n-------------------------------------------------------\n")
		fmt.Print(itemKey)
		hits := 0
		visited, counter := map[string]int{}, map[string]int{}

		// Even					Odd
		// 0.0.0.0.0.0  .0.0.0.0.0.
		// .0.0.0.0.0.  0.0.0.0.0.0
		// 0.0.0.0.0.0  .0.0.0.0.0.
		// .0.0.0.0.0.  0.0.0.0.0.0
		// 0.0.0.0.0.0  .0.0.0.0.0.
		// .0.0.S.0.0.  0.0.0S0.0.0
		// 0.0.0.0.0.0  .0.0.0.0.0.
		// .0.0.0.0.0.  0.0.0.0.0.0
		// 0.0.0.0.0.0  .0.0.0.0.0.
		// .0.0.0.0.0.  0.0.0.0.0.0
		// 0.0.0.0.0.0  .0.0.0.0.0.
		//
		// for full tiles we need to alternate between starting with an even number of steps left vs an odd
		// given our map length is odd we need to alternate patterns which ends up in different results
		if itemKey == "oddTile" {
			// given an odd tile we want to have an odd number of steps starting from our origin
			// but a stepLenght wide enough to cover the full area
			hits, visited, counter = getPossibleCount(m, start, size*2+1)
		} else if itemKey == "evenTile" {
			// given an even tile we want to have an even number of steps starting from our origin
			// also a stepLenght wide enough to cover the full area
			hits, visited, counter = getPossibleCount(m, start, size*2)
		}

		// Top End			Bottom End	 Left End			Right End
		// .....0.....  .0.0.0.0.0.  .....0.0.0.  .0.0.0.....
		// ....0.0....  0.0.0.0.0.0  ....0.0.0.0  0.0.0.0....
		// ...0.0.0...  .0.0.0.0.0.  ...0.0.0.0.  .0.0.0.0...
		// ..0.0.0.0..  0.0.0.0.0.0  ..0.0.0.0.0  0.0.0.0.0..
		// .0.0.0.0.0.  .0.0.0.0.0.  .0.0.0.0.0.  .0.0.0.0.0.
		// 0.0.0S0.0.0  0.0.0S0.0.0  0.0.0S0.0.0  0.0.0S0.0.0
		// .0.0.0.0.0.  .0.0.0.0.0.  .0.0.0.0.0.  .0.0.0.0.0.
		// 0.0.0.0.0.0  ..0.0.0.0..  ..0.0.0.0.0  0.0.0.0.0..
		// .0.0.0.0.0.  ...0.0.0...  ...0.0.0.0.  .0.0.0.0...
		// 0.0.0.0.0.0  ....0.0....  ....0.0.0.0  0.0.0.0....
		// .0.0.0.0.0.  .....0.....  .....0.0.0.  .0.0.0.....
		//
		// we can assume that start is always in the middle of our map
		// we also assume that on ends the "final" reachable point will be on the rightmost/leftmost/topmost/bottommost/ point from our start
		// lastly we assume that we always reach the ends/corners (te, re, be, le) with enough steps left to reach the opposite side
		if itemKey == "te" {
			// for the top end we expect the starting point to be the last row's center column
			hits, visited, counter = getPossibleCount(m, []int{size - 1, start[1]}, size-1)
		} else if itemKey == "re" {
			// for the right end we expect the starting point to be the center row and the first column
			hits, visited, counter = getPossibleCount(m, []int{start[0], 0}, size-1)
		} else if itemKey == "be" {
			// v case
			// for the bottom end we expect the starting point to be the top rows center column
			hits, visited, counter = getPossibleCount(m, []int{0, start[1]}, size-1)
		} else if itemKey == "le" {
			// < case
			// for the left end we expect the starting point to be the center row and the last column
			hits, visited, counter = getPossibleCount(m, []int{start[0], size - 1}, size-1)
		}

		// now we get to the "halfs". These are the "bigger" elements of slopes while "quarters" the are smaller
		// Right Top		Left Top		 Right Btm		Left Btm
		// .0.0.0.....  .....0.0.0.  .0.0.0.0.0.  .0.0.0.0.0.
		// 0.0.0.0....  ....0.0.0.0  0.0.0.0.0.0  0.0.0.0.0.0
		// .0.0.0.0...  ...0.0.0.0.  .0.0.0.0.0.  .0.0.0.0.0.
		// 0.0.0.0.0..  ..0.0.0.0.0  0.0.0.0.0.0  0.0.0.0.0.0
		// .0.0.0.0.0.  .0.0.0.0.0.  .0.0.0.0.0.  .0.0.0.0.0.
		// 0.0.0S0.0.0  0.0.0S0.0.0  0.0.0S0.0.0  0.0.0S0.0.0
		// .0.0.0.0.0.  .0.0.0.0.0.  .0.0.0.0.0.  .0.0.0.0.0.
		// 0.0.0.0.0.0  0.0.0.0.0.0  0.0.0.0.0..  ..0.0.0.0.0
		// .0.0.0.0.0.  .0.0.0.0.0.  .0.0.0.0...  ...0.0.0.0.
		// 0.0.0.0.0.0  0.0.0.0.0.0  0.0.0.0....  ....0.0.0.0
		// .0.0.0.0.0.  .0.0.0.0.0.  .0.0.0.....  .....0.0.0.
		//
		// the step length is set reach 3/4 of the tile
		if itemKey == "rth" {
			// for the right top half we expect the starting point to be the last row's first column
			hits, visited, counter = getPossibleCount(m, []int{size - 1, 0}, size*3/2-1)
		} else if itemKey == "lth" {
			// for the left top half we expect the starting point to be the last row's last column
			hits, visited, counter = getPossibleCount(m, []int{size - 1, size - 1}, size*3/2-1)
		} else if itemKey == "rbh" {
			// for the right bottom half we expect the starting point to be the first row's first column
			hits, visited, counter = getPossibleCount(m, []int{0, 0}, size*3/2-1)
		} else if itemKey == "lbh" {
			// for the left bottom half we expect the starting point to be the first row's last column
			hits, visited, counter = getPossibleCount(m, []int{0, size - 1}, size*3/2-1)
		}

		// now we handle the quarters
		// Right Top 		Left Top		 Right Btm		Left Btm
		// ...........  ...........  0.0.0......  ......0.0.0
		// ...........  ...........  .0.0.......  .......0.0.
		// ...........  ...........  0.0........  ........0.0
		// ...........  ...........  .0.........  .........0.
		// ...........  ...........  0..........  ..........0
		// .....S.....  .....S.....  .....S.....  .....S.....
		// 0..........  ..........0  ...........  ...........
		// .0.........  .........0.  ...........  ...........
		// 0.0........  ........0.0  ...........  ...........
		// .0.0.......  .......0.0.  ...........  ...........
		// 0.0.0......  ......0.0.0  ...........  ...........
		//
		// the step length is set reach 1/4 of the tile
		if itemKey == "rtq" {
			// for the right top quarter we expect the starting point to be the last row's first column
			hits, visited, counter = getPossibleCount(m, []int{size - 1, 0}, size/2-1)
		} else if itemKey == "ltq" {
			// for the left top quarter we expect the starting point to be the last row's last column
			hits, visited, counter = getPossibleCount(m, []int{size - 1, size - 1}, size/2-1)
		} else if itemKey == "rbq" {
			// for the right bottom quarter we expect the starting point to be the first row's first column
			hits, visited, counter = getPossibleCount(m, []int{0, 0}, size/2-1)
		} else if itemKey == "lbq" {
			// for the left bottom quarter we expect the starting point to be the first row's last column
			hits, visited, counter = getPossibleCount(m, []int{0, size - 1}, size/2-1)
		}

		printMap(m, counter, visited)
		tileResults[itemKey] = hits
	}

	oddAmt := int(math.Pow(float64(rowWidth/2*2+1), 2))
	evenAmt := int(math.Pow(float64(((rowWidth + 1) / 2 * 2)), 2))
	amtOfHalfs := rowWidth
	amtOfQuarters := amtOfHalfs + 1

	sum := 0

	for resultKey, result := range tileResults {
		if resultKey == "evenTile" {
			sum += evenAmt * result
		} else if resultKey == "oddTile" {
			sum += oddAmt * result
		} else if resultKey == "rth" || resultKey == "rbh" || resultKey == "lth" || resultKey == "lbh" {
			sum += amtOfHalfs * result
		} else if resultKey == "rtq" || resultKey == "rbq" || resultKey == "ltq" || resultKey == "lbq" {
			sum += amtOfQuarters * result
		} else {
			sum += result
		}
	}

	fmt.Printf("Sum: %+v --------------------- \n", sum)
	return sum
}

func printMap(m []string, counter map[string]int, visited map[string]int) {
	for rI, row := range m {
		fmt.Print("\n")
		for cI, col := range row {
			nodeKey := fmt.Sprintf("%v", []int{rI, cI})
			char := string(col)
			if char == "S" {
				fmt.Print(secondary(char))
			} else if _, ok := counter[nodeKey]; ok {
				fmt.Print(primary("0"))
			} else if _, ok := visited[nodeKey]; ok {
				fmt.Print(tertiary(char))
			} else {
				fmt.Print(white(char))
			}
		}
	}
}
