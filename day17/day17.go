package day17

import (
	"container/heap"
	"fmt"
	"strconv"

	"github.com/fatih/color"
)

var primary = color.New(color.FgHiBlue).SprintFunc()
var tertiary = color.New(color.FgHiRed).SprintFunc()

var directions = []int{
	0, 1, 2, 3,
}

var directionOppositionMap = map[int]int{
	0: 2,
	1: 3,
	3: 1,
	2: 0,
}

type Node struct {
	next          []int
	direction     int
	localDistance int
	heat          int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].heat < pq[j].heat
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
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func Travel(input []string, start []int, minDistance, maxDistance int) int {

	var pq = PriorityQueue{
		&Node{
			next:          start,
			direction:     1,
			localDistance: 0,
			heat:          0,
		},
	}

	heap.Init(&pq)

	visited := make(map[string]int)

	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*Node)

		var directionToNext = map[int][]int{
			0: {node.next[0] - 1, node.next[1]},
			1: {node.next[0], node.next[1] + 1},
			2: {node.next[0] + 1, node.next[1]},
			3: {node.next[0], node.next[1] - 1},
		}

		char := input[node.next[0]][node.next[1]]
		charAsInt, _ := strconv.Atoi(string(char))
		heat := node.heat + charAsInt

		if node.next[0] == len(input)-1 && node.next[1] == len(input[0])-1 {
			fmt.Printf("Found End at %+v, with heat %+v \n", node.next, heat)
			return heat
		}

		for dirIndex := 0; dirIndex < len(directions); dirIndex++ {
			next := directionToNext[directions[dirIndex]]
			nodeKey := fmt.Sprintf("%v%v", next, directions[dirIndex])

			if previousHeat, ok := visited[nodeKey]; ok {
				if previousHeat <= heat {
					continue
				}
			}

			// if direction is opposite of previous direction, skip
			if directions[dirIndex] == directionOppositionMap[node.direction] {
				continue
			}

			// if direction is same as previous direction and local distance is greater than max distance, skip
			if node.localDistance > maxDistance {
				continue
			}

			// if input borders are reached, skip
			if directions[dirIndex] == 0 && next[0] < 0 || directions[dirIndex] == 3 && next[1] < 0 ||
				directions[dirIndex] == 1 && next[1] > len(input[0])-1 ||
				directions[dirIndex] == 2 && next[0] > len(input)-1 {
				continue
			}

			visited[nodeKey] = heat
			localDistance := node.localDistance

			if directions[dirIndex] != node.direction {
				localDistance = 0
			} else {
				localDistance += 1
			}

			node := &Node{
				next:          next,
				direction:     directions[dirIndex],
				localDistance: localDistance,
				heat:          heat,
			}

			heap.Push(&pq, node)
		}

	}

	panic("No path found")
}
