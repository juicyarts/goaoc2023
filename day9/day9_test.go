package day9

import (
	"aoc2023/utils"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
)

var testInput = []string{
	"0 3 6 9 12 15",
	"1 3 6 10 15 21",
	"10 13 16 21 30 45",
}

func TestFindNextInSequence(t *testing.T) {
	var input = []int{0, 3, 6, 9, 12, 15}

	expected := 18
	actual := FindNextInSequence(input)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestFindNextInSequence2(t *testing.T) {
	var input = []int{1, 3, 6, 10, 15, 21}

	expected := 28
	actual := FindNextInSequence(input)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestFindNextInSequence3(t *testing.T) {
	var input = []int{10, 13, 16, 21, 30, 45}

	expected := 68
	actual := FindNextInSequence(input)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestGetSequenceSum(t *testing.T) {
	expected := 114
	actual := GetSequenceSum(testInput, false)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestGetSequenceSumFromInput(t *testing.T) {
	godotenv.Load()
	expectedResult, _ := strconv.Atoi(os.Getenv("result_1"))

	Input, _ := utils.ReadInputFile("input.txt")
	actual := GetSequenceSum(Input, false)

	if actual != expectedResult {
		t.Errorf("Expected actual to be less than %d, got %d", expectedResult, actual)
	}
}

// Part 2
func TestFindNextInSequenceReverse(t *testing.T) {
	var input = []int{45, 30, 21, 16, 13, 10}

	expected := 5
	actual := FindNextInSequence(input)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestGetSequenceSumFromInputPart2(t *testing.T) {
	godotenv.Load()
	expectedResult, _ := strconv.Atoi(os.Getenv("result_2"))

	Input, _ := utils.ReadInputFile("input.txt")
	actual := GetSequenceSum(Input, true)

	if actual != expectedResult {
		t.Errorf("Expected actual to be less than %d, got %d", expectedResult, actual)
	}
}
