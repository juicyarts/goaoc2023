package day17

import (
	"testing"
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
	"111111111111",
	"999999999991",
	"999999999991",
	"999999999991",
	"999999999991",
}

func TestTravelBasic(t *testing.T) {
	expected := 94
	heat := Travel(testInput, []int{0, 0}, 4, 10, 3)
	// heatB := Travel(testInput, []int{0, 0}, 4, 10, 2)
	// fmt.Print("T", heatB)
	// foo := Travel(testInput2, []int{0, 0}, 4, 10, 3)

	if heat != expected {
		t.Errorf("Expected %v, got %v", expected, heat)
	}
}

// func TestTravelInput(t *testing.T) {
// 	godotenv.Load()
// 	expected, _ := strconv.Atoi(os.Getenv("result_1"))

// 	Input, _ := utils.ReadInputFile("input.txt")
// 	result := Travel(Input, []int{0, 0}, 4, 10)

// 	if result <= expected {
// 		t.Errorf("Expected %+v to be higher than %+v", result, expected)
// 	}
// }

// func TestTravelBasicUltra(t *testing.T) {
// 	expected := 94 // remove the start point
// 	heat := Travel(testInput, []int{0, 0}, 4, 10)

// 	if heat != expected {
// 		t.Errorf("Expected %v, got %v", expected, heat)
// 	}
// }

// func TestTravelBasicUltra2(t *testing.T) {
// 	expected := 71 // remove the start point
// 	heat := Travel(testInput2, []int{0, 0}, 4, 10)

// 	// fmt.Printf("Map: %+v \n", actual)
// 	fmt.Printf("Heat: %+v \n", heat)
// 	if heat != expected {
// 		t.Errorf("Expected %v, got %v", expected, heat)
// 	}
// }
