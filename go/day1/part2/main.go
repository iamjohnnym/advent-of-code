package main

import (
	"github.com/iamjohnnym/advent-of-code/go/lib"
)

func main() {
	content := lib.LoadFile("../challenge.txt")
	frequencies := lib.NumberConverter(content)
	// Iterate through frequency array until the duplicate is found
	lib.FindDuplicateFrequency(frequencies)
}
