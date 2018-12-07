package lib

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// file name of the input file
var fileName string
var iterator []int
var numbers []int
var frequency_results map[int]bool
var last_frequency int

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

// NumbersArray
func NumberConverter(stringContent string) []int {
	for _, line := range strings.Split(stringContent, "\n") {
		if n, err := strconv.Atoi(line); err == nil {
			numbers = append(numbers, n)
		}
	}
	return numbers
}

// FindDuplicateFrequency
func FindDuplicateFrequency(frequencies []int) []int {
	frequency_results = make(map[int]bool)
	for {
		for _, frequency := range frequencies {
			last_frequency += frequency
			if frequency_results[last_frequency] {
				fmt.Printf("Duplicate Frequency: %d\n", last_frequency)
				os.Exit(0)
			}
			frequency_results[last_frequency] = true
		}
	}
}
