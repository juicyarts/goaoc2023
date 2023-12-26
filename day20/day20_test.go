package day20

import (
	"aoc2023/utils"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
)

// var testInputEz = []string{
// 	"broadcaster -> a, b, c",
// 	"%a -> b",
// 	"%b -> c",
// 	"%c -> inv",
// 	"&inv -> a",
// }

// var testInput = []string{
// 	"broadcaster -> a",
// 	"%a -> inv, con",
// 	"&inv -> b",
// 	"%b -> con",
// 	"&con -> output",
// }

// func TestMeasurePulsesEz(t *testing.T) {
// 	expected := 32000000
// 	actual, _ := MeasurePulses(testInputEz)

// 	if actual != expected {
// 		t.Errorf("Expected %+v, got %+v", expected, actual)
// 	}
// }

// func TestMeasurePulses(t *testing.T) {
// 	expected := 11687500
// 	actual, _ := MeasurePulses(testInput)

// 	if actual != expected {
// 		t.Errorf("Expected %+v, got %+v", expected, actual)
// 	}
// }

// func TestMeasurePulseWithInput(t *testing.T) {
// 	godotenv.Load()
// 	expected, _ := strconv.Atoi(os.Getenv("result_1"))

// 	Input, _ := utils.ReadInputFile("input.txt")
// 	actual, _ := MeasurePulses(Input)

// 	if actual != expected {
// 		t.Errorf("Expected to equal %+v, got %+v", expected, actual)
// 	}
// }

func TestAmountOfRepNeeded(t *testing.T) {
	godotenv.Load()
	expected, _ := strconv.Atoi(os.Getenv("result_2"))

	Input, _ := utils.ReadInputFile("input.txt")
	actual := MeasurePulses(Input)

	if actual != expected {
		t.Errorf("Expected to equal %+v, got %+v", expected, actual)
	}
}
