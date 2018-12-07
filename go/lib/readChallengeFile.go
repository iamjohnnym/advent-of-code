package lib

import (
	"fmt"
	"io/ioutil"
)

// file name of the input file
var fileName string

// LoadFile Load input file as a string
func LoadFile(fileName string) string {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Unable to read file", err)
		return "Unable to read file"
	}
	return string(content)
}
