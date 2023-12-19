package day16

import (
	"aoc2023/utils"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
)

var testInput = []string{
	`.|...\....`,
	`|.-.\.....`,
	`.....|-...`,
	`........|.`,
	`..........`,
	`.........\`,
	`..../.\\..`,
	`.-.-/..|..`,
	`.|....-|.\`,
	`..//.|....`,
}

// func TestMeasureEnergizedTiles(t *testing.T) {
// 	expected := 46
// 	actual := MeasureEnergizedTiles(testInput)

// 	if !reflect.DeepEqual(expected, actual) {
// 		t.Errorf("Expected %v, got %v", expected, actual)
// 	}
// }

func TestMeasureEnergizedTilesWithInput(t *testing.T) {
	godotenv.Load()
	expected, _ := strconv.Atoi(os.Getenv("result_1"))

	Input, _ := utils.ReadInputFile("input.txt")
	result := MeasureEnergizedTiles(Input)

	if result != expected {
		t.Errorf("Expected to be higher than %+v, got %+v", expected, result)
	}
}
