package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/joshuabl97/Advent-of-Code/2023/utils"
)

var filePath = flag.String("f", "input.txt", "path to the input file")

func main() {
	flag.Parse()
	input := utils.GetInputStrings(*filePath)

	var engineMap [][]string
	for _, row := range input {
		engineMap = append(engineMap, strings.Split(row, ""))
	}

	fmt.Printf("Problem 1: %v\n", p1(engineMap))
}

/*
The engine schematic (your puzzle input) consists of a visual representation of the engine.
There are lots of numbers and symbols you don't really understand, but apparently any number
adjacent to a symbol, even diagonally, is a "part number" and should be included in your sum.
(Periods (.) do not count as a symbol.)

What is the sum of all of the part numbers in the engine schematic?
*/
func p1(engineMap [][]string) int {
	var total int

	for i, row := range engineMap {
		var currentNum []string
		numIsSpecial := false

		for j, symbol := range row {
			if _, err := strconv.Atoi(symbol); err != nil {
				// we can add the current num and add it to the total only if it is special
				if len(currentNum) > 0 && numIsSpecial {
					numStr := strings.Join(currentNum, "")
					num, _ := strconv.Atoi(numStr)
					fmt.Printf("Special number: %v\n", num)
					total += num
				}

				numIsSpecial = false
				currentNum = []string{}
			} else {
				// this character is a number
				currentNum = append(currentNum, symbol)
				fmt.Printf("Row: %v Current char: %v\n", row, symbol)
				// check top left corner
				if i != 0 && j != 0 {
					topleftcorner := engineMap[i-1][j-1]
					fmt.Printf("top left: %v\n", topleftcorner)
					if topleftcorner != "." {
						_, err := strconv.Atoi(topleftcorner)
						if err != nil {
							numIsSpecial = true
						}
					}
				}
				// check top right corner
				if i != 0 && j != len(row)-1 {
					toprightcorner := engineMap[i-1][j+1]
					fmt.Printf("top right: %v\n", toprightcorner)
					if toprightcorner != "." {
						_, err := strconv.Atoi(toprightcorner)
						if err != nil {
							numIsSpecial = true
						}
					}
				}

				// check bottom left corner
				if i != len(engineMap)-1 && j != 0 {
					bottomleftcorner := engineMap[i+1][j-1]
					fmt.Printf("bottom left: %v\n", bottomleftcorner)
					if bottomleftcorner != "." {
						_, err := strconv.Atoi(bottomleftcorner)
						if err != nil {
							numIsSpecial = true
						}
					}
				}

				// check bottom right corner
				if i != len(engineMap)-1 && j != len(row)-1 {
					bottomrightcorner := engineMap[i+1][j+1]
					fmt.Printf("bottom right: %v\n", bottomrightcorner)
					if bottomrightcorner != "." {
						_, err := strconv.Atoi(bottomrightcorner)
						if err != nil {
							numIsSpecial = true
						}
					}
				}

				// check left
				if j != 0 {
					left := engineMap[i][j-1]
					fmt.Printf("left: %v\n", left)
					if left != "." {
						_, err := strconv.Atoi(left)
						if err != nil {
							numIsSpecial = true
						}
					}
					// check above
				}
				if i != 0 {
					top := engineMap[i-1][j]
					fmt.Printf("top: %v\n", top)
					if top != "." {
						_, err := strconv.Atoi(top)
						if err != nil {
							numIsSpecial = true
						}
					}
				}

				// check right
				if j != len(row)-1 {
					right := engineMap[i][j+1]
					fmt.Printf("right: %v\n", right)
					if right != "." {
						_, err := strconv.Atoi(right)
						if err != nil {
							numIsSpecial = true
						}
					}
				}

				// check below
				if i != len(engineMap)-1 {
					below := engineMap[i+1][j]
					fmt.Printf("below: %v\n", below)
					_, err := strconv.Atoi(below)
					if below != "." {
						if err != nil {
							numIsSpecial = true
						}
					}
				}

			}
		}
	}

	return total
}
