package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"edchja.de/tic-tac-toe/fonts"
)

const (
	player1 = 1
	player2 = 2
)

var gameBoard = make([][]int, 3)
var fontArr = make([][]rune, 7)
var gameBoardArr = make([][]rune, 6)
var drawValues = []string{" ", "X", "O"}

var reader = bufio.NewReader(os.Stdin)
var xWon, oWon, tie bool

// TODO:	Cursor resetting einfügen
// TODO:	Refactor to structs
// TODO:	Refactor to DRY
// TODO:	AI never loses

func main() {
	var gameState bool

	fillArray()

	printBoard()

	for gameState != true {
		input := setInput()

		processInput(input, player1)

		computerInput := getRandomNumber(1, 9)
		processInput(computerInput, player2)

		gameState = gamelogic()

		printBoard()
		printWinningScreen()
	}
}

func gamelogic() bool {
	var (
		// Check all rows.
		x = (gameBoard[0][0] == 1 && gameBoard[0][1] == 1 && gameBoard[0][2] == 1 ||
			gameBoard[1][0] == 1 && gameBoard[1][1] == 1 && gameBoard[1][2] == 1 ||
			gameBoard[2][0] == 1 && gameBoard[2][1] == 1 && gameBoard[2][2] == 1 ||

			// Check all columns.
			gameBoard[0][0] == 1 && gameBoard[1][0] == 1 && gameBoard[2][0] == 1 ||
			gameBoard[0][1] == 1 && gameBoard[1][1] == 1 && gameBoard[2][1] == 1 ||
			gameBoard[0][2] == 1 && gameBoard[1][2] == 1 && gameBoard[2][2] == 1 ||

			// Check diaGonals.
			gameBoard[0][0] == 1 && gameBoard[1][1] == 1 && gameBoard[2][2] == 1 ||
			gameBoard[0][2] == 1 && gameBoard[1][1] == 1 && gameBoard[2][0] == 1)

		o = (gameBoard[0][0] == 2 && gameBoard[0][1] == 2 && gameBoard[0][2] == 2 ||
			gameBoard[1][0] == 2 && gameBoard[1][1] == 2 && gameBoard[1][2] == 2 ||
			gameBoard[2][0] == 2 && gameBoard[2][1] == 2 && gameBoard[2][2] == 2 ||

			// Check all columns.
			gameBoard[0][0] == 2 && gameBoard[1][0] == 2 && gameBoard[2][0] == 2 ||
			gameBoard[0][1] == 2 && gameBoard[1][1] == 2 && gameBoard[2][1] == 2 ||
			gameBoard[0][2] == 2 && gameBoard[1][2] == 2 && gameBoard[2][2] == 2 ||

			// Check diaGonals.
			gameBoard[0][0] == 2 && gameBoard[1][1] == 2 && gameBoard[2][2] == 2 ||
			gameBoard[0][2] == 2 && gameBoard[1][1] == 2 && gameBoard[2][0] == 2)

		// Check if cells are empty.
		freeCellsLeft = (gameBoard[0][0] == 0 || gameBoard[0][1] == 0 || gameBoard[0][2] == 0 ||
			gameBoard[1][0] == 0 || gameBoard[1][1] == 0 || gameBoard[1][2] == 0 ||
			gameBoard[2][0] == 0 || gameBoard[2][1] == 0 || gameBoard[2][2] == 0)
	)

	switch {
	case x && !o:
		xWon = true
		return xWon
	case o && !x:
		oWon = true
		return oWon
	case !freeCellsLeft:
		tie = true
		return tie
	}
	return false
}

func processInput(input int, player int) {
	if input >= 1 && input <= 9 {
		switch input {
		case 1:
			checkWinner(2, 0, player)
		case 2:
			checkWinner(2, 1, player)
		case 3:
			checkWinner(2, 2, player)
		case 4:
			checkWinner(1, 0, player)
		case 5:
			checkWinner(1, 1, player)
		case 6:
			checkWinner(1, 2, player)
		case 7:
			checkWinner(0, 0, player)
		case 8:
			checkWinner(0, 1, player)
		case 9:
			checkWinner(0, 2, player)
		}
	} else if player == 1 {
		fmt.Println("wrong input!")

		playerInput := setInput()
		processInput(playerInput, player1)
	} else if player == 2 {
		computer := getRandomNumber(1, 9)
		processInput(computer, player2)
	}
}

func fillArray() {
	// propagates array with 0.
	for row := range gameBoard {
		gameBoard[row] = make([]int, 3)
		for col := range gameBoard[row] {
			gameBoard[row][col] = 0
		}
	}
}

func checkWinner(row, col, player int) {
	switch {
	case gameBoard[row][col] == 0:
		gameBoard[row][col] = player

	case gameBoard[row][col] != 0 && player != 2:
		fmt.Println("field already set")

		playerInput := setInput()
		processInput(playerInput, player1)

	case gameBoard[row][col] != 0 && player != 1 && gamelogic() == false:
		computer := getRandomNumber(1, 9)
		// fmt.Println("{computer: if already set: ", computer, "}")
		processInput(computer, player2)
	}
}

func setInput() int {
	fmt.Println("Enter a number: ")

	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\r\n", "", -1)

	inputInt, _ := strconv.Atoi(input)

	return inputInt
}

func printBoard() {
	for row := 0; row < len(gameBoard); row++ {
		for col := range gameBoard[row] {
			fmt.Print("| ", drawValues[gameBoard[row][col]], " ")
		}
		fmt.Println("|")
		if row != len(gameBoard)-1 {
			fmt.Println("-------------")
		}
	}
}

func printWinningScreen() {
	switch {
	case xWon:
		fonts.PrintFont(fontArr, fonts.XWonFont, 45)
	case oWon:
		fonts.PrintFont(fontArr, fonts.OWonFont, 45)
	case tie:
		fonts.PrintFont(fontArr, fonts.TieFont, 25)
	}
}

func getRandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn((max - min + 1) + min)

	if randomNum != 0 {
		return randomNum
	} else {
		return getRandomNumber(min, max)
	}
}
