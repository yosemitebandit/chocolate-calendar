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

type decompressTestInstance struct {
	input        string
	outputLength int
}

var decompressTestInstances = []decompressTestInstance{
	{"(3x3)XYZ", 9},
	{"X(8x2)(3x3)ABCY", 20},
	{"(27x12)(20x12)(13x14)(7x10)(1x12)A", 241920},
	{"(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN", 445},
}

func TestDecompress(t *testing.T) {
	for _, test := range decompressTestInstances {
		result := decompress(test.input, 1)
		if len(result) != test.outputLength {
			t.Error("decompress fail, input:", test.input, "Got:", result, "Len:", len(result))
		}
	}
}
