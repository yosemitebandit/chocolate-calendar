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

func generateBab(text string) string {
	return fmt.Sprintf("%s%s%s", string(text[1]), string(text[0]), string(text[1]))
}

func sliceOfStringsContains(input []string, target string) bool {
	for _, v := range input {
		if v == target {
			return true
		}
	}
	return false
}

func main() {
	validIPs := 0
	validSSLs := 0
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
		// Part Two.
		var vanillaAbaSequences []string
		var bracketedAbaSequences []string
		for _, v := range dip.vanilla {
			vanillaAbaSequences = append(splitAba(v), vanillaAbaSequences...)
		}
		for _, v := range dip.bracketed {
			for _, b := range splitAba(v) {
				bracketedAbaSequences = append(bracketedAbaSequences, generateBab(b))
			}
		}
		for _, v := range vanillaAbaSequences {
			if sliceOfStringsContains(bracketedAbaSequences, v) {
				validSSLs += 1
				break
			}
		}
	}
	fmt.Println("Part 1 Solution:", validIPs)
	fmt.Println("Part 2 Solution:", validSSLs)
}
