package day12

import (
	"aoc2023/utils"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
)

var testInput = []string{
	"???.### 1,1,3",
	".??..??...?##. 1,1,3",
	"?#?#?#?#?#?#?#? 1,3,1,6",
	"????.#...#... 4,1,1",
	"????.######..#####. 1,6,5",
	"?###???????? 3,2,1",
}

func TestSumOfArrangements(t *testing.T) {
	result := GetNumberOfArrangements(testInput[0])
	expectedResult := 1

	if result != expectedResult {
		t.Errorf("Expected %+v, got %+v", expectedResult, result)
	}
}

func TestSumOfArrangements1(t *testing.T) {
	result := GetNumberOfArrangements(testInput[1])
	expectedResult := 4

	if result != expectedResult {
		t.Errorf("Expected %+v, got %+v", expectedResult, result)
	}
}

func TestSumOfArrangements2(t *testing.T) {
	result := GetNumberOfArrangements(testInput[2])
	expectedResult := 1

	if result != expectedResult {
		t.Errorf("Expected %+v, got %+v", expectedResult, result)
	}
}

func TestSumOfArrangements3(t *testing.T) {
	result := GetNumberOfArrangements(testInput[3])
	expectedResult := 1

	if result != expectedResult {
		t.Errorf("Expected %+v, got %+v", expectedResult, result)
	}
}

func TestSumOfArrangements4(t *testing.T) {
	result := GetNumberOfArrangements(testInput[4])
	expectedResult := 4

	if result != expectedResult {
		t.Errorf("Expected %+v, got %+v", expectedResult, result)
	}
}

func TestSumOfArrangements5(t *testing.T) {
	result := GetNumberOfArrangements(testInput[5])
	expectedResult := 10

	if result != expectedResult {
		t.Errorf("Expected %+v, got %+v", expectedResult, result)
	}
}

func TestTotalNumberOfArrangements(t *testing.T) {
	result := GetTotalNumberOfArrangements(testInput)
	expectedResult := 21

	if result != expectedResult {
		t.Errorf("Expected %+v, got %+v", expectedResult, result)
	}
}

func TestTotalNumberOfArrangementsWithInput(t *testing.T) {
	godotenv.Load()
	expectedResult, _ := strconv.Atoi(os.Getenv("result_1"))

	Input, _ := utils.ReadInputFile("input.txt")
	result := GetTotalNumberOfArrangements(Input)

	if result != expectedResult {
		t.Errorf("Expected to be lower than %+v, got %+v", expectedResult, result)
	}
}
