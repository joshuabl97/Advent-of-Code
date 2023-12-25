package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/joshuabl97/Advent-of-Code/2023/utils"
)

var filePath = flag.String("f", "input.txt", "path to the input file")
var winningNumLen = flag.Int("t", 10, "the number of winning numbers per lotto card")

type Card struct {
	card     int
	winning  []int
	userNums []int
}

func main() {
	flag.Parse()
	input := utils.GetInputStrings(*filePath)

	cards := parse(input, *winningNumLen)
	fmt.Printf("Problem 1: %v\n", problem1(cards))
	fmt.Printf("Problem 2: %v\n", problem2(cards))
}

// it looks like each card has two lists of numbers:
// a list of winning numbers and then a list of numbers you have
// you have to figure out which of the numbers you have appear in the list of winning numbers.
// The first match makes the card worth one point and each match after the first doubles the point value of that card.
func problem1(cards []Card) int {
	var total int
	for _, card := range cards {
		var count int
		for _, userNum := range card.userNums {
			for _, winningNum := range card.winning {
				if userNum == winningNum {
					if count == 0 {
						count++
					} else {
						count = count * 2
					}
				}
			}
		}
		total += count
	}

	return total
}

// you win copies of the scratchcards below the winning card equal to the number of matches.
// So, if card 10 were to have 5 matching numbers, you would win one copy each of cards 11, 12, 13, 14, and 15.
// Copies of scratchcards are scored like normal scratchcards and have the same card number as the card they copied.
// So, if you win a copy of card 10 and it has 5 matching numbers, it would then win a copy of the same cards that the original
// card 10 won: cards 11, 12, 13, 14, and 15. This process repeats until none of the copies cause you to win any more cards.
// (Cards will never make you copy a card past the end of the table.)
func problem2(cards []Card) int {
	var total int

	return total
}

func parse(rows []string, userNumLen int) []Card {
	var cards []Card

	for i, row := range rows {
		strs := strings.Fields(row)

		var (
			numCount int
			winning  []int
			userNums []int
		)
		for _, str := range strs {
			num, err := strconv.Atoi(str)
			if err == nil && numCount < userNumLen {
				numCount++
				winning = append(winning, num)
			} else if err == nil {
				numCount++
				userNums = append(userNums, num)
			}
		}
		cards = append(cards, Card{
			card:     i + 1,
			winning:  winning,
			userNums: userNums,
		})
	}

	return cards
}
