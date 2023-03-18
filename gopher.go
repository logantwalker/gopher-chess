package main

import (
	"fmt"

	"github.com/logantwalker/gopher-chess/board"
)

func main(){
	boardArr := board.ParseFen("rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2")
	for i := 0x70; i >= 0x00; i -= 0x10 {
		for j := 0; j < 8; j++ {
			square := i + j
			fmt.Printf("%v ", board.GetPieceSymbol(boardArr[square]))
		}
		fmt.Printf("\n")
	}
}