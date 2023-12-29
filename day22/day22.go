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
	name              int
	isGrounded        bool
	supportingBricks  []Brick
	bricksSupportedBy []Brick
	start             Position
	end               Position
}

var primary = color.New(color.FgRed).SprintFunc()
var secondary = color.New(color.FgYellow).SprintFunc()
var high = color.New(color.FgCyan).SprintFunc()
var empty = color.New(color.FgBlue).SprintFunc()
var ground = color.New(color.FgGreen).SprintFunc()

func CollectBricks(input []string) (int, int) {
	initialBricks := map[string]Brick{}

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

		initialBricks[fmt.Sprintf("%+v", rowIndex)] = Brick{
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

	fmt.Print("BEFORE\n")
	printMap(initialBricks, maxX, maxY, maxZ, "zx", 0)

	bricks, _ := Fall(initialBricks, maxX, maxY, maxZ)

	sumOfRemovableBricks := 0
	sumOfFallingBricks := 0

	for brickId, brick := range bricks {
		// count amounts of bricks supported by this brick
		supports := 0
		// unsure if needed, cache length and count should be same
		fallingBricks := 0

		// Part 1
		// check the amount of supports for each supporting brick of the current brick
		// a brick can be supported by only one brick, so every supported brick that has more than one supporter
		// can be skipped
		for _, supported := range brick.supportingBricks {
			if len(supported.bricksSupportedBy) == 1 {
				supports++
			}
		}
		if supports == 0 {
			sumOfRemovableBricks++
		}

		bricksWithoutMe := map[string]Brick{}
		for k, b := range bricks {
			if k != brickId {
				bricksWithoutMe[k] = b
			}
		}

		// lazy & sloow solution, just remove the brick and simulate falling
		_, movedBricks := Fall(bricksWithoutMe, maxX, maxY, maxZ)
		fallingBricks += len(movedBricks)

		fmt.Printf("When %+v is removed, %+v other Bricks will fall\n", brick.name, fallingBricks)
		sumOfFallingBricks += fallingBricks
	}

	fmt.Printf("\n%+v Bricks can be savely removed, %+v bricks have moved \n", sumOfRemovableBricks, sumOfFallingBricks)
	printMap(bricks, maxX, maxY, maxZ, "zx", 0)
	printMap(bricks, maxX, maxY, maxZ, "zy", 0)

	return sumOfRemovableBricks, sumOfFallingBricks
}

// Alternative solution ideas
// just make it faster by using goroutines
// instead of using Fall create a method that memorizes how many bricks fall due to a single one
// and allow looking up

func Fall(bricks map[string]Brick, maxX, maxY, maxZ int) (map[string]Brick, map[string]int) {
	// define amount of moves/cycles to begin with
	// each cycle moves all movable bricks one z level down until
	// it cannot move any further. In case a brick has moved
	// increase the amount of moves/cycles to determine the final state
	moves := 1
	movedBricks := map[string]int{}

	for i := 0; i < moves; i++ {
		// instead of adding a cycle with each brick move
		// we only want to add one tick after all brick changes of an iteration
		movesThisTick := 0

		for brickId, brick := range bricks {
			// for each brick determine if it can move, is supporting any other bricks
			// or is grounded
			canMove := true
			isGrounded := brick.isGrounded
			jump := 1

			// we save the bricks we support to later estimate if we can safely remove a brick
			// depending the condition explained above
			supportingBricks := []Brick{}
			// we also save the bricks we are supported by for later checks
			bricksSupportedBy := []Brick{}

			// in case z is 1 we know that the brick is on ground level
			if brick.start.z == 1 || brick.end.z == 1 {
				isGrounded = true
				jump = 0
			}

			// compare each brick with the other bricks
			for _, otherBrick := range bricks {
				// continue in case we compare to ourselves
				if otherBrick.name == brick.name || (otherBrick.start.z > brick.start.z+1 && otherBrick.end.z > brick.end.z+1) {
					continue
				}

				if rangesDoCollide(
					[]int{brick.start.x, brick.start.y},
					[]int{brick.end.x, brick.end.y},
					[]int{otherBrick.start.x, otherBrick.start.y},
					[]int{otherBrick.end.x, otherBrick.end.y},
				) {
					// compare with bricks on top of the current brick
					// if bricks collide we can assume the current brick is a support brick
					// and add the colliding brick to the current bricks supporting list
					if otherBrick.start.z == brick.start.z+1 || otherBrick.end.z == brick.end.z+1 {
						supportingBricks = append(supportingBricks, otherBrick)
					} else if otherBrick.start.z == brick.start.z-1 || otherBrick.end.z == brick.end.z-1 {
						// compare with bricks below the current brick
						// if they collide we can assume the current brick can not move further down
						// we can also memoize that the current brick has supports for later
						// comparisons
						bricksSupportedBy = append(bricksSupportedBy, otherBrick)
						jump = 0
						canMove = false
					}

					// find distance to next colliding brick below and jump
					// won't work :/
					// if otherBrick.start.z < brick.start.z {
					// 	// find collider with highest z-index and take its z-index+1
					// 	condition := (brick.start.z - (otherBrick.start.z + 1))
					// 	if jump > condition {
					// 		if condition == 0 {
					// 			jump = 1
					// 		} else {
					// 			jump = condition
					// 		}
					// 	}
					// }
				}
			}

			// in case the brick can move and is not on ground level
			// we decrease the z index of both start and end
			// and add another cycle
			if canMove && !isGrounded && jump > 0 {
				// fmt.Printf("Brick %+v will move %+v this tick! \n", brick.name, jump)
				bricks[brickId] = Brick{
					name:              brick.name,
					isGrounded:        isGrounded,
					supportingBricks:  supportingBricks,
					bricksSupportedBy: bricksSupportedBy,
					start: Position{
						x: brick.start.x,
						y: brick.start.y,
						z: brick.start.z - jump,
					},
					end: Position{
						x: brick.end.x,
						y: brick.end.y,
						z: brick.end.z - jump,
					},
				}

				movedBricks[brickId]++
				movesThisTick++
			} else {
				// if we cannot move we just update the state of our current brick without moving
				bricks[brickId] = Brick{
					name:              brick.name,
					isGrounded:        isGrounded,
					start:             brick.start,
					end:               brick.end,
					supportingBricks:  supportingBricks,
					bricksSupportedBy: bricksSupportedBy,
				}
			}
		}

		// wonky as hell
		if movesThisTick > 0 {
			moves += 2 // add another final move to ensure supports are set properly
		}
	}

	return bricks, movedBricks
}

func rangesDoCollide(ss []int, se []int, ds []int, de []int) bool {
	return ss[0] <= de[0] && se[0] >= ds[0] && ss[1] <= de[1] && se[1] >= ds[1]
}

// fancy printing is fancy spaghetti
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
						fmt.Printf("| %04v |", j)
					}
					continue
				} else if j == x+1 {
					fmt.Printf("   %+v  ", i)
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
						fmt.Printf("| %04v |", j)
					}
					continue
				} else if j == y+1 {
					fmt.Printf("   %+v  ", i)
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

		}
	}

	fmt.Printf("\n--------------- \n")
}
