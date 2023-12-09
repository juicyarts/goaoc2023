package day7_2

import (
	"aoc2023/utils"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
)

var testInput = []string{
	"32T3K 765",
	"T55J5 684",
	"KK677 28",
	"KTJJT 220",
	"QQQJA 483",
}

func TestFiveOfAKind(t *testing.T) {
	hands := InputToHands([]string{"AAAAA 765"})
	expected := 6

	fmt.Printf("%+v \n", hands[0])

	if hands[0].handType != expected {
		t.Errorf("Expected handType to be %v, got %v", expected, hands[0].handType)
	}
}

func TestFourOfAKind(t *testing.T) {
	hands := InputToHands([]string{"AAAA2 765"})
	expected := 5

	fmt.Printf("%+v \n", hands[0])

	if hands[0].handType != expected {
		t.Errorf("Expected handType to be %v, got %v", expected, hands[0].handType)
	}
}

func TestFullHouse(t *testing.T) {
	hands := InputToHands([]string{"AAA22 765"})
	expected := 4

	fmt.Printf("%+v \n", hands[0])

	if hands[0].handType != expected {
		t.Errorf("Expected handType to be %v, got %v", expected, hands[0].handType)
	}
}

func TestThreeOfKind(t *testing.T) {
	hands := InputToHands([]string{"AAA32 765"})
	expected := 3

	fmt.Printf("%+v \n", hands[0])

	if hands[0].handType != expected {
		t.Errorf("Expected handType to be %v, got %v", expected, hands[0].handType)
	}
}

func TestTwoPair(t *testing.T) {
	hands := InputToHands([]string{"A23A2 765"})
	expected := 2

	fmt.Printf("%+v \n", hands[0])

	if hands[0].handType != expected {
		t.Errorf("Expected handType to be %v, got %v", expected, hands[0].handType)
	}
}

func TestOnePair(t *testing.T) {
	hands := InputToHands([]string{"A23A4 765"})
	expected := 1

	fmt.Printf("%+v \n", hands[0])

	if hands[0].handType != expected {
		t.Errorf("Expected handType to be %v, got %v", expected, hands[0].handType)
	}
}

func TestHighCard(t *testing.T) {
	hands := InputToHands([]string{"23456 765"})
	expected := 0

	fmt.Printf("%+v \n", hands[0])

	if hands[0].handType != expected {
		t.Errorf("Expected handType to be %v, got %v", expected, hands[0].handType)
	}
}


func TestInputToHand(t *testing.T) {
	hands := InputToHands(testInput)
	result := GetTotalWinnings(hands)
	expected := 5905

	if len(hands) != 5 {
		t.Errorf("Expected hands to have length of 5, got %v", len(hands))
	}

	if result != expected {
		t.Errorf("Expected result to be %v, got %v", expected, result)
	}

}


// Part 2

func TestJokerCard(t *testing.T) {
	hands := InputToHands([]string{"QJJQ2 765"})
	expected := 5

	fmt.Printf("%+v \n", hands[0])

	if hands[0].handType != expected {
		t.Errorf("Expected handType to be %v, got %v", expected, hands[0].handType)
	}
}

func TestJokerCard1(t *testing.T) {
	hands := InputToHands([]string{"32J4K 765"})
	expected := 1

	fmt.Printf("%+v \n", hands[0])

	if hands[0].handType != expected {
		t.Errorf("Expected handType to be %v, got %v", expected, hands[0].handType)
	}
}

func TestJokerCard2(t *testing.T) {
	hands := InputToHands([]string{"32J42 765"})
	expected := 3

	fmt.Printf("%+v \n", hands[0])

	if hands[0].handType != expected {
		t.Errorf("Expected handType to be %v, got %v", expected, hands[0].handType)
	}
}

func TestJokerCard3(t *testing.T) {
	hands := InputToHands([]string{"32J32 765"})
	expected := 4

	fmt.Printf("%+v \n", hands[0])

	if hands[0].handType != expected {
		t.Errorf("Expected handType to be %v, got %v", expected, hands[0].handType)
	}
}

func TestJokerCard4(t *testing.T) {
	hands := InputToHands([]string{"JJJJJ 765"})
	expected := 6

	fmt.Printf("%+v \n", hands[0])

	if hands[0].handType != expected {
		t.Errorf("Expected handType to be %v, got %v", expected, hands[0].handType)
	}
}

func TestJokerCard5(t *testing.T) {
	hands := InputToHands([]string{"JJJJQ 765"})
	expected := 6

	fmt.Printf("%+v \n", hands[0])

	if hands[0].handType != expected {
		t.Errorf("Expected handType to be %v, got %v", expected, hands[0].handType)
	}
}

func TestJokerCard6(t *testing.T) {
	hands := InputToHands([]string{"JJJQQ 765"})
	expected := 6

	fmt.Printf("%+v \n", hands[0])

	if hands[0].handType != expected {
		t.Errorf("Expected handType to be %v, got %v", expected, hands[0].handType)
	}
}

func TestJokerCard7(t *testing.T) {
	hands := InputToHands([]string{"JJQQQ 765"})
	expected := 6

	fmt.Printf("%+v \n", hands[0])

	if hands[0].handType != expected {
		t.Errorf("Expected handType to be %v, got %v", expected, hands[0].handType)
	}
}

func TestGetTotalWinningsForInput(t *testing.T) {
	godotenv.Load()
	expectedResult, _ := strconv.Atoi(os.Getenv("result_1"))

	Input, _ := utils.ReadInputFile("input.txt")

	hands := InputToHands(Input)
	result := GetTotalWinnings(hands)

	if result != expectedResult {
		t.Errorf("Expected result not to be %v, got %v", expectedResult, result)
	}
}
