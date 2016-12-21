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

var rotateTests = []testInstance{
	{"rotate column x=46 by 2", []int{46, 2}},
	{"rotate column x=1 by 3", []int{1, 3}},
	{"rotate row y=12 by 1", []int{12, 1}},
	{"rotate column x=18 by 17", []int{18, 17}},
}

func TestRotateCommand(t *testing.T) {
	for _, test := range rotateTests {
		if !reflect.DeepEqual(parseRotateCommand(test.input), test.output) {
			t.Error("Fail, input:", test.input, "Got:", parseRotateCommand(test.input))
		}
	}
}
