package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type marker struct {
	letters int
	repeats int
	length  int // the number of chars in the marker
}

func findMarker(input string) marker {
	var newMarker marker
	values := strings.Split(input, "x")
	letters, _ := strconv.Atoi(strings.Replace(values[0], "(", "", 1))
	repeats, _ := strconv.Atoi(strings.Split(values[1], ")")[0])
	length := len(fmt.Sprintf("(%dx%d)", letters, repeats))
	newMarker.letters = letters
	newMarker.repeats = repeats
	newMarker.length = length
	return newMarker
}

func decompress(input string) string {
	var buffer bytes.Buffer
	charIndex := 0
	for {
		char := input[charIndex]
		if char == '(' {
			nextMarker := findMarker(input[charIndex:])
			charIndex += nextMarker.length
			lettersToRepeat := decompress(input[charIndex : charIndex+nextMarker.letters])
			for i := 0; i < nextMarker.repeats; i++ {
				buffer.WriteString(lettersToRepeat)
			}
			charIndex += nextMarker.letters
		} else {
			buffer.WriteString(string(char))
			charIndex++
		}
		if charIndex > len(input)-1 {
			break
		}
	}
	return buffer.String()
}

func main() {
	// Read the file and join all the lines into one long string.
	data, _ := ioutil.ReadFile("input.txt")
	input := ""
	for _, line := range strings.Split(string(data), "\n") {
		for _, char := range line {
			input = fmt.Sprintf("%s%c", input, char)
		}
	}
	output := decompress(input)
	fmt.Println("Part 2 Solution:", len(output))
}
