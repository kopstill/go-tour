package main

import (
	"fmt"
	"strings"
)

func main() {
	board := [][]string{
		//[]string{"X", "_", "X"},
		//[]string{"O", "_", "X"},
		//[]string{"_", "_", "O"},
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
