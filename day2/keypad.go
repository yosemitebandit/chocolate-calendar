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

func main() {
	data, _ := ioutil.ReadFile("input.txt")

	runeConverter := map[rune]position{
		'U': {x: 0, y: 1},
		'R': {x: 1, y: 0},
		'D': {x: 0, y: -1},
		'L': {x: -1, y: 0},
	}

	allowedMoves := map[position]string{
		{x: 0, y: 2}:   "D",
		{x: -1, y: 1}:  "RD",
		{x: 0, y: 1}:   "URDL",
		{x: 1, y: 1}:   "LD",
		{x: -2, y: 0}:  "R",
		{x: -1, y: 0}:  "URDL",
		{x: 0, y: 0}:   "URDL",
		{x: 1, y: 0}:   "URDL",
		{x: 2, y: 0}:   "L",
		{x: -1, y: -1}: "UR",
		{x: 0, y: -1}:  "URDL",
		{x: 1, y: -1}:  "UL",
		{x: 0, y: -2}:  "U",
	}

	positionConverter := map[position]rune{
		{x: 0, y: 2}:   '1',
		{x: -1, y: 1}:  '2',
		{x: 0, y: 1}:   '3',
		{x: 1, y: 1}:   '4',
		{x: -2, y: 0}:  '5',
		{x: -1, y: 0}:  '6',
		{x: 0, y: 0}:   '7',
		{x: 1, y: 0}:   '8',
		{x: 2, y: 0}:   '9',
		{x: -1, y: -1}: 'A',
		{x: 0, y: -1}:  'B',
		{x: 1, y: -1}:  'C',
		{x: 0, y: -2}:  'D',
	}

	currentPosition := position{x: -2, y: 0}
	var doorCode []string

	for _, element := range strings.Fields(string(data)) {
		for _, value := range element {
			// Check if the move is allowed based on the current position.
			if strings.Contains(allowedMoves[currentPosition], string(value)) {
				currentPosition.x += runeConverter[value].x
				currentPosition.y += runeConverter[value].y
			}
		}
		doorCode = append(doorCode, string(positionConverter[currentPosition]))
	}

	fmt.Println("Part 1 Solution:", doorCode)
}
