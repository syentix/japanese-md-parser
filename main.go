package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gojp/kana"
)

// Used Regex for files
var reg = regexp.MustCompile("/j\"[a-zA-Z ]+\"")

func main() {

	args := os.Args[1:]

	romaji := false

	for i, arg := range args {
		if arg == "--romaji" {
			romaji = true
			args[i] = args[len(args)-1] // Copy last element to index i.
			args[len(args)-1] = ""      // Erase last element (write zero value).
			args = args[:len(args)-1]   // Truncate slice.
		}
	}

	for _, file := range args {
		ab, err := filepath.Abs(file)
		if err != nil {
			panic(err)
		}
		parseFile(ab, romaji)
	}

}

func parseFile(fileName string, romaji bool) {
	// Read file
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("File not found.")
		return
	}

	// How many matches are in the Text.
	occurs := reg.FindAllString(string(file), len(file))
	fmt.Printf("There were %d matches found.\n", len(occurs))

	// Loop through the matches and replacing the text with Hiragana
	i := 0
	for i < len(occurs) {
		// Gives starting and ending index
		occur := reg.FindIndex(file)
		start := occur[0]
		end := occur[1]

		// Converting to hiragana
		word := kana.RomajiToHiragana(string(file[start:end]))

		// Replacing placeholder with parsed text with form: <Hiragana>(<Romaji>)
		if romaji == true {
			word = word[3:len(word)-1] + "(" + string(file[start+3:end-1]) + ")"
		} else {
			word = word[3 : len(word)-1]
		}
		file = bytes.Replace(file, file[start:end], []byte(word), len([]byte(word)))
		i++
	}

	newFileName := strings.Split(fileName, ".")[0] + "(parsed).md"
	newFile, err := os.Create(newFileName)

	if err != nil {
		panic(err)
	}

	_, err = newFile.Write(file)
	if err != nil {
		panic(err)
	}

	fmt.Println("File successfully parsed and can be found here:", newFileName)
}
