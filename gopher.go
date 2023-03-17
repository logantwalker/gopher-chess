package main

import (
	"fmt"

	"github.com/logantwalker/gopher-chess/board"
)

func main(){
	boardArr := make([]int8, 128)

	boardArr[0x00] = board.WhiteRook
	boardArr[0x01] = board.WhiteKnight
	boardArr[0x02] = board.WhiteBishop
	boardArr[0x03] = board.WhiteQueen
	boardArr[0x04] = board.WhiteKing
	boardArr[0x05] = board.WhiteBishop
	boardArr[0x06] = board.WhiteKnight
	boardArr[0x07] = board.WhiteRook

	for i := 0x10; i < 0x18; i++ {
		boardArr[i] = board.WhitePawn
	}

	boardArr[0x70] = board.BlackRook
	boardArr[0x71] = board.BlackKnight
	boardArr[0x72] = board.BlackBishop
	boardArr[0x73] = board.BlackQueen
	boardArr[0x74] = board.BlackKing
	boardArr[0x75] = board.BlackBishop
	boardArr[0x76] = board.BlackKnight
	boardArr[0x77] = board.BlackRook

	for i := 0x60; i < 0x68; i++ {
		boardArr[i] = board.BlackPawn
	}

	for i := 0x70; i >= 0x00; i -= 0x10 {
		for j := 0; j < 8; j++ {
			square := i + j
			fmt.Printf("%v ", board.GetPieceSymbol(boardArr[square]))
		}
		fmt.Printf("\n")
	}
}