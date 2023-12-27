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

func TestTravel(t *testing.T) {
	expected := 16
	actual := Travel(testInput, 6)

	if actual != expected {
		t.Errorf("Expected %+v, Got %+v", expected, actual)
	}
}

func TestTravelWithExpand(t *testing.T) {
	expected := 50
	actual := Travel(testInput, 10)

	if actual != expected {
		t.Errorf("Expected %+v, Got %+v", expected, actual)
	}
}

func TestTravelWithExpand2(t *testing.T) {
	expected := 1594
	actual := Travel(testInput, 50)

	if actual != expected {
		t.Errorf("Expected %+v, Got %+v", expected, actual)
	}
}

func TestTravelWithExpand3(t *testing.T) {
	expected := 6536
	actual := Travel(testInput, 100)

	if actual != expected {
		t.Errorf("Expected %+v, Got %+v", expected, actual)
	}
}

// func TestTravelWithExpand4(t *testing.T) {
// 	expected := 167004
// 	actual := Travel(testInput, 500)

// 	if actual != expected {
// 		t.Errorf("Expected %+v, Got %+v", expected, actual)
// 	}
// }

func TestTravelWithInput(t *testing.T) {
	godotenv.Load()
	expected, _ := strconv.Atoi(os.Getenv("result_1"))

	Input, _ := utils.ReadInputFile("input.txt")
	actual := Travel(Input, 64)

	if actual != expected {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}

// func TestNormalize(t *testing.T) {
// 	expected := []int{1, 1}
// 	actual := Normalize(1, 1, 10, 10)

// 	if expected[0] != actual[0] || expected[1] != actual[1] {
// 		t.Errorf("Expected %+v, got %+v", expected, actual)
// 	}
// }

// func TestNormalize00(t *testing.T) {
// 	expected := []int{0, 0}
// 	actual := Normalize(0, 0, 10, 10)

// 	if expected[0] != actual[0] || expected[1] != actual[1] {
// 		t.Errorf("Expected %+v, got %+v", expected, actual)
// 	}
// }

// func TestNormalize0(t *testing.T) {
// 	expected := []int{0, 0}
// 	actual := Normalize(10, 10, 10, 10)

// 	if expected[0] != actual[0] || expected[1] != actual[1] {
// 		t.Errorf("Expected %+v, got %+v", expected, actual)
// 	}
// }

// func TestNormalize1(t *testing.T) {
// 	expected := []int{1, 1}
// 	actual := Normalize(11, 11, 10, 10)

// 	if expected[0] != actual[0] || expected[1] != actual[1] {
// 		t.Errorf("Expected %+v, got %+v", expected, actual)
// 	}
// }

// func TestNormalize2(t *testing.T) {
// 	expected := []int{9, 9}
// 	actual := Normalize(-11, -11, 10, 10)

// 	if expected[0] != actual[0] || expected[1] != actual[1] {
// 		t.Errorf("Expected %+v, got %+v", expected, actual)
// 	}
// }

// func TestNormalize22(t *testing.T) {
// 	expected := []int{0, 0}
// 	actual := Normalize(-10, -10, 10, 10)

// 	if expected[0] != actual[0] || expected[1] != actual[1] {
// 		t.Errorf("Expected %+v, got %+v", expected, actual)
// 	}
// }

// func TestNormalize3(t *testing.T) {
// 	expected := []int{9, 2}
// 	actual := Normalize(-1, -8, 10, 10)

// 	if expected[0] != actual[0] || expected[1] != actual[1] {
// 		t.Errorf("Expected %+v, got %+v", expected, actual)
// 	}
// }

// func TestNormalize4(t *testing.T) {
// 	expected := []int{4, 10}
// 	actual := Normalize(4, -1, 11, 11)

// 	if expected[0] != actual[0] || expected[1] != actual[1] {
// 		t.Errorf("Expected %+v, got %+v", expected, actual)
// 	}
// }

// func TestNormalize5(t *testing.T) {
// 	expected := []int{0, 0}
// 	actual := Normalize(-11, -11, 11, 11)

// 	if expected[0] != actual[0] || expected[1] != actual[1] {
// 		t.Errorf("Expected %+v, got %+v", expected, actual)
// 	}
// }

// func TestNormalize6(t *testing.T) {
// 	expected := []int{0, 0}
// 	actual := Normalize(-22, -22, 11, 11)

// 	if expected[0] != actual[0] || expected[1] != actual[1] {
// 		t.Errorf("Expected %+v, got %+v", expected, actual)
// 	}
// }

// func TestNormalize7(t *testing.T) {
// 	expected := []int{1, 9}
// 	actual := Normalize(-12, -10, 11, 11)

// 	if expected[0] != actual[0] || expected[1] != actual[1] {
// 		t.Errorf("Expected %+v, got %+v", expected, actual)
// 	}
// }
