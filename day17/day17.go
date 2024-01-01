package day17

import (
	"container/heap"
	"fmt"
	"strconv"

	"github.com/fatih/color"
)

var primary = color.New(color.FgHiBlue).SprintFunc()
var tertiary = color.New(color.FgHiRed).SprintFunc()

type Node struct {
	pos   []int
	dir   int
	steps int
	heat  int
	path  [][]int
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
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

func Travel(m []string, start []int, minSteps, maxSteps, dir int) int {

	var pq = PriorityQueue{
		&Node{
			pos:   start,
			dir:   dir,
			steps: 0,
			heat:  0,
		},
	}

	heap.Init(&pq)

	visited := make(map[string]int)

	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*Node)
		nodeKey := fmt.Sprintf("%v", node.pos)

		if _, ok := visited[nodeKey]; ok || node.pos[0] < 0 || node.pos[1] < 0 || node.pos[0] >= len(m) || node.pos[1] >= len(m[0]) {
			continue
		}

		if node.steps > maxSteps {
			continue
		}

		char := m[node.pos[0]][node.pos[1]]
		charAsInt, _ := strconv.Atoi(string(char))
		heat := node.heat + charAsInt
		visited[nodeKey] = heat
		fmt.Print("TEST: ", heat, " ", charAsInt, " ", string(char), "\n")

		if node.pos[0] == len(m)-1 && node.pos[1] == len(m[0])-1 {
			fmt.Printf("-----------------------\n")
			fmt.Printf("Found End at %+v, with heat %+v, travel distance: %+v \n", node.pos, heat, node.steps)
			drawTileMap(m, visited)
			return heat
		}

		directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

		// keep moving in same dir
		if node.steps < minSteps {
			next := []int{node.pos[0] + directions[node.dir][0], node.pos[1] + directions[node.dir][1]}
			heap.Push(&pq, &Node{pos: next, heat: heat, steps: node.steps + 1, dir: node.dir})
			continue
		}

		for dirIndex, dir := range directions {
			next := []int{node.pos[0] + dir[0], node.pos[1] + dir[1]}
			if node.dir == dirIndex && node.steps > maxSteps {
				continue
			}

			if node.dir == dirIndex && node.steps < minSteps {
				heap.Push(&pq, &Node{pos: next, heat: heat, steps: node.steps + 1, dir: dirIndex})
				continue
			}

			heap.Push(&pq, &Node{pos: next, heat: heat, steps: 1, dir: dirIndex})

		}

		drawTileMap(m, visited)
	}

	panic("No path found")
}

func drawTileMap(input []string, visited map[string]int) {

	// fmt.Printf("-----------------------\n")
	printString := "\n"
	evenTab := "	"

	highlightCount := 0

	for rI := 0; rI < len(input); rI++ {
		line := input[rI]
		printString += evenTab
		for cI := 0; cI < len(line); cI++ {
			char := line[cI]
			if _, ok := visited[fmt.Sprintf("%v", []int{rI, cI})]; ok {
				printString += tertiary(string(char))
				highlightCount += 1
			} else {
				printString += primary(string(char))
			}
		}
		printString += "\n"
	}
	fmt.Print(printString)
	fmt.Print("\n")
	// fmt.Print(visited, "\n")
	fmt.Printf("Highlight Counter (to show diff between viz and actual distance): %+v \n", highlightCount)
	fmt.Printf("-----------------------\n")
}
