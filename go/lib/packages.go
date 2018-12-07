package lib

import (
	"fmt"
	"io/ioutil"
)

// file name of the input file
var fileName string
var iterator []int

// LoadFile Load input file as a string
func LoadFile(fileName string) string {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Unable to read file", err)
		return "Unable to read file"
	}
	return string(content)
}

// Sum iterate through array of ints to determine total sum
func Sum(iterator []int) int {
	sum := 0
	// iterate through array of ints to determine sum
	for _, value := range iterator {
		sum += value
	}
	return sum
}
