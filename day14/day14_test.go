package day14

import (
	"aoc2023/utils"
	"fmt"
	"os"
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

var sortTestInput = []string{
	"...OO....O",
	".O...#O..O",
	".O.#......",
	".#.O......",
	"#.#..O#.##",
	"..#...O.#.",
	"....O#.O#.",
	"....#.....",
	".#.O.#O...",
}

var sortTestExpected = []string{
	"OOO.......",
	"O....#OO..",
	"O..#......",
	".#O.......",
	"#.#O..#.##",
	"..#O....#.",
	"O....#O.#.",
	"....#.....",
	".#O..#O...",
}

func TestSortRow(t *testing.T) {
	for i, row := range sortTestInput {
		tName := fmt.Sprintf("Sorting %s", row)
		t.Run(tName, func(t *testing.T) {
			expected := sortTestExpected[i]
			result := SortRow(row)

			if expected != result {
				t.Errorf("Expected %s, got %s", expected, result)
			}
		})
	}
}

func BenchmarkSortRow(b *testing.B) {
	for _, v := range sortTestInput {
		b.Run(fmt.Sprintf("input_size_%s", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SortRow(v)
			}
		})
	}
}

// only part 2 left here
func TestCycle(t *testing.T) {
	expected := 64
	result := Main(testInput)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestMainWithInput(t *testing.T) {
	godotenv.Load()
	expected, _ := strconv.Atoi(os.Getenv("result_1"))

	Input, _ := utils.ReadInputFile("input.txt")
	result := Main(Input)

	if result != expected {
		t.Errorf("Expected to be higher than %+v, got %+v", expected, result)
	}
}
