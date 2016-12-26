package main

import (
	"fmt"
	"math"
	"strconv"
)

func coordIsWall(x int, y int, n int) bool {
	value := x*x + 3*x + 2*x*y + y + y*y + n
	inBaseTwo := strconv.FormatInt(int64(value), 2)
	sum := 0
	for _, bit := range inBaseTwo {
		value, _ = strconv.Atoi(string(bit))
		sum += value
	}
	if math.Mod(float64(sum), 2) == 0 {
		return false
	} else {
		return true
	}
}

func main() {
	magicNumber := 1364
	targetCoords := [2]int{31, 39}

	side := 45
	var coords = make([]string, side*side)
	for i := 0; i < side*side; i++ {
		x := math.Mod(float64(i), float64(side))
		y := math.Mod(float64(i/side), float64(side))
		if coordIsWall(int(x), int(y), magicNumber) {
			coords[i] = "#"
		} else {
			coords[i] = " "
		}
		if int(x) == targetCoords[0] && int(y) == targetCoords[1] {
			coords[i] = "o"
		}
	}

	// Plot.
	for i := 0; i < side; i++ {
		fmt.Println(coords[side*i : side*(i+1)])
	}
}
