package day10

import (
	"aoc2023/utils"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
)

var testInput = []string{
	"..F7.",
	".FJ|.",
	"SJ.L7",
	"|F--J",
	"LJ...",
}

func TestStepsToFarthestLocation(t *testing.T) {
	var result, _ = StepsToFarthestLocation(testInput)
	var expectedResult = 8

	if result != expectedResult {
		t.Errorf("Expected %d, got %d", expectedResult, result)
	}
}

var testInput2 = []string{
	".....",
	".S-7.",
	".|.|.",
	".L-J.",
	".....",
}

func TestStepsToFarthestLocation2(t *testing.T) {
	var result, _ = StepsToFarthestLocation(testInput2)
	var expectedResult = 4

	if result != expectedResult {
		t.Errorf("Expected %d, got %d", expectedResult, result)
	}
}

func TestStepsToFarthersLocationWithInput(t *testing.T) {
	godotenv.Load()
	expectedResult, _ := strconv.Atoi(os.Getenv("result_1"))

	Input, _ := utils.ReadInputFile("input.txt")
	result, _ := StepsToFarthestLocation(Input)

	if result != expectedResult {
		t.Errorf("Expected result to be greater than %d, got %d", expectedResult, result)
	}
}

var testInput3 = []string{
	"...........",
	".S-------7.",
	".|F-----7|.",
	".||.....||.",
	".||.....||.",
	".|L-7.F-J|.",
	".|..|.|..|.",
	".L--J.L--J.",
	"...........",
}

// Part 2
func TestTilesEnclosedByLoop(t *testing.T) {
	var _, result = StepsToFarthestLocation(testInput3)
	var expectedResult = 4

	if result != expectedResult {
		t.Errorf("Expected %d, got %d", expectedResult, result)
	}
}

var testInput4 = []string{
	".F----7F7F7F7F-7....",
	".|F--7||||||||FJ....",
	".||.FJ||||||||L7....",
	"FJL7L7LJLJ||LJ.L-7..",
	"L--J.L7...LJS7F-7L7.",
	"....F-J..F7FJ|L7L7L7",
	"....L7.F7||L7|.L7L7|",
	".....|FJLJ|FJ|F7|.LJ",
	"....FJL-7.||.||||...",
	"....L---J.LJ.LJLJ...",
}

// Part 2
func TestTilesEnclosedByLoop2(t *testing.T) {
	var _, result = StepsToFarthestLocation(testInput4)
	var expectedResult = 8

	if result != expectedResult {
		t.Errorf("Expected %d, got %d", expectedResult, result)
	}
}


func TestTilesEnclosedByLoopWithInput(t *testing.T) {
	godotenv.Load()
	expectedResult, _ := strconv.Atoi(os.Getenv("result_2"))

	Input, _ := utils.ReadInputFile("input.txt")
	_, result := StepsToFarthestLocation(Input)

	if result != expectedResult {
		t.Errorf("Expected result to be greater than %d, got %d", expectedResult, result)
	}
}
