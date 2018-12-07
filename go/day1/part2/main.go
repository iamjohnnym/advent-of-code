package main

import (
	"fmt"
	"os"

	"github.com/iamjohnnym/advent-of-code/go/lib"
)

var frequency_results map[int]bool
var last_frequency int

func main() {
	content := lib.LoadFile("../challenge.txt")
	frequencies := lib.NumberConverter(content)
	frequency_results = make(map[int]bool)
	for {
		for _, frequency := range frequencies {
			last_frequency += frequency
			frequency_results[frequency] = true
			if frequency_results[frequency] {
				fmt.Printf("Duplicate Frequency: %d\n", frequency)
				os.Exit(0)
			}
		}
	}
}
