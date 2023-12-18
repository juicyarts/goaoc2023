package day13

import (
	"aoc2023/utils"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
)

var testInput = []string{
	"#.##..##.",
	"..#.##.#.",
	"##......#",
	"##......#",
	"..#.##.#.",
	"..##..##.",
	"#.#.##.#.",
	"",
	"#...##..#",
	"#....#..#",
	"..##..###",
	"#####.##.",
	"#####.##.",
	"..##..###",
	"#....#..#",
}

func TestFindReflections(t *testing.T) {
	result := FindReflections(testInput)
	expectedResult := 400

	if result != expectedResult {
		t.Errorf("Expected %+v, got %+v", expectedResult, result)
	}
}

func TestFindReflectionsWithInput(t *testing.T) {
	godotenv.Load()
	expectedResult, _ := strconv.Atoi(os.Getenv("result_2"))

	Input, _ := utils.ReadInputFile("input.txt")
	result := FindReflections(Input)

	if result != expectedResult {
		t.Errorf("Expected to be higher than %+v, got %+v", expectedResult, result)
	}
}
