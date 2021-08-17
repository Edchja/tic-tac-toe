package main

import (
	"edchja.de/tic-tac-toe/fonts"
	"edchja.de/tic-tac-toe/game"
)

var gameBoardArr = make([][]int, 3)
var playerArr = []int{1, 2}
var player = playerArr[0]
var computer = playerArr[1]

var gameState bool

func main() {
	fillArray()

	fonts.PrintBoard(gameBoardArr)

	for gameState != true {
		input := game.SetInput()

		game.ProcessInput(gameBoardArr, input, player)

		// computerInput := getRandomNumber(1, 9)
		// processInput(computerInput, computer)
		game.ProcessComputerMove(gameBoardArr, player)

		for _, player := range playerArr {
			gameState = game.Condition(gameBoardArr, player)
		}

		fonts.PrintBoard(gameBoardArr)
		game.PrintWinningAnimation()
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
