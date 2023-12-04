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
			if symbol == "." {
				// we can add the current num and add it to the total only if it is special
				if len(currentNum) > 0 {
					if numIsSpecial {
						numStr := ""
						for _, numChar := range currentNum {
							numStr += numChar
						}
						num, _ := strconv.Atoi(numStr)
						total += num
					}

					numIsSpecial = false
					currentNum = []string{}

				}
			} else if _, err := strconv.Atoi(symbol); err == nil {
				// this character is a number
				currentNum = append(currentNum, symbol)

				// check top left corner
				if i != 0 && j != 0 {
					topleftcorner := engineMap[i-1][i-1]
					if topleftcorner != "." {
						_, err := strconv.Atoi(topleftcorner)
						if err != nil {
							// we found a symbol attached to this number and can skip to constructing the rest of the number
							numIsSpecial = true
						}
					}
					// check top right corner
				} else if i != 0 && j != len(row) {
					toprightcorner := engineMap[i-1][j+1]
					if toprightcorner != "." {
						_, err := strconv.Atoi(toprightcorner)
						if err != nil {
							numIsSpecial = true
						}
					}
					// check bottom left corner
				} else if i != len(engineMap) && j != 0 {
					bottomleftcorner := engineMap[i+1][j-1]
					if bottomleftcorner != "." {
						_, err := strconv.Atoi(bottomleftcorner)
						if err != nil {
							numIsSpecial = true
						}
					}
					// check bottom right corner
				} else if i != len(engineMap) && j != len(row) {
					bottomrightcorner := engineMap[i+1][j+1]
					if bottomrightcorner != "." {
						_, err := strconv.Atoi(bottomrightcorner)
						if err != nil {
							numIsSpecial = true
						}
					}
					// check left
				} else if j != 0 {
					left := engineMap[i][j-1]
					if left != "." {
						_, err := strconv.Atoi(left)
						if err != nil {
							numIsSpecial = true
						}
					}
					// check above
				} else if i != 0 {
					top := engineMap[i-1][j]
					if top != "." {
						_, err := strconv.Atoi(top)
						if err != nil {
							numIsSpecial = true
						}
					}
					// check right
				} else if j != len(row) {
					right := engineMap[i][j+1]
					_, err := strconv.Atoi(right)
					if err != nil {
						numIsSpecial = true
					}
					// check below
				} else if i != len(engineMap) {
					below := engineMap[i-1][j]
					_, err := strconv.Atoi(below)
					if err != nil {
						numIsSpecial = true
					}
				}
			}
		}
	}

	return total
}
