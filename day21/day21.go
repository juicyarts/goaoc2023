package day21

import (
	"container/heap"
	"fmt"
	"math"

	"github.com/fatih/color"
)

var primary = color.New(color.FgRed).SprintFunc()
var secondary = color.New(color.FgYellow).SprintFunc()

var directions = []int{
	0, 1, 2, 3,
}

type Node struct {
	next      []int
	stepsLeft int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].stepsLeft < pq[j].stepsLeft
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Node)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
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

func Walk(m []string, start []int, stepsLeft int) int {
	var pq = PriorityQueue{
		&Node{
			next:      start,
			stepsLeft: stepsLeft,
		},
	}

	heap.Init(&pq)

	visited := make(map[string]int)
	total := 0
	counter := make(map[string]int)

	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*Node)
		sl := node.stepsLeft - 1
		nodeKey := fmt.Sprintf("%v|%v", node.next, sl)

		if _, ok := visited[nodeKey]; ok {
			continue
		}

		validNeighbours := 0

		var directionToNext = map[int][]int{
			0: {node.next[0] - 1, node.next[1]},
			1: {node.next[0], node.next[1] + 1},
			2: {node.next[0] + 1, node.next[1]},
			3: {node.next[0], node.next[1] - 1},
		}

		for dirIndex := 0; dirIndex < len(directions); dirIndex++ {
			next := directionToNext[directions[dirIndex]]
			normNext := Normalize(next[0], next[1], len(m), len(m[0]))
			nextNodeKey := fmt.Sprintf("%v", next)
			nextChar := m[normNext[0]][normNext[1]]

			if nextChar != '#' {
				if sl == 0 {
					if _, ok := counter[nextNodeKey]; !ok {
						validNeighbours++
						total++
						counter[nextNodeKey] = 1
					}
				} else {
					heap.Push(&pq, &Node{next: next, stepsLeft: sl})
				}
			}
		}

		visited[nodeKey] = validNeighbours
	}

	for lI, l := range m {
		fmt.Print("\n")
		for cI, c := range l {
			nodeKey := fmt.Sprintf("%v", []int{lI, cI})
			if _, ok := counter[nodeKey]; ok {
				fmt.Print(primary(string("0")))
			} else {
				fmt.Print(secondary(string(c)))
			}
		}
	}

	fmt.Print("\n---------------\n")
	return total
}

func Normalize(x, y, xLen, yLen int) []int {
	rx, ry := x, y

	if x >= xLen {
		rx = x % xLen
	}

	if x < 0 {
		if int(math.Abs(float64(x))) >= xLen {
			if (x % xLen) == 0 {
				rx = 0
			} else {
				rx = xLen + (x % xLen)
			}
		} else {
			rx = xLen + x
		}
	}

	if y >= yLen {
		ry = y % yLen
	}

	if y < 0 {
		if int(math.Abs(float64(y))) >= yLen {
			if (y % yLen) == 0 {
				ry = 0
			} else {
				ry = yLen + (y % yLen)
			}
		} else {
			ry = yLen + y
		}
	}

	return []int{rx, ry}
}
