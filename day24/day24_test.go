package day24

import (
	"aoc2023/utils"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
)

var testInput = []string{
	"19, 13, 30 @ -2,  1, -2", // 0
	"18, 19, 22 @ -1, -1, -2", // 1
	"20, 25, 34 @ -2, -2, -4", // 2
	"12, 31, 28 @ -1, -2, -1", // 3
	"20, 19, 15 @  1, -5, -3", // 4
	// "4, 14, 12 @ -4, -1, 0",   // 5
}

// Collisions
// 0, 1 -> 14,15
// 0, 2 -> 11,16
// 0, 3 6, 19

func TestEzReadInput(t *testing.T) {
	expected := 2
	actual := ReadInput(testInput, 8, 26)

	if actual != expected {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}

func TestMainCollect(t *testing.T) {
	godotenv.Load()
	expected, _ := strconv.Atoi(os.Getenv("result_1"))

	Input, _ := utils.ReadInputFile("input.txt")
	actual := ReadInput(Input, 200000000000000, 400000000000000)

	// must be higher than 12019 & lower than 17363
	// not 11567
	if actual != expected {
		t.Errorf("Expected other than %+v, got %+v", expected, actual)
	}
}

// func TestDoIntersectDummyP(t *testing.T) {
// 	// lines are parrallel ||
// 	l1 := [][]float64{{1, 1}, {1, 10}}
// 	l2 := [][]float64{{3, 1}, {3, 10}}

// 	expected := false
// 	actual, _ := doIntersectDummy(
// 		l1[0][0],
// 		l1[1][0],
// 		l2[0][0],
// 		l2[1][0],
// 		l1[0][1],
// 		l1[1][1],
// 		l2[0][1],
// 		l2[1][1],
// 	)

// 	if expected != actual {
// 		t.Errorf("Expected %+v, got %+v", expected, actual)
// 	}
// }

// func TestDoIntersectDummy1(t *testing.T) {
// 	// lines not parallel and do cross _|_
// 	l1 := [][]float64{{2, 1}, {2, 10}}
// 	l2 := [][]float64{{1, 2}, {10, 2}}

// 	expected := true
// 	actual, _ := doIntersectDummy(
// 		l1[0][0],
// 		l1[1][0],
// 		l2[0][0],
// 		l2[1][0],
// 		l1[0][1],
// 		l1[1][1],
// 		l2[0][1],
// 		l2[1][1],
// 	)

// 	if expected != actual {
// 		t.Errorf("Expected %+v, got %+v", expected, actual)
// 	}
// }

// func TestDoIntersectDummy2(t *testing.T) {
// 	// lines not parallel but don't cross _|
// 	l1 := [][]float64{{1, 1}, {1, 10}}
// 	l2 := [][]float64{{1, 1}, {10, 1}}

// 	expected := false
// 	actual, _ := doIntersectDummy(
// 		l1[0][0],
// 		l1[1][0],
// 		l2[0][0],
// 		l2[1][0],
// 		l1[0][1],
// 		l1[1][1],
// 		l2[0][1],
// 		l2[1][1],
// 	)

// 	if expected != actual {
// 		t.Errorf("Expected %+v, got %+v", expected, actual)
// 	}
// }
