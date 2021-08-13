package fonts

import (
	"fmt"
	"time"

	"edchja.de/tic-tac-toe/game"
)

const (
	refreshRate = 100 * time.Millisecond
)

var (
	XWonFont = `
██╗  ██╗    ██╗    ██╗ ██████╗ ███╗   ██╗
╚██╗██╔╝    ██║    ██║██╔═══██╗████╗  ██║
 ╚███╔╝     ██║ █╗ ██║██║   ██║██╔██╗ ██║
 ██╔██╗     ██║███╗██║██║   ██║██║╚██╗██║
██╔╝ ██╗    ╚███╔███╔╝╚██████╔╝██║ ╚████║
╚═╝  ╚═╝     ╚══╝╚══╝  ╚═════╝ ╚═╝  ╚═══╝    `

	OWonFont = `
 ██████╗     ██╗    ██╗ ██████╗ ███╗   ██╗
██╔═══██╗    ██║    ██║██╔═══██╗████╗  ██║
██║   ██║    ██║ █╗ ██║██║   ██║██╔██╗ ██║
██║   ██║    ██║███╗██║██║   ██║██║╚██╗██║
╚██████╔╝    ╚███╔███╔╝╚██████╔╝██║ ╚████║
 ╚═════╝      ╚══╝╚══╝  ╚═════╝ ╚═╝  ╚═══╝   `

	TieFont = `
████████╗██╗███████╗
╚══██╔══╝██║██╔════╝
   ██║   ██║█████╗
   ██║   ██║██╔══╝
   ██║   ██║███████╗
   ╚═╝   ╚═╝╚══════╝`
)

func PrintFont(arr [][]rune, font string, size int) {
	row := 0
	col := 0
	arr[0] = make([]rune, size)
	for _, r := range font {
		if r == '\n' {
			row++
			arr[row] = make([]rune, size)
			col = 0
			continue
		}
		arr[row][col] = r
		col++
	}

	go func() {
		first := true
		for {
			if !first {
				for i := 0; i < len(arr); i++ {
					up()
					clearLine()
				}
			}
			for _, line := range arr {
				fmt.Println("\u001b[34;1m", string(line), "\u001b[0m")
			}
			first = false
			time.Sleep(refreshRate)
		}
	}()

	go flow(arr)

	for {
	}
}

func flow(arr [][]rune) {
	for {
		for row := range arr {
			tempLastIndex := arr[row][len(arr[row])-1]
			for col := len(arr[row]) - 1; col > 0; col-- {
				arr[row][col] = arr[row][col-1]
			}
			arr[row][0] = tempLastIndex
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func PrintWinningAnimation(fontArr [][]rune) {
	switch {
	case game.XWon:
		PrintFont(fontArr, XWonFont, 45)
	case game.OWon:
		PrintFont(fontArr, OWonFont, 45)
	case game.Tie:
		PrintFont(fontArr, TieFont, 25)
	}
}

func PrintErrorMessage(text string) {
	fmt.Printf("\u001b[31;1m%s\u001b[0m\n", text)
}

// Bewegt den Cursor eine Zeile nach oben.
func up() {
	fmt.Print("\033[A")
}

// Bewegt den Cursor eine Zeile runter.
func down() {
	fmt.Print("\033[B")
}

// Löscht die aktuelle Zeile.
func clearLine() {
	fmt.Print("\033[G\033[K")
}
