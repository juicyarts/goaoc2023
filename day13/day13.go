package day13

import (
	"fmt"

	"github.com/fatih/color"
)

var blue = color.New(color.FgBlue).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()

func compareStrings(a, b string, allowedMisses int) (bool, int) {
	misses := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			misses++
		}
		if misses > allowedMisses {
			return false, misses
		}
	}
	return true, misses
}

func findReflections(rows []string, smudge int) (int, int, bool) {
	mirrors := make(map[int][]int)
	printString := ""
	for rowIndex := 0; rowIndex < len(rows); rowIndex++ {
		var l = rows[rowIndex]
		if rowIndex >= 0 && rowIndex < len(rows)-1 {
			var linesAreEqual, misses = compareStrings(l, rows[rowIndex+1], smudge)
			if linesAreEqual {
				printString += fmt.Sprintf("%+s \n", green(string(l), " < Found Potential Mirror at ", rowIndex+1))
				for i := 0; i <= rowIndex; i++ {
					if len(rows) > rowIndex+i+1 && rowIndex-i >= 0 {
						var linesAreEqualInner, missesInner = compareStrings(rows[rowIndex-i], rows[rowIndex+i+1], smudge)
						if linesAreEqualInner {
							if _, ok := mirrors[rowIndex]; !ok {
								mirrors[rowIndex] = []int{rowIndex, 1, misses + missesInner}
							} else {
								mirrors[rowIndex][1] += 1
								mirrors[rowIndex][2] += misses + missesInner
							}
						} else {
							delete(mirrors, rowIndex)
						}
					}
				}
			} else {
				printString += fmt.Sprintf("%+s \n", blue(string(l)))
			}
		} else {
			printString += fmt.Sprintf("%+s \n", blue(string(l)))
		}
	}

	fmt.Print("\n", printString, "\n")

	mostValuableMirrorIndex := 0
	mostValuableMirrorAmount := 0
	mirrorHasSmudge := false

	if len(mirrors) == 0 {
		return 0, 0, false
	}

	for _, mirror := range mirrors {
		if smudge > 0 {
			if !mirrorHasSmudge && mirror[2] > 0 {
				mirrorHasSmudge = true
				mostValuableMirrorIndex = mirror[0]
				mostValuableMirrorAmount = mirror[1]
			} else if (mirrorHasSmudge && mirror[2] > 0 && mirror[1] > mostValuableMirrorAmount) ||
				(!mirrorHasSmudge && mirror[2] == 0 && mirror[1] > mostValuableMirrorAmount) {
				mostValuableMirrorIndex = mirror[0]
				mostValuableMirrorAmount = mirror[1]
			}

		} else {
			if mirror[1] > mostValuableMirrorAmount {
				mostValuableMirrorIndex = mirror[0]
				mostValuableMirrorAmount = mirror[1]
			}
		}
	}

	if mostValuableMirrorAmount == 0 || !mirrorHasSmudge {
		return 0, 0, false
	}

	return mostValuableMirrorIndex + 1, mostValuableMirrorAmount, mirrorHasSmudge
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

		fmt.Printf("-------START %v---------\n", i)
		horizontalMirrorIndex, horizontalMirrorAmount, _ := findReflections(pattern[0], 1)
		verticalMirrorsIndex, verticalMirrorsAmount, _ := findReflections(pattern[1], 1)
		if verticalMirrorsAmount == 0 && horizontalMirrorAmount == 0 {
			continue
		}

		sum += 100 * horizontalMirrorIndex
		sum += verticalMirrorsIndex
		fmt.Print("-------END--------------\n")
	}

	return sum
}
