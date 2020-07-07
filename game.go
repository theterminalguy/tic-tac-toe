package main

import "fmt"

func main() {
	InvertMarker := map[string]string {
		"X": "O",
		"O": "X",
	}
	Board := [][]string {
		{"0", "1", "2"},
		{"3", "4", "5"},
		{"6", "7", "8"},
	}
	Wins := [][]uint {
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}
	Choices := []uint {0, 1, 2, 3, 4, 5, 6, 7, 8}

	
}