package day11

import (
	"fmt"
	"sort"
	"strings"
)

func GetSumOfSteps(input []string, expansionFactor int) int {
	var galaxiesInColumn = make([]int, len(input))
	var galaxiesInRow = make([]int, len(input))
	var galaxies = map[int][]int{}

	fmt.Printf("GALAXY SEARCH")
	fmt.Printf("\n------------------------------------------\n")

	for lineIndex := 0; lineIndex < len(input); lineIndex++ {
		newLine := strings.Split(input[lineIndex], "")

		for columnIndex, char := range newLine {
			if char == "#" {
				galaxiesInRow[lineIndex] += 1
				galaxiesInColumn[columnIndex] += 1
				location := []int{lineIndex, columnIndex}
				galaxies[len(galaxies)+1] = location
			} else {
				galaxiesInRow[lineIndex] += 0
				galaxiesInColumn[columnIndex] += 0
			}
		}

		fmt.Printf("%+v %+v\n", newLine, galaxiesInRow[lineIndex])
	}

	var traveledList = map[string]int{}
	var totalSteps = 0

	for sourceGalaxyKey, sourceGalaxy := range galaxies {
		for targetGalaxyKey, targetGalaxy := range galaxies {
			var hasNotBeenTraveledTo = traveledList[fmt.Sprintf("%d,%d", sourceGalaxyKey, targetGalaxyKey)] == 0 && traveledList[fmt.Sprintf("%d,%d", targetGalaxyKey, sourceGalaxyKey)] == 0
			if sourceGalaxyKey != targetGalaxyKey && hasNotBeenTraveledTo {
				var routeX = []int{sourceGalaxy[0], targetGalaxy[0]}
				var routeY = []int{sourceGalaxy[1], targetGalaxy[1]}
				sort.Ints(routeX)
				sort.Ints(routeY)

				var xRange = routeX[1] - routeX[0]
				var yRange = routeY[1] - routeY[0]

				xDist, yDist := 0, 0

				for i := 0; i < xRange+1; i++ {
					var x = routeX[0] + i

					if i < xRange {
						if galaxiesInRow[x] == 0 {
							xDist += expansionFactor
						} else {
							xDist += 1
						}
					}
				}

				for i := yRange; i > 0; i-- {
					if galaxiesInColumn[routeY[1]-i] == 0 {
						yDist += expansionFactor
					} else {
						yDist += 1
					}
				}

				steps := xDist + yDist
				traveledList[fmt.Sprintf("%d,%d", sourceGalaxyKey, targetGalaxyKey)] = steps
				totalSteps += steps
			}
		}
	}

	fmt.Printf("------------------------------------------\n")
	fmt.Printf("Travel list has length of %+v \n", len(traveledList))
	fmt.Printf("Travelled total distance of %+v \n", totalSteps)
	fmt.Printf("------------------------------------------\n")

	return totalSteps

}
