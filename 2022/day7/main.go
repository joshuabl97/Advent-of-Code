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

type Directory struct {
	cwd      string
	files    map[string]int
	dirSize  int
	parent   *Directory
	children []*Directory
}

func main() {
	flag.Parse()
	filesystem := parseInput(*filePath)
	getDirSizes(&filesystem)
	problem1(&filesystem)

	fmt.Printf("Problem 1: %v\n", p1count)

	freeDiskSpace := 70000000 - filesystem.dirSize

	problem2(&filesystem, freeDiskSpace)
	/*
		Find the smallest directory that, if deleted, would free up enough space on the
		filesystem to run the update. What is the total size of that directory?
	*/
	smallestNumber := p2sizes[0]

	for _, num := range p2sizes {
		if num < smallestNumber {
			smallestNumber = num
		}
	}

	fmt.Printf("Problem 2: %v\n", smallestNumber)
}

// populates the total directory size for each Directory in tree
// (including the value of the file sizes in their subdirectories recussively)
func getDirSizes(d *Directory) {
	for _, v := range d.files {
		d.dirSize += v
	}

	for _, child := range d.children {
		getDirSizes(child)
		d.dirSize += child.dirSize
	}
}

/*
find all of the directories with a total size of at
most 100000, then calculate the sum of their total sizes
*/
var p1count int

func problem1(d *Directory) {
	if d.dirSize <= 100000 {
		p1count += d.dirSize
	}

	for _, child := range d.children {
		problem1(child)
	}
}

/*
The total disk space available to the filesystem is 70000000. To run the update,
you need unused space of at least 30000000. You need to find a directory you can
delete that will free up enough space to run the update.
*/
var p2sizes []int

func problem2(d *Directory, freeDiskSpace int) {
	if d.dirSize > 30000000-freeDiskSpace {
		p2sizes = append(p2sizes, d.dirSize)
	}

	for _, child := range d.children {
		problem2(child, freeDiskSpace)
	}
}

func parseInput(filePath string) Directory {
	f, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var (
		root    Directory
		current *Directory
	)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		if strings.HasPrefix(scanner.Text(), "$ cd") {
			if current == nil {
				current = &root
				current.cwd = line[2]
				continue
			}

			switch line[2] {
			case "..":
				current = current.parent
			default:
				newDir := &Directory{
					parent: current,
					cwd:    line[2],
				}
				current.children = append(current.children, newDir)
				current = newDir
			}
			continue
		}

		if current.files == nil {
			m := make(map[string]int)
			current.files = m
		}

		if s, err := strconv.Atoi(line[0]); err == nil {
			current.files[line[1]] = s
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return root
}
