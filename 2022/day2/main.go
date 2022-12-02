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
	plays := parseInput(*filePath)
	fmt.Printf("Problem 1: %v\n", problem1(plays))
	fmt.Printf("Problem 2: %v\n", problem2(plays))
}

// plays a game of rock, paper, scissors and returns the score
// the opponent is stored at index 0 and the player at index 1
// 0 if you lost, 3 if the round was a draw, and 6 if you won
// 1 for Rock, 2 for Paper, and 3 for Scissors
func playRPS(p []int) int {
	if p[0] == p[1] {
		return p[1] + 3
	} else if (p[1] == 1 && p[0] == 3) || (p[1] == 2 && p[0] == 1) || (p[1] == 3 && p[0] == 2) {
		return p[1] + 6
	}
	return p[1]
}

/*
The winner of the whole tournament is the player with the highest score.
Your total score is the sum of your scores for each round. The score for a single
round is the score for the shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors)
plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won)
What would your total score be if everything goes exactly according to your strategy guide?
*/
func problem1(plays [][]int) int {
	var score int

	for _, p := range plays {
		score += playRPS(p)
	}

	return score
}

/*
the second column says how the round needs to end: X (1) means you need to lose,
Y (2) means you need to end the round in a draw, and Z (3) means you need to win
The total score is still calculated in the same way, but now you need to figure out
what shape to choose so the round ends as indicated. what would your total score be if
everything goes exactly according to your strategy guide?
*/
func problem2(plays [][]int) int {
	var score int

	for _, p := range plays {
		if p[1] == 1 {
			if p[0] == 1 {
				p[1] = 3
			} else if p[0] == 2 {
				p[1] = 1
			} else {
				p[1] = 2
			}
		} else if p[1] == 2 {
			p[1] = p[0]
		} else if p[1] == 3 {
			if p[0] == 1 {
				p[1] = 2
			} else if p[0] == 2 {
				p[1] = 3
			} else {
				p[1] = 1
			}
		}
		score += playRPS(p)
	}

	return score
}

func parseInput(filePath string) [][]int {
	f, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var plays [][]int

	for scanner.Scan() {
		game := strings.Split(scanner.Text(), " ")

		for i, g := range game {
			if g == "A" || g == "X" {
				game[i] = "1"
			} else if g == "B" || g == "Y" {
				game[i] = "2"
			} else if g == "C" || g == "Z" {
				game[i] = "3"
			}
		}

		ints := make([]int, len(game))

		for i, g := range game {
			ints[i], _ = strconv.Atoi(g)
		}

		plays = append(plays, ints)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return plays
}
