package main

import (
	"flag"
	"fmt"

	"github.com/joshuabl97/Advent-of-Code/2023/utils"
)

var filePath = flag.String("f", "input.txt", "path to the input file")

func main() {
	flag.Parse()
	input := utils.ParseInput(*filePath)
	fmt.Printf("%v\n", input)
}
