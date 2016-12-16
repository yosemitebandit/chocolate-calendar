package main

import "testing"

func TestGenerateChecksum(t *testing.T) {
	room := "aaaaabbbzyx"
	if generateChecksum(room) != "abxyz" {
		t.Error("fail")
	}
	room = "abcdefgh"
	if generateChecksum(room) != "abcde" {
		t.Error("fail")
	}
}
