package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/iamjohnnym/advent-of-code/go/lib"
)

var numbers []int

func main() {
	content := lib.LoadFile("../challenge.txt")
	for _, line := range strings.Split(content, "\n") {
		// convert `line`string to `n`int and append []int{numbers}
		if n, err := strconv.Atoi(line); err == nil {
			numbers = append(numbers, n)
		}
	}
	fmt.Println(lib.Sum(numbers))
}
