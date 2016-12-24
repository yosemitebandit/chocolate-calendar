package main

import "testing"

type findMarkerTestInstance struct {
	input  string
	output marker
}

var findMarkerTestInstances = []findMarkerTestInstance{
	{"(32x4)ABCDE", marker{32, 4, 6}},
	{"(123x456)ABCDEFG", marker{123, 456, 9}},
	{"(3x4)ASDF", marker{3, 4, 5}},
	{"(3x4)(3x2)ASDF", marker{3, 4, 5}},
}

func TestFindMarker(t *testing.T) {
	for _, test := range findMarkerTestInstances {
		result := findMarker(test.input)
		if result.letters != test.output.letters ||
			result.repeats != test.output.repeats ||
			result.length != test.output.length {
			t.Error("Fail.")
		}
	}
}
