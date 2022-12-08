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
var level int

type Directory struct {
	cwd       string
	files     map[string]int
	fileSizes int
	parent    *Directory
	children  []*Directory
}

func main() {
	flag.Parse()
	filesystem := parseInput(*filePath)
	problem1(&filesystem, nil)
	fmt.Printf("%T %v\n", filesystem.parent, filesystem.parent)
	fmt.Printf("%T %v\n", filesystem.children[0], filesystem.children[0])
}

func problem1(root *Directory, q []*Directory) {
	var queue []*Directory
	var nextLevel []*Directory

	if len(q) > 0 {
		queue = append(queue, q...)
	}

	if level == 0 {
		queue = append(queue, root)
	}

	for len(queue) > 0 {
		nextLevel = append(nextLevel, queue[0].children...)
		fmt.Printf("%v\n", queue[0].fileSizes)
		for _, s := range queue[0].children {
			addUp(s, s.fileSizes)
		}
		fmt.Printf("%v\n", queue[0].fileSizes)

		queue = queue[1:]
		level++
	}
	if len(nextLevel) > 0 {
		fmt.Printf("%v\n", nextLevel)
		problem1(root, nextLevel)
	}
}

func addUp(dir *Directory, addMe int) {
	if dir.parent != nil {
		dir.parent.fileSizes += addMe
		addUp(dir.parent, addMe)
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
			current.fileSizes += s
			current.files[line[1]] = s
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return root
}
