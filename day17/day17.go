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
	localDistance int
	distance      int
	heat          int
	path          [][]int
}

func pathContainsLocation(path [][]int, location []int) bool {
	for i := 0; i < len(path); i++ {
		if path[i][0] == location[0] && path[i][1] == location[1] {
			return true
		}
	}

	return false
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
			next:          append(start, 1), // one for left to right
			localDistance: 0,
			distance:      0,
			heat:          0,
			path:          [][]int{},
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
		node.path = append(node.path, append(node.next, []int{node.localDistance, heat}...))

		if node.next[0] == len(input)-1 && node.next[1] == len(input[0])-1 {
			fmt.Printf("-----------------------\n")
			fmt.Printf("Found End at %+v, with heat %+v, travel distance: %+v \n", node.next, heat, node.distance)
			fmt.Print(node.path)
			drawTileMap(input, node.path)
			return heat
		}

		for dirIndex := 0; dirIndex < len(directions); dirIndex++ {
			next := append(directionToNext[directions[dirIndex]], directions[dirIndex])
			nodeKey := fmt.Sprintf("%v", next)

			fmt.Print("KEY: ", nodeKey)

			isOutOfBounds := next[0] < 0 || next[1] < 0 || next[0] > len(input)-1 || next[1] > len(input[0])-1
			isInOppositeDirection := directions[dirIndex] == directionOppositionMap[node.next[2]]

			if previousHeat, ok := visited[nodeKey]; ok {
				if previousHeat <= heat {
					continue
				}
			}

			if isOutOfBounds || isInOppositeDirection {
				continue
			}

			localDistance := node.localDistance

			if !pathContainsLocation(node.path, next) {
				node.distance += 1
			}

			if directions[dirIndex] != node.next[2] {
				localDistance = 1
			} else {
				localDistance += 1
			}

			// If localDistance is less than minDistance or greater than maxDistance, skip this direction
			if localDistance < minDistance || localDistance > maxDistance {
				continue
			}

			visited[nodeKey] = heat

			if localDistance > maxDistance && directions[dirIndex] != node.next[2] {
				heap.Push(&pq, &Node{
					next:          next, // change direction
					localDistance: 1,    // reset localDistance
					heat:          heat,
					path:          node.path,
					distance:      node.distance,
				})
			} else {
				heap.Push(&pq, &Node{
					next:          []int{next[0], next[1], node.next[2]}, // keep same direction
					localDistance: localDistance,
					heat:          heat,
					path:          node.path,
					distance:      node.distance,
				})
			}
		}
	}

	panic("No path found")
}

// vis unreliable again, but fun :/
func drawTileMap(input []string, visisted [][]int) {

	// fmt.Printf("-----------------------\n")
	printString := "\n"
	evenTab := "	"

	highlightCount := 0

	for lineIndex := 0; lineIndex < len(input); lineIndex++ {
		line := input[lineIndex]
		printString += evenTab
		for charIndex := 0; charIndex < len(line); charIndex++ {
			char := line[charIndex]
			if pathContainsLocation(visisted, []int{lineIndex, charIndex}) {
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
	// fmt.Print(visisted, "\n")
	fmt.Printf("Highlight Counter (to show diff between viz and actual distance): %+v \n", highlightCount)
	fmt.Printf("-----------------------\n")
}
