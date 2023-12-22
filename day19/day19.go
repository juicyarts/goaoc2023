package day19

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

var cRegex = regexp.MustCompile(`\<|\>|\=`)

func Main(input []string) (map[string][][]interface{}, []map[string]int, int) {
	workflows := make(map[string][][]interface{})
	parts := []map[string]int{}

	target := "wf"

	for _, line := range input {
		if line == "" {
			target = "re"
			continue
		}

		if target == "wf" {
			wfKey, wf := makeWorkflows(line)
			workflows[wfKey] = wf
		} else {
			pt := makeParts(line)
			parts = append(parts, pt)
		}
	}

	sum := 0
	for _, part := range parts {
		result := SortPart(part, "in", workflows)
		if result == "A" {
			for _, c := range part {
				sum += c
			}
		}
	}

	minMaxMap := map[string][]int{}

	for key := range workflows {
		minMaxMap = findMinMax(key, workflows, minMaxMap)
	}

	diffs := []int{}

	for _, entry := range minMaxMap {
		diffs = append(diffs, entry[0]-entry[1])
	}

	spew.Dump(diffs)
	return workflows, parts, sum
}

func makeWorkflows(l string) (string, [][]interface{}) {
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

func findMinMax(dest string, workflows map[string][][]interface{}, minMaxMap map[string][]int) map[string][]int {
	if dest == "A" || dest == "R" {
		return minMaxMap
	}

	for _, wf := range workflows {
		for _, rule := range wf {
			if len(rule) == 1 {
				return minMaxMap
			} else {
				nextDest := rule[0]
				key := rule[1]
				condition := rule[2]
				value := rule[3]
				if _, ok := minMaxMap[key.(string)]; !ok {
					minMaxMap[key.(string)] = []int{0, 4000}
				}

				if condition == "<" {
					if minMaxMap[key.(string)][1] >= value.(int) {
						minMaxMap[key.(string)][1] = value.(int) - 1
					}
				} else {
					if minMaxMap[key.(string)][0] <= value.(int) {
						minMaxMap[key.(string)][0] = value.(int) - 1
					}
				}

				minMaxMap = findMinMax(nextDest.(string), workflows, minMaxMap)
			}
		}
	}

	return minMaxMap
}

func conditionMatches(condition string, a int, b int) bool {
	if condition == "<" {
		return a < b
	} else {
		return a > b
	}
}
