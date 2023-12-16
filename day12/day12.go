package day12

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/fatih/color"
)

var yellow = color.New(color.FgYellow).SprintFunc()
var white = color.New(color.FgWhite).SprintFunc()

func replaceAtIndex(s string, replacement string, index int) string {
	return s[:index] + replacement + s[index+1:]
}

var jokerRegex = regexp.MustCompile(`\?`)

func variantHasValidAmountOfGroups(variant string, damagedSpringGroups []string) bool {
	normalizedVariant := strings.ReplaceAll(variant, "?", ".")
	var variantAsArray = strings.Split(normalizedVariant, ".")

	hasValidAmountOfGroups := false
	variantPartsWithoutEmpty := []string{}

	for _, variantPart := range variantAsArray {
		if variantPart == "" {
			continue
		}

		variantPartsWithoutEmpty = append(variantPartsWithoutEmpty, variantPart)
	}

	if len(variantPartsWithoutEmpty) != len(damagedSpringGroups) {
		return false
	}

	for variantIndex, variantPart := range variantPartsWithoutEmpty {
		var damagedSpringGroupSize, _ = strconv.Atoi(damagedSpringGroups[variantIndex])
		if len(variantPart) != damagedSpringGroupSize {
			hasValidAmountOfGroups = false
			break
		} else {
			hasValidAmountOfGroups = true
		}
	}

	return hasValidAmountOfGroups
}

func CalculateMatches(record string, damagedSpringGroups []string) int {
	var stringHasJoker = jokerRegex.MatchString(record)

	if !stringHasJoker {
		if variantHasValidAmountOfGroups(record, damagedSpringGroups) {
			fmt.Print(yellow(record), " ", white(damagedSpringGroups), "\n")
			return 1
		} else {
			return 0
		}
	}

	jokerIndex := jokerRegex.FindStringIndex(record)[0]
	recordA := replaceAtIndex(record, "#", jokerIndex)
	a := CalculateMatches(recordA, damagedSpringGroups)
	recordB := replaceAtIndex(record, ".", jokerIndex)
	b := CalculateMatches(recordB, damagedSpringGroups)

	return a + b
}

func GetNumberOfArrangements(input string) int {
	var conditionRecord = strings.Split(input, " ")[0]
	var damagedSpringGroups = strings.Split(strings.Split(input, " ")[1], ",")

	fmt.Printf("------------------------------------------------------\n")
	fmt.Printf("Record: %s %s \n", conditionRecord, damagedSpringGroups)
	var matches = CalculateMatches(conditionRecord, damagedSpringGroups)
	fmt.Printf("%d Total Matches \n", matches)
	fmt.Printf("--------------------\n")

	return matches
}

func GetTotalNumberOfArrangements(input []string) int {
	var totalNumberOfArrangements = 0
	var wg sync.WaitGroup

	for _, record := range input {
		wg.Add(1)
		go func(record string) {
			defer wg.Done()
			totalNumberOfArrangements += GetNumberOfArrangements(record)
		}(record)
	}

	wg.Wait()

	return totalNumberOfArrangements
}
