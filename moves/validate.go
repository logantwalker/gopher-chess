package moves

import (
	"errors"

	"github.com/logantwalker/gopher-chess/board"
)

func ValidateUserMove(b *board.Board, move board.Move) (board.Move, error){
	gen := NewGenerator(b)
	validMoves := gen.GenerateMoves()

	for _, validMove := range validMoves{
		if move.Type == movePromote{
			if validMove.To == move.To && validMove.From == move.From && validMove.Promotion == move.Promotion{
				return validMove, nil
			}
		}else{
			if validMove.To == move.To && validMove.From == move.From{
				return validMove, nil
			}
		}
	}

	return board.Move{}, errors.New("invalid move")
}


func checkRepititions(b *board.Board) {
	r := 0
	first := len(b.History) - b.HalfMoveClock
	if first >= 0 {
		for i := first; i < len(b.History)-1; i++ {
			if b.History[i].ZobristHash == b.ZobristHash {
				r++
			}
		}
	}

	if r >= 3 {
		b.Status = board.StatusThreeFoldRep
	}
}