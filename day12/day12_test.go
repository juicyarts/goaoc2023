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

func TestSumOfArrangements5(t *testing.T) {
	result := GetNumberOfArrangements(testInput[5])
	expectedResult := 10

	if result != expectedResult {
		t.Errorf("Expected %+v, got %+v", expectedResult, result)
	}
}

func TestCount5(t *testing.T) {
	result := Count("????", []int{2, 1}, make(map[string]int))
	expectedResult := 1

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
