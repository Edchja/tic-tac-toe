package fonts

import (
	"fmt"
	"time"
)

const (
	refreshRate = 100 * time.Millisecond
)

var fontArr = make([][]rune, 7)
var drawValues = []string{" ", "X", "O"}

func PrintFont(font string, size int) {
	row := 0
	col := 0
	fontArr[0] = make([]rune, size)
	for _, r := range font {
		if r == '\n' {
			row++
			fontArr[row] = make([]rune, size)
			col = 0
			continue
		}
		fontArr[row][col] = r
		col++
	}

	go func() {
		first := true
		for {
			if !first {
				for i := 0; i < len(fontArr); i++ {
					up()
					clearLine()
				}
			}
			for _, line := range fontArr {
				fmt.Println("\u001b[34;1m", string(line), "\u001b[0m")
			}
			first = false
			time.Sleep(refreshRate)
		}
	}()

	go flow()

	for {
	}
}

func flow() {
	for {
		for row := range fontArr {
			tempLastIndex := fontArr[row][len(fontArr[row])-1]
			for col := len(fontArr[row]) - 1; col > 0; col-- {
				fontArr[row][col] = fontArr[row][col-1]
			}
			fontArr[row][0] = tempLastIndex
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func PrintBoard(gameBoardArr [][]int) {
	for row := 0; row < len(gameBoardArr); row++ {
		for col := range gameBoardArr[row] {
			fmt.Print("\u001b[37m| ", drawValues[gameBoardArr[row][col]], " \u001b[0m")
		}
		fmt.Println("\u001b[37m|\u001b[0m")
		if row != len(gameBoardArr)-1 {
			fmt.Println("\u001b[37m-------------\u001b[0m")
		}
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
