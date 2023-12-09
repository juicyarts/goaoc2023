package day5

import (
	"aoc2023/utils"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
)

var testInput = []string{
	"seeds: 79 14 55 13",
	"",
	"seed-to-soil map:",
	"50 98 2",
	"52 50 48",
	"",
	"soil-to-fertilizer map:",
	"0 15 37",
	"37 52 2",
	"39 0 15",
	"",
	"fertilizer-to-water map:",
	"49 53 8",
	"0 11 42",
	"42 0 7",
	"57 7 4",
	"",
	"water-to-light map:",
	"88 18 7",
	"18 25 70",
	"",
	"light-to-temperature map:",
	"45 77 23",
	"81 45 19",
	"68 64 13",
	"",
	"temperature-to-humidity map:",
	"0 69 1",
	"1 0 69",
	"",
	"humidity-to-location map:",
	"60 56 37",
	"56 93 4",
}

func TestGetLowestSeedLocation(t *testing.T) {
	almanac := InputToAlmanac(testInput)
	result := almanac.GetLowestSeedLocation()

	expectedResult := 35

	if result != expectedResult {
		t.Errorf("Expected %d, got %d", expectedResult, result)
	}
}

func TestGetLowestSeedLocationWithInput(t *testing.T) {
	godotenv.Load()
	expectedResult, _ := strconv.Atoi(os.Getenv("result_1"))

	Input, _ := utils.ReadInputFile("input.txt")

	almanac := InputToAlmanac(Input)
	result := almanac.GetLowestSeedLocation()

	if result != expectedResult {
		t.Errorf("Expected %d, got %d", expectedResult, result)
	}
}

// // Part 2

func TestGetLowestSeedLocationBySeedRange(t *testing.T) {
	almanac := InputToAlmanac(testInput)
	result := almanac.GetLowestSeedLocationBySeedRange()

	expectedResult := 46

	if result != expectedResult {
		t.Errorf("Expected %d, got %d", expectedResult, result)
	}
}

func TestGetLowestSeedLocationBySeedRangeWithInput(t *testing.T) {
	godotenv.Load()
	expectedResult, _ := strconv.Atoi(os.Getenv("result_2"))

	Input, _ := utils.ReadInputFile("input.txt")

	almanac := InputToAlmanac(Input)
	result := almanac.GetLowestSeedLocationBySeedRange()

	if result != expectedResult {
		t.Errorf("Expected %d, got %d", expectedResult, result)
	}
}
