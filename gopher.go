package main

import (
	"fmt"

	"github.com/logantwalker/gopher-chess/board"
	"github.com/logantwalker/gopher-chess/moves"
)

func main(){
	boardObj, _  := board.ParseFen("8/8/8/8/4b3/8/8/8 b kq - 1 2")
	boardObj.CastlingRights = "KQkq"

	for i := 0x70; i >= 0x00; i -= 0x10 {
		for j := 0; j < 8; j++ {
			square := i + j
			fmt.Printf("%v ", board.GetPieceSymbol(boardObj.State[square]))
		}
		fmt.Printf("\n")
	}

	moves := moves.GenerateMovesList(boardObj)

	for _, move := range moves {
		pieceSymbol := board.GetPieceSymbol(move.MovedPiece)
		moveString := board.SquareHexToString[move.From] + board.SquareHexToString[move.To]

		if move.Capture != board.Empty{
			captureSymbol := board.GetPieceSymbol(move.Capture)
			fmt.Println(pieceSymbol + " " + moveString + " " + captureSymbol)
		}
		fmt.Println(pieceSymbol + " " + moveString)
	}
}