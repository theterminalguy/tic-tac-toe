package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// Information => Green
	// Instruction =>
	// colorRed := "\033[31m"
	// colorGreen := "\033[32m"

	board := [][]string{
		{"0", "1", "2"},
		{"3", "4", "5"},
		{"6", "7", "8"},
	}
	choices := []uint{0, 1, 2, 3, 4, 5, 6, 7, 8}

	var (
		userMarkedSpots     []uint
		computerMarkedSpots []uint
		markedSpots         []uint
		unMarkedSpots       = []uint{0, 1, 2, 3, 4, 5, 6, 7, 8}
	)

	userMarker, computerMarker := assignMarkers()

	printInstructions()

	drawBoard(board)

	// start game

	for true {
		if gameOver(markedSpots) {
			fmt.Println("Game Over, it's a tie")

			break
		}

		fmt.Printf("Your turn(%s), pick a spot\n", userMarker)
		userChoice := awaitUserChoice(unMarkedSpots)

		updateBoard(board, userChoice, userMarker)
		drawBoard(board)

		userMarkedSpots = append(userMarkedSpots, userChoice)
		if hasWon(userMarkedSpots) {
			fmt.Printf("User(%s) won!\n", userMarker)

			break
		}

		// computer turn

		markedSpots = append(userMarkedSpots, computerMarkedSpots...)
		unMarkedSpots = diff(choices, markedSpots)

		fmt.Println("Computer is thinking...")
		computerChoice := awaitComputerChoice(unMarkedSpots)

		updateBoard(board, computerChoice, computerMarker)
		drawBoard(board)

		computerMarkedSpots = append(computerMarkedSpots, computerChoice)
		if hasWon(computerMarkedSpots) {
			fmt.Printf("Computer(%s) won!\n", computerMarker)

			break
		}

		markedSpots = append(userMarkedSpots, computerMarkedSpots...)
		unMarkedSpots = diff(choices, markedSpots)

	}
}

func hasWon(markedSpots []uint) bool {
	const (
		SPOTS_REQUIRED_TO_WIN = 3
	)
	wins := [][]uint{
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
		for _, win := range wins {
			if isSubset(win, markedSpots) {
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

func awaitUserChoice(unMarkedSpots []uint) uint {
	var choice uint
	for true {
		fmt.Scanf("%d", &choice)

		if Contains(unMarkedSpots, choice) {
			break
		}

		fmt.Println("Invalid choice. Please pick an available spot.")
	}

	return choice
}

func awaitComputerChoice(unmarkedSpots []uint) uint {
	time.Sleep(1000 * time.Millisecond) // sleep for 2 seconds
	rand.Seed(time.Now().Unix())
	randomPosition := rand.Intn(len(unmarkedSpots))

	return unmarkedSpots[randomPosition]
}

func assignMarkers() (string, string) {
	marker := map[string]string{
		"X": "O",
		"O": "X",
	}

	var (
		userMarker     string
		computerMarker string
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
	fmt.Println("\n")
	fmt.Println("*****************************")
	fmt.Println("Initializing game board.....")
	fmt.Println("*****************************")
	fmt.Println("\n")
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

func updateBoard(board [][]string, pos uint, marker string) {
	r := row(pos)
	c := column(pos)

	board[r][c] = marker
}

func diff(superSet []uint, subset []uint) []uint {
	var res []uint

	for _, choice := range superSet {
		if !Contains(subset, choice) {
			res = append(res, choice)
		}
	}

	return res
}

func isSubset(subset []uint, superSet []uint) bool {
	// TOOD: Use contains

	foundCount := 0

	for _, s1 := range subset {
		for _, s2 := range superSet {
			if s1 == s2 {
				foundCount += 1

				continue
			}
		}
	}

	return foundCount >= len(subset)
}

func Contains(set []uint, element uint) bool {
	// TODO rename to memeber

	for _, n := range set {
		if element == n {
			return true
		}
	}

	return false
}

func row(pos uint) uint {
	return pos / 3
}

func column(pos uint) uint {
	return pos % 3
}
