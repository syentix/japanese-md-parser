package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/gojp/kana"
)

// Used Regex for files
var reg = regexp.MustCompile("/j\"[a-zA-Z ]+\"")

func main() {

	// TODO: Dehardcode this bit
	fileName := "./examples/Lesson-C"

	// Read file
	file, err := ioutil.ReadFile(fileName + ".md")
	if err != nil {
		panic(err)
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
		word = word[3:len(word)-1] + "(" + string(file[start+3:end-1]) + ")"
		file = bytes.Replace(file, file[start:end], []byte(word), len([]byte(word)))
		i++
	}

	newFile, err := os.Create(fileName + "(parsed).md")
	if err != nil {
		panic(err)
	}

	_, err = newFile.Write(file)
	if err != nil {
		panic(err)
	}

	fmt.Println("File successfully parsed and can be found here:", fileName+"(parsed).md")
}
