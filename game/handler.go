package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"edchja.de/tic-tac-toe/fonts"
)

var (
	XWon, OWon, Tie bool
	reader          = bufio.NewReader(os.Stdin)
)

const (
	player = 1
)

func ProcessPlayerMove(gameBoardArr [][]int, input int) {
	if input >= 1 && input <= 9 {
		row, col := getPostition(input)
		if gameBoardArr[row][col] == 0 {
			gameBoardArr[row][col] = player
		} else {
			fonts.PrintErrorMessage("field already set!")
			fonts.PrintBoard(gameBoardArr)

			playerInput := SetInput()
			ProcessPlayerMove(gameBoardArr, playerInput)
		}
	} else {
		fonts.PrintErrorMessage("wrong input!")

		fonts.PrintBoard(gameBoardArr)

		playerInput := SetInput()
		ProcessPlayerMove(gameBoardArr, playerInput)
	}
}

func SetInput() int {
	fmt.Println("\u001b[37mEnter a number: \u001b[0m")

	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\r\n", "", -1)

	inputInt, _ := strconv.Atoi(input)

	return inputInt
}

func getPostition(playerInput int) (int, int) {
	switch playerInput {
	case 1:
		return 2, 0
	case 2:
		return 2, 1
	case 3:
		return 2, 2
	case 4:
		return 1, 0
	case 5:
		return 1, 1
	case 6:
		return 1, 2
	case 7:
		return 0, 0
	case 8:
		return 0, 1
	case 9:
		return 0, 2
	}
	return -1, -1
}

func CheckCondition(gameBoardArr [][]int, player int) bool {
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

func PrintWinningAnimation() {
	switch {
	case XWon:
		fonts.PrintFont(fonts.XWonFont, 45)
	case OWon:
		fonts.PrintFont(fonts.OWonFont, 45)
	case Tie:
		fonts.PrintFont(fonts.TieFont, 25)
	}
}
