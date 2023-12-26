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

func enqueue(queue []Inst, element Inst) []Inst {
	queue = append(queue, element)
	return queue
}

// func dequeue(queue []Inst) []Inst {
// 	return queue[:1]
// }

func MeasurePulses(input []string) int {
	ffmodules := map[string]FFModule{}
	cmodules := map[string]CModule{}

	broadcastInstructions := []Inst{}

	for _, l := range input {
		s, t := strings.Split(l, " -> ")[0], strings.Split(l, " -> ")[1]
		mt, mk := s[:1], s[1:]
		targets := strings.Split(t, ", ")

		if s == "broadcaster" {
			for _, target := range targets {
				broadcastInstructions = enqueue(broadcastInstructions, Inst{source: s, target: target, pulse: false})
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

	for n, mod := range ffmodules {
		for _, target := range mod.targets {
			if _, exists := cmodules[target]; exists {
				cmodules[target].memory[n] = false
			}
		}
	}

	for n, mod := range cmodules {
		for _, target := range mod.targets {
			if _, exists := cmodules[target]; exists {
				cmodules[target].memory[n] = false
			}
		}
	}

	var hi, lo int = 0, 0
	repetition := 1000

	// fmt.Print("FF: ", ffmodules, "\n")
	// fmt.Print("CM: ", cmodules, "\n")

	for i := 0; i <= repetition; i++ {
		lo++
		instructionQueue := broadcastInstructions

		for j := 0; j < len(instructionQueue); j++ {
			instruction := instructionQueue[j]

			if instruction.pulse {
				hi++
			} else {
				lo++
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
	}

	return hi * lo
}
