package day11

import (
	"aoc2023/utils"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
)

var testInput = []string{
	"...#......",
	".......#..",
	"#.........",
	"..........",
	"......#...",
	".#........",
	".........#",
	"..........",
	".......#..",
	"#...#.....",
}

func TestGetSumOfSteps(t *testing.T) {
	result := GetSumOfSteps(testInput, 2)
	expectedResult := 374

	if result != expectedResult {
		t.Errorf("Expected %+v, got %+v", expectedResult, result)
	}
}

func TestGetSumOfStepsWithInput(t *testing.T) {
	godotenv.Load()
	expectedResult, _ := strconv.Atoi(os.Getenv("result_1"))

	Input, _ := utils.ReadInputFile("input.txt")
	result := GetSumOfSteps(Input, 2)

	if result != expectedResult {
		t.Errorf("Expected result to be greater than %d, got %d", expectedResult, result)
	}
}

func TestGetSumOfStepsWithInputOneMillionMulti(t *testing.T) {
	godotenv.Load()
	expectedResult, _ := strconv.Atoi(os.Getenv("result_2"))

	Input, _ := utils.ReadInputFile("input.txt")
	result := GetSumOfSteps(Input, 1000000)

	if result != expectedResult {
		t.Errorf("Expected result to be greater than %d, got %d", expectedResult, result)
	}
}
