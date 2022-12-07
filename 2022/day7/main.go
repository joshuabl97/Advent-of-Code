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

type directoryContents struct {
	cwd         string
	directories []string
	files       map[string]int
}

func main() {
	flag.Parse()
	filesystem := parseInput(*filePath)
	fmt.Printf("Problem 1: %v\n", problem1(filesystem))
}

func dirSum(filesystem []directoryContents, folder string) int {
	var count int

	for _, dir := range filesystem {
		if dir.cwd == folder {
			for _, v := range dir.files {
				count += v
			}
			for _, d := range dir.directories {
				count += dirSum(filesystem, d)
			}
		}
	}

	return count
}

func problem1(filesystem []directoryContents) int {
	results := make(map[string]int)
	var count int

	for _, dir := range filesystem {
		results[dir.cwd] = dirSum(filesystem, dir.cwd)
	}

	for _, v := range results {
		if v <= 100000 {
			count += v
		}
	}

	return count
}

func parseInput(filePath string) []directoryContents {
	f, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var (
		filesystem       []directoryContents
		directoryTracker []string
		isExist          bool
	)

filescan:
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "$") {
			switch scanner.Text() {
			case "$ cd /":
				directoryTracker = []string{"/"}
			case "$ cd ..":
				if len(directoryTracker) > 0 {
					directoryTracker = directoryTracker[:len(directoryTracker)-1]
				}
			case "$ ls":
				continue filescan
			default:
				directoryTracker = append(directoryTracker, strings.TrimPrefix(scanner.Text(), "$ cd "))
			}
			continue
		}

		for _, f := range filesystem {
			if f.cwd == directoryTracker[len(directoryTracker)-1] {
				isExist = true
			}
		}

		if !isExist {
			pwd := directoryContents{
				cwd: directoryTracker[len(directoryTracker)-1],
			}

			filesystem = append(filesystem, pwd)
		}
		isExist = false

		for i, f := range filesystem {
			if f.cwd == directoryTracker[len(directoryTracker)-1] {
				if strings.Contains(scanner.Text(), "dir") {
					filesystem[i].directories = append(filesystem[i].directories, strings.Split(scanner.Text(), " ")[1])
					continue filescan
				}

				fileStrs := strings.Split(scanner.Text(), " ")

				if f.files == nil {
					m := make(map[string]int)
					filesystem[i].files = m
				}

				if num, err := strconv.Atoi(fileStrs[0]); err == nil {
					filesystem[i].files[fileStrs[1]] = num
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return filesystem
}
