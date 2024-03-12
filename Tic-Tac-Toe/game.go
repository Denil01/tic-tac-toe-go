package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var board [3][3]string
var currentPlayer string

func main() {
	initializeBoard()
	displayBoard()

	for {
		makeMove()
		displayBoard()

		if checkWin() {
			fmt.Printf("Player %s wins!\n", currentPlayer)
			break
		}

		if checkDraw() {
			fmt.Println("It is a draw!")
			break
		}
		switchPlayer()
	}
}

func initializeBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = " "
		}
	}
	currentPlayer = "X"
}

func displayBoard() {
	clearScreen()
	for i := 0; i < 3; i++ {
		fmt.Println(strings.Join(board[i][:], " | "))
		if i < 2 {
			fmt.Println("---------")
		}
	}
}

func clearScreen() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd = exec.Command("clear") // Linux or macOS
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls") // Windows
	default:
		fmt.Println("Unsupported platform")
		return
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func makeMove() {
	var row, col int
	for {
		fmt.Printf("Player %s, enter your move (row[1-3] clo[1-3]): ", currentPlayer)
		fmt.Scan(&row, &col)
		row--
		col--
		if row >= 0 && row < 3 && col >= 0 && col < 3 && board[row][col] == " " {
			board[row][col] = currentPlayer
			break
		}
		fmt.Println("Invalid move. Try again.")
	}
}

func checkWin() bool {
	// Check rows
	for i := 0; i < 3; i++ {
		if board[i][0] != " " && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return true
		}
	}

	// Check columns
	for j := 0; j < 3; j++ {
		if board[0][j] != " " && board[0][j] == board[1][j] && board[1][j] == board[2][j] {
			return true
		}
	}

	// Check diagonals
	if board[0][0] != " " && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return true
	}
	if board[0][2] != " " && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return true
	}
	return false
}

func checkDraw() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == " " {
				return false
			}
		}
	}
	return true
}

func switchPlayer() {
	if currentPlayer == "X" {
		currentPlayer = "O"
	} else {
		currentPlayer = "X"
	}
}
