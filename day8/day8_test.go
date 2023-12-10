package day8

import (
	"aoc2023/utils"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
)

// var testInput = []string{
// 	"RL",
// 	"",
// 	"AAA = (BBB, CCC)",
// 	"BBB = (DDD, EEE)",
// 	"CCC = (ZZZ, GGG)",
// 	"DDD = (DDD, DDD)",
// 	"EEE = (EEE, EEE)",
// 	"GGG = (GGG, GGG)",
// 	"ZZZ = (ZZZ, ZZZ)",
// }

// func TestInputToMap(t *testing.T) {
// 	result := InputToMap(testInput)

// 	fmt.Printf("%+v \n", result.m["AAA"][0])
// 	if result.m["AAA"][0] != "BBB" && result.m["AAA"][1] != "CCC" {
// 		t.Errorf("AAA should be (BBB, CCC), got %s \n %+v", result.m["AAA"], result)
// 	}
// }

// func TestStepsToZZZ(t *testing.T) {
// 	gameMap := InputToMap(testInput)
// 	result := gameMap.FindStepsToEnd("AAA", 0, 0, "ZZZ")

// 	expectedResult := 2

// 	if result.steps != expectedResult {
// 		t.Errorf("Expected %d steps, got %d", expectedResult, result.steps)
// 	}
// }

// var testInput2 = []string{
// 	"LLR",
// 	"",
// 	"AAA = (BBB, BBB)",
// 	"BBB = (AAA, ZZZ)",
// 	"ZZZ = (ZZZ, ZZZ)",
// }

// func TestStepsToZZZ2(t *testing.T) {
// 	gameMap := InputToMap(testInput2)
// 	result := gameMap.FindStepsToEnd("AAA", 0, 0, "ZZZ")

// 	expectedResult := 6

// 	if result.steps != expectedResult {
// 		t.Errorf("Expected %d steps, got %d", expectedResult, result.steps)
// 	}
// }

// func TestStepsToZZZWithInput(t *testing.T) {
// 	godotenv.Load()
// 	expectedResult, _ := strconv.Atoi(os.Getenv("result_1"))

// 	Input, _ := utils.ReadInputFile("input.txt")

// 	gameMap := InputToMap(Input)
// 	result := gameMap.FindStepsToEnd("AAA", 0, 0, "ZZZ")

// 	if result.steps != expectedResult {
// 		t.Errorf("Expected %d steps, got %d", expectedResult, result.steps)
// 	}

// }

// Part 2

var testInput3 = []string{
	"LR",
	"",
	"11A = (11B, XXX)",
	"11B = (XXX, 11Z)",
	"11Z = (11B, XXX)",
	"22A = (22B, XXX)",
	"22B = (22C, 22C)",
	"22C = (22Z, 22Z)",
	"22Z = (22B, 22B)",
	"XXX = (XXX, XXX)",
}

func TestStepsToZZZWithInputPart2(t *testing.T) {
	gameMap := InputToMap(testInput3)
	result := gameMap.FindStepsToEndForEachEndingWithA()

	expectedResult := 6

	if result != expectedResult {
		t.Errorf("Expected %d steps, got %d", expectedResult, result)
	}
}

func TestStepsToZZZWithInputPartWithInput(t *testing.T) {
	godotenv.Load()
	expectedResult, _ := strconv.Atoi(os.Getenv("result_2"))

	Input, _ := utils.ReadInputFile("input.txt")

	gameMap := InputToMap(Input)
	result := gameMap.FindStepsToEndForEachEndingWithA()

	if result != expectedResult {
		t.Errorf("Expected %d steps, got %d", expectedResult, result)
	}
}
