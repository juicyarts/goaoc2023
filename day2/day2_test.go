package day2

import (
	"bufio"
	"os"
	"reflect"
	"testing"
)

func ReadInputFile(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

var testInput = []string{
	"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
	"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
	"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
	"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
	"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
}

func TestInputToGame(t *testing.T) {
	gameOne := InputToGame(testInput[0])
	wantOne := Game{
		id: 1,
		sets: []CubeSet{
			{blue: 3, red: 4},
			{red: 1, green: 2, blue: 6},
			{green: 2},
		},
	}

	if !reflect.DeepEqual(gameOne, wantOne) {
		t.Errorf("Expected game to be %v, got %v", wantOne, gameOne)
	}

	gameTwo := InputToGame(testInput[1])
	wantTwo := Game{
		id: 2,
		sets: []CubeSet{
			{blue: 1, green: 2},
			{green: 3, blue: 4, red: 1},
			{green: 1, blue: 1},
		},
	}
	if !reflect.DeepEqual(gameTwo, wantTwo) {
		t.Errorf("Expected game to be %v, got %v", wantTwo, gameTwo)
	}

	gameThree := InputToGame(testInput[2])
	wantThree := Game{
		id: 3,
		sets: []CubeSet{
			{green: 8, blue: 6, red: 20},
			{blue: 5, red: 4, green: 13},
			{green: 5, red: 1},
		},
	}

	if !reflect.DeepEqual(gameThree, wantThree) {
		t.Errorf("Expected game to be %v, got %v", wantThree, gameThree)
	}
}

func TestIsGamePossible(t *testing.T) {
	gameOne := InputToGame(testInput[0])
	gameTwo := InputToGame(testInput[1])
	gameThree := InputToGame(testInput[2])
	gameFour := InputToGame(testInput[3])
	gameFive := InputToGame(testInput[4])

	var availableCubes = CubeSet{blue: 14, red: 12, green: 13}

	if !gameOne.IsPossible(availableCubes) {
		t.Errorf("Expected game One to be possible, got impossible")
	}

	if !gameTwo.IsPossible(availableCubes) {
		t.Errorf("Expected game Two to be possible, got impossible")
	}

	if !gameFive.IsPossible(availableCubes) {
		t.Errorf("Expected game Five to be possible, got impossible")
	}

	if gameThree.IsPossible(availableCubes) {
		t.Errorf("Expected game Three to be impossible, got possible")
	}

	if gameFour.IsPossible(availableCubes) {
		t.Errorf("Expected game Four to be impossible, got possible")
	}
}

func TestSumOfPossibleGameIds(t *testing.T) {
	var availableCubes = CubeSet{blue: 14, red: 12, green: 13}

	var sumOfPossibleGameIds = GetSumOfPossibleGameIds(testInput, availableCubes)
	var expectedSum = 8

	if sumOfPossibleGameIds != expectedSum {
		t.Errorf("Expected sum of possible game ids to be %d, got %d", expectedSum, sumOfPossibleGameIds)
	}
}

func TestSumOfPossibleGameIdsWithInput(t *testing.T) {
	Input, _ := ReadInputFile("input.txt")
	var availableCubes = CubeSet{blue: 14, red: 12, green: 13}

	var sumOfPossibleGameIds = GetSumOfPossibleGameIds(Input, availableCubes)
	var expectedSum = 2810

	if sumOfPossibleGameIds != expectedSum {
		t.Errorf("Expected sum of possible game ids to be %d, got %d", expectedSum, sumOfPossibleGameIds)
	}
}
