package day1

import (
	"fmt"
	"strconv"
	"strings"
)

// Part 1
func FindFirstAndLastDigit(value string) int {
	var firstDigit, lastDigit int

	for _, char := range value {
		if char >= '0' && char <= '9' {
			if firstDigit == 0 {
				firstDigit = int(char - '0')
			}
			lastDigit = int(char - '0')
		}
	}

	concatenatedDigits, _ := strconv.Atoi(strconv.Itoa(firstDigit) + strconv.Itoa(lastDigit))
	return concatenatedDigits
}

func SumCalbirationValues(values []string) int {
	var sum int
	for _, value := range values {
		sum += FindFirstAndLastDigit(value)
	}
	return sum
}

// Part 2
var legalWords = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func FindAllIndexesOfSubstringInString(s, sub string) []int {
	var indexes []int
	cumulativeIndex := 0
	for {
		index := strings.Index(s, sub)
		if index == -1 {
			break
		}
		indexes = append(indexes, index+cumulativeIndex)
		cumulativeIndex += index + len(sub)
		s = s[index+len(sub):]
	}
	return indexes
}

func FindFirstAndLastDigitOrWord(value string) int {
	var firstDigit, firstDigitIndex, lastDigit, lastDigitIndex int

	for index, char := range value {
		if char >= '0' && char <= '9' {
			if firstDigit == 0 {
				firstDigit = int(char - '0')
				firstDigitIndex = index
			}

			lastDigit = int(char - '0')
			lastDigitIndex = index
		}
	}

	for word := range legalWords {
		foundIndexes := FindAllIndexesOfSubstringInString(value, word)
		for _, index := range foundIndexes {
			if index <= firstDigitIndex {
				firstDigit = legalWords[word]
				firstDigitIndex = index
			}
			if index >= lastDigitIndex {
				lastDigit = legalWords[word]
				lastDigitIndex = index
			}
		}
	}

	concatenatedDigits, _ := strconv.Atoi(strconv.Itoa(firstDigit) + strconv.Itoa(lastDigit))
	return concatenatedDigits
}

func SumCalbirationValuesWithWords(values []string) int {
	var sum int
	for _, value := range values {
		res := FindFirstAndLastDigitOrWord(value)
		fmt.Println(res)
		sum += res
	}
	return sum
}
