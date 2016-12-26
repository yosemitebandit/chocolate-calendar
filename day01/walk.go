package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type position struct {
	x float64
	y float64
}

func contains(history []position, p position) bool {
	for _, element := range history {
		if p == element {
			return true
		}
	}
	return false
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")

	// Define the starting position and angle.
	currentPosition := position{x: 0, y: 0}
	angle := 90
	// Store position history.
	var positionHistory []position

Outer:
	for _, element := range strings.Fields(string(data)) {
		// Breakdown to "L3" and "R2" pieces.
		value := strings.Replace(element, ",", "", 1)
		// Update the angle based on L vs R.
		if string(value[0]) == "L" {
			angle += 90
		} else if string(value[0]) == "R" {
			angle -= 90
		}
		// Get the total number of blocks to move.
		blocks, _ := strconv.Atoi(value[1:])
		// Walk each block.
		i := 1
		for i <= blocks {
			// Get position delta.
			xdelta := math.Cos(float64(angle) * math.Pi / 180)
			ydelta := math.Sin(float64(angle) * math.Pi / 180)
			// Update the position.
			currentPosition.x += xdelta
			currentPosition.y += ydelta
			// Check to see if the current position has been visited before.
			if contains(positionHistory, currentPosition) {
				fmt.Println("Deja vu!")
				break Outer
			}
			// Store the new position.
			positionHistory = append(positionHistory, currentPosition)
			i += 1
		}
	}

	fmt.Println("Part 2 Solution:", int(currentPosition.x+currentPosition.y))
}
