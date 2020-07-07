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
	choices := []uint {0, 1, 2, 3, 4, 5, 6, 7, 8}

	var (
		userMarkedSpots = uint[]
		computerMarkedSpots = uint[]
	)

	userMarker, computerMarker := assignMarkers()

  printInstructions()

	drawBoard() 

	// start game

	for true {
		markedSpots = append(userMarkedSpots, computerMarkedSpots...)
		unMarkedSpots = diff(choices, markedSpots)

		if gameOver(markedSpots) {
			fmt.Println("Game Over, it's a tie")

			break
		}

		fmt.Printf("Your turn(%s), pick a spot\n", userMarker)
		userChoice = awaitUserChoice()

		userMarkedSpots = append(userMarkedSpots, userChoice)
		updateBoard(board, userChoice, userMarker)
		drawBoard(board)

		if hasWon(userMarkedSpots) {
			fmt.Printf ("User(%s) won!\n", userMarker)
		}

	}
}

func hasWon(markedSpots uint) bool {
	const (
		SPOTS_REQUIRED_TO_WIN = 3
	)
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

	if len(markedSpots) >= SPOTS_REQUIRED_TO_WIN {
		for _, win := wins{
			if subset(win, markedSpots) {
				return true 
			}
		}
	}

	return false
}

func gameOver(markedSpots []uint) bool {
	const (
		AVAILABLE_SPOTS = 9
	)

	return len(markedSpots) == AVAILABLE_SPOTS
}

func awaitUserChoice() string {
	var choice uint
	for true {
		fmt.Scanf("%d", &choice)
		
		if(Contains(unMarkedSpots, choice)) {
			break
		} 

		fmt.Println("Invalid choice. Please pick an available spot.")
	}

	return choice
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
		} 
		
		fmt.Println("Invalid marker. Please choose either X or O!")
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

// should be in the board package
func drawBoard(board [][]string) {
  displayBoard := ""

  for _, row := range board {
    displayBoard += " "
    displayBoard += strings.Join(row, " | ")
    displayBoard += "\n-----------\n"
  }

  fmt.Println(displayBoard)
}
