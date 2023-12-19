package day16

import (
	"fmt"

	"github.com/fatih/color"
)

var highlight = color.New(color.FgRed).SprintFunc()

func Travel(input []string, start []int, direction int, tilesVisited map[string]int) map[string]int {
	next := []int{}

	if direction == 0 {
		next = []int{start[0] - 1, start[1]}
	} else if direction == 1 {
		next = []int{start[0], start[1] + 1}
	} else if direction == 2 {
		next = []int{start[0] + 1, start[1]}
	} else if direction == 3 {
		next = []int{start[0], start[1] - 1}
	}

	if _, ok := tilesVisited[fmt.Sprintf("%d|%d", start[0], start[1])]; ok {
		if tilesVisited[fmt.Sprintf("%d|%d", start[0], start[1])] == direction {
			// fmt.Print("Travel endedn since we are in the same loop")
			return tilesVisited
		}
	} else {
		tilesVisited[fmt.Sprintf("%d|%d", start[0], start[1])] = direction
	}

	if next[0] < 0 || next[0] > len(input)-1 || next[1] < 0 || next[1] > len(input[0])-1 {
		// fmt.Printf("Coming to an end at: %+v \n", start)
		return tilesVisited
	}

	char := input[next[0]][next[1]]

	// fmt.Printf("Current: %+v %+v \n", start, string(input[start[0]][start[1]]))
	// fmt.Printf("Next: %+v %+v \n", next, string(char))

	if char == '.' || direction == 1 && char == '-' ||
		direction == 3 && char == '-' ||
		direction == 0 && char == '|' ||
		direction == 2 && char == '|' {
		tilesVisited = Travel(input, next, direction, tilesVisited)
	} else if char == '/' {
		// fmt.Print("Next Char forces a direction change! \n")
		// fmt.Printf("Old: %+v \n", direction)
		if direction == 0 {
			direction = 1
		} else if direction == 1 {
			direction = 0
		} else if direction == 2 {
			direction = 3
		} else if direction == 3 {
			direction = 2
		}
		tilesVisited = Travel(input, next, direction, tilesVisited)
	} else if char == '\\' {
		// fmt.Print("Next Char forces a direction change! \n")
		// fmt.Printf("Old: %+v \n", direction)
		if direction == 0 {
			direction = 3
		} else if direction == 1 {
			direction = 2
		} else if direction == 2 {
			direction = 1
		} else if direction == 3 {
			direction = 0
		}
		tilesVisited = Travel(input, next, direction, tilesVisited)
	} else if char == '|' {
		// fmt.Printf("Branching up/down \n")

		// fmt.Printf("Branch Up | : %+v %+v %+v \n", next, 0, tilesVisited)
		// fmt.Printf("Branch Down | : %+v %+v %+v \n", next, 2, tilesVisited)

		tilesVisited = Travel(input, next, 0, tilesVisited)
		tilesVisited = Travel(input, next, 2, tilesVisited)
	} else if char == '-' {
		// fmt.Printf("Branching left/right \n")

		// fmt.Printf("Branch Left | : %+v %+v %+v \n", next, 3, tilesVisited)
		// fmt.Printf("Branch Right | : %+v %+v %+v \n", next, 1, tilesVisited)

		tilesVisited = Travel(input, next, 3, tilesVisited)
		tilesVisited = Travel(input, next, 1, tilesVisited)
	}

	return tilesVisited
}

func MeasureEnergizedTiles(input []string) int {
	tilesVisited := Travel(input, []int{0, 0}, 2, map[string]int{})

	fmt.Print("\n")
	for lineIndex, line := range input {
		for charIndex, char := range line {
			if tilesVisited[fmt.Sprintf("%d|%d", lineIndex, charIndex)] > 0 {
				fmt.Print(highlight(string(char)))
			} else {
				fmt.Print(string(char))
			}
		}
		fmt.Print("\n")
	}

	fmt.Print("\n")
	fmt.Printf("Energized: %+v Tiles ------ \n", len(tilesVisited))
	fmt.Print("\n")

	return len(tilesVisited)
}
