package day6

import (
	"aoc2023/utils"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
)

var testInput = []string{
	"Time:      7  15   30",
	"Distance:  9  40  200",
}

func TestInputToRaces(t *testing.T) {
	result := InputToRaces(testInput)
	expectedResult := 3

	if len(result) != expectedResult {
		t.Errorf("Expected %d, got %d", expectedResult, result)
	}
}

func TestNumberOfWaysToBeatRace(t *testing.T) {
	races := InputToRaces(testInput)
	resultOne := races[0].NumberOfWaysToBeatRace()
	expectedResultOne := 4

	if resultOne != expectedResultOne {
		t.Errorf("Expected %d, got %d", expectedResultOne, resultOne)
	}

	resultTwo := races[1].NumberOfWaysToBeatRace()
	expectedResultTwo := 8

	if resultTwo != expectedResultTwo {
		t.Errorf("Expected %d, got %d", expectedResultTwo, resultTwo)
	}

	resultThree := races[2].NumberOfWaysToBeatRace()
	expectedResultThree := 9

	if resultThree != expectedResultThree {
		t.Errorf("Expected %d, got %d", expectedResultThree, resultThree)
	}
}

func TestMultiplyNumberOfWaysToBeatRace(t *testing.T) {
	races := InputToRaces(testInput)
	result := MultiplyNumberOfWaysToBeatRace(races)
	expectedResult := 288

	if result != expectedResult {
		t.Errorf("Expected %d, got %d", expectedResult, result)
	}
}

func TestMultiplyNumberOfWaysToBeatRaceWithInput(t *testing.T) {
	godotenv.Load()
	expectedResult, _ := strconv.Atoi(os.Getenv("result_1"))

	Input, _ := utils.ReadInputFile("input.txt")
	races := InputToRaces(Input)
	result := MultiplyNumberOfWaysToBeatRace(races)

	if result != expectedResult {
		t.Errorf("Expected %d, got %d", expectedResult, result)
	}
}

// Part 2
func TestMultiplyNumberOfWaysToBeatRaceWithInputAndSingleRace(t *testing.T) {
	godotenv.Load()
	expectedResult, _ := strconv.Atoi(os.Getenv("result_2"))

	Input, _ := utils.ReadInputFile("input.txt")
	races := InputToRace(Input)
	result := MultiplyNumberOfWaysToBeatRace(races)

	if result != expectedResult {
		t.Errorf("Expected %d, got %d", expectedResult, result)
	}
}
