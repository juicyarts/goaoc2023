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

func TestTravel(t *testing.T) {
	expected := 51
	travelMap := Travel(testInput, []int{0, 3}, 2, map[string]int{})
	actual := len(travelMap)

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestMeasureEnergizedTiles(t *testing.T) {
	godotenv.Load()
	expected := 51

	result := MeasureEnergizedTiles(testInput)

	if result != expected {
		t.Errorf("Expected to equal %+v, got %+v", expected, result)
	}
}

func TestMeasureEnergizedTilesWithInput(t *testing.T) {
	godotenv.Load()
	expected, _ := strconv.Atoi(os.Getenv("result_1"))

	Input, _ := utils.ReadInputFile("input.txt")
	result := MeasureEnergizedTiles(Input)

	if result != expected {
		t.Errorf("Expected to equal %+v, got %+v", expected, result)
	}
}
