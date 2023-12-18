package day14

import (
	"aoc2023/utils"
	"regexp"
	"strings"
)

var solidRockRegex = regexp.MustCompile(`\#`)
var dotRegex = regexp.MustCompile(`\.`)

func Main(input []string) int {
	// rotate input anti clockwise two times initially to make sorting easier in general
	// and because Cycle will do one Clockwise rotation before sorting
	// might be a bit of overhead but it works
	rotated := utils.RotateAntiClockWise(utils.RotateAntiClockWise(input))
	// multiply with 4 for rotations, 1 will only do the first rotation to get the array sideways for easier sorting
	cycles := 4 * 1000000000
	cycled := Cycle(rotated, (cycles))

	weight := 0

	for _, row := range utils.RotateClockWise(cycled) {
		weight += calculateWeight(row)
	}

	return weight
}

// cache each row that was sorted once
var cache = make(map[string]string)
// TODO: instead of keeping two caches one could reduce to one here
// cache each pattern that we came across
var patternCache = make(map[string][]string)
// cache each cycle length until a pattern is reached again
var patternCycleCache = make(map[string]int)

func Cycle(input []string, times int) []string {
	output := make([]string, len(input))
	copy(output, input)

	for i := 0; i < times; i++ {
		output = sortRows(utils.RotateClockWise(output))
		pattern := strings.Join(output, " ")
		if _, ok := patternCycleCache[pattern]; ok {
			remainingIterations := times - i
			cycleLength := i - patternCycleCache[pattern]
			iterationsToSkip := remainingIterations % cycleLength
			i = times - iterationsToSkip
		}
		patternCycleCache[pattern] = i
	}

	return output
}

func SortRow(row string) string {
	if cached, ok := cache[row]; ok {
		return cached
	}

	for i := 0; i < len(row); i++ {
		if row[i] == '.' || row[i] == '#' {
			continue
		}
		dotLocations := dotRegex.FindAllStringIndex(row[:i], -1)
		brickLocations := solidRockRegex.FindAllStringIndex(row[:i], -1)
		if row[i] == 'O' {
			if len(dotLocations) == 0 {
				continue
			}
			if len(brickLocations) > 0 {
				targetIndex := brickLocations[len(brickLocations)-1][0] + 1
				row = row[:targetIndex] + "O" + row[targetIndex:i] + row[i+1:]
			} else if len(dotLocations) > 0 {
				targetIndex := dotLocations[0][0]
				row = row[:targetIndex] + "O" + row[targetIndex:i] + row[i+1:]
			}
		}
	}

	cache[row] = row
	return row
}

func sortRows(input []string) []string {
	output := make([]string, len(input))
	if cached, ok := patternCache[strings.Join(input, " ")]; ok {
		return cached
	}
	for i, row := range input {
		output[i] = SortRow(row)
	}
	patternCache[strings.Join(input, " ")] = output
	return output
}

func calculateWeight(row string) int {
	weight := 0
	for charIndex, char := range row {
		if char == 'O' {
			weight += len(row) - charIndex
		}
	}
	return weight
}
