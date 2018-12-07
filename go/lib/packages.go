package lib

import (
	"fmt"
	"io/ioutil"
)

// file name of the input file
var fileName string
var iterator []float64

// LoadFile Load input file as a string
func LoadFile(fileName string) string {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Unable to read file", err)
		return "Unable to read file"
	}
	return string(content)
}

func Sum(iterator []float64) float64 {
	total := 0.0
	for _, value := range iterator {
		total += value
	}
	return total
}
