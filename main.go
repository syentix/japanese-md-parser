package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

var reg = regexp.MustCompile("/j\"[a-zA-Z ]+\"")

func main() {
	file, err := ioutil.ReadFile("Lesson-C.md")
	if err != nil {
		panic(err)
	}
	// Go line by line
	// Take the array of found matches
	// Translate and then replace
	first := reg.Match(file)

	fmt.Println(first)
}
