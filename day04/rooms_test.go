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

func TestRotate(t *testing.T) {
	if rotateName("a", 3) != "d" {
		t.Error("fail!")
	}
	if rotateName("z", 1) != "a" {
		t.Error("fail!")
	}
	if rotateName("ab-c", 2) != "cd e" {
		t.Error("fail!")
	}
}
