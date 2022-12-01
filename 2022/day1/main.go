package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var filePath = flag.String("f", "input.txt", "path to the input path")

/*
In case the Elves get hungry and need extra snacks,
they need to know which Elf to ask: they'd like to know how many
Calories are being carried by the Elf carrying the most Calories.
*/
func problem1(input []string) int {
	var mostCalories int
	var currentElf int
	for _, cals := range input {
		if cals == "" {
			currentElf = 0
			continue
		}
		calInt, err := strconv.Atoi(cals)
		if err != nil {
			log.Fatal(err)
		}
		currentElf += calInt
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
func problem2(input []string) int {
	// this is a slice of the total number of calories that each elf possesses
	var calorieCounts []int
	var currentElf int
	for _, cals := range input {
		if cals == "" {
			calorieCounts = append(calorieCounts, currentElf)
			currentElf = 0
			continue
		}
		calInt, err := strconv.Atoi(cals)
		if err != nil {
			log.Fatal(err)
		}
		currentElf += calInt
	}

	var topThreeCalories []int
	for j := 0; j <= 2; j++ {
		maxNum := calorieCounts[0]
		var sliceIndex int
		for i := 0; i < len(calorieCounts); i++ {
			if calorieCounts[i] > maxNum {
				maxNum = calorieCounts[i]
				sliceIndex = i
			}
		}
		topThreeCalories = append(topThreeCalories, maxNum)
		calorieCounts = append(calorieCounts[:sliceIndex], calorieCounts[sliceIndex+1:]...)
	}

	var totalCals int
	for _, cals := range topThreeCalories {
		totalCals += cals
	}

	return totalCals
}

func parseInput(filePath string) []string {
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	text := string(fileContent)
	input := strings.Split(text, "\n")
	input = append(input, "")
	return input
}

func main() {
	flag.Parse()
	input := parseInput(*filePath)
	fmt.Printf("Problem 1: %v \n", problem1(input))
	fmt.Printf("Problem 2: %v \n", problem2(input))
}
