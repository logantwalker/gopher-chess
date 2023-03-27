package main

import (
	"github.com/logantwalker/gopher-chess/board"
	"github.com/logantwalker/gopher-chess/game"
)

func main(){
	boardObj, _  := board.ParseFen("rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2")
	boardObj.CastlingRights = "KQkq"

	g := game.NewGame()
	g.Run()

	// moves := moves.GenerateMovesList(boardObj)

	// for _, move := range moves {
	// 	pieceSymbol := board.GetPieceSymbol(move.MovedPiece)
	// 	moveString := board.SquareHexToString[move.From] + board.SquareHexToString[move.To]

	// 	if move.Capture != board.Empty{
	// 		captureSymbol := board.GetPieceSymbol(move.Capture)
	// 		fmt.Println(pieceSymbol + " " + moveString + " " + captureSymbol)
	// 	}
	// 	fmt.Println(pieceSymbol + " " + moveString)
	// }
}