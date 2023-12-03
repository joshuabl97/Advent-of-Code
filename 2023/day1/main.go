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

	input := parseInput(*filePath)
	fmt.Printf("Problem 1: %v\n", p1(input))
	fmt.Printf("Problem 2: %v\n", p2(input))
}

/*
The newly-improved calibration document consists of lines of text;
each line originally contained a specific calibration value that the Elves
now need to recover. On each line, the calibration value can be found by combining
the first digit and the last digit (in that order) to form a single two-digit number.
What is the sum of all of the calibration values?
*/
func p1(input []string) int {
	var total int

	for _, str := range input {
		s := strings.Split(str, "")

		var (
			firstdigit string
			secondigit string
		)

		// grab the first digit and add it to the total
		for _, char := range s {
			if _, err := strconv.Atoi(char); err == nil {
				firstdigit = char
				break
			}
		}

		// grab the last digit and add it to the total
		for i := len(s) - 1; i >= 0; i-- {
			if _, err := strconv.Atoi(s[i]); err == nil {
				secondigit = s[i]
				break
			}
		}

		strNum, _ := strconv.Atoi(fmt.Sprintf("%v%v", firstdigit, secondigit))
		total += strNum
	}

	return total
}

/*
Your calculation isn't quite right. It looks like some of the digits are actually
spelled out with letters: one, two, three, four, five, six, seven, eight, and nine
also count as valid "digits".
*/
func p2(input []string) int {
	var total int

	for _, str := range input {
		var num1 int
		var num2 int

		for i := 0; i < len(str); i++ {
			currentForward := ""
			for j := i; j < len(str); j++ {
				currentForward += string(str[j])

				if num, isNum := spelledOutToNumber(currentForward); isNum == true {
					num2 = num
				}

				if num, err := strconv.Atoi(currentForward); err == nil {
					num2 = num
				}
			}
		}

		for i := len(str) - 1; i >= 0; i-- {
			currentBackward := ""
			for j := i; j >= 0; j-- {
				currentBackward = string(str[j]) + currentBackward

				if num, isNum := spelledOutToNumber(currentBackward); isNum == true {
					num1 = num
				}

				if num, err := strconv.Atoi(currentBackward); err == nil {
					num1 = num
				}
			}
		}

		strNum, _ := strconv.Atoi(fmt.Sprintf("%v%v", num1, num2))
		total += strNum
	}

	return total
}

var spelledOutNumbers = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func spelledOutToNumber(input string) (int, bool) {
	input = strings.ToLower(input)
	input = strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || r == ' ' {
			return r
		}
		return -1
	}, input)

	words := strings.Fields(input)

	var number int
	var found bool

	for i := 0; i < len(words); i++ {
		val, exists := spelledOutNumbers[words[i]]
		if exists {
			// Accumulate the numeric value
			number += val

			// Check for subsequent words to form larger numbers
			if i+1 < len(words) {
				nextVal, nextExists := spelledOutNumbers[words[i+1]]
				if nextExists && nextVal < 1000 {
					number += nextVal
					i++ // Skip the next word since it's part of the number
				}
			}
			found = true
			break
		}
	}

	return number, found
}

// each line is a string
func parseInput(filePath string) []string {
	f, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
