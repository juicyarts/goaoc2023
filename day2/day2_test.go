package day2

import (
	"aoc2023/utils"
	"reflect"
	"slices"
	"testing"
)

var testInput = []string{
	"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
	"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
	"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
	"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
	"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
}

func TestInputToGame(t *testing.T) {
	var expectedResults = []Game{
		{
			id: 1,
			sets: []CubeSet{
				{blue: 3, red: 4},
				{red: 1, green: 2, blue: 6},
				{green: 2},
			},
			minmumCubesNeeded: CubeSet{
				red:   4,
				green: 2,
				blue:  6,
			},
		},
		{
			id: 2,
			sets: []CubeSet{
				{blue: 1, green: 2},
				{green: 3, blue: 4, red: 1},
				{green: 1, blue: 1},
			},
			minmumCubesNeeded: CubeSet{
				red:   1,
				green: 3,
				blue:  4,
			},
		},
		{
			id: 3,
			sets: []CubeSet{
				{green: 8, blue: 6, red: 20},
				{blue: 5, red: 4, green: 13},
				{green: 5, red: 1},
			},
			minmumCubesNeeded: CubeSet{
				red:   20,
				green: 13,
				blue:  6,
			},
		},
		{
			id: 4,
			sets: []CubeSet{
				{green: 1, red: 3, blue: 6},
				{green: 3, red: 6},
				{green: 3, blue: 15, red: 14},
			},
			minmumCubesNeeded: CubeSet{
				red:   14,
				green: 3,
				blue:  15,
			},
		},
		{
			id: 5,
			sets: []CubeSet{
				{red: 6, blue: 1, green: 3},
				{blue: 2, red: 1, green: 2},
			},
			minmumCubesNeeded: CubeSet{
				red:   6,
				green: 3,
				blue:  2,
			},
		},
	}

	for index, input := range testInput {
		game := InputToGame(input)
		if !reflect.DeepEqual(game, expectedResults[index]) {
			t.Errorf("Expected game to be %v, got %v", expectedResults[index], game)
		}
	}
}

func TestIsGamePossible(t *testing.T) {
	var availableCubes = CubeSet{blue: 14, red: 12, green: 13}
	for _, input := range testInput {
		game := InputToGame(input)
		if slices.Contains([]int{1, 2, 5}, game.id) {
			if !game.IsPossible(availableCubes) {
				t.Errorf("Expected game %d to be possible, got impossible", game.id)
			}
		} else {
			if game.IsPossible(availableCubes) {
				t.Errorf("Expected game %d to be impossible, got impossible", game.id)
			}
		}
	}
}

func TestSumOfPossibleGameIds(t *testing.T) {
	var availableCubes = CubeSet{blue: 14, red: 12, green: 13}

	var sumOfPossibleGameIds = SumOfPossibleGameIds(testInput, availableCubes)
	var expectedSum = 8

	if sumOfPossibleGameIds != expectedSum {
		t.Errorf("Expected sum of possible game ids to be %d, got %d", expectedSum, sumOfPossibleGameIds)
	}
}

func TestSumOfPossibleGameIdsWithInput(t *testing.T) {
	Input, _ := utils.ReadInputFile("input.txt")
	var availableCubes = CubeSet{blue: 14, red: 12, green: 13}

	var sumOfPossibleGameIds = SumOfPossibleGameIds(Input, availableCubes)
	var expectedSum = 2810

	if sumOfPossibleGameIds != expectedSum {
		t.Errorf("Expected sum of possible game ids to be %d, got %d", expectedSum, sumOfPossibleGameIds)
	}
}

// Part 2
func TestPowerOfMinimumSetOfCubes(t *testing.T) {
	var expectedPowers = []int{48, 12, 1560, 630, 36}

	for index, input := range testInput {
		game := InputToGame(input)
		actualPower := game.PowerOfMinimumSetOfCubes()

		if actualPower != expectedPowers[index] {
			t.Errorf("Expected power of minimum set of cubes to be %d, got %d", expectedPowers[index], actualPower)
		}
	}
}

func TestSumOfPowerOfMinimumCubeSetsOfGames(t *testing.T) {
	var sumOfPowerOfGames = SumOfPowerOfMinimumCubeSetsOfGames(testInput)
	var expectedSum = 2286

	if sumOfPowerOfGames != expectedSum {
		t.Errorf("Expected sum of power of games to be %d, got %d", expectedSum, sumOfPowerOfGames)
	}
}

func TestSumOfPowerOfMinimumCubeSetsOfGamesWithInput(t *testing.T) {
	Input, _ := utils.ReadInputFile("input.txt")
	var sumOfPowerOfGames = SumOfPowerOfMinimumCubeSetsOfGames(Input)
	var expectedSum = 69110

	if sumOfPowerOfGames != expectedSum {
		t.Errorf("Expected sum of power of games to be %d, got %d", expectedSum, sumOfPowerOfGames)
	}
}
