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
	str := parseInput(*filePath)
	fmt.Printf("Problem 1: %v\n", problem1(str))
	fmt.Printf("Problem 2: %v\n", problem2(str))
}

/*
To fix the communication system, you need to add a subroutine to the device that detects a start-of-packet
marker in the datastream. In the protocol being used by the Elves, the start of a packet is indicated by a
sequence of four characters that are all different.
How many characters need to be processed before the first start-of-packet marker is detected?
*/
func problem1(str string) int {
	s := strings.Split(str, "")

	for i := 0; i < len(s); i++ {
		if !dupeExist(s[i : i+4]) {
			return i + 4
		}
	}

	return 0
}

/*
Your device's communication system is correctly detecting packets, but still isn't working. It looks like it also
needs to look for messages. A start-of-message marker is just like a start-of-packet marker, except it consists of
14 distinct characters rather than 4.
How many characters need to be processed before the first start-of-message marker is detected?
*/
func problem2(str string) int {
	s := strings.Split(str, "")

	for i := 0; i < len(s); i++ {
		if !dupeExist(s[i : i+14]) {
			return i + 14
		}
	}

	return 0
}

// checks if there are multiple of the same value in the same array of strings
// i.e ['a','a'] -> true ['a','b'] -> false
func dupeExist(arr []string) bool {
	visited := make(map[string]bool, 0)

	for i := 0; i < len(arr); i++ {
		if visited[arr[i]] {
			return true
		} else {
			visited[arr[i]] = true
		}
	}

	return false
}

func parseInput(filePath string) string {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var str string

	for scanner.Scan() {
		str = str + scanner.Text()
	}

	return str
}
