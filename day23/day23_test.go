package day23

import (
	"aoc2023/utils"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
)

var testInput = []string{
	"#.#####################",
	"#.......#########...###",
	"#######.#########.#.###",
	"###.....#.>.>.###.#.###",
	"###v#####.#v#.###.#.###",
	"###.>...#.#.#.....#...#",
	"###v###.#.#.#########.#",
	"###...#.#.#.......#...#",
	"#####.#.#.#######.#.###",
	"#.....#.#.#.......#...#",
	"#.#####.#.#.#########v#",
	"#.#...#...#...###...>.#",
	"#.#.#v#######v###.###v#",
	"#...#.>.#...>.>.#.###.#",
	"#####v#.#.###v#.#.###.#",
	"#.....#...#...#.#.#...#",
	"#.#########.###.#.#.###",
	"#...###...#...#...#.###",
	"###.###.#.###v#####v###",
	"#...#...#.#.>.>.#.>.###",
	"#.###.###.#.###.#.#v###",
	"#.....###...###...#...#",
	"#####################.#",
}

// func TestEzTravel(t *testing.T) {
// 	expected := 82
// 	actual := Travel(testInput)

// 	if actual != expected {
// 		t.Errorf("Expected %+v, got %+v", expected, actual)
// 	}
// }

func TestTravel(t *testing.T) {
	godotenv.Load()
	expected, _ := strconv.Atoi(os.Getenv("result_1"))

	Input, _ := utils.ReadInputFile("input.txt")
	actual := Travel(Input)

	if actual != expected {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}

func TestEzTravelB(t *testing.T) {
	expected := 154
	actual := TravelB(testInput)

	if actual != expected {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}

func TestMainTravelB(t *testing.T) {
	godotenv.Load()
	expected, _ := strconv.Atoi(os.Getenv("result_2"))

	// not 4922

	Input, _ := utils.ReadInputFile("input.txt")
	actual := TravelB(Input)

	if actual != expected {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}
