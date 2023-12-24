package day19

import (
	"maps"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var cRegex = regexp.MustCompile(`\<|\>|\=`)

func Main(input []string) (map[string][][]interface{}, []map[string]int, int) {
	workflows := make(map[string][][]interface{})
	parts := []map[string]int{}
	target := "wf"

	sum := 0

	for _, line := range input {
		if line == "" {
			target = "re"
			continue
		}

		if target == "wf" {
			wfKey, wf := makeWorkflow(line)
			workflows[wfKey] = wf
		} else {
			pt := makeParts(line)
			parts = append(parts, pt)
		}
	}

	for _, part := range parts {
		result := SortPart(part, "in", workflows)
		if result == "A" {
			for _, c := range part {
				sum += c
			}
		}
	}

	combinations := findMinMax("in", workflows, map[string][]int{
		"x": {1, 4000},
		"m": {1, 4000},
		"a": {1, 4000},
		"s": {1, 4000},
	})

	return workflows, parts, combinations
}

func makeWorkflow(l string) (string, [][]interface{}) {
	k := strings.Split(l, "{")[0]
	v := strings.Replace(strings.Split(l, "{")[1], "}", "", -1)
	c := strings.Split(v, ",")
	conditions := [][]interface{}{}

	for _, con := range c {
		if len(strings.Split(con, ":")) > 1 {
			source, dest := strings.Split(con, ":")[0], strings.Split(con, ":")[1]
			key, val := cRegex.Split(source, -1)[0], cRegex.Split(source, -1)[1]
			valAsInt, _ := strconv.Atoi(val)
			condition := cRegex.FindString(source)

			conditions = append(conditions, []interface{}{dest, key, condition, valAsInt})
		} else {
			conditions = append(conditions, []interface{}{con})
		}
	}

	return k, conditions
}

func makeParts(l string) map[string]int {
	v := strings.Replace(strings.Split(l, "{")[1], "}", "", -1)
	c := strings.Split(v, ",")
	conditions := map[string]int{}

	for _, con := range c {
		key, val := cRegex.Split(con, -1)[0], cRegex.Split(con, -1)[1]
		conditions[key], _ = strconv.Atoi(val)
	}

	return conditions
}

func SortPart(part map[string]int, dest string, workflows map[string][][]interface{}) string {
	if workflows[dest] != nil {
		for _, rule := range workflows[dest] {
			if len(rule) == 1 {
				if rule[0] == 'A' || rule[0] == 'R' {
					return rule[0].(string)
				}

				return SortPart(part, rule[0].(string), workflows)
			} else {
				newDest := rule[0]
				keyToCheck := rule[1]
				condition := rule[2]
				valueToCheck := rule[3]
				doesMatch := conditionMatches(condition.(string), part[keyToCheck.(string)], valueToCheck.(int))
				if doesMatch {
					return SortPart(part, newDest.(string), workflows)
				}
				if newDest == 'A' || newDest == 'R' {
					return newDest.(string)
				}
			}
		}
	} else {
		return dest
	}

	panic("No Result found for part")
}

func findMinMax(dest string, workflows map[string][][]interface{}, minMaxMap map[string][]int) int {
	if dest == "R" {
		return 0
	}

	if dest == "A" {
		combinations := 1
		for _, entry := range minMaxMap {
			lo, hi := entry[0], entry[1]
			combinations *= int(math.Abs(float64(hi - lo + 1)))
		}

		return combinations
	}

	rules := workflows[dest]
	result := 0

	for i := 0; i < len(rules); i++ {
		rule := rules[i]
		nextDest := rule[0]
		if len(rule) == 1 {
			result += findMinMax(nextDest.(string), workflows, minMaxMap)
		} else {
			nextDest := rule[0]
			key := rule[1]
			condition := rule[2]
			value := rule[3]
			lo, hi := minMaxMap[key.(string)][0], minMaxMap[key.(string)][1]
			var a, b []int

			if condition == "<" {
				a = []int{lo, min(value.(int)-1, hi)}
				b = []int{max(value.(int), lo), hi}
			} else {
				a = []int{max(value.(int)+1, lo), hi}
				b = []int{lo, min(value.(int), hi)}
			}

			if a[0] <= a[1] {
				copy := maps.Clone(minMaxMap)
				copy[key.(string)] = a
				result += findMinMax(nextDest.(string), workflows, copy)
			}

			if b[0] <= b[1] {
				minMaxMap[key.(string)] = b
			} else {
				break
			}
		}
	}

	return result
}

func conditionMatches(condition string, a int, b int) bool {
	if condition == "<" {
		return a < b
	} else {
		return a > b
	}
}
