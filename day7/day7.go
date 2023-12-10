package day7

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var knownCards = map[string]int{
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"J": 10,
	"Q": 11,
	"K": 12,
	"A": 13,
}

var numberRegex = regexp.MustCompile("[0-9A-Z]")

type Hand struct {
	sorter     []int
	cards      string
	cardValues []int
	bid        int
	rank       int
	handType   int
}

func InputToHand(input string) Hand {
	var bid, _ = strconv.Atoi(strings.Split(input, " ")[1])
	var cards = strings.Split(input, " ")[0]

	hand := Hand{
		cards: cards,
		bid:   bid,
		rank:  0,
	}

	// var handType int
	var cardsAsInt = numberRegex.FindAllString(hand.cards, -1)
	var counter = make(map[string]int)
	var highest int
	var pairs int

	for _, card := range cardsAsInt {
		counter[card]++

		if counter[card]%2 == 0 {
			pairs++
		}

		if counter[card] > highest {
			highest = counter[card]
		}

		hand.cardValues = append(hand.cardValues, knownCards[card])
	}

	hand.handType = highest

	if pairs < 2 && hand.handType == 2 {
		hand.handType = 1
	} else if pairs >= 2 && hand.handType == 3 {
		hand.handType = 4
	}

	if highest >= 4 {
		hand.handType += 1
	}

	if highest == 1 {
		hand.handType = 0
	}

	hand.sorter = append([]int{hand.handType}, hand.cardValues...)

	return hand
}

func InputToHands(input []string) []Hand {
	var hands []Hand
	for _, line := range input {
		hands = append(hands, InputToHand((line)))
	}

	for i := 0; i < len(hands); i++ {
		for j := i + 1; j < len(hands); j++ {
			for k := 0; k < len(hands[i].sorter); k++ {
				if hands[i].sorter[k] > hands[j].sorter[k] {
					hands[i], hands[j] = hands[j], hands[i]
					break
				} else if hands[i].sorter[k] == hands[j].sorter[k] {
					continue
				} else {
					break
				}
			}
		}
	}

	return hands
}

func GetTotalWinnings(hands []Hand) int {
	totalWin := 0

	for handIndex, hand := range hands {
		hand.rank = handIndex + 1
		// fmt.Printf("%+v \n", hand)
		totalWin += hand.rank * hand.bid
	}

	fmt.Print(totalWin, "\n")
	return totalWin
}
