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
	"18, 19, 30 @ -1, -1, -2", // 1
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
	actual := ReadInput(testInput, 7, 27)

	if actual != expected {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}

func TestMainCollect(t *testing.T) {
	godotenv.Load()
	expected, _ := strconv.Atoi(os.Getenv("result_1"))

	Input, _ := utils.ReadInputFile("input.txt")
	actual := ReadInput(Input, 200000000000000, 400000000000000)

	if actual != expected {
		t.Errorf("Expected other than %+v, got %+v", expected, actual)
	}
}


func TestEzThrow(t *testing.T) {
	expected := 47
	actual := ReadInput(testInput, 7, 27)

	if actual != expected {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}
