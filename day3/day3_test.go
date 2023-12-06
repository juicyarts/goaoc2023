package day3

import (
	"aoc2023/utils"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
)

func TestSumOfPartNumbers(t *testing.T) {

	var testInput = []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	var expectedResult = 4361

	result := SumOfPartNumbers(testInput)

	if result != expectedResult {
		t.Errorf("Expected %d, got %d", expectedResult, result)
	}

}

func TestSumOfPartNumbersWithInput(t *testing.T) {
	godotenv.Load()
	expectedResult, _ := strconv.Atoi(os.Getenv("result_1"))

	Input, _ := utils.ReadInputFile("input.txt")

	result := SumOfPartNumbers(Input)

	if result != expectedResult {
		t.Errorf("Expected %d, got %d", expectedResult, result)
	}
}

// // Part 2

func TestSumOfGearRatios(t *testing.T) {

	var testInput = []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	var expectedResult = 467835

	result := SumOfGearRatios(testInput)

	if result != expectedResult {
		t.Errorf("Expected %d, got %d", expectedResult, result)
	}

}

func TestSumOfGearRatiosWithInput(t *testing.T) {

	godotenv.Load()
	expectedResult, _ := strconv.Atoi(os.Getenv("result_2"))

	Input, _ := utils.ReadInputFile("input.txt")

	result := SumOfGearRatios(Input)

	if result != expectedResult {
		t.Errorf("Expected %d, got %d", expectedResult, result)
	}

}
