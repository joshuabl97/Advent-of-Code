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
	group1, group2 := parseInput(*filePath)
	fmt.Printf("Problem 1: %v \n", problem1(group1, group2))
	fmt.Printf("Problem 2: %v \n", problem2(group1, group2))
}

/*
In pairs where one assignment fully contains the other, one Elf in the
pair would be exclusively cleaning sections their partner will already be cleaning,
so these seem like the most in need of reconsideration.
In how many assignment pairs does one range fully contain the other?
*/
func problem1(group1 [][]int, group2 [][]int) int {
	var count int

	for i := 0; i < len(group1); i++ {
		if group1[i][0] >= group2[i][0] && group1[i][1] <= group2[i][1] ||
			group1[i][0] <= group2[i][0] && group1[i][1] >= group2[i][1] {
			count++
		}
	}

	return count
}

/*
It seems like there is still quite a bit of duplicate work planned. Instead, the Elves
would like to know the number of pairs that overlap at all.
In how many assignment pairs do the ranges overlap?
*/
func problem2(group1 [][]int, group2 [][]int) int {
	var count int

	for i := 0; i < len(group1); i++ {
		if (group1[i][0] >= group2[i][0] && group1[i][0] <= group2[i][1]) ||
			(group1[i][1] <= group2[i][1] && group1[i][1] >= group2[i][0]) ||
			(group2[i][0] >= group1[i][0] && group2[i][0] <= group1[i][1]) ||
			(group2[i][1] <= group1[i][1] && group2[i][1] >= group1[i][0]) {
			count++
		}
	}

	return count
}

// takes input and converts to a list of list of ints
// i.e 2-4,6-8\n2-3,4-5 -> [[2,4][2,3]] [[6,8][4,5]]
func parseInput(filePath string) ([][]int, [][]int) {
	f, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var (
		group1 [][]int
		group2 [][]int
	)

	for scanner.Scan() {
		x := strings.Split(scanner.Text(), ",")
		for i, str := range x {
			intStr := strings.Split(str, "-")
			var ints []int
			for _, c := range intStr {
				x, _ := strconv.Atoi(c)
				ints = append(ints, x)
			}
			if i%2 == 0 || i == 0 {
				group1 = append(group1, ints)
			} else {
				group2 = append(group2, ints)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return group1, group2
}
