package day3

import (
	"regexp"
	"strconv"
)

// Welcome to spaghetti code land

var specialCharsRegex = regexp.MustCompile(`[^a-zA-Z0-9.\s]`)
var numberRegex = regexp.MustCompile("[0-9]+")
var asteriskRegex = regexp.MustCompile(`\*`)

func SumOfPartNumbers(input []string) (result int) {
	var validPartNumbers []int

	for lineIndex, line := range input {
		numbers := numberRegex.FindAllString(line, -1)
		numbersIndexes := numberRegex.FindAllStringIndex(line, -1)

		for nIndex, number := range numbers {
			numberIndex := numbersIndexes[nIndex]
			numberAsInt, _ := strconv.Atoi(number)

			var startIndex = 0
			var endIndex = len(line)

			if numberIndex[0] > 0 {
				startIndex = numberIndex[0] - 1
			}

			if numberIndex[1] < len(line) {
				endIndex = numberIndex[1] + 1
			}

			stringToCheck := line[startIndex:endIndex]

			if lineIndex > 0 {
				stringToCheck += input[lineIndex-1][startIndex:endIndex]
			}

			if lineIndex < len(input)-1 {
				stringToCheck += input[lineIndex+1][startIndex:endIndex]
			}

			foundSpecialChars := specialCharsRegex.FindString(stringToCheck)

			if foundSpecialChars != "" {
				validPartNumbers = append(validPartNumbers, numberAsInt)
			}
		}
	}

	var sum int
	for _, value := range validPartNumbers {
		sum += value
	}

	return sum
}

// Part 2
func SumOfGearRatios(input []string) (result int) {

	var validGearRatios []int

	for lineIndex, line := range input {
		asterisks := asteriskRegex.FindAllString(line, -1)
		asterisksIndexes := asteriskRegex.FindAllStringIndex(line, -1)

		for aIndex := range asterisks {
			asteriskIndex := asterisksIndexes[aIndex]

			var startIndex = 0
			var endIndex = len(line)

			if asteriskIndex[0] > 3 {
				startIndex = asteriskIndex[0] - 3
			}

			if asteriskIndex[1] < len(line)-3 {
				endIndex = asteriskIndex[1] + 3
			}

			numbers := numberRegex.FindAllString(line[startIndex:endIndex], -1)
			numbersIndexes := numberRegex.FindAllStringIndex(line[startIndex:endIndex], -1)

			if lineIndex > 0 {
				numbers = append(numbers, numberRegex.FindAllString(input[lineIndex-1][startIndex:endIndex], -1)...)
				numbersIndexes = append(numbersIndexes, numberRegex.FindAllStringIndex(input[lineIndex-1][startIndex:endIndex], -1)...)
			}

			if lineIndex < len(input)-1 {
				numbers = append(numbers, numberRegex.FindAllString(input[lineIndex+1][startIndex:endIndex], -1)...)
				numbersIndexes = append(numbersIndexes, numberRegex.FindAllStringIndex(input[lineIndex+1][startIndex:endIndex], -1)...)
			}

			var validNumbers []string
			var ratio int = 1

			newStartIndex := 2
			newEndIndex := 4

			for numberIndex, value := range numbers {
				if (numbersIndexes[numberIndex][0] >= newStartIndex && numbersIndexes[numberIndex][0] <= newEndIndex) ||
					(numbersIndexes[numberIndex][1]-1 >= newStartIndex && numbersIndexes[numberIndex][1]-1 <= newEndIndex) {
					validNumbers = append(validNumbers, value)
				}
			}

			if len(validNumbers) == 2 {
				for _, value := range validNumbers {
					valueAsNumber, _ := strconv.Atoi(value)
					ratio *= valueAsNumber
				}
				validGearRatios = append(validGearRatios, ratio)
			}
		}
	}

	var sum int
	for _, value := range validGearRatios {
		sum += value
	}

	return sum
}
