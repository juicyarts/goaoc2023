package day23

import (
	"container/heap"
	"fmt"

	"github.com/fatih/color"
)

var primary = color.New(color.FgHiBlue).SprintFunc()
var tertiary = color.New(color.FgHiRed).SprintFunc()
var secondary = color.New(color.FgYellow).SprintFunc()
var white = color.New(color.FgWhite).SprintFunc()

type Node struct {
	pos   []int
	steps int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].steps > pq[j].steps
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

func Travel(input []string) int {
	return PQTravel(input, []int{0, 1}, []int{len(input) - 1, len(input[0]) - 2}, 0)
}

func TravelB(input []string) int {
	return DFS(input, []int{0, 1}, []int{len(input) - 1, len(input[0]) - 2}, map[string]int{})
}

func DFS(m []string, pos []int, end []int, visited map[string]int) int {
	steps := 0
	nodeKey := fmt.Sprintf("%v", pos)

	if _, ok := visited[nodeKey]; ok {
		visited[nodeKey]++
		return 0
	}

	if pos[0] < 0 || pos[1] < 0 || pos[0] >= len(m) || pos[1] >= len(m[0]) || m[pos[0]][pos[1]] == '#' {
		return 0
	}

	visited[nodeKey] = 1

	if pos[0] == end[0] && pos[1] == end[1] {
		printMap(m, visited)
	}

	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for _, dir := range directions {
		next := []int{pos[0] + dir[0], pos[1] + dir[1]}
		steps = max(steps, DFS(m, next, end, visited))
	}

	return steps + 1
}

// part 1
func PQTravel(m []string, start []int, end []int, steps int) int {
	var pq = PriorityQueue{
		&Node{
			pos:   start,
			steps: steps,
		},
	}

	visited := map[string]int{}
	heap.Init(&pq)
	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*Node)
		nodeKey := fmt.Sprintf("%v", node.pos)

		if node.pos[0] == end[0] && node.pos[1] == end[1] {
			visited[nodeKey] = 1
			printMap(m, visited)
			return node.steps
		}

		if _, ok := visited[nodeKey]; ok || node.pos[0] < 0 || node.pos[1] < 0 || node.pos[0] >= len(m) || node.pos[1] >= len(m[0]) || m[node.pos[0]][node.pos[1]] == '#' {
			continue
		}

		visited[nodeKey] = 1

		directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

		for _, dir := range directions {
			next := []int{node.pos[0] + dir[0], node.pos[1] + dir[1]}
			heap.Push(&pq, &Node{pos: next, steps: node.steps + 1})
		}

	}

	panic("Cannot cross it")
}

func printMap(m []string, visited map[string]int) {
	for rI, row := range m {
		fmt.Print("\n")
		for cI, col := range row {
			nodeKey := fmt.Sprintf("%v", []int{rI, cI})
			char := string(col)
			if char == "S" {
				fmt.Print(secondary(char))
			} else if _, ok := visited[nodeKey]; ok {
				if visited[nodeKey] > 0 {
					fmt.Print(tertiary("0"))
				} else {
					fmt.Print(tertiary(char))
				}
			} else {
				fmt.Print(primary(char))
			}
		}
	}
	fmt.Print("\n")
}
