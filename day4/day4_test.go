package day4

import (
	"aoc2023/utils"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
)

var testInput = []string{
	"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
	"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
	"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
	"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
	"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
	"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
}


func TestInputToCard(t *testing.T) {
	card := InputToCard(testInput[0])

	if card.id != 1 {
		t.Errorf("Expected Card id be 1, got %d", card.id)
	}

	if card.Value() != 8 {
		t.Errorf("Expected Card to be worth 8, got %d", card.Value())
	}
}

func TestTotalWorthOfScratchcards(t *testing.T) {
	var totalPoints = TotalWorthOfScratchcards(testInput)

	if totalPoints != 13 {
		t.Errorf("Expected total worth to be 13, got %d", totalPoints)
	}
}

func TestTotalWorthOfScratchcardsWithInput(t *testing.T) {
	godotenv.Load()
	expectedResult, _ := strconv.Atoi(os.Getenv("result_1"))

	Input, _ := utils.ReadInputFile("input.txt")

	result := TotalWorthOfScratchcards(Input)

	if result != expectedResult {
		t.Errorf("Expected %d, got %d", expectedResult, result)
	}
}

// Part2
func TestTotalAmountOfScratchcards(t *testing.T) {
	var totalPoints = TotalAmountOfScratchcards(testInput)

	if totalPoints != 30 {
		t.Errorf("Expected total amount to be 30, got %d", totalPoints)
	}
}

func TestTotalAmountOfScratchcardsWithInput(t *testing.T) {
	godotenv.Load()
	expectedResult, _ := strconv.Atoi(os.Getenv("result_2"))

	Input, _ := utils.ReadInputFile("input.txt")

	result := TotalAmountOfScratchcards(Input)

	if result != expectedResult {
		t.Errorf("Expected %d, got %d", expectedResult, result)
	}
}
