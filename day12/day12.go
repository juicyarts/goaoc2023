package day12

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

// shamefully translated from https://github.com/hyper-neutrino/advent-of-code/blob/main/2023/day12p2.py
// much more efficient than generating all the possible combinations as in my initial attempt
// even tho it looked funnier
func Count(record string, defects []int, cache map[string]int) int {
	key := record + strings.Join(strings.Fields(fmt.Sprint(defects)), "")

	if record == "" {
		if len(defects) == 0 {
			return 1
		}
		return 0
	}

	if len(defects) == 0 {
		if strings.Contains(record, "#") {
			return 0
		}
		return 1
	}

	if val, ok := cache[key]; ok {
		return val
	}

	result := 0

	if strings.Contains(".?", string(record[0])) {
		result += Count(record[1:], defects, cache)
	}

	if strings.Contains("#?", string(record[0])) {
		if defects[0] <= len(record) &&
			!strings.Contains(record[:defects[0]], ".") &&
			(defects[0] == len(record) || record[defects[0]] != '#') {
			if len(defects) == 1 && (len(record) == defects[0]+1 ||
				len(record) == defects[0]) {
				result += Count("", []int{}, cache)
			} else {
				if len(record) > defects[0]+1 && len(defects) > 1 || len(record) >= defects[0]+1 {
					result += Count(record[defects[0]+1:], defects[1:], cache)
				}
			}
		}
	}

	cache[key] = result
	return result
}

func stringListToArr(input string) []int {
	var result []int
	var inputAsArray = strings.Split(input, ",")

	for _, item := range inputAsArray {
		var itemAsInt, _ = strconv.Atoi(item)
		result = append(result, itemAsInt)
	}

	return result
}

func GetNumberOfArrangements(input string) int {
	var record = strings.Split(input, " ")[0]
	var damagedSpringGroups = stringListToArr(strings.Split(input, " ")[1])
	var damagedSpringGroupsRepeated []int
	var cache = make(map[string]int)

	record = strings.Join([]string{record, record, record, record, record}, "?")
	for i := 0; i < 5; i++ {
		damagedSpringGroupsRepeated = append(damagedSpringGroupsRepeated, damagedSpringGroups...)
	}

	var numberOfArragements = Count(record, damagedSpringGroupsRepeated, cache)

	return numberOfArragements
}

func GetTotalNumberOfArrangements(input []string) int {
	var totalMatches = 0
	var wg sync.WaitGroup

	for _, record := range input {
		wg.Add(1)
		go func(record string) {
			defer wg.Done()
			totalMatches += GetNumberOfArrangements(record)

		}(record)
	}

	wg.Wait()

	return totalMatches
}
