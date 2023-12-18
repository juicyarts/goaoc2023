package day13

import (
	"aoc2023/utils"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
)

var testInput = []string{
	// "#.##..##.",
	// "..#.##.#.",
	// "##......#",
	// "##......#",
	// "..#.##.#.",
	// "..##..##.",
	// "#.#.##.#.",
	// "",
	// "#...##..#",
	// "#....#..#",
	// "..##..###",
	// "#####.##.",
	// "#####.##.",
	// "..##..###",
	// "#....#..#",
	// "",
	// "#.#.#..##",
	// ".#..#####",
	// "##..#####",
	// "#.#.#..##",
	// ".#..#.#..",
	// ".##.###..",
	// "#..#...##",
	// "#....####",
	// "###.#.###",
	// "..#.##...",
	// "###.##.##",
	// "",
	// THOSE WHERE NO MATCHES COULD BE FOUND
	// "####..#",
	// "####...",
	// "..#.#..",
	// "##.....",
	// "##.###.",
	// "..##.##",
	// "...#.##",
	// "##..#.#",
	// "..#####",
	// "..#.#.#",
	// "...#...",
	// "",
	// "##...#.",
	// "...#.##",
	// "..#.#..",
	// "##.....",
	// "..##.##",
	// "####...",
	// "###..##",
	// "####...",
	// "###.###",
	// ".....##",
	// "..#.#..",
	// "...##..",
	// "##...##",
	// "",
	// "##.#.....",
	// "....##..#",
	// "..#.#....",
	// "..#.#.##.",
	// "..##.#..#",
	// "..##..##.",
	// "##...#..#",
	// "###..####",
	// "##...#.##",
	// "",
	// "#..#.#.##..",
	// "#..#.#.##..",
	// "##..##..###",
	// "###.#..#...",
	// "###.#.#..##",
	// "......###..",
	// "...#..#.#..",
	// "#.###.###..",
	// ".##.#.##.##",
	// ".##...#..##",
	// "#.###......",
	// "###...##.##",
	// "..##...###.",
	// "...##.#.#..",
	// ".#..##.#.##",
	// "",
	// "#....##.##..#.#.#",
	// "#....##.##..#.#.#",
	// ".#...#.##.####.#.",
	// "..######..#.#.#..",
	// "###.##..##.......",
	// "....##.####...###",
	// "....#..#.###.###.",
	// ".#.######.###....",
	// ".#.######.###...#",
	// "....#..#.###.###.",
	// "....##.####...###",
	// "",
	// "...###..###",
	// "###...##...",
	// "...#......#",
	// "##..#..#.#.",
	// "....##..##.",
	// "..###....##",
	// "#####.##.##",
}

// func TestFindReflections(t *testing.T) {
// 	result := FindReflections(testInput)
// 	expectedResult := 405

// 	if result != expectedResult {
// 		t.Errorf("Expected %+v, got %+v", expectedResult, result)
// 	}
// }

func TestFindReflectionsWithInput(t *testing.T) {
	godotenv.Load()
	expectedResult, _ := strconv.Atoi(os.Getenv("result_1"))

	Input, _ := utils.ReadInputFile("input.txt")
	result := FindReflections(Input)

	if result != expectedResult {
		t.Errorf("Expected to be lower than %+v, got %+v", expectedResult, result)
	}
}
