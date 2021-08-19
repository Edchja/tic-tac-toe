package game

const (
	computer = 2
)

func ProcessComputerMove(gameBoardArr [][]int, input int) {
	row, col := getPostition(input)

	if row != -1 && col != -1 {
		for i, c := range gameBoardArr[row] {
			if i != col && c == 0 {
				gameBoardArr[row][i] = computer
				return
			}
		}
		for i, r := range gameBoardArr {
			if i != row && r[i] == 0 {
				gameBoardArr[i][col] = computer
				return
			}
		}
		for row := range gameBoardArr {
			for col := range gameBoardArr[row] {
				if gameBoardArr[row][col] == 0 {
					gameBoardArr[row][col] = computer
					return
				}
			}
		}
	} else {
		for row := range gameBoardArr {
			for col := range gameBoardArr[row] {
				if gameBoardArr[row][col] == 0 {
					gameBoardArr[row][col] = computer
					return
				}
			}
		}
	}
}
