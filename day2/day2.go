package day2

import (
	"fmt"
	"strconv"
	"strings"
)

type CubeSet struct {
	blue  int
	red   int
	green int
}

type Game struct {
	id                int
	sets              []CubeSet
	minmumCubesNeeded CubeSet
}

func getCubeAmountByColorSubstring(s string, color string) (int, error) {
	return strconv.Atoi(strings.Replace(s, fmt.Sprintf(" %s", color), "", -1))
}

func InputToGame(input string) (game Game) {
	var id int
	id, _ = strconv.Atoi(strings.Replace(strings.Split(input, ":")[0], "Game ", "", -1))

	game = Game{
		id: id,
	}

	inputWithoutId := strings.Replace(input, "Game "+strconv.Itoa(id)+": ", "", -1)
	sets := strings.Split(inputWithoutId, "; ")

	for _, set := range sets {
		var newSet CubeSet
		for _, colorSubstring := range strings.Split(set, ", ") {
			if strings.Contains(colorSubstring, "blue") {
				newSet.blue, _ = getCubeAmountByColorSubstring(colorSubstring, "blue")
			}
			if strings.Contains(colorSubstring, "red") {
				newSet.red, _ = getCubeAmountByColorSubstring(colorSubstring, "red")
			}
			if strings.Contains(colorSubstring, "green") {
				newSet.green, _ = getCubeAmountByColorSubstring(colorSubstring, "green")
			}
		}

		if newSet.blue > game.minmumCubesNeeded.blue {
			game.minmumCubesNeeded.blue = newSet.blue
		}
		if newSet.red > game.minmumCubesNeeded.red {
			game.minmumCubesNeeded.red = newSet.red
		}
		if newSet.green > game.minmumCubesNeeded.green {
			game.minmumCubesNeeded.green = newSet.green
		}

		game.sets = append(game.sets, newSet)
	}

	return game
}

func (game Game) IsPossible(availableCubes CubeSet) bool {
	for _, set := range game.sets {
		if availableCubes.blue < set.blue || availableCubes.red < set.red || availableCubes.green < set.green {
			return false
		}
	}
	return true
}

func (game Game) PowerOfMinimumSetOfCubes() int {
	return game.minmumCubesNeeded.blue * game.minmumCubesNeeded.red * game.minmumCubesNeeded.green
}

func SumOfPossibleGameIds(input []string, availableCubes CubeSet) (sum int) {
	for _, gameInput := range input {
		game := InputToGame(gameInput)
		if game.IsPossible(availableCubes) {
			sum += game.id
		}
	}

	return sum
}

// Part 2
func SumOfPowerOfMinimumCubeSetsOfGames(input []string) (sum int) {
	for _, gameInput := range input {
		game := InputToGame(gameInput)
		sum += game.PowerOfMinimumSetOfCubes()
	}

	return sum
}
