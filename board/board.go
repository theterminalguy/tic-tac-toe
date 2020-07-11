package board

func Draw(board [][]string) string {
	displayBoard := ""

	for _, row := range board {
		displayBoard += " "
		displayBoard += strings.Join(row, " | ")
		displayBoard += "\n-----------\n"
	}

	return displayBoard
}

func Update(board [][]string, pos uint, marker string) [][]string{
	row := pos / 3
	column := pos % 3

	// TODO: might want to treturn the new board
	board[row][column] = marker

	return board
}
