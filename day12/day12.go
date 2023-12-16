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

func CalculateMatches(record string, damagedSpringGroups []string, validVariantMap map[string]bool, shouldAddToMap bool) int {
	var stringHasJoker = jokerRegex.MatchString(record)

	if !stringHasJoker {
		if variantHasValidAmountOfGroups(record, damagedSpringGroups) {
			fmt.Print(yellow(record), " ", white(damagedSpringGroups), "\n")
			if shouldAddToMap {
				validVariantMap[record] = true
			}
			return 1
		} else {
			return 0
		}
	}

	fmt.Printf("%+v \n", white(record))

	jokerIndex := jokerRegex.FindStringIndex(record)[0]
	recordA := replaceAtIndex(record, "#", jokerIndex)
	a := CalculateMatches(recordA, damagedSpringGroups, validVariantMap, shouldAddToMap)
	recordB := replaceAtIndex(record, ".", jokerIndex)
	b := CalculateMatches(recordB, damagedSpringGroups, validVariantMap, shouldAddToMap)

	return a + b
}

func makeVariantCopies(variant string, copyAmount int, validVariants map[string]bool, copies []string, damagedSpringGroups []string) int {
	matches := 0

	if copyAmount == 0 {
		matches += CalculateMatches(variant, damagedSpringGroups, validVariants, false)
		return matches
	}

	for validVariantOut := range validVariants {
		matches += makeVariantCopies(variant+"?"+validVariantOut, copyAmount-1, validVariants, copies, damagedSpringGroups)
	}

	return matches
}

func GetNumberOfArrangements(input string) (int, int) {
	var conditionRecord = strings.Split(input, " ")[0]
	var damagedSpringGroups = strings.Split(strings.Split(input, " ")[1], ",")
	var validVariantMap = make(map[string]bool)

	var simpleMatches = CalculateMatches(conditionRecord, damagedSpringGroups, validVariantMap, true)
	var complexMatches = 0
	var newDamagedSpringGroups = []string{}
	for i := 0; i < 5; i++ {
		newDamagedSpringGroups = append(newDamagedSpringGroups, damagedSpringGroups...)
	}

	for key := range validVariantMap {
		complexMatches += makeVariantCopies(key, 4, validVariantMap, []string{}, newDamagedSpringGroups)
	}

	fmt.Printf("%+v %+v \n", simpleMatches, complexMatches)

	return simpleMatches, complexMatches
}

func GetTotalNumberOfArrangements(input []string) (int, int) {
	var result = []int{}
	var wg sync.WaitGroup

	for _, record := range input {
		wg.Add(1)
		go func(record string) {
			defer wg.Done()
			simpleMatches, complexMatches := GetNumberOfArrangements(record)
			result[0] += simpleMatches
			result[1] += complexMatches
		}(record)
	}

	wg.Wait()

	return result[0], result[1]
}
