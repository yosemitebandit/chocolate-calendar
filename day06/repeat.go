package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func contains(values []int, target int) bool {
	for _, element := range values {
		if element == target {
			return true
		}
	}
	return false
}

type letterCount struct {
	letter byte
	count  int
}

func main() {
	var code []string
	for col := 0; col < 8; col += 1 {
		letterFrequency := map[byte]int{}
		data, _ := ioutil.ReadFile("input.txt")

		for _, row := range strings.Split(string(data), "\n") {
			if row == "" {
				continue
			}
			_, ok := letterFrequency[row[col]]
			if ok {
				letterFrequency[row[col]] += 1
			} else {
				letterFrequency[row[col]] = 1
			}
		}

		var mostFrequentLetter letterCount
		mostFrequentLetter.letter = 'a'
		mostFrequentLetter.count = 100
		for letter, count := range letterFrequency {
			if count < mostFrequentLetter.count {
				mostFrequentLetter.letter = letter
				mostFrequentLetter.count = count
			}
		}
		code = append(code, string(mostFrequentLetter.letter))
	}
	fmt.Println("Part 2 Solution:", strings.Join(code, ""))
}
