package main

import (
	"fmt"

	"github.com/iamjohnnym/advent-of-code/go/lib"
)

var numbers []int

func main() {
	// load contents of file
	content := lib.LoadFile("../challenge.txt")
	// convert content to an array of `int`
	numbers := lib.NumberConverter(content)
	// return sum of array of `int`
	fmt.Println(lib.Sum(numbers))
}
