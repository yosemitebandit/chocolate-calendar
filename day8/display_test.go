package main

import (
	"reflect"
	"testing"
)

type rectParseTestInstance struct {
	input  string
	output []int
}

var rectParseCommandTests = []rectParseTestInstance{
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
