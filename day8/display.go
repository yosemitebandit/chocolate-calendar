package main

import (
	"fmt"
	"io/ioutil"
	"math"
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

func parseRotateCommand(command string) []int {
	elements := strings.Fields(command)
	field, _ := strconv.Atoi(strings.Split(elements[2], "=")[1])
	amount, _ := strconv.Atoi(elements[4])
	return []int{field, amount}
}

func rotateRow(row []string, amount int) []string {
	newRow := make([]string, len(row))
	for index, value := range row {
		newRow[int(math.Mod(float64(index+amount), float64(len(row))))] = value
	}
	return newRow
}

func rotateCol(matrix [][]string, targetCol int, amount int) [][]string {
	var elements []string
	for _, row := range matrix {
		elements = append(elements, row[targetCol])
	}
	result := rotateRow(elements, amount)
	// Make a copy of the matrix.
	var matrixCopy [][]string
	for row := 0; row < len(matrix); row++ {
		matrixCopy = append(matrixCopy, []string{})
		for col := 0; col < len(matrix[0]); col++ {
			if col == targetCol {
				matrixCopy[row] = append(matrixCopy[row], result[row])
			} else {
				matrixCopy[row] = append(matrixCopy[row], ".")
			}
		}
	}
	return matrixCopy
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
	for commandIndex, command := range strings.Split(string(data), "\n") {
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
		} else if strings.Contains(command, "rotate row") {
			result := parseRotateCommand(command)
			targetRow := result[0]
			amount := result[1]
			originalRow := display[targetRow]
			newRow := rotateRow(originalRow[0:], amount)
			// Convert to an array.
			var updatedRow [50]string
			for index, value := range newRow {
				updatedRow[index] = value
			}
			display[targetRow] = updatedRow
		} else if strings.Contains(command, "rotate column") {
			fmt.Println("col!")
		}

		// Plot it.
		var output []string
		for i := 0; i < displayRows; i++ {
			output = append(output, strings.Join(display[i][:], ""))
			output = append(output, "\n")
		}
		fmt.Printf("\033[0;0H")
		fmt.Println(strings.Join(output, ""))
		time.Sleep(time.Second / 32)
		fmt.Println("commandIndex:", commandIndex)

	}

	// Count the 'on' pixels.
	activatedPixels := 0
	for row := 0; row < displayRows; row++ {
		for col := 0; col < displayCols; col++ {
			if display[row][col] == "#" {
				activatedPixels++
			}
		}
	}
	fmt.Println("Activated Pixels:", activatedPixels)

}
