package main

import (
	"reflect"
	"testing"
)

type testInstance struct {
	input  string
	output []int
}

var rectParseCommandTests = []testInstance{
	{"rect 3x2", []int{3, 2}},
	{"rect 1x1", []int{1, 1}},
	{"rect 10x9", []int{10, 9}},
}

func TestRectCommandParsing(t *testing.T) {
	for _, test := range rectParseCommandTests {
		if !reflect.DeepEqual(parseRectCommand(test.input), test.output) {
			t.Error("Fail, input:", test.input, "Got:", parseRectCommand(test.input))
		}
	}
}

var rotateParsingTests = []testInstance{
	{"rotate column x=46 by 2", []int{46, 2}},
	{"rotate column x=1 by 3", []int{1, 3}},
	{"rotate row y=12 by 1", []int{12, 1}},
	{"rotate column x=18 by 17", []int{18, 17}},
}

func TestRotateCommandParsing(t *testing.T) {
	for _, test := range rotateParsingTests {
		if !reflect.DeepEqual(parseRotateCommand(test.input), test.output) {
			t.Error("Fail, input:", test.input, "Got:", parseRotateCommand(test.input))
		}
	}
}

type rowRotationTest struct {
	input  []string
	output []string
	amount int
}

var rowRotationTests = []rowRotationTest{
	{[]string{".", ".", ".", ".", ".", "#", "."}, []string{".", ".", ".", ".", ".", ".", "#"}, 1},
	{[]string{".", ".", "#", "."}, []string{".", "#", ".", "."}, 3},
}

func TestRowRotation(t *testing.T) {
	for _, test := range rowRotationTests {
		result := rotateRow(test.input, test.amount)
		if !reflect.DeepEqual(result, test.output) {
			t.Error("Rotation failed:", test.input, "Got:", result)
		}
	}
}

type colRotationTest struct {
	input     [][]string
	output    [][]string
	targetCol int
	amount    int
}

var colRotationTests = []colRotationTest{
	{[][]string{{".", "#", "."}, {".", ".", "."}, {".", ".", "."}},
		[][]string{{".", ".", "."}, {".", ".", "."}, {".", "#", "."}},
		1,
		2,
	},
	{[][]string{{".", ".", "."}, {"#", ".", "."}, {".", ".", "."}},
		[][]string{{".", ".", "."}, {".", ".", "."}, {"#", ".", "."}},
		0,
		4,
	},
}

func TestColRotation(t *testing.T) {
	for _, test := range colRotationTests {
		result := rotateCol(test.input, test.targetCol, test.amount)
		for resultRowIndex, resultRow := range result {
			if !reflect.DeepEqual(resultRow, test.output[resultRowIndex]) {
				t.Error("ColRotation failed:", test.input, "Got:", result)
				break
			}
		}
	}
}
