package main

import "fmt"

func main(){
	board := make([]int8, 128)

	// Set the starting position of the pieces
	board[0x00] = 'R'
	board[0x01] = 'N'
	board[0x02] = 'B'
	board[0x03] = 'Q'
	board[0x04] = 'K'
	board[0x05] = 'B'
	board[0x06] = 'N'
	board[0x07] = 'R'

	for i := 0x10; i < 0x18; i++ {
		board[i] = 'P'
	}

	board[0x70] = 'r'
	board[0x71] = 'n'
	board[0x72] = 'b'
	board[0x73] = 'q'
	board[0x74] = 'k'
	board[0x75] = 'b'
	board[0x76] = 'n'
	board[0x77] = 'r'

	for i := 0x60; i < 0x68; i++ {
		board[i] = 'p'
	}

	// Print out the board
	for i := 0x70; i >= 0x00; i -= 0x10 {
		for j := 0; j < 8; j++ {
			square := i + j
			if board[square] == 0 {
				fmt.Printf(". ")
			} else {
				fmt.Printf("%c ", board[square])
			}
		}
		fmt.Printf("\n")
	}
}