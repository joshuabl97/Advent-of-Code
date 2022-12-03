package main

import (
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

// inputs a slice of strings and returns a string slice containing only unique characters
// i.e ["a","b","b","b","c","c"] returns ["a","b","c"]
func unique(s []string) []string {
	inResult := make(map[string]bool)
	var result []string
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}

// returns value in alphabet if given a single digit string
// a = 1, b = 2, c = 3, A = 27, B = 28, C = 29, etc
func charNum(s string) int {
	char := rune(s[0])
	if char >= 97 && char <= 122 {
		return int(char) - 96
	} else if char >= 65 && char <= 90 {
		return int(char) - 38
	}
	return 0
}

// iterates through a slice of a slice of strings and only returns only unique values on the inner slice
// i.e [['a','a','b','c','c'], ['a','b','b','c']] returns [['a','b','c'],['a','b','c']]
func nestedStringUnique(nestedS [][]string) [][]string {
	for i, s := range nestedS {
		nestedS[i] = unique(s)
	}

	return nestedS
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

	comp1 = nestedStringUnique(comp1)
	comp2 = nestedStringUnique(comp2)

	for i := 0; i < len(comp1); i++ {
		for _, c1 := range comp1[i] {
			for _, c2 := range comp2[i] {
				if c1 == c2 {
					count += charNum(c1)
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
	bags = nestedStringUnique(bags)

	for i := 0; i < len(bags); i += 3 {
		for _, b1 := range bags[i] {
			for _, b2 := range bags[i+1] {
				for _, b3 := range bags[i+2] {
					if b1 == b2 && b2 == b3 {
						count += charNum(b1)
					}

				}
			}
		}
	}

	return count
}

// splits the elves bags (input) and contents into two compartments (comp1 & comp2)
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
