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
	"edchja.de/tic-tac-toe/game"
)

var gameBoardArr = make([][]int, 3)
var fontArr = make([][]rune, 7)
var drawValues = []string{" ", "X", "O"}
var playerArr = []int{1, 2}
var player = playerArr[0]
var computer = playerArr[1]

var gameState bool
var reader = bufio.NewReader(os.Stdin)

// TODO:	Cursor resetting einfügen
// TODO:	AI never loses

func main() {
	fillArray()

	printBoard()

	for gameState != true {
		input := setInput()

		processInput(input, player)

		computerInput := getRandomNumber(1, 9)
		processInput(computerInput, computer)

		for _, player := range playerArr {
			gameState = game.Condition(gameBoardArr, player)
		}

		printBoard()
		fonts.PrintWinningAnimation(fontArr)
	}
}

func processInput(input, player int) {
	if input >= 1 && input <= 9 {
		switch input {
		case 1:
			processMove(2, 0, player)
		case 2:
			processMove(2, 1, player)
		case 3:
			processMove(2, 2, player)
		case 4:
			processMove(1, 0, player)
		case 5:
			processMove(1, 1, player)
		case 6:
			processMove(1, 2, player)
		case 7:
			processMove(0, 0, player)
		case 8:
			processMove(0, 1, player)
		case 9:
			processMove(0, 2, player)
		}
	} else if player == 1 {
		fmt.Println("\u001b[31mwrong input!\u001b[0m")

		playerInput := setInput()
		processInput(playerInput, player)
	} else if player == 2 {
		computer := getRandomNumber(1, 9)
		processInput(computer, player)
	}
}

func processMove(row, col, player int) {
	switch {
	case gameBoardArr[row][col] == 0:
		gameBoardArr[row][col] = player

	case gameBoardArr[row][col] != 0 && player != 2:
		fmt.Println("\u001b[31mfield already set!\u001b[0m")

		playerInput := setInput()
		processInput(playerInput, player)

	case gameBoardArr[row][col] != 0 && player != 1 && game.Condition(gameBoardArr, player) != true:
		computer := getRandomNumber(1, 9)
		processInput(computer, player)
	}
}

func fillArray() {
	for row := range gameBoardArr {
		gameBoardArr[row] = make([]int, 3)
		for col := range gameBoardArr[row] {
			gameBoardArr[row][col] = 0
		}
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
	for row := 0; row < len(gameBoardArr); row++ {
		for col := range gameBoardArr[row] {
			fmt.Print("| ", drawValues[gameBoardArr[row][col]], " ")
		}
		fmt.Println("|")
		if row != len(gameBoardArr)-1 {
			fmt.Println("-------------")
		}
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
