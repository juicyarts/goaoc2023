package day22

import (
	"aoc2023/utils"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
)

var testInput = []string{
	"1,0,1~1,2,1",
	"0,0,2~2,0,2",
	"0,2,3~2,2,3",
	"0,0,4~0,2,4",
	"2,0,5~2,2,5",
	"0,1,6~2,1,6",
	"1,1,8~1,1,9",
}

func TestCollectBricks(t *testing.T) {
	expected := 5
	actual := CollectBricks(testInput)

	if expected != actual {
		t.Errorf("Expected %+v , got %+v", expected, actual)
	}
}

func TestMain(t *testing.T) {
	godotenv.Load()
	expected, _ := strconv.Atoi(os.Getenv("result_1"))

	Input, _ := utils.ReadInputFile("input.txt")
	actual := CollectBricks(Input)

	if actual != expected {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}

func TestRangesDoCollide1(t *testing.T) {
	expected := true
	a := [][]int{{1, 0}, {1, 2}}
	b := [][]int{{0, 0}, {2, 0}}

	actual := rangesDoCollide(a[0], a[1], b[0], b[1])

	if actual != expected {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}

func TestRangesDoCollide2(t *testing.T) {
	expected := true
	a := [][]int{{0, 0}, {2, 2}}
	b := [][]int{{1, 0}, {2, 0}}

	actual := rangesDoCollide(a[0], a[1], b[0], b[1])

	if actual != expected {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}

func TestRangesDoCollide3(t *testing.T) {
	expected := false
	a := [][]int{{1, 1}, {1, 1}}
	b := [][]int{{0, 0}, {2, 0}}

	actual := rangesDoCollide(a[0], a[1], b[0], b[1])

	if actual != expected {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}

func TestRangesDoCollide4(t *testing.T) {
	expected := false
	a := [][]int{{0, 0}, {1, 2}}
	b := [][]int{{2, 0}, {2, 0}}

	actual := rangesDoCollide(a[0], a[1], b[0], b[1])

	if actual != expected {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}

func TestRangesDoCollide5(t *testing.T) {
	expected := true
	a := [][]int{{2, 0}, {2, 2}}
	b := [][]int{{0, 1}, {2, 1}}

	actual := rangesDoCollide(a[0], a[1], b[0], b[1])

	if actual != expected {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}
