package day2

import (
	"strconv"
	"strings"
)

type CubeSet struct {
	blue  int
	red   int
	green int
}

type Game struct {
	id   int
	sets []CubeSet
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
		var blue, red, green int
		for _, color := range strings.Split(set, ", ") {
			if strings.Contains(color, "blue") {
				blue, _ = strconv.Atoi(strings.Replace(color, " blue", "", -1))
			}
			if strings.Contains(color, "red") {
				red, _ = strconv.Atoi(strings.Replace(color, " red", "", -1))
			}
			if strings.Contains(color, "green") {
				green, _ = strconv.Atoi(strings.Replace(color, " green", "", -1))
			}
		}
		game.sets = append(game.sets, CubeSet{
			blue:  blue,
			red:   red,
			green: green,
		})
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

func GetSumOfPossibleGameIds(input []string, availableCubes CubeSet) (sum int) {
	for _, gameInput := range input {
		game := InputToGame(gameInput)
		if game.IsPossible(availableCubes) {
			sum += game.id
		}
	}

	return sum
}