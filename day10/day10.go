package day10

import (
	"fmt"
	"sort"
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

func getNextStep(maze [][]string, next Step, totalSteps int, stepMap map[string]int, rowMap map[int][]int) (int, Step, map[string]int, map[int][]int) {
	totalSteps += 1

	if next.instruction == "|" || next.instruction == "F" || next.instruction == "7" {
		if rowMap[next.location[0]] == nil {
			rowMap[next.location[0]] = []int{next.location[1]}
		} else {
			rowMap[next.location[0]] = append(rowMap[next.location[0]], next.location[1])
		}
	}

	stepMap[fmt.Sprintf("%d|%d", next.location[0], next.location[1])] = 1

	if next.instruction == "S" {
		return totalSteps, next, stepMap, rowMap
	} else {

		var nextSourceDir int

		var nextLocation = []int{
			next.location[0],
			next.location[1],
		}

		for i, value := range instructionMap[next.instruction] {
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

		var nextStep = Step{
			instruction: maze[nextLocation[0]][nextLocation[1]],
			location:    nextLocation,
			sourceDir:   nextSourceDir,
		}

		return getNextStep(maze, nextStep, totalSteps, stepMap, rowMap)
	}
}

func getTotalSteps(maze [][]string, start []int) (int, int) {
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

	var initialRowMap = make(map[int][]int, 0)
	initialRowMap[start[0]] = []int{start[1]}

	stepsNeeded, _, stepMap, rowMap := getNextStep(maze, nextStep, 0, make(map[string]int, 0), initialRowMap)
	var enclosedTiles = 0

	for rowIndex, value := range maze {
		for charIndex, char := range value {
			foo := fmt.Sprintf("%d|%d", rowIndex, charIndex)
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
				if rowMap[rowIndex] != nil {
					sort.Ints(rowMap[rowIndex])

					for rowMapIndex := range rowMap[rowIndex] {
						if rowMapIndex > 0 {
							if rowMapIndex%2 != 0 {

								if charIndex > rowMap[rowIndex][rowMapIndex-1] && charIndex < rowMap[rowIndex][rowMapIndex] {
									enclosedTiles += 1
									char = "0"
									break
								}
							}
						}
					}
				}
				fmt.Printf("%s", white(char))
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
	return stepsNeeded / 2, enclosedTiles
}

func StepsToFarthestLocation(input []string) (int, int) {
	var maze [][]string
	var start []int
	for lineIndex, line := range input {
		newLine := strings.Split(line, "")
		for charIndex, char := range newLine {
			if char == "S" {
				start = []int{lineIndex, charIndex}
			}
		}
		maze = append(maze, newLine)
	}
	return getTotalSteps(maze, start)
}
