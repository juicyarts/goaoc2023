package day17

import (
	"aoc2023/utils"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
)

var testInput = []string{
	"2413432311323",
	"3215453535623",
	"3255245654254",
	"3446585845452",
	"4546657867536",
	"1438598798454",
	"4457876987766",
	"3637877979653",
	"4654967986887",
	"4564679986453",
	"1224686865563",
	"2546548887735",
	"4322674655533",
}

func TestTravelBasic(t *testing.T) {
	expected := 102 // remove the start point
	heat := Travel(testInput, []int{0, 0}, 0, 3) - 2

	// fmt.Printf("Map: %+v \n", actual)
	fmt.Printf("Heat: %+v \n", heat)
	if heat != expected {
		t.Errorf("Expected %v, got %v", expected, heat)
	}
}

func TestTravelInput(t *testing.T) {
	godotenv.Load()
	expectedRangeOne, _ := strconv.Atoi(os.Getenv("result_1"))
	expectedRangeTwo, _ := strconv.Atoi(os.Getenv("result_2"))

	Input, _ := utils.ReadInputFile("input.txt")
	result := Travel(Input, []int{0, 0}, 0, 3)

	if !(result > expectedRangeTwo && result < expectedRangeOne) {
		t.Errorf("Expected %+v to be between %+v and %+v", result, expectedRangeTwo, expectedRangeOne)
	}
}
