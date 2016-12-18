package main

import (
	"reflect"
	"testing"
)

type sequence struct {
	text   string
	isAbba bool
}

var abbaTests = []sequence{
	{"asdfabbaxyz", true},
	{"abbaxyz", true},
	{"ababxyz", false},
	{"aaaaxyz", false},
}

func TestSequenceDetection(t *testing.T) {
	for _, test := range abbaTests {
		if abbaSequenceIsPresent(test.text) != test.isAbba {
			t.Error("fail on", test.text)
		}
	}
}

type bracketTestInstance struct {
	input  string
	output dividedIP
}

var bracketTests = []bracketTestInstance{
	{"abc[def]", dividedIP{[]string{"abc"}, []string{"def"}}},
	{"abc[def]ghi", dividedIP{[]string{"abc", "ghi"}, []string{"def"}}},
	{"abc[def]ghi[lmn]", dividedIP{[]string{"abc", "ghi"}, []string{"def", "lmn"}}},
	{"abc[def]ghi[lmn]opq", dividedIP{[]string{"abc", "ghi", "opq"}, []string{"def", "lmn"}}},
}

func TestBracketSplit(t *testing.T) {
	for _, test := range bracketTests {
		t.Log(splitIP(test.input))
		if !reflect.DeepEqual(splitIP(test.input).vanilla, test.output.vanilla) {
			t.Error("fail:", test.input)
		}
		if !reflect.DeepEqual(splitIP(test.input).bracketed, test.output.bracketed) {
			t.Error("fail:", test.input)
		}
	}
}

// Part Two.

type abaTestInstance struct {
	input  string
	output []string
}

var abaTests = []abaTestInstance{
	{"abc", []string{}},
	{"aba", []string{"aba"}},
	{"aaa", []string{}},
	{"asdfxyx", []string{"xyx"}},
	{"asdfxyxioi", []string{"xyx", "ioi"}},
	{"asdfxyxioixyzmnm", []string{"xyx", "ioi", "mnm"}},
	{"abab", []string{"aba", "bab"}},
}

func TestAbaSplit(t *testing.T) {
	for _, test := range abaTests {
		// Urgh, nil slice vs empty slice is different according to DeepEqual.
		if len(test.output) == 0 && len(splitAba(test.input)) != 0 {
			t.Error("Failing on:", test.input, "Got:", splitAba(test.input))
		}
		if len(test.output) > 0 && !reflect.DeepEqual(splitAba(test.input), test.output) {
			t.Error("Failing on:", test.input, "Got:", splitAba(test.input))
		}
	}
}
