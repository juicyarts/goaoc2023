package day10

import (
	"fmt"
	"strings"
)

var instructionMap = map[string][]int{
	"|": {1, 0, 1, 0},
	"-": {0, 1, 0, 1},
	"L": {1, 1, 0, 0},
	"J": {1, 0, 0, 1},
	"7": {0, 0, 1, 1},
	"F": {0, 1, 1, 0},
}

const (
	DirectionTop    int = 0
	DirectionRight  int = 1
	DirectionBottom int = 2
	DirectionLeft   int = 3
)

type Step struct {
	instruction string
	location    []int
	sourceDir   int
}

func getNextStep(maze [][]string, next Step, totalSteps int) (int, Step, [][]string) {
	totalSteps += 1

	if next.instruction == "S" {
		return totalSteps, next, maze
	} else {

		var nextSourceDir int
		var inst = instructionMap[next.instruction]

		var nextLocation = []int{
			next.location[0],
			next.location[1],
		}

		maze[next.location[0]][next.location[1]] = "0"

		for i, value := range inst {
			if value != 0 && i != next.sourceDir {
				if i == DirectionTop {
					nextLocation[0] -= value
					nextSourceDir = DirectionBottom
				} else if i == DirectionRight {
					nextLocation[1] += value
					nextSourceDir = DirectionLeft
				} else if i == DirectionBottom {
					nextLocation[0] += value
					nextSourceDir = DirectionTop
				} else if i == DirectionLeft {
					nextLocation[1] -= value
					nextSourceDir = DirectionRight
				}
				break
			}
		}

		var nextInstruction = maze[nextLocation[0]][nextLocation[1]]
		var nextStep = Step{
			instruction: nextInstruction,
			location:    nextLocation,
			sourceDir:   nextSourceDir,
		}

		// fmt.Printf("nextStep: %+v \n", nextStep)

		return getNextStep(maze, nextStep, totalSteps)
	}
}

func getTotalSteps(maze [][]string, start []int) int {
	var nextStep Step

	var topNeighbour, rightNeighbour, bottomNeighbour, leftNeighbour Step
	if start[0] > 0 {
		var key = maze[start[0]-1][start[1]]
		if key != "." {
			topNeighbour = Step{
				instruction: key,
				location:    []int{start[0] - 1, start[1]},
				sourceDir:   DirectionBottom,
			}
			nextStep = topNeighbour
		}
	}
	if start[0] < len(maze)-1 {
		var key = maze[start[0]+1][start[1]]
		if key != "." {
			bottomNeighbour = Step{
				instruction: key,
				location:    []int{start[0] + 1, start[1]},
				sourceDir:   DirectionTop,
			}
			nextStep = bottomNeighbour
		}
	}

	if start[1] > 0 {
		var key = maze[start[0]][start[1]-1]
		if key != "." {
			leftNeighbour = Step{
				instruction: key,
				location:    []int{start[0], start[1] - 1},
				sourceDir:   DirectionRight,
			}
			nextStep = leftNeighbour
		}
	}

	if start[1] < len(maze[start[0]])-1 {
		var key = maze[start[0]][start[1]+1]
		if key != "." {
			rightNeighbour = Step{
				instruction: key,
				location:    []int{start[0], start[1] + 1},
				sourceDir:   DirectionLeft,
			}

			nextStep = rightNeighbour
		}
	}

	stepsNeeded, _, newMaze := getNextStep(maze, nextStep, 0)

	var totalCount = 0

	for _, row := range newMaze {
		fmt.Printf("%s \n", row)
		for _, col := range row {
			if col == "0" {
				totalCount += 1
			}
		}
	}

	// divide by 2 because we're counting steps to the farthest location and back
	fmt.Printf("%+v %+v \n", stepsNeeded/2, totalCount)
	return stepsNeeded / 2
}

func StepsToFarthestLocation(input []string) int {
	var maze [][]string
	var start []int
	for lineIndex, line := range input {
		newLine := strings.Split(line, "")
		for charIndex, char := range newLine {
			if char == "S" {
				start = []int{lineIndex, charIndex}
			}
		}
		fmt.Printf("%+v \n", newLine)
		maze = append(maze, newLine)
	}

	fmt.Print("found start at: ", start, "\n")
	return getTotalSteps(maze, start)
}
