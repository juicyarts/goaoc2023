package day10

import (
	"testing"
)

// var testInput = []string{
// 	"..F7.",
// 	".FJ|.",
// 	"SJ.L7",
// 	"|F--J",
// 	"LJ...",
// }

// func TestStepsToFarthestLocation(t *testing.T) {
// 	var result = StepsToFarthestLocation(testInput)
// 	var expectedResult = 8

// 	if result != expectedResult {
// 		t.Errorf("Expected %d, got %d", expectedResult, result)
// 	}
// }

// var testInput2 = []string{
// 	".....",
// 	".S-7.",
// 	".|.|.",
// 	".L-J.",
// 	".....",
// }

// func TestStepsToFarthestLocation2(t *testing.T) {
// 	var result = StepsToFarthestLocation(testInput2)
// 	var expectedResult = 4

// 	if result != expectedResult {
// 		t.Errorf("Expected %d, got %d", expectedResult, result)
// 	}
// }

// func TestStepsToFarthersLocationWithInput(t *testing.T) {
// 	godotenv.Load()
// 	expectedResult, _ := strconv.Atoi(os.Getenv("result_1"))

// 	Input, _ := utils.ReadInputFile("input.txt")
// 	actual := StepsToFarthestLocation(Input)

// 	if actual <= expectedResult {
// 		t.Errorf("Expected actual to be greater than %d, got %d", expectedResult, actual)
// 	}
// }

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
	var result = StepsToFarthestLocation(testInput3)
	var expectedResult = 4

	if result != expectedResult {
		t.Errorf("Expected %d, got %d", expectedResult, result)
	}
}
