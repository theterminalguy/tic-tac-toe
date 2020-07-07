package main

import (
	"fmt"
	"strings"
)

func main() {
	board := [][]string {
		{"0", "1", "2"},
		{"3", "4", "5"},
		{"6", "7", "8"},
	}
	wins := [][]uint {
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}
	choices := []uint {0, 1, 2, 3, 4, 5, 6, 7, 8}

	userMarker, computerMarker := assignMarkers()

  printInstructions()

}

func assignMarkers() (string, string) {
	marker := map[string]string {
		"X": "O",
		"O": "X",
	}

	var (
		userMarker
		computerMarker
	)

	fmt.Println("Please choose a marker.", "X or O ?")
	for true {
		fmt.Scanf("%s", &userMarker)
		userMarker = strings.Title(userMarker)

		if marker, ok := marker[userMarker]; ok {
			computerMarker = marker

			break
		} else {
			fmt.Println("Invalid marker. Please choose either X or O!")
		}
	}

	fmt.Printf("Your marker is: %s\n", userMarker)
  fmt.Printf("The computer's marker is: %s\n", computerMarker)

	return userMarker, computerMarker
}

func printInstructions() {
  fmt.Println("Initializing game board.....")
  fmt.Println("*****************************")
  fmt.Println("You can mark a location by entering any of the numbers shown on the board.")
}
