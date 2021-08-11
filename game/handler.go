package game

var (
	XWon, OWon, Tie bool
)

func Condition(gameBoardArr [][]int, player int) bool {
	var (
		// Check all rows.
		state = (gameBoardArr[0][0] == player && gameBoardArr[0][1] == player && gameBoardArr[0][2] == player ||
			gameBoardArr[1][0] == player && gameBoardArr[1][1] == player && gameBoardArr[1][2] == player ||
			gameBoardArr[2][0] == player && gameBoardArr[2][1] == player && gameBoardArr[2][2] == player ||

			// Check all columns.
			gameBoardArr[0][0] == player && gameBoardArr[1][0] == player && gameBoardArr[2][0] == player ||
			gameBoardArr[0][1] == player && gameBoardArr[1][1] == player && gameBoardArr[2][1] == player ||
			gameBoardArr[0][2] == player && gameBoardArr[1][2] == player && gameBoardArr[2][2] == player ||

			// Check diagonals.
			gameBoardArr[0][0] == player && gameBoardArr[1][1] == player && gameBoardArr[2][2] == player ||
			gameBoardArr[0][2] == player && gameBoardArr[1][1] == player && gameBoardArr[2][0] == player)

		// Check if cells are empty.
		emptyCellsLeft = (gameBoardArr[0][0] == 0 || gameBoardArr[0][1] == 0 || gameBoardArr[0][2] == 0 ||
			gameBoardArr[1][0] == 0 || gameBoardArr[1][1] == 0 || gameBoardArr[1][2] == 0 ||
			gameBoardArr[2][0] == 0 || gameBoardArr[2][1] == 0 || gameBoardArr[2][2] == 0)
	)

	switch {
	case state && player != 2:
		XWon = true
		return XWon
	case state && player != 1:
		OWon = true
		return OWon
	case !emptyCellsLeft:
		Tie = true
		return Tie
	}

	return false
}
