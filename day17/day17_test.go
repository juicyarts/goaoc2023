package day17

import (
	"aoc2023/utils"
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

var testInput2 = []string{
	"2413432",
	"3215453",
	"3255245",
	"4546657",
}

// func TestTravel(t *testing.T) {
// 	expected := 102 - 2
// 	heat := Travel(testInput, []int{0, 0}, []int{len(testInput) - 1, len(testInput[0]) - 1}, 0, 0, 0, "")

// 	// fmt.Printf("Map: %+v \n", actual)
// 	fmt.Printf("Heat: %+v \n", heat-2)
// 	if heat != expected {
// 		t.Errorf("Expected %v, got %v", expected, heat-2)
// 	}
// }

func TestTravel2(t *testing.T) {
	godotenv.Load()
	expected, _ := strconv.Atoi(os.Getenv("result_1"))

	Input, _ := utils.ReadInputFile("input.txt")
	result := Travel(Input, []int{0, 0}, []int{len(Input) - 1, len(Input[0]) - 1}, 0, 0, 0, "")

	if result != expected {
		t.Errorf("Expected to equal %+v, got %+v", expected, result)
	}
}
