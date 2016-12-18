package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func abbaSequenceIsPresent(text string) bool {
	for index, _ := range text {
		if index == len(text)-3 {
			return false
		}
		if text[index] == text[index+3] && text[index+1] == text[index+2] {
			if text[index+1] != text[index+3] {
				return true
			}
		}
	}
	return false
}

type dividedIP struct {
	vanilla   []string
	bracketed []string
}

func splitIP(text string) dividedIP {
	var dip dividedIP
	dividedText := strings.FieldsFunc(text, func(r rune) bool {
		return r == '[' || r == ']'
	})
	for index, element := range dividedText {
		if math.Mod(float64(index), 2) == 0 {
			dip.vanilla = append(dip.vanilla, element)
		} else {
			dip.bracketed = append(dip.bracketed, element)
		}
	}
	return dip
}

func splitAba(text string) []string {
	var sequences []string
	for index, _ := range text {
		if index+2 >= len(text) {
			continue
		}
		if text[index] == text[index+2] && text[index] != text[index+1] {
			sequences = append(sequences, text[index:index+3])
		}
	}
	return sequences
}

func main() {
	validIPs := 0
	data, _ := ioutil.ReadFile("input.txt")
	for _, row := range strings.Split(string(data), "\n") {
		if row == "" {
			continue
		}
		dip := splitIP(row)
		validIP := false
		for _, v := range dip.vanilla {
			if abbaSequenceIsPresent(v) {
				validIP = true
			}
		}
		for _, b := range dip.bracketed {
			if abbaSequenceIsPresent(b) {
				validIP = false
			}
		}
		if validIP {
			validIPs += 1
		}
	}
	fmt.Println("Part 1 Solution:", validIPs)
}
