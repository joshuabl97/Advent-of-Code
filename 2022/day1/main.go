package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

var filePath = flag.String("f", "input.txt", "path to the input file")

/*
In case the Elves get hungry and need extra snacks,
they need to know which Elf to ask: they'd like to know how many
Calories are being carried by the Elf carrying the most Calories.
*/
func problem1(input []int) int {
	var mostCalories int
	var currentElf int
	for _, cals := range input {
		if cals == 0 {
			currentElf = 0
			continue
		}
		currentElf += cals
		if currentElf > mostCalories {
			mostCalories = currentElf
		}
	}

	return mostCalories
}

/*
The Elves would instead like to know the total Calories carried by
the top three Elves carrying the most Calories.
(The elf to english dictionary is missing a few key words)
*/
func problem2(input []int) int {
	// this is a slice of the total number of calories that each elf possesses
	var calorieCounts []int
	var currentElf int
	for _, cals := range input {
		if cals == 0 {
			calorieCounts = append(calorieCounts, currentElf)
			currentElf = 0
			continue
		}
		currentElf += cals
	}

	sort.Ints(calorieCounts)

	var totalCals int
	for i := 1; i <= 3; i++ {
		totalCals += calorieCounts[len(calorieCounts)-i]
	}

	return totalCals
}

func parseInput(filePath string) []int {
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	text := string(fileContent)
	input := strings.Split(text, "\n")
	input = append(input, "")

	var calorieSlice []int
	for _, in := range input {
		if in == "" {
			calorieSlice = append(calorieSlice, 0)
			continue
		}
		calInt, err := strconv.Atoi(in)
		if err != nil {
			log.Fatal(err)
		}
		calorieSlice = append(calorieSlice, calInt)
	}

	return calorieSlice
}

func main() {
	flag.Parse()
	input := parseInput(*filePath)
	fmt.Printf("Problem 1: %v \n", problem1(input))
	fmt.Printf("Problem 2: %v \n", problem2(input))
}
