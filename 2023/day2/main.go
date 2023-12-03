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
	//fmt.Printf("%v\n", input)
	cubeMaxCounts := countMaxCubes(input)
	fmt.Printf("Problem 1: %v\n", p1(cubeMaxCounts))
	fmt.Printf("Problem 2: %v\n", p2(cubeMaxCounts))
}

/*
To get information, once a bag has been loaded with cubes, the Elf will reach into the bag
, grab a handful of random cubes, show them to you, and then put them back in the bag.
He'll do this a few times per game.

In game 1, three sets of cubes are revealed from the bag (and then put back again).
The first set is 3 blue cubes and 4 red cubes; the second set is 1 red cube, 2 green
cubes, and 6 blue cubes; the third set is only 2 green cubes.

Determine which games would have been possible if the bag had been loaded with only
12 red cubes, 13 green cubes, and 14 blue cubes. What is the sum of the IDs of those games?
*/
func p1(cubeCounts map[int]cubes) int {
	var total int

	for k, v := range cubeCounts {
		//fmt.Printf("Game: %v - Red: %v Green: %v Blue: %v \n", k, v.red, v.green, v.blue)
		if v.red > 12 || v.green > 13 || v.blue > 14 {
			continue
		}
		total += k
	}

	return total
}

/*
s you continue your walk, the Elf poses a second question: in each game you played, what is the
fewest number of cubes of each color that could have been in the bag to make the game possible?

For each game, find the minimum set of cubes that must have been present. What is the sum of the
power of these sets?
*/
func p2(cubeCounts map[int]cubes) int {
	var total int

	for _, v := range cubeCounts {
		total += v.red * v.green * v.blue
	}

	return total
}

type cubes struct {
	red   int
	green int
	blue  int
}

// returns a map[game_number int]maxCubesPerColor type cubes
func countMaxCubes(input []string) map[int]cubes {
	gameCounts := make(map[int]cubes)

	for _, s := range input {
		colonStr := strings.Split(s, ":")
		gameNumStr, _ := strings.CutPrefix(colonStr[0], "Game ")
		gameNum, _ := strconv.Atoi(gameNumStr)
		//fmt.Printf("Game number: %v\n", gameNum)

		reveals := strings.Split(colonStr[1], ";")
		currentCubes := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, reveal := range reveals {
			//fmt.Printf("Reveal: %v\n", reveal)
			revealSlice := (strings.Split(reveal, " "))[1:]
			for i := 0; i < len(revealSlice); i += 2 {
				currentColor := strings.Split(revealSlice[i+1], ",")[0]
				//fmt.Printf("Color: %v\n", currentColor)
				cubeCount, _ := strconv.Atoi(revealSlice[i])
				//fmt.Printf("Quantity: %v\n", cubeCount)

				if currentCubes[currentColor] < cubeCount {
					currentCubes[currentColor] = cubeCount
				}
			}
		}
		//fmt.Printf("Max Cube Count Game: \n%v\n", currentCubes)
		gameCounts[gameNum] = cubes{
			red:   currentCubes["red"],
			green: currentCubes["green"],
			blue:  currentCubes["blue"],
		}
	}

	return gameCounts
}
