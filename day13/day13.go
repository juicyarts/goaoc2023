package day13

import (
	"fmt"
	"slices"

	"github.com/fatih/color"
)

var Ash = "."
var Fog = "#"

var blue = color.New(color.FgBlue).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()
var cyan = color.New(color.FgCyan).SprintFunc()

func findReflections(rows []string) (int, int) {
	mirrors := make(map[int][]int)
	printString := ""

	for rowIndex := 0; rowIndex < len(rows); rowIndex++ {
		var l = rows[rowIndex]
		if rowIndex >= 0 && rowIndex < len(rows)-1 {
			if l == rows[rowIndex+1] {
				if _, ok := mirrors[rowIndex]; !ok {
					mirrors[rowIndex] = []int{rowIndex, 0}
				}

				printString += fmt.Sprintf("	%+s \n", green(string(l), " < Found Mirror at ", rowIndex+1))

				var depth = 0

				for i := 0; i <= rowIndex; i++ {
					// fmt.Print("Doin stuff \n", rows[i], " ", rowIndex, i, rowIndex-i, rowIndex+i, "\n")
					if len(rows) > rowIndex+i+1 && rowIndex-i >= 0 {
						// fmt.Printf("%+v \n%+v \n\n", rows[rowIndex-i], rows[rowIndex+i+1])
						if rows[rowIndex-i] == rows[rowIndex+i+1] {
							depth++
							// printString += fmt.Sprintf("%+s \n", cyan(string(rows[rowIndex+i]), " < Found Reflection"))
							mirrors[rowIndex][1] = depth
						} else {
							break
						}
					}
				}

			} else {
				printString += fmt.Sprintf("	%+s \n", blue(string(l)))
			}
		} else {
			printString += fmt.Sprintf("	%+s \n", blue(string(l)))
		}
	}

	fmt.Print("\n", printString, "\n")

	mostValuableMirrorIndex := 0
	mostValuableMirrorAmount := 0

	if len(mirrors) == 0 {
		return 0, 0
	}

	for _, mirror := range mirrors {
		if mirror[1] > mostValuableMirrorAmount ||
			mirror[1] == mostValuableMirrorAmount && mirror[0] > mostValuableMirrorIndex {
			mostValuableMirrorIndex = mirror[0]
			mostValuableMirrorAmount = mirror[1]
		}
	}

	if mostValuableMirrorAmount == 0 {
		fmt.Print(red("No mirrors found \n"))
		return 0, 0
	}

	return mostValuableMirrorIndex + 1, mostValuableMirrorAmount
}

func FindReflections(input []string) int {
	sum := 0
	patternIndex := 0

	var patterns = make(map[int][][]string)

	for i := 0; i < len(input); i++ {
		var line = input[i]
		if line == "" {
			patternIndex++
			continue
		}

		if _, ok := patterns[patternIndex]; !ok {
			patterns[patternIndex] = [][]string{{}, make([]string, len(line))}
		}

		patterns[patternIndex][0] = append(patterns[patternIndex][0], line)

		for j := 0; j < len(line); j++ {
			patterns[patternIndex][1][j] += string(line[j])
		}
	}

	for i := 0; i < len(patterns); i++ {
		pattern := patterns[i]

		fmt.Print("-------START---------\n")
		fmt.Print("-------Horizontal Test-------\n")
		horizontalMirrorIndex, horizontalMirrorAmount := findReflections(pattern[0])
		fmt.Print("-------Vertical Test---------\n")
		verticalMirrorsIndex, verticalMirrorsAmount := findReflections(pattern[1])

		var test1 = mirror(pattern[0])
		var test2 = mirror(pattern[1])

		fmt.Printf("TestH: %+v TestV: %+v \n", test1, test2)
		fmt.Printf("OwnH: %+v (%+v) OwnV: %+v (%+v) \n", horizontalMirrorIndex, horizontalMirrorAmount, verticalMirrorsIndex, verticalMirrorsAmount)

		if verticalMirrorsAmount == 0 && horizontalMirrorAmount == 0 {
			fmt.Print(red("No mirrors found \n"))
			continue
		}

		if horizontalMirrorAmount > verticalMirrorsAmount {
			if horizontalMirrorIndex != test1 {
				fmt.Print(red("Values don't match! \n"))
			}
			fmt.Printf("Chose horizontal %v vs %v \n", horizontalMirrorIndex, test1)
			sum += 100 * horizontalMirrorIndex
		} else {
			if verticalMirrorsIndex != test2 {
				fmt.Print(red("Values don't match! \n"))
			}

			fmt.Printf("Chose vertical %v vs %v \n", verticalMirrorsIndex, test2)
			sum += verticalMirrorsIndex
		}

		fmt.Print("-------END---------\n")
	}

	return sum
}

func mirror(input []string) int {
	for i := 1; i < len(input); i++ {
		l := slices.Min([]int{i, len(input) - i})
		s1, s2 := slices.Clone(input[i-l:i]), input[i:i+l]
		slices.Reverse(s1)
		if slices.Equal(s1, s2) {
			return i
		}
	}
	return 0
}
