package moves

import (
	"testing"

	"github.com/logantwalker/gopher-chess/board"
)

func TestCheck(t *testing.T) {
	b := board.NewBoard("4k3/8/4q3/8/8/4n2B/3N4/3RK3 b - - 0 1")
	MakeMove(&b,Move{From: board.E3, To: board.C2})

	moves := GenerateMovesList(&b)
	if len(moves) != 2{
		t.Errorf("expected 2 moves, got %d", len(moves))
	}

	if moves[1].To != board.F1 && moves[0].To != board.F2{
		t.Error("generating invalid moves: ", board.SquareHexToString[moves[0].To], board.F2)
	}
}