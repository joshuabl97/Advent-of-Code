package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var filePath = flag.String("f", "input.txt", "path to the input file")

func main() {
	flag.Parse()
	instructions := inputParse(*filePath)
	problem1(instructions)
}

func problem1(instructions [][]string) {
	var (
		ropeBridge [][]string
		lastMove   string
	)

	ropeBridge = append(ropeBridge, []string{"H"})

	for _, instruction := range instructions {
		fmt.Printf("INSTRUCTION %v\n", instruction)
		steps, err := strconv.Atoi(instruction[1])

		if err != nil {
			log.Fatal(err)
		}

		direction := instruction[0]

		var (
			currentHead []int
		)

		for i, ropeBridgeRow := range ropeBridge {
			for j, ropeBridgeTile := range ropeBridgeRow {
				if ropeBridgeTile == "H" {
					currentHead = append(currentHead, i, j)
				}
			}
		}

		for i := 0; i < steps; i++ {
			switch direction {
			case "R":
				if ropeBridge[currentHead[0]][currentHead[1]] == ropeBridge[currentHead[0]][len(ropeBridge[currentHead[0]])-1] {
					for i, rbRow := range ropeBridge {
						ropeBridge[i] = append([]string{"."}, rbRow...)
					}
				}
				ropeBridge[currentHead[0]][currentHead[1]] = "."
				currentHead[1] = currentHead[1] + 1
				ropeBridge[currentHead[0]][currentHead[1]] = "H"
				lastMove = "R"

			case "U":
				if currentHead[0] == 0 {
					var newRow []string
					for i := 0; i < len(ropeBridge[0]); i++ {
						newRow = append(newRow, ".")
					}
					ropeBridge = append(ropeBridge, newRow)
				} else {
					ropeBridge[currentHead[0]][currentHead[1]] = "."
					currentHead[0] = currentHead[0] - 1
					ropeBridge[currentHead[0]][currentHead[1]] = "H"
				}
				lastMove = "U"

			case "L":
				if currentHead[1] == 0 {
					for i, rbRow := range ropeBridge {
						ropeBridge[i] = append(rbRow, ".")
					}
				} else {
					ropeBridge[currentHead[0]][currentHead[1]] = "."
					currentHead[1] = currentHead[1] - 1
					ropeBridge[currentHead[0]][currentHead[1]] = "H"
				}
				lastMove = "L"

			case "D":
				if currentHead[0] == len(ropeBridge)-1 {
					var newRow []string
					for i := 0; i < len(ropeBridge[0]); i++ {
						newRow = append(newRow, ".")
					}
					ropeBridge = append([][]string{newRow}, ropeBridge...)
				}
				ropeBridge[currentHead[0]][currentHead[1]] = "."
				currentHead[0] = currentHead[0] + 1
				ropeBridge[currentHead[0]][currentHead[1]] = "H"
				lastMove = "D"

			}
		}
		fmt.Printf("current head position %v last move %v\n", currentHead, lastMove)
		for _, ropeBridgeRow := range ropeBridge {
			fmt.Printf("%v\n", ropeBridgeRow)
		}
	}
}

func inputParse(filepath string) [][]string {
	f, err := os.Open(*filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var inputLines [][]string

	for scanner.Scan() {
		inputLines = append(inputLines, strings.Split(scanner.Text(), " "))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return inputLines
}
