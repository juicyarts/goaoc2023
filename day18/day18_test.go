package day18

import (
	"testing"
)

var testInput = []string{
	"R 6 (#70c710)",
	"D 5 (#0dc571)",
	"L 2 (#5713f0)",
	"D 2 (#d2c081)",
	"R 2 (#59c680)",
	"D 2 (#411b91)",
	"L 5 (#8ceee2)",
	"U 2 (#caa173)",
	"L 1 (#1b58a2)",
	"U 2 (#caa171)",
	"R 2 (#7807d2)",
	"U 3 (#a77fa3)",
	"L 2 (#015232)",
	"U 2 (#7a21e3)",
}

var testResult1 = []string{
	"#######",
	"#.....#",
	"###...#",
	"..#...#",
	"..#...#",
	"###.###",
	"#...#..",
	"##..###",
	".#....#",
	".######",
}

func TestMakeGrid(t *testing.T) {
	expected := 38
	result, _ := Travel(testInput)

	if result != expected {
		t.Errorf("Expected to equal %+v, got %+v", expected, result)
	}
}

func TestGridDimension(t *testing.T) {
	expected := 62
	_, result := Travel(testInput)

	if result != expected {
		t.Errorf("Expected to equal %+v, got %+v", expected, result)
	}
}

// func TestTravelWithInput(t *testing.T) {

// 	expected := [][]int{}

// 	Input, _ := utils.ReadInputFile("input.txt")
// 	result, _ := Travel(Input)

// 	if result == 0 {
// 		t.Errorf("Expected to equal %+v, got %+v", expected, result)
// 	}
// }
