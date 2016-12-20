package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func parseRectCommand(command string) []int {
	elements := strings.Split(command, "rect ")
	values := strings.Split(elements[1], "x")
	rows, _ := strconv.Atoi(values[0])
	cols, _ := strconv.Atoi(values[1])
	result := []int{rows, cols}
	return result
}

func main() {
	// Init the display.
	const displayRows int = 6
	const displayCols int = 50
	var display [displayRows][displayCols]string
	for row := 0; row < displayRows; row++ {
		for col := 0; col < displayCols; col++ {
			display[row][col] = "."
		}
	}

	// Work through the instructions.
	data, _ := ioutil.ReadFile("input.txt")
	for command_index, command := range strings.Split(string(data), "\n") {
		if command == "" {
			continue
		}
		if command[0:4] == "rect" {
			dims := parseRectCommand(command)
			for row := 0; row < dims[1]; row++ {
				for col := 0; col < dims[0]; col++ {
					display[row][col] = "#"
				}
			}
			// Plot it.
			var output []string
			for i := 0; i < displayRows; i++ {
				output = append(output, strings.Join(display[i][:], ""))
				output = append(output, "\n")
			}
			fmt.Printf("\033[0;0H")
			fmt.Println(strings.Join(output, ""))
			fmt.Println(command_index, dims)
			time.Sleep(time.Second)
		}
	}

}
