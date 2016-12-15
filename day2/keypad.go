package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type position struct {
	x float64
	y float64
}

func clamp(value float64) float64 {
	if value > 1 {
		return 1
	}
	if value < -1 {
		return -1
	}
	return value
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")

	runeConverter := map[rune]position{
		'U': {x: 0, y: 1},
		'R': {x: 1, y: 0},
		'D': {x: 0, y: -1},
		'L': {x: -1, y: 0},
	}

	positionConverter := map[position]int{
		{x: -1, y: 1}:  1,
		{x: 0, y: 1}:   2,
		{x: 1, y: 1}:   3,
		{x: -1, y: 0}:  4,
		{x: 0, y: 0}:   5,
		{x: 1, y: 0}:   6,
		{x: -1, y: -1}: 7,
		{x: 0, y: -1}:  8,
		{x: 1, y: -1}:  9,
	}

	currentPosition := position{x: 0, y: 0}
	var doorCode []int

	for _, element := range strings.Fields(string(data)) {
		for _, value := range element {
			currentPosition.x = clamp(currentPosition.x + runeConverter[value].x)
			currentPosition.y = clamp(currentPosition.y + runeConverter[value].y)
		}
		doorCode = append(doorCode, positionConverter[currentPosition])
	}

	fmt.Println("Part 1 Solution:", doorCode)
}
