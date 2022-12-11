package main

import (
	"aoc2022/reusable"
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
	trees := parseInput(*filePath)
	fmt.Printf("Problem 1: %v\n", problem1(trees))
	fmt.Printf("Problem 2: %v\n", problem2(trees))
}

/*
A tree is visible if all of the other trees between it and an edge of the grid are shorter than it
how many trees are visible from outside the grid?
*/
func problem1(trees [][]int) int {
	canSee := len(trees)*2 + len(trees[0])*2 - 4

	for i1, t := range trees {
	treeSearch:
		for i2, tree := range t {
			if i1 > 0 && i1 != len(trees)-1 && i2 != 0 && i2 != len(t)-1 {

				var (
					top    bool
					bottom bool
					right  bool
					left   bool
				)

				for _, topScan := range trees[:i1] {
					if tree <= topScan[i2] {
						top = true
						break
					}
				}

				if !top {
					canSee++
					continue treeSearch
				}

				for _, bottomScan := range trees[i1+1:] {
					if tree <= bottomScan[i2] {
						bottom = true
						break
					}
				}

				if !bottom {
					canSee++
					continue treeSearch
				}

				for i3, s1 := range t {
					if i2 < i3 { //right
						if tree <= s1 {
							right = true
						}
					}

					if i2 > i3 { //left
						if tree <= s1 {
							left = true
						}

					}
				}

				if !left || !right {
					canSee++
					continue treeSearch
				}
			}
		}
	}

	return canSee
}

/*
To measure the viewing distance from a given tree, look up, down, left, and right
from that tree; stop if you reach an edge or at the first tree that is the same height
or taller than the tree under consideration. (If a tree is right on the edge, at least one
of its viewing distances will be zero.)
A tree's scenic score is found by multiplying together
its viewing distance in each of the four directions
Consider each tree on your map. What is the highest scenic score possible for any tree?
*/
func problem2(trees [][]int) int {
	var count int

	for i1, t := range trees {
		for i2, tree := range t {

			var top []int
			var topCount int
			if i1 > 0 {
				for _, topScan := range trees[:i1] {
					top = append(top, topScan[i2])
				}
				if len(top) > 0 {
					topTrees := reusable.ReverseIntSlice(top)
					for _, t := range topTrees {
						if tree <= t {
							topCount++
							break
						}
						topCount++
					}
				}
			} else {
				topCount = 1
			}

			var belowCount int
			if i1 != len(trees)-1 {
				for _, bottomScan := range trees[i1+1:] {
					if tree <= bottomScan[i2] {
						belowCount++
						break
					}
					belowCount++
				}
			} else {
				belowCount = 1
			}

			var rightCount int
			if i2 != len(t)-1 {
				for i3, s1 := range t {
					if i2 < i3 {
						if tree <= s1 {
							rightCount++
							break
						}
						rightCount++
					}
				}
			} else {
				rightCount = 1
			}

			var leftSide []int
			var leftCount int
			if i2 != 0 {
				for i3, s1 := range t {
					if i2 > i3 {
						leftSide = append(leftSide, s1)
					}
				}
				if len(leftSide) > 0 {
					s := reusable.ReverseIntSlice(leftSide)
					for _, otherTree := range s {
						fmt.Printf("left %v\n", otherTree)
						if tree <= otherTree {
							leftCount++
							break
						}
						leftCount++
					}
				}
			} else {
				leftCount = 1
			}

			x := belowCount * topCount * leftCount * rightCount
			if x > count {
				count = x
			}
		}
	}

	return count
}

func parseInput(filepath string) [][]int {
	f, err := os.Open(*filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var trees [][]int

	for scanner.Scan() {
		treeSlice := strings.Split(scanner.Text(), "")
		var tree []int
		for _, t := range treeSlice {
			if s, err := strconv.Atoi(t); err == nil {
				tree = append(tree, s)
			}
		}
		trees = append(trees, tree)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return trees
}
