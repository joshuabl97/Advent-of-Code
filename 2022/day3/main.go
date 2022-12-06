package main

import (
	"aoc2022/reusable"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var filePath = flag.String("f", "input.txt", "path to the input file")

func main() {
	flag.Parse()
	bags := parseInput(*filePath)
	fmt.Printf("Problem 1: %v\n", problem1(bags))
	fmt.Printf("Problem 2: %v\n", problem2(bags))
}

// Find the item type that appears in both compartments of each rucksack.
// What is the sum of the priorities of those item types?
// a = 1, b = 2, c = 3, A = 27, B = 28, C = 29, etc
func problem1(bags [][]string) int {
	var (
		count int
		comp1 [][]string
		comp2 [][]string
	)

	for _, b := range bags {
		comp1 = append(comp1, b[:len(b)/2])
		comp2 = append(comp2, b[len(b)/2:])
	}

	comp1 = reusable.NestedStringUnique(comp1)
	comp2 = reusable.NestedStringUnique(comp2)

	for i := 0; i < len(comp1); i++ {
		for _, c1 := range comp1[i] {
			for _, c2 := range comp2[i] {
				if c1 == c2 {
					count += reusable.CharNum(c1)
				}
			}
		}
	}

	return count
}

// Find the item type that corresponds to the
// badges of each three-Elf group. What is the sum of the priorities of
// those item types?
// a = 1, b = 2, c = 3, A = 27, B = 28, C = 29, etc
func problem2(bags [][]string) int {
	var count int
	bags = reusable.NestedStringUnique(bags)

	for i := 0; i < len(bags); i += 3 {
		for _, b1 := range bags[i] {
			for _, b2 := range bags[i+1] {
				for _, b3 := range bags[i+2] {
					if b1 == b2 && b2 == b3 {
						count += reusable.CharNum(b1)
					}

				}
			}
		}
	}

	return count
}

// takes stock of the inventory of each elf bag in a nested slice of strings
func parseInput(filePath string) [][]string {
	f, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var bags [][]string

	for scanner.Scan() {
		bags = append(bags, strings.Split(scanner.Text(), ""))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return bags
}
