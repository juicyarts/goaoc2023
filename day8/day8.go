package day8

import (
	"regexp"
	"strings"
)

var parensRegex = regexp.MustCompile(`\(|\)`)
var charRegex = regexp.MustCompile("[A-Z]")

type Game struct {
	directions []string
	m          map[string][]string
}

type StepInstruction struct {
	location       string
	directionIndex int
	steps          int
}

func InputToMap(input []string) Game {
	m := make(map[string][]string)
	var directions []string

	for _, line := range input {
		if line == "" {
			continue
		}

		parts := strings.Split(line, " = ")
		if len(parts) == 1 {
			directions = charRegex.FindAllString(parts[0], -1)
			continue
		}

		var foo = strings.Split(strings.TrimSpace(parensRegex.ReplaceAllString(parts[1], "")), ", ")
		m[parts[0]] = foo
	}

	var gameMap = Game{
		directions: directions,
		m:          m,
	}

	return gameMap
}

func (game Game) FindStepsToEnd(location string, directionIndex int, steps int, suffix string) StepInstruction {

	if strings.HasSuffix(location, suffix) {
		return StepInstruction{
			location:       location,
			directionIndex: directionIndex,
			steps:          steps,
		}
	}
	steps++

	if len(game.directions) == directionIndex {
		directionIndex = 0
	}

	var dir = game.directions[directionIndex]

	if dir == "R" {
		return game.FindStepsToEnd(game.m[location][1], directionIndex+1, steps, suffix)
	}

	return game.FindStepsToEnd(game.m[location][0], directionIndex+1, steps, suffix)
}

func (game Game) FindnextStep(games map[string][]string, nextStep StepInstruction, gameStatus []int, directionIndex int, totalSteps int) []int {

	var gameIndex = 0
	for name := range games {
		nextStep = game.FindStepsToEnd(name, directionIndex, 0, "Z")
		gameStatus[gameIndex] = nextStep.steps
		gameIndex++
		totalSteps *= nextStep.steps
	}

	return gameStatus
}

// Part 2
func (game Game) FindStepsToEndForEachEndingWithA() int {

	var instructionsEndingWithA = make(map[string][]string)
	for name, instruction := range game.m {
		if strings.HasSuffix(name, "A") {
			instructionsEndingWithA[name] = instruction
		}
	}
	var gameStatus = make([]int, len(instructionsEndingWithA))
	var maxSteps = game.FindnextStep(instructionsEndingWithA, StepInstruction{}, gameStatus, 0, 0)

	gcd := maxSteps[0]
	for i := 1; i < len(maxSteps); i++ {
		gcd = gcd * maxSteps[i] / int(GCD(int(gcd), int(maxSteps[i])))
	}

	return gcd
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
