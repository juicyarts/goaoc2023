package day15

import (
	"regexp"
	"strconv"
	"strings"
)

// is there a no magic numbers rule for go ?
var amountOfBoxes = 256
var multiplier = 17
var operatorRegex = regexp.MustCompile(`\=|\-+`)
var digitRegex = regexp.MustCompile(`\d+`)

func Hash(input string) int {
	sum := 0
	for _, v := range input {
		sum += int(v)
		sum *= multiplier
		sum %= amountOfBoxes
	}
	return sum
}

func HashSum(input string) (int, int) {
	inputs := strings.Split(strings.Split(input, "\n")[0], ",")
	hashSum := 0
	totalFocusPower := 0
	boxToLensMap := make(map[int][][]string, amountOfBoxes)

	for _, v := range inputs {
		label := operatorRegex.Split(v, -1)[0]
		operation := operatorRegex.FindString(v)
		focalLength := digitRegex.FindString(v)

		if _, ok := boxToLensMap[Hash(label)]; !ok {
			boxToLensMap[Hash(label)] = [][]string{}
		}

		lenseEntryExists := false

		for i := 0; i < len(boxToLensMap[Hash(label)]); i++ {
			if len(boxToLensMap[Hash(label)][i]) > 0 {
				if boxToLensMap[Hash(label)][i][0] == label {
					lenseEntryExists = true
					break
				}
			}
		}

		if !lenseEntryExists {
			if _, ok := boxToLensMap[Hash(label)]; ok {
				boxToLensMap[Hash(label)] = append(boxToLensMap[Hash(label)], []string{label})
			}
		}

		for i := 0; i < len(boxToLensMap[Hash(label)]); i++ {
			if boxToLensMap[Hash(label)][i][0] == label {
				if operation == "=" {
					boxToLensMap[Hash(label)][i] = []string{label, focalLength}
				} else if operation == "-" {
					boxToLensMap[Hash(label)] = append(boxToLensMap[Hash(label)][:i], boxToLensMap[Hash(label)][i+1:]...)
				}
			}
		}

		hashSum += Hash(v)
	}

	for key, value := range boxToLensMap {
		box := value
		for j := 0; j < len(box); j++ {
			focalLengthAsInt, _ := strconv.Atoi(box[j][1])
			focusPower := (1 + key) * (j + 1) * focalLengthAsInt
			box[j] = []string{box[j][0], box[j][1], strconv.Itoa(focusPower)}
			totalFocusPower += focusPower
		}
	}

	return hashSum, totalFocusPower
}
