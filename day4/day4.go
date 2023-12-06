package day4

import (
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	id             int
	winningNumbers []string
	ownNumbers     []string
}

func InputToCard(input string) (card Card) {
	id, _ := strconv.Atoi(strings.Trim(strings.Replace(strings.Split(input, ":")[0], "Card", "", -1), " "))
	return Card{
		id:             id,
		winningNumbers: strings.Split(strings.Split(strings.Split(input, ": ")[1], " | ")[0], " "),
		ownNumbers:     strings.Split(strings.Split(strings.Split(input, ": ")[1], " | ")[1], " "),
	}
}

func (card Card) Value() int {
	var pointValue int
	for _, number := range card.ownNumbers {
		numberAsInt, _ := strconv.Atoi(number)
		if slices.Contains(card.winningNumbers, strconv.Itoa(numberAsInt)) {
			if pointValue == 0 {
				pointValue = 1
			} else {
				pointValue *= 2
			}
		}
	}
	return pointValue
}

func TotalWorthOfScratchcards(input []string) (result int) {
	var sum int
	for _, gameInput := range input {
		sum += InputToCard(gameInput).Value()
	}
	return sum
}

// Part 2
func (card Card) NumberOfWinningCards() int {
	var numberOfWinningCards int
	for _, number := range card.ownNumbers {
		numberAsInt, _ := strconv.Atoi(number)
		if slices.Contains(card.winningNumbers, strconv.Itoa(numberAsInt)) {
			numberOfWinningCards++
		}
	}
	return numberOfWinningCards
}

func TotalAmountOfScratchcards(input []string) (result int) {
	for i := 0; i < len(input); i++ {
		card := InputToCard(input[i])
		numberOfWinningCards := card.NumberOfWinningCards()
		if numberOfWinningCards > 0 {
			input = append(input, input[card.id:card.id+numberOfWinningCards]...)
		}
	}

	return len(input)
}
