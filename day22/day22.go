package day22

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type Position struct {
	x int
	y int
	z int
}
type Brick struct {
	name             int
	isGrounded       bool
	supports         int
	supportingBricks []Brick
	start            Position
	end              Position
}

var primary = color.New(color.FgRed).SprintFunc()
var secondary = color.New(color.FgYellow).SprintFunc()
var high = color.New(color.FgCyan).SprintFunc()
var empty = color.New(color.FgBlue).SprintFunc()
var ground = color.New(color.FgGreen).SprintFunc()

func CollectBricks(input []string) int {
	bricks := map[string]Brick{}

	maxX, maxY, maxZ := 0, 0, 0

	for rowIndex, row := range input {
		edges := strings.Split(row, "~")
		start, end := strings.Split(edges[0], ","), strings.Split(edges[1], ",")
		startAsInt := []int{}
		endAsInt := []int{}
		for _, s := range start {
			n, _ := strconv.Atoi(s)
			startAsInt = append(startAsInt, n)
		}

		for _, e := range end {
			n, _ := strconv.Atoi(e)
			endAsInt = append(endAsInt, n)
		}

		bricks[fmt.Sprintf("%+v", rowIndex)] = Brick{
			name: rowIndex,
			start: Position{
				x: startAsInt[0],
				y: startAsInt[1],
				z: startAsInt[2],
			},
			end: Position{
				x: endAsInt[0],
				y: endAsInt[1],
				z: endAsInt[2],
			},
		}

		// only for drawing purposes
		if maxX < endAsInt[0] {
			maxX = endAsInt[0]
		}

		if maxY < endAsInt[1] {
			maxY = endAsInt[1]
		}

		if maxZ < endAsInt[2] {
			maxZ = endAsInt[2]
		}
	}

	fmt.Print("BEFORE-----------------\n")
	printMap(bricks, maxX, maxY, maxZ, "zx", 0)

	bricks = Fall(bricks, maxX, maxY, maxZ)
	sum := 0

	fmt.Print("After-----------------\n")
	printMap(bricks, maxX, maxY, maxZ, "zx", 0)

	for _, brick := range bricks {
		supports := 0

		// check for the amount of supports for each supporting brick of the current brick
		// a brick can be supported by only one other brick, so in case one supporting
		// brick has more than one support we can safely assume that we can remove the current one
		for _, supported := range brick.supportingBricks {
			if supported.supports <= 1 {
				supports++
			}
		}

		if supports == 0 {
			sum++
		}
	}

	return sum
}

func Fall(bricks map[string]Brick, maxX, maxY, maxZ int) map[string]Brick {
	// define amount of moves/cycles to begin with
	// each cycle moves all movable bricks one z level down until
	// it cannot move any further. In case a brick has moved
	// increase the amount of moves/cycles to determine the final state
	// of ig other bricks can move now
	moves := 1

	for i := 0; i < moves; i++ {
		// instead of adding a cycle with each brick move
		// we only want to add one tick after all brick changes of an iteration
		movesThisTick := 0

		for brickId, brick := range bricks {

			// for each brick determine if it can move, is supporting any other bricks
			// or is grounded
			canMove := true
			isGrounded := brick.isGrounded
			// we save the amount of supporters for this brick to
			// be able to determine if we can safely remove one when a brick has more than one
			// supporter
			brickSupports := 0
			// we save the bricks we support to later estimate if we can safely remove a brick
			// depending the condition explained above
			supportingBricks := []Brick{}

			// in case z is 1 we know that the brick is on ground level
			if brick.start.z == 1 || brick.end.z == 1 {
				isGrounded = true
			}

			// compare each brick with the other bricks
			for _, otherBrick := range bricks {
				// continue in case we compare to ourselves
				if otherBrick.name == brick.name {
					continue
				}
				// compare with bricks on top of the current brick
				if otherBrick.start.z == brick.start.z+1 || otherBrick.end.z == brick.end.z+1 {
					// if they collide we can assume the current brick is a support brick
					// and add the colliding brick to the current bricks supporting list
					if rangesDoCollide(
						[]int{brick.start.x, brick.start.y},
						[]int{brick.end.x, brick.end.y},
						[]int{otherBrick.start.x, otherBrick.start.y},
						[]int{otherBrick.end.x, otherBrick.end.y},
					) {
						// fmt.Printf("Supporting other Brick \n")
						// fmt.Printf("Source %+v\n", brick)
						// fmt.Printf("Target %+v\n", otherBrick)
						// fmt.Printf("\n")
						supportingBricks = append(supportingBricks, otherBrick)
					}
				}

				// compare with bricks below the current brick
				if otherBrick.start.z == brick.start.z-1 || otherBrick.end.z == brick.end.z-1 {
					// if they collide we can assume the current brick can not move further down
					// we can also memoize that the current brick has supports for later
					// comparisons
					if rangesDoCollide(
						[]int{brick.start.x, brick.start.y},
						[]int{brick.end.x, brick.end.y},
						[]int{otherBrick.start.x, otherBrick.start.y},
						[]int{otherBrick.end.x, otherBrick.end.y},
					) {
						// fmt.Printf("Brick Cannot move because of Collision \n")
						// fmt.Printf("Source %+v\n", brick)
						// fmt.Printf("Target %+v\n", otherBrick)
						// fmt.Printf("\n")
						brickSupports++
						canMove = false
					}
				}
			}

			// in case the brick can move and is not on ground level
			// we decrease the z index of both start and end
			// and add another cycle
			if canMove && !isGrounded {
				// fmt.Printf("This Brick can move downwards this tick! %+v \n", brick)
				bricks[brickId] = Brick{
					name:             brick.name,
					isGrounded:       isGrounded,
					supports:         brickSupports,
					supportingBricks: supportingBricks,
					start: Position{
						x: brick.start.x,
						y: brick.start.y,
						z: brick.start.z - 1,
					},
					end: Position{
						x: brick.end.x,
						y: brick.end.y,
						z: brick.end.z - 1,
					},
				}

				movesThisTick++
			} else {
				// if we cannot move we just update the state of our current brick without moving
				bricks[brickId] = Brick{
					name:             brick.name,
					isGrounded:       isGrounded,
					start:            brick.start,
					end:              brick.end,
					supports:         brickSupports,
					supportingBricks: supportingBricks,
				}
			}
		}

		if movesThisTick > 0 {
			moves += 2 // add another final move to ensure supports are set properly
		}

		// fmt.Printf("AFTER Move %+v-----------------\n", i)
		// printMap(bricks, maxX, maxY, maxZ, "zx", 0)
		// printMap(bricks, maxX, maxY, maxZ, "zy", 0)
		// printMap(bricks, maxX, maxY, maxZ, "xy", 2)
	}

	return bricks
}

func rangesDoCollide(ss []int, se []int, ds []int, de []int) bool {
	return ss[0] <= de[0] && se[0] >= ds[0] && ss[1] <= de[1] && se[1] >= ds[1]
}

// fancy printing is fancy
func printMap(bricks map[string]Brick, x, y, z int, view string, currentZ int) {
	fmt.Printf("--------------- \n")
	fmt.Printf("Drawing: %+v,%+v,%+v,%+v at level: %+v \n", x, y, z, view, currentZ)
	fmt.Printf("--------------- \n")

	if view == "xy" {
		for i := x + 1; i > -1; i-- {
			fmt.Print("\n")
			for j := 0; j <= y+1; j++ {
				foundBrick := 0
				lastFoundName := ""
				foundBrickZ := 0

				if i == x+1 {
					if j != y+1 {
						fmt.Printf("   %+v   ", j)
					}
					continue
				} else if j == y+1 {
					fmt.Printf("  %+v  ", i)
					continue
				}

				for _, brick := range bricks {
					if currentZ >= brick.start.z && currentZ <= brick.end.z {

						isInX := brick.start.x <= i && brick.end.x >= i
						isinY := brick.start.y <= j && brick.end.y >= j

						if isInX && isinY {
							foundBrick++
							lastFoundName = strconv.Itoa(brick.name)
							foundBrickZ = brick.start.y
						}
					} else {
						continue
					}
				}

				if foundBrick <= 0 {
					fmt.Print(empty("......."))
				} else {
					if foundBrick > 1 {
						fmt.Print(secondary("[  ?  ]"))
					} else {
						if foundBrickZ != currentZ {
							fmt.Print(primary(fmt.Sprintf("[%03v|%v]", lastFoundName, foundBrickZ)))
						} else {
							fmt.Print(high(fmt.Sprintf("[%03v|%v]", lastFoundName, currentZ)))
						}
					}
				}
			}
			if i == 0 {
				fmt.Print("\n")
			}

			if i == x+1 {
				fmt.Print("\n")
			}
		}
	}
	if view == "zx" {
		for i := z + 1; i > -1; i-- {
			fmt.Print("\n")
			for j := 0; j <= x+1; j++ {
				foundBrick := 0
				foundBrickZ := 0
				lastFoundName := ""

				if i == z+1 {
					if j != x+1 {
						fmt.Printf("  %+v ", j)
					}
					continue
				} else if j == x+1 {
					fmt.Printf("  %+v ", i)
					continue
				}

				for _, brick := range bricks {
					isInZ := brick.start.z <= i && brick.end.z >= i
					isInX := brick.start.x <= j && brick.end.x >= j

					if isInX && isInZ {
						foundBrick++
						lastFoundName = strconv.Itoa(brick.name)
						foundBrickZ = brick.start.y
					}
				}

				if foundBrick <= 0 {
					if i == 0 {
						fmt.Print(ground("[------]"))
					} else {
						fmt.Print(empty("........"))
					}
				} else {
					if foundBrick > 1 {
						fmt.Print(secondary("[   ?  ]"))
					} else {
						if foundBrickZ != currentZ {
							fmt.Print(primary(fmt.Sprintf("[%04v|%v]", lastFoundName, foundBrickZ)))
						} else {
							fmt.Print(high(fmt.Sprintf("[%04v|%v]", lastFoundName, currentZ)))
						}
					}
				}
			}

			if i == 0 {
				fmt.Print("\n")
			}

			if i == z+1 {
				fmt.Print("\n")
			}
		}
	}
	if view == "zy" {
		for i := z + 1; i > -1; i-- {
			fmt.Print("\n")
			for j := 0; j <= y+1; j++ {
				foundBrick := 0
				lastFoundName := ""
				foundBrickZ := 0

				if i == z+1 {
					if j != y+1 {
						fmt.Printf("   %+v   ", j)
					}
					continue
				} else if j == y+1 {
					fmt.Printf("  %+v  ", i)
					continue
				}

				for _, brick := range bricks {
					isInZ := brick.start.z <= i && brick.end.z >= i
					isInY := brick.start.y <= j && brick.end.y >= j

					if isInY && isInZ {
						foundBrick++
						lastFoundName = strconv.Itoa(brick.name)
						foundBrickZ = brick.start.x
					}
				}

				if foundBrick <= 0 {
					if i == 0 {
						fmt.Print(ground("[-----]"))
					} else {
						fmt.Print(empty("......."))
					}
				} else {
					if foundBrick > 1 {
						fmt.Print(secondary("[  ?  ]"))
					} else {
						if foundBrickZ != currentZ {
							fmt.Print(primary(fmt.Sprintf("[%03v|%v]", lastFoundName, foundBrickZ)))
						} else {
							fmt.Print(high(fmt.Sprintf("[%03v|%v]", lastFoundName, currentZ)))
						}
					}
				}
			}

			if i == 0 {
				fmt.Print("\n")
			}

			if i == z+1 {
				fmt.Print("\n")
			}
		}
	}

	fmt.Printf("\n--------------- \n")
}
