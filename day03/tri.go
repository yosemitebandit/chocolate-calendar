package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func convert(v []string) []int {
	r := make([]int, len(v))
	r[0], _ = strconv.Atoi(v[0])
	r[1], _ = strconv.Atoi(v[1])
	r[2], _ = strconv.Atoi(v[2])
	return r
}

func main() {
	// For part two we have to read columns..so we'll build a large vector first.
	var sideLengths []int
	for _, col := range [3]int{0, 1, 2} {
		data, _ := ioutil.ReadFile("input.txt")
		for _, row := range strings.Split(string(data), "\n") {
			if len(strings.Fields(row)) != 3 {
				continue
			}
			sides := convert(strings.Fields(row))
			sideLengths = append(sideLengths, sides[col])
		}
	}

	// Move through this vector, grabbing groups of 3 side lengths.
	validTriangles := 0
	for index := 0; index < len(sideLengths)/3; index++ {
		a := sideLengths[3*index]
		b := sideLengths[3*index+1]
		c := sideLengths[3*index+2]
		if a+b > c && a+c > b && b+c > a {
			validTriangles += 1
		}
	}

	fmt.Println("Part 2 Solution:", validTriangles)
}
