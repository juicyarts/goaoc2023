package day14

import (
	"aoc2023/utils"
	"os"
	"reflect"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
)

var testInput = []string{
	"O....#....",
	"O.OO#....#",
	".....##...",
	"OO.#O....O",
	".O.....O#.",
	"O.#..O.#.#",
	"..O..#O..O",
	".......O..",
	"#....###..",
	"#OO..#....",
}

func TestSortRow2(t *testing.T) {
	expected := "OOO......."
	actual := SortRow("...OO....O")

	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestSortRow3(t *testing.T) {
	expected := "O....#OO.."
	actual := SortRow(".O...#O..O")

	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestSortRow4(t *testing.T) {
	expected := "O..#......"
	actual := SortRow(".O.#......")

	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestSortRow5(t *testing.T) {
	expected := ".#O......."
	actual := SortRow(".#.O......")

	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestSortRow6(t *testing.T) {
	expected := "#.#O..#.##"
	actual := SortRow("#.#..O#.##")

	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestSortRow7(t *testing.T) {
	expected := "..#O....#."
	actual := SortRow("..#...O.#.")

	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestSortRow8(t *testing.T) {
	expected := "O....#O.#."
	actual := SortRow("....O#.O#.")

	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestSortRow9(t *testing.T) {
	expected := "....#....."
	actual := SortRow("....#.....")

	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestSortRow10(t *testing.T) {
	expected := ".#O..#O..."
	actual := SortRow(".#.O.#O...")

	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

// only part 2 left here
func TestCycle(t *testing.T) {
	result := Main(testInput)
	// expectedResult := 136
	expectedResult := 64

	if result != expectedResult {
		t.Errorf("Expected %d, got %d", expectedResult, result)
	}
}

func TestMainWithInput(t *testing.T) {
	godotenv.Load()
	expectedResult, _ := strconv.Atoi(os.Getenv("result_2"))

	Input, _ := utils.ReadInputFile("input.txt")
	result := Main(Input)

	if result != expectedResult {
		t.Errorf("Expected to be higher than %+v, got %+v", expectedResult, result)
	}
}

func TestRotateAntiClockwise(t *testing.T) {
	localTestInput := []string{
		"ABCD",
		"EFGH",
	}

	expected := []string{
		"DH",
		"CG",
		"BF",
		"AE",
	}

	actual := utils.RotateAntiClockWise(localTestInput)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestRotateClockwise(t *testing.T) {
	localTestInput := []string{
		"ABCD",
		"EFGH",
	}

	expected := []string{
		"EA",
		"FB",
		"GC",
		"HD",
	}

	actual := utils.RotateClockWise(localTestInput)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
