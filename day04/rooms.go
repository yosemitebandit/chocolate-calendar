package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

func splitRoomAndChecksum(text string) (string, string, int, string) {
	// Some spurious text wrangling..
	values := strings.Split(text, "-")
	roomWithDashes := strings.Join(values[0:len(values)-1], "-")
	room := strings.Join(values[0:len(values)-1], "")
	sectorAndChecksum := strings.Split(values[len(values)-1], "[")
	sector, _ := strconv.Atoi(sectorAndChecksum[0])
	checksum := strings.Replace(sectorAndChecksum[1], "]", "", 1)
	return roomWithDashes, room, sector, checksum
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

func rotateName(text string, number int) string {
	letters := "abcdefghijklmnopqrstuvwxyz"
	var output []string
	for _, letter := range text {
		if letter == '-' {
			output = append(output, " ")
			continue
		}
		startingIndex := strings.Index(letters, string(letter))
		endingIndex := math.Mod(float64(startingIndex+number), float64(len(letters)))
		output = append(output, string(letters[int(endingIndex)]))
	}
	return strings.Join(output, "")
}

func main() {
	var validSectors []int
	var validNames []string

	data, _ := ioutil.ReadFile("input.txt")
	for _, room := range strings.Split(string(data), "\n") {
		if room == "" {
			continue
		}
		roomWithDashes, room, sector, checksum := splitRoomAndChecksum(room)
		if checksum == generateChecksum(room) {
			validSectors = append(validSectors, sector)
			validNames = append(validNames, rotateName(roomWithDashes, sector))
		}
	}

	sum := 0
	for _, sector := range validSectors {
		sum += sector
	}
	fmt.Println("Part 1 Solution:", sum)
	fmt.Println("Part 2 Solution:")
	for index, _ := range validNames {
		if !strings.Contains(validNames[index], "storage") {
			continue
		}
		fmt.Print(validSectors[index])
		fmt.Print(":")
		fmt.Println(validNames[index])
	}
}
