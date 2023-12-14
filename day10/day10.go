package day10

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
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

func getNextStep(maze [][]string, next Step, totalSteps int, stepMap map[string]int) (int, Step, map[string]int) {
	totalSteps += 1

	stepMap[fmt.Sprintf("%d|%d", next.location[0], next.location[1])] = 1

	if next.instruction == "S" {
		return totalSteps, next, stepMap
	} else {

		var nextSourceDir int
		var inst = instructionMap[next.instruction]

		var nextLocation = []int{
			next.location[0],
			next.location[1],
		}

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

		return getNextStep(maze, nextStep, totalSteps, stepMap)
	}
}

func getTotalSteps(maze [][]string, start []int) int {
	var nextStep Step

	var topNeighbour, rightNeighbour, bottomNeighbour, leftNeighbour Step
	if start[0] > 0 {
		var key = maze[start[0]-1][start[1]]
		if key != "." && key != "-" && key != "L" && key != "J" {
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
		if key != "." && key != "-" && key != "F" && key != "J" && key != "7" {
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
		if key != "." && key != "|" && key != "7" && key != "J" {
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
		if key != "." && key != "|" && key != "F" && key != "L" {
			rightNeighbour = Step{
				instruction: key,
				location:    []int{start[0], start[1] + 1},
				sourceDir:   DirectionLeft,
			}

			nextStep = rightNeighbour
		}
	}

	stepsNeeded, _, stepMap := getNextStep(maze, nextStep, 0, make(map[string]int, 0))

	// Just for visualizing the maze
	for key, value := range maze {
		for index, char := range value {
			foo := fmt.Sprintf("%d|%d", key, index)
			yellow := color.New(color.FgYellow).SprintFunc()
			red := color.New(color.FgRed).SprintFunc()
			white := color.New(color.FgWhite).SprintFunc()

			if char == "S" {
				fmt.Printf("%s", red(char))
			} else if stepMap[foo] == 1 {
				if char == "L" {
					char = "└"
				}
				if char == "J" {
					char = "┘"
				}
				if char == "7" {
					char = "┐"
				}
				if char == "F" {
					char = "┌"
				}
				if char == "-" {
					char = "─"
				}
				if char == "|" {
					char = "│"
				}

				fmt.Printf("%s", yellow(char))
			} else {
				fmt.Printf("%s", white(char))

			}

		}

		fmt.Print("\n")
	}

	// divide by 2 because we're counting steps to the farthest location and back
	fmt.Printf("---- \n steps needed: %+v \n", stepsNeeded/2)
	fmt.Printf("---- \n steps map: %+v \n", stepMap)
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
		// fmt.Printf("%+v \n", newLine)
		maze = append(maze, newLine)
	}
	return getTotalSteps(maze, start)
}
