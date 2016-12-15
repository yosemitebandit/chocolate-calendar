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
	data, _ := ioutil.ReadFile("input.txt")
	validTriangles := 0

	for _, triangle := range strings.Split(string(data), "\n") {
		if len(strings.Fields(triangle)) != 3 {
			continue
		}
		sides := convert(strings.Fields(triangle))
		if sides[0]+sides[1] > sides[2] && sides[0]+sides[2] > sides[1] && sides[1]+sides[2] > sides[0] {
			validTriangles += 1
		}
	}

	fmt.Println("Part 1 Solution:", validTriangles)
}
