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
	fmt.Printf("%+v\n", filesystem)
	getDirSizes(&filesystem)
	fmt.Printf("%+v\n", filesystem)
	fmt.Printf("%+v\n", filesystem.children[0])
	fmt.Printf("%+v\n", filesystem.children[0].children[0])
	fmt.Printf("%+v\n", filesystem.children[1])
	problem1(&filesystem)
	fmt.Printf("Problem 1: %v\n", p1count)
}

var p1count int

func problem1(d *Directory) {
	if d.dirSize <= 100000 {
		p1count += d.dirSize
	}

	for _, child := range d.children {
		problem1(child)
	}
}

func getDirSizes(d *Directory) {
	for _, v := range d.files {
		d.dirSize += v
	}

	for _, child := range d.children {
		getDirSizes(child)
		d.dirSize += child.dirSize
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
