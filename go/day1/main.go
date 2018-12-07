package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/iamjohnnym/advent-of-code/go/lib"
)

var numbers []float64

func main() {
	content := lib.LoadFile("challenge.txt")
	for _, line := range strings.Split(string(content), "\n") {
		if n, err := strconv.ParseFloat(line, 64); err == nil {
			numbers = append(numbers, n)
		}
	}
	total := lib.Sum(numbers)
	fmt.Println(total)
}
