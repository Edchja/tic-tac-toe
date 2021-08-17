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

func SetInput() int {
	fmt.Println("\u001b[37mEnter a number: \u001b[0m")

	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\r\n", "", -1)

	inputInt, _ := strconv.Atoi(input)

	return inputInt
}

func ProcessInput(gameBoardArr [][]int, input, player int) {
	if input >= 1 && input <= 9 {
		switch input {
		case 1:
			ProcessMove(gameBoardArr, 2, 0, player)
		case 2:
			ProcessMove(gameBoardArr, 2, 1, player)
		case 3:
			ProcessMove(gameBoardArr, 2, 2, player)
		case 4:
			ProcessMove(gameBoardArr, 1, 0, player)
		case 5:
			ProcessMove(gameBoardArr, 1, 1, player)
		case 6:
			ProcessMove(gameBoardArr, 1, 2, player)
		case 7:
			ProcessMove(gameBoardArr, 0, 0, player)
		case 8:
			ProcessMove(gameBoardArr, 0, 1, player)
		case 9:
			ProcessMove(gameBoardArr, 0, 2, player)
		}
	} else if player == 1 {
		fonts.PrintErrorMessage("wrong input!")

		fonts.PrintBoard(gameBoardArr)

		playerInput := SetInput()
		ProcessInput(gameBoardArr, playerInput, player)
	} else if player == 2 {
		// computer := getRandomNumber(1, 9)
		// processInput(computer, player)
	}
}

func ProcessMove(gameBoardArr [][]int, row, col, player int) {
	switch {
	case gameBoardArr[row][col] == 0:
		gameBoardArr[row][col] = player

	case gameBoardArr[row][col] != 0 && player != 2:
		fonts.PrintErrorMessage("field already set!")

		fonts.PrintBoard(gameBoardArr)

		playerInput := SetInput()
		ProcessInput(gameBoardArr, playerInput, player)

	case gameBoardArr[row][col] != 0 && player != 1 && Condition(gameBoardArr, player) != true:
		computer := GetRandomNumber(1, 9)
		ProcessInput(gameBoardArr, computer, player)
	}
}

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
