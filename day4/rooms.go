package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func splitRoomAndChecksum(text string) (string, int, string) {
	// Some spurious text wrangling..
	values := strings.Split(text, "-")
	room := strings.Join(values[0:len(values)-1], "")
	sectorAndChecksum := strings.Split(values[len(values)-1], "[")
	sector, _ := strconv.Atoi(sectorAndChecksum[0])
	checksum := strings.Replace(sectorAndChecksum[1], "]", "", 1)
	return room, sector, checksum
}

func contains(values []int, target int) bool {
	for _, element := range values {
		if element == target {
			return true
		}
	}
	return false
}

func generateChecksum(text string) string {
	// First build a map to track frequency of letters.
	letterFrequency := map[string]int{}
	for _, letter := range text {
		_, ok := letterFrequency[string(letter)]
		if ok {
			letterFrequency[string(letter)] += 1
		} else {
			letterFrequency[string(letter)] = 1
		}
	}
	var counts []int
	for _, v := range letterFrequency {
		if !contains(counts, v) {
			counts = append(counts, v)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))
	// Now map the counts -> letters.
	countsToLetters := map[int][]string{}
	for letter, count := range letterFrequency {
		countsToLetters[count] = append(countsToLetters[count], letter)
	}
	// Within each letter set, alphabetize.
	for key, _ := range countsToLetters {
		sort.Strings(countsToLetters[key])
	}
	// Glue them all together.
	var values []string
	for _, count := range counts {
		for _, letter := range countsToLetters[count] {
			values = append(values, letter)
		}
	}
	// Take the first five.
	return strings.Join(values[0:5], "")
}

func main() {
	var validSectors []int

	data, _ := ioutil.ReadFile("input.txt")
	for _, room := range strings.Split(string(data), "\n") {
		if room == "" {
			continue
		}
		room, sector, checksum := splitRoomAndChecksum(room)
		if checksum == generateChecksum(room) {
			validSectors = append(validSectors, sector)
		}
	}

	sum := 0
	for _, sector := range validSectors {
		sum += sector
	}
	fmt.Println("Part 1 Solution:", sum)
}
