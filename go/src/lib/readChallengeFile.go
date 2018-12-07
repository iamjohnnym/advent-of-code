package lib

import (
	"fmt"
	"io/ioutil"
)

var fileName string

func ReadChallengeFile(fileName string) string {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Unable to read file", err)
		return "Unable to read file"
	}
	return string(content)
}
