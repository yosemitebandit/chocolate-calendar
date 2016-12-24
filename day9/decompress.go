package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
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
	fmt.Println("single round decompression input length:", len(input))
	parensInInput := 0
	for _, char := range input {
		if char == '(' {
			parensInInput++
		}
	}
	fmt.Println("Parens in input:", parensInInput)
	// Walk through the string and build the output.
	var buffer bytes.Buffer
	charIndex := 0
	for {
		char := input[charIndex]
		if char == '(' {
			nextMarker := findMarker(input[charIndex:])
			charIndex += nextMarker.length
			lettersToRepeat := input[charIndex : charIndex+nextMarker.letters]
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
		if math.Mod(float64(charIndex), 1000) == 0 {
			fmt.Println(float64(100*charIndex) / float64(len(input)))
		}
	}
	return buffer.String()
}

func decompressV2(input string) string {
	// Keep decompressing until there are no more markers.
	rounds := 0
	for {
		//fmt.Println("Using input:", input)
		input = decompress(input)
		// Check for a marker.
		noParenFound := true
		for _, char := range input {
			if char == '(' {
				noParenFound = false
				//fmt.Println("paren!")
				break
			}
		}
		if noParenFound {
			break
		}
		fmt.Println(rounds)
		rounds++
	}
	return input
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
	output := decompressV2(input)
	fmt.Println("Part 2 Solution:", len(output))
}
