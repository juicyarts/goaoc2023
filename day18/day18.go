package day18

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

var primary = color.New(color.FgHiWhite).SprintFunc()
var secondary = color.New(color.FgHiBlack).SprintFunc()
var tertiary = color.New(color.FgHiRed).SprintFunc()

type DirCfg struct {
	d int
	a int
}

var dirCfg = map[string]DirCfg{
	"R": {
		d: 0,
		a: 1,
	},
	"L": {
		d: 0,
		a: -1,
	},
	"D": {
		d: 1,
		a: 1,
	},
	"U": {
		d: 1,
		a: -1,
	},
}

func Travel(input []string) (int, int) {
	grid := make(map[string]int, len(input))
	x, y := 0, 0
	loX, hiX, loY, hiY := 0, 0, 0, 0
	for i := 0; i < len(input); i++ {
		inst := strings.Split(input[i], " ")
		dir, a, _ := inst[0], inst[1], inst[2]
		amount, _ := strconv.Atoi(a)
		dirConfig := dirCfg[dir]

		for j := 0; j < amount; j++ {
			if dirConfig.d == 0 {
				if dirConfig.a > 0 {
					x++
				}
				if dirConfig.a < 0 {
					x--
				}

				if x < loX {
					loX = x
				}

				if x > hiX {
					hiX = x
				}
			}

			if dirConfig.d == 1 {
				if dirConfig.a < 0 {
					y--
				}
				if dirConfig.a > 0 {
					y++
				}

				if y < loY {
					loY = y
				}

				if y > hiY {
					hiY = y
				}
			}
			gridKey := fmt.Sprintf("%v|%v", y, x)
			grid[gridKey] += 1

		}
	}

	widthDiff, heightDiff := int(math.Abs(float64(loX))), int(math.Abs(float64(loY)))
	fieldWidth, fieldHeight := hiX+widthDiff, hiY+heightDiff
	printString := "\n"

	counter := 0

	rowMap := make(map[int][]int, 0)
	for i := 0; i <= fieldHeight; i++ {
		printString += "\n"
		countToggle := false

		for j := 0; j <= fieldWidth; j++ {

			gridItemKey := fmt.Sprintf("%v|%v", i+loY, j+loX)
			if grid[gridItemKey] > 0 {
				if rowMap[i] == nil {
					rowMap[i] = []int{j}
				} else {

					if j > 0 {
						previousGridItemKey := fmt.Sprintf("%v|%v", i+loY, (j-1)+loX)
						previousGridItem := grid[previousGridItemKey]
						if previousGridItem == 0 {
							rowMap[i] = append(rowMap[i], j)
						} else {

							for rowMapIndex := 0; rowMapIndex < len(rowMap[i]); rowMapIndex++ {
								if rowMapIndex > 0 {
									if rowMapIndex%2 != 0 {
										fmt.Print("WWWWWWWWWWWW---------------------------- \n")
										fmt.Printf("WWWWWWWWWWWW-----%v %v %v %v----------------------- \n", j, rowMap[i], j, rowMap[i][rowMapIndex])
										if j > rowMap[i][rowMapIndex-1] && j < rowMap[i][rowMapIndex] {
											printString += "\nWURST\n"
											break
										}
									}
								}
							}

						}
					}
				}

				countToggle = !countToggle
				counter++

				printString += primary("#")

			} else if countToggle {
				counter++
				printString += tertiary("X")
			} else {
				printString += secondary(".")
			}
		}
	}

	count := 0
	for rowIndex := 0; rowIndex < len(rowMap); rowIndex++ {
		sort.Ints(rowMap[rowIndex])
		for rowMapIndex := range rowMap[rowIndex] {
			if rowMapIndex > 0 {
				if rowMapIndex%2 != 0 {
					count += rowMap[rowIndex][rowMapIndex-1] + rowMap[rowIndex][rowMapIndex]
				}
			}
		}
	}

	fmt.Print(printString)

	fmt.Printf("\n COUNT: %v | %v \n", count, counter)
	fmt.Printf("\nloHi: %v|%v, %v|%v", loX, hiX, loY, hiY)
	fmt.Printf("\nAbs Width/Height: %v|%v\n", fieldWidth, fieldHeight)
	return len(grid), counter
}

func Draw(grid map[string]int) {

	printString := "\n----------------"
	printString += "\nGridWORX -------"

	// for

}
