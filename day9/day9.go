package day9

import (
	"regexp"
	"strconv"
)

var numberRegex = regexp.MustCompile(`[\-0-9]+`)

func IsLastRow(cols []int) bool {
	for _, col := range cols {
		if col != 0 {
			return false
		}
	}
	return true
}

func FindNextInSequence(sequence []int) int {
	var rows = [][]int{sequence}
	var nextNumberInSequence int

	for i := 0; i < len(rows); i++ {
		if IsLastRow(rows[i]) {
			return nextNumberInSequence
		}

		var nextInSequence []int
		for j := 0; j < len(rows[i]); j++ {
			if j < len(rows[i])-1 {
				var result = rows[i][j+1] - rows[i][j]
				nextInSequence = append(nextInSequence, result)
			}

			if j == len(rows[i])-1 {
				nextNumberInSequence += rows[i][j]
			}
		}

		if len(nextInSequence) > 0 {
			rows = append(rows, nextInSequence)
		}
	}
	return nextNumberInSequence
}

func GetSequenceSum(input []string, revert bool) int {
	var rev = revert || false
	var sum int

	for _, line := range input {
		sequence := numberRegex.FindAllString(line, -1)
		var intSequence []int
		if !rev {
			for _, number := range sequence {
				num, _ := strconv.Atoi(number)
				intSequence = append(intSequence, num)
			}
		} else {
			for i := len(sequence) - 1; i > -1; i-- {
				num, _ := strconv.Atoi(sequence[i])
				intSequence = append(intSequence, num)
			}
		}
		sum += FindNextInSequence(intSequence)
	}

	return sum
}
