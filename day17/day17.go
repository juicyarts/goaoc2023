package day17

import (
	"fmt"
	"strconv"
	"strings"

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


var maxDistanceInSameDir = 3

func Travel(input []string, start []int, end []int, direction int, distanceInSameDir int, totalHeat int, tileKey string) int {

	var tilesVisited = make(map[string]int)

	tileKey += fmt.Sprintf("%d|%d#", start[0], start[1])
	ownKey := fmt.Sprintf("%d|%d|%d", start[0], start[1], direction)

	char := input[start[0]][start[1]]
	charAsInt, _ := strconv.Atoi(string(char))
	totalHeat += charAsInt

	if _, ok := tilesVisited[tileKey]; ok {
		return tilesVisited[tileKey]
	} else {
		tilesVisited[tileKey] = totalHeat
	}

	if _, ok := tilesVisited[ownKey]; ok {
		return tilesVisited[ownKey]
	} else {
		tilesVisited[ownKey] = totalHeat
	}

	var directionToNext = map[int][]int{
		0: {start[0] - 1, start[1]},
		1: {start[0], start[1] + 1},
		2: {start[0] + 1, start[1]},
		3: {start[0], start[1] - 1},
	}

	if start[0] == end[0] && start[1] == end[1] {
		// fmt.Print("----------------------- \n")
		// fmt.Printf("The End Has Been Reached With Heat:%+v \n", totalHeat)
		drawTileMap(tileKey, input)
		return totalHeat
	}

	// fmt.Printf("Coming from %+v, Trying To Reach:%+v, Direction: %+v, DistanceInsameDir: %+v \n", start, end, oldDirection, distanceInSameDir)

	minHeat := 0

	for dirIndex := 0; dirIndex < len(directions)-1; dirIndex++ {
		next := directionToNext[directions[dirIndex]]

		if directions[dirIndex] == directionOppositionMap[direction] ||
			distanceInSameDir >= maxDistanceInSameDir && directions[direction] == direction ||
			directions[dirIndex] == 0 && next[0] < 0 || directions[dirIndex] == 3 && next[1] < 0 ||
			directions[dirIndex] == 1 && next[1] > len(input[0])-1 ||
			directions[dirIndex] == 2 && next[0] > len(input)-1 {
			continue
		}

		localDistance := distanceInSameDir

		if directions[dirIndex] != direction {
			localDistance = 1
		} else {
			localDistance += 1
		}

		heat := Travel(input, next, end, directions[dirIndex], localDistance, totalHeat, tileKey)

		if heat < minHeat && heat > 0 || minHeat == 0 {
			minHeat = heat
		}
	}

	return minHeat
}

func drawTileMap(name string, input []string) {

	foo := strings.Split(name, "#")

	// fmt.Printf("-----------------------\n")
	printString := "\n"
	evenTab := "	"

	for lineIndex := 0; lineIndex < len(input); lineIndex++ {
		line := input[lineIndex]
		printString += evenTab

		for charIndex := 0; charIndex < len(line); charIndex++ {
			char := line[charIndex]
			foundMatch := false

			for i := 0; i < len(foo)-1; i++ {
				splitFoo := strings.Split(foo[i], "|")
				x, _ := strconv.Atoi(splitFoo[0])
				y, _ := strconv.Atoi(splitFoo[1])
				if x == lineIndex && y == charIndex {
					printString += tertiary(string(char))
					foundMatch = true
					break
				}
			}

			if !foundMatch {
				printString += primary(string(char))
			}
		}
		printString += "\n"
	}
	fmt.Print(printString)
	fmt.Print("\n")
	// fmt.Print(name, "\n")
	// fmt.Printf("Passed Through: %+v Tiles \n", len(foo))
	// fmt.Printf("-----------------------\n")
}
