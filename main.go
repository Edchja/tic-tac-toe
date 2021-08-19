package main

import (
	"edchja.de/tic-tac-toe/fonts"
	"edchja.de/tic-tac-toe/game"
)

var gameBoardArr = make([][]int, 3)
var playerArr = []int{1, 2}

var gameState bool

func main() {
	fillArray()

	fonts.PrintBoard(gameBoardArr)

	for gameState != true {
		input := game.SetInput()

		game.ProcessPlayerMove(gameBoardArr, input)
		game.ProcessComputerMove(gameBoardArr, input)

		for _, player := range playerArr {
			gameState = game.CheckCondition(gameBoardArr, player)
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
