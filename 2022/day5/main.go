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
	stacks, instructions := parseInput(*filePath)
	fmt.Printf("Problem 1: %v \n", problem1(stacks, instructions))
	stacks, instructions = parseInput(*filePath)
	fmt.Printf("Problem 2: %v \n", problem2(stacks, instructions))
}

/*
The Elves have a drawing of the starting stacks of crates and the rearrangement procedure
The Elves just need to know which crate will end up on top of each stack
You can only move one crate at a time
*/
func problem1(stacks map[int]string, instructions [][]int) string {
	for _, instruction := range instructions {
		var reversedStr string

		for _, v := range string(stacks[instruction[1]][:instruction[0]]) {
			reversedStr = string(v) + reversedStr
		}

		stacks[instruction[2]] = reversedStr + stacks[instruction[2]]
		stacks[instruction[1]] = strings.TrimPrefix(stacks[instruction[1]], string(stacks[instruction[1]][:instruction[0]]))
	}

	var topStacks string

	for i := 1; i <= len(stacks); i++ {
		topStacks = topStacks + string(stacks[i][0])
	}

	return topStacks
}

//Same as problem 1, but you can move multiple crates at once
func problem2(stacks map[int]string, instructions [][]int) string {
	for _, instruction := range instructions {
		stacks[instruction[2]] = string(stacks[instruction[1]][:instruction[0]]) + stacks[instruction[2]]
		stacks[instruction[1]] = strings.TrimPrefix(stacks[instruction[1]], string(stacks[instruction[1]][:instruction[0]]))
	}

	var topStacks string

	for i := 1; i <= len(stacks); i++ {
		topStacks = topStacks + string(stacks[i][0])
	}

	return topStacks
}

// Each slice in the instructions are parsed like 'move 1 from 2 to 1' -> [1,2,1]
// Each string in the map goes from top to bottom of the stack
func parseInput(filePath string) (map[int]string, [][]int) {
	f, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	stacks := make(map[int]string)
	var instructions [][]int

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "[") {
			strSlice := strings.Split(scanner.Text(), "")
			m := make(map[int]string)

			for i, s := range strSlice {
				if s == "]" {
					m[i] = strSlice[i-1]
				}
			}

			for k, v := range m {
				k = k - 2
				if k != 0 {
					_, isPresent := stacks[(k/4)+1]
					if isPresent {
						stacks[(k/4)+1] = stacks[(k/4)+1] + v
					} else {
						stacks[(k/4)+1] = v
					}
				} else {
					_, isPresent := stacks[1]
					if isPresent {
						stacks[1] = (stacks[1] + v)
					} else {
						stacks[1] = v
					}
				}
			}
		} else if strings.Contains(scanner.Text(), "move") {
			var instruction []int

			for _, str := range strings.Split(scanner.Text(), "") {
				if i, err := strconv.Atoi(str); err == nil {
					instruction = append(instruction, i)
				}
			}

			if len(instruction) > 3 {
				var num string

				for i := 0; i < len(instruction)-2; i++ {
					num = num + strconv.Itoa(instruction[i])
				}

				i, _ := strconv.Atoi(num)
				instruction = []int{i, instruction[len(instruction)-2], instruction[len(instruction)-1]}
			}

			instructions = append(instructions, instruction)
		}
	}

	return stacks, instructions
}
