package day15

import (
	"strings"
)

func Hash(input string) int {
	sum := 0
	for _, v := range input {
		sum += int(v)
		sum *= 17
		sum %= 256
	}
	return sum
}

func HashSum(input string) int {
	inputs := strings.Split(strings.Split(input, "\n")[0], ",")
	sum := 0

	for _, v := range inputs {
		sum += Hash(v)
	}

	return sum
}
