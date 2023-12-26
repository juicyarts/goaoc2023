package day20

import (
	"strings"
)

type Inst struct {
	source string
	target string
	pulse  bool
}

type FFModule struct {
	targets []string
	memory  bool
}

type CModule struct {
	targets []string
	memory  map[string]bool
}

func MeasurePulses(input []string) int {
	// having two collections of modules here cuz too lazy to use interface type

	// FlipFlop module
	ffmodules := map[string]FFModule{}
	// conjunction module
	cmodules := map[string]CModule{}

	broadcastInstructions := []Inst{}

	for _, l := range input {
		s, t := strings.Split(l, " -> ")[0], strings.Split(l, " -> ")[1]
		mt, mk := s[:1], s[1:]
		targets := strings.Split(t, ", ")

		if s == "broadcaster" {
			for _, target := range targets {
				broadcastInstructions = append(broadcastInstructions, Inst{source: s, target: target, pulse: false})
			}
		} else if mt == "%" {
			ffmodules[mk] = FFModule{targets: targets, memory: false}
		} else if mt == "&" {
			cmodules[mk] = CModule{
				targets: targets,
				memory:  map[string]bool{},
			}
		}
	}

	// we keep a cache of all modules pointing to the module "rx" (only "cs")
	// since cs is a cmodule ("&") and we want "rx" to receive a "low" Pulse (false)
	// all inputs need to have passed a "high" Pulse (true) to reach our final condition
	visistedCache := map[string]int{}

	// for each module that is connected to a conjunction module
	// pass a falsy pulse value to the module cache of that conjunction module
	for n, mod := range ffmodules {
		for _, target := range mod.targets {
			if _, exists := cmodules[target]; exists {
				cmodules[target].memory[n] = false
			}

			// keep track of modules that point to "cs"
			if target == "cs" {
				visistedCache[n] = 0
			}
		}
	}

	for n, mod := range cmodules {
		for _, target := range mod.targets {
			if _, exists := cmodules[target]; exists {
				cmodules[target].memory[n] = false
			}

			if target == "cs" {
				visistedCache[n] = 0
			}
		}
	}

	repetition := 1
	cycleLengths := map[string]int{}

	for i := 0; i < repetition; i++ {
		instructionQueue := broadcastInstructions
		for j := 0; j < len(instructionQueue); j++ {
			instruction := instructionQueue[j]

			// shamelessly copied from https://github.com/hyper-neutrino/advent-of-code/blob/main/2023/day20p2.py
			// when our target is "cs" (which will finally point to "rx") and we have a high pulse (true)
			// we increment the times we have seen this module
			if instruction.target == "cs" && instruction.pulse {
				visistedCache[instruction.source] += 1

				if _, ok := cycleLengths[instruction.source]; !ok {
					cycleLengths[instruction.source] = repetition
				}

				allTruthy := true

				// when all origins pointing to "cs" have been visisted at least once when having a
				// high pulse (instruction.pulse == true) ...
				for _, value := range visistedCache {
					if value <= 0 {
						allTruthy = false
						break
					}
				}

				// .. we want to find the lowest common divisor of the cycle lengths of each origin
				// giving us the total amount of repetitions necessary to reach our final condition
				// giving only "high" pulses to "cs" which allows giving a "low" pulse to "rx"
				if allTruthy {
					x := 1
					for _, l := range cycleLengths {
						x = x * l / GCD(x, l)
					}
					return x
				}
			}

			if _, exists := ffmodules[instruction.target]; exists {
				if !instruction.pulse {
					ffmodules[instruction.target] = FFModule{targets: ffmodules[instruction.target].targets, memory: !ffmodules[instruction.target].memory}
					for _, target := range ffmodules[instruction.target].targets {
						instructionQueue = append(instructionQueue, Inst{source: instruction.target, target: target, pulse: ffmodules[instruction.target].memory})
					}
				}
			} else if _, exists := cmodules[instruction.target]; exists {
				cmodules[instruction.target].memory[instruction.source] = instruction.pulse
				out := false

				for _, mem := range cmodules[instruction.target].memory {
					if !mem {
						out = true
						break
					}
				}

				for _, target := range cmodules[instruction.target].targets {
					instructionQueue = append(instructionQueue, Inst{source: instruction.target, target: target, pulse: out})
				}
			} else {
				continue
			}
		}

		repetition++
	}

	return repetition
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
