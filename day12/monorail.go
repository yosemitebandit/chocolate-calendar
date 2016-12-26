package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("input.txt")

	registers := make(map[string]int)
	registers["a"] = 0
	registers["b"] = 0
	registers["c"] = 1
	registers["d"] = 0

	lines := strings.Split(string(data), "\n")
	lineIndex := 0
	for {
		if lineIndex >= len(lines) {
			break
		}
		line := lines[lineIndex]
		if line == "" {
			lineIndex++
			continue
		}

		values := strings.Fields(line)
		switch values[0] {
		case "cpy":
			// See if we're copying a register or a value.
			value, err := strconv.Atoi(values[1])
			if err != nil {
				registers[values[2]] = registers[values[1]]
			} else {
				registers[values[2]] = value
			}
		case "jnz":
			// Check if the value is zero.
			value := 0
			value, err := strconv.Atoi(values[1])
			if err != nil {
				value = registers[values[1]]
			}
			if value == 0 {
				lineIndex++
				continue
			} else {
				jumpValue, _ := strconv.Atoi(values[2])
				lineIndex += jumpValue
				continue
			}
		case "inc":
			registers[values[1]]++
		case "dec":
			registers[values[1]]--
		}

		lineIndex++
	}
	fmt.Println("Part 2 Solution:", registers)
}
