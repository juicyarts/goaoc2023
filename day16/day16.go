package day16

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

var primary = color.New(color.FgHiWhite).SprintFunc()
var secondary = color.New(color.FgHiBlack).SprintFunc()
var tertiary = color.New(color.FgHiRed).SprintFunc()

var forwardSlashMap = map[int]int{
	0: 1,
	1: 0,
	2: 3,
	3: 2,
}

var backwardSlashMap = map[int]int{
	0: 3,
	1: 2,
	2: 1,
	3: 0,
}

func Travel(input []string, start []int, direction int, tilesVisited map[string]int) map[string]int {
	var directionToNext = map[int][]int{
		0: {start[0] - 1, start[1]},
		1: {start[0], start[1] + 1},
		2: {start[0] + 1, start[1]},
		3: {start[0], start[1] - 1},
	}

	next := directionToNext[direction]
	tileKey := fmt.Sprintf("%d|%d", start[0], start[1])

	if _, ok := tilesVisited[tileKey]; ok {
		if tilesVisited[tileKey] == direction {
			return tilesVisited
		}
	} else {
		tilesVisited[tileKey] = direction
	}

	if next[0] < 0 || next[0] > len(input)-1 || next[1] < 0 || next[1] > len(input[0])-1 {
		return tilesVisited
	}

	char := input[next[0]][next[1]]

	if char == '.' ||
		direction == 1 && char == '-' ||
		direction == 3 && char == '-' ||
		direction == 0 && char == '|' ||
		direction == 2 && char == '|' {
		tilesVisited = Travel(input, next, direction, tilesVisited)
	} else if char == '/' {
		tilesVisited = Travel(input, next, forwardSlashMap[direction], tilesVisited)
	} else if char == '\\' {
		tilesVisited = Travel(input, next, backwardSlashMap[direction], tilesVisited)
	} else if char == '|' {
		tilesVisited = Travel(input, next, 0, tilesVisited)
		tilesVisited = Travel(input, next, 2, tilesVisited)
	} else if char == '-' {
		tilesVisited = Travel(input, next, 3, tilesVisited)
		tilesVisited = Travel(input, next, 1, tilesVisited)
	}

	return tilesVisited
}

func MeasureEnergizedTiles(input []string) int {
	configurationMap := map[string]map[string]int{}
	maxTilesVisited := 0
	winningConfiguration := ""
	winningConfigurationTiles := map[string]int{}

	conditions := [][][]int{
		{{0, 0}, {1}},
		{{0, len(input[0]) - 1}, {3, 2}},
		{{len(input) - 1, 0}, {0, 1}},
		{{len(input) - 1, len(input[0]) - 1}, {3, 0}},
	}

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {

			var directionToStartPositionMap = map[int][]int{
				2: {i - 1, j},
				3: {i, j + 1},
				0: {i + 1, j},
				1: {i, j - 1},
			}

			configurationKey := fmt.Sprintf("%d|%d", i, j)
			for condition := 0; condition < len(conditions); condition++ {
				if i == conditions[condition][0][0] && j == conditions[condition][0][1] {
					directions := conditions[condition][1]
					for i := 0; i < len(directions); i++ {
						ownKey := configurationKey + "|" + strconv.Itoa(directions[i])
						tiles := Travel(input, directionToStartPositionMap[directions[i]], directions[i], map[string]int{})
						configurationMap[ownKey] = tiles
						if len(tiles) > maxTilesVisited {
							maxTilesVisited = len(tiles)
							winningConfiguration = ownKey
							winningConfigurationTiles = tiles
						}
					}
				}
			}

			direction := -1

			// First row From Top to Bottom
			if i == 0 && j > 0 && j < len(input[0])-1 {
				direction = 2
			}
			// Last column from Left to Right
			if i > 0 && i < len(input)-1 && j == len(input[0])-1 {
				direction = 3
			}
			// Last row from Bottom to Top
			if i == len(input)-1 && j > 0 && j < len(input[0])-1 {
				direction = 0
			}
			// First column from Right to Left
			if i > 0 && i < len(input)-1 && j == 0 {
				direction = 1
			}

			if direction == -1 {
				continue
			}

			ownKey := configurationKey + "|" + strconv.Itoa(direction)
			tiles := Travel(input, directionToStartPositionMap[direction], direction, map[string]int{})
			configurationMap[ownKey] = tiles
			if len(tiles) > maxTilesVisited {
				maxTilesVisited = len(tiles)
				winningConfiguration = ownKey
				winningConfigurationTiles = tiles
			}
		}
	}

	fmt.Printf("Winner: %+v \n", winningConfiguration)
	drawTileMap(winningConfiguration, input, winningConfigurationTiles)
	return maxTilesVisited - 1
}

// drawing is off ¯\_(ツ)_/¯
func drawTileMap(name string, input []string, tilesVisited map[string]int) {
	startX, _ := strconv.Atoi(strings.Split(name, "|")[0])
	startY, _ := strconv.Atoi(strings.Split(name, "|")[1])

	fmt.Printf("Route For: %+v -----------------------\n", name)
	printString := "\n"
	eventTab := "	"

	for lineIndex := 0; lineIndex < len(input); lineIndex++ {
		line := input[lineIndex]
		printString += eventTab
		for charIndex := 0; charIndex < len(line); charIndex++ {
			char := line[charIndex]
			if tilesVisited[fmt.Sprintf("%d|%d", lineIndex, charIndex)] > 0 {
				if lineIndex == startX && charIndex == startY {
					printString += tertiary(string(char))
				} else {
					printString += primary(string(char))
				}
			} else {
				printString += secondary(string(char))
			}
		}

		printString += "\n"
	}
	fmt.Print(printString)
	fmt.Print("\n")
	fmt.Printf("Energized: %+v Tiles ------ \n", len(tilesVisited))
}
