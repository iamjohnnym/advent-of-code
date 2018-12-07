package main

import (
	"fmt"
	"lib"
	"strconv"
	"strings"
)

var numbers []float64

func sum(frequencies []float64) float64 {
	total := 0.0
	for _, value := range frequencies {
		total += value
	}
	return total
}

func main() {
	content := lib.ReadChallengeFile("challenge.txt")
	for _, line := range strings.Split(string(content), "\n") {
		if n, err := strconv.ParseFloat(line, 64); err == nil {
			numbers = append(numbers, n)
		}
	}
	total := sum(numbers)
	fmt.Println(total)
}
