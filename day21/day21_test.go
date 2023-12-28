package day21

import (
	"aoc2023/utils"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
)

var testInput = []string{
	"...........",
	".....###.#.",
	".###.##..#.",
	"..#.#...#..",
	"....#.#....",
	".##..S####.",
	".##..#...#.",
	".......##..",
	".##.#.####.",
	".##..##.##.",
	"...........",
}

var testInput1 = []string{
	"...........",
	"...........",
	"...........",
	"...........",
	"...........",
	".....S.....",
	"...........",
	"...........",
	"...........",
	"...........",
	"...........",
}

func TestEdges(t *testing.T) {
	expected := 3712
	actual := Walk(testInput1, []int{5, 5}, 65)

	// printMap(testInput1, counter, visited)

	if actual != expected {
		t.Errorf("Expected %+v, Got %+v", expected, actual)
	}
}

func TestCount1(t *testing.T) {
	expected := 6
	actual, _, _ := getPossibleCount(testInput, []int{5, 5}, 3)

	if actual != expected {
		t.Errorf("Expected %+v, Got %+v", expected, actual)
	}
}

func TestCount(t *testing.T) {
	expected := 16
	actual, _, _ := getPossibleCount(testInput, []int{5, 5}, 6)

	if actual != expected {
		t.Errorf("Expected %+v, Got %+v", expected, actual)
	}
}

func TestTravelWithInput(t *testing.T) {
	godotenv.Load()
	expected, _ := strconv.Atoi(os.Getenv("result_1"))

	Input, _ := utils.ReadInputFile("input.txt")
	actual := Travel(Input, 26501365)

	if actual != expected {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}
