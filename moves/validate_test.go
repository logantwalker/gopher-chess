package moves

import (
	"fmt"
	"testing"

	"github.com/logantwalker/gopher-chess/board"
)

func TestCheck(t *testing.T) {

	// testing double attacks - white
	b := board.NewBoard("4k3/8/4q3/8/8/4n2B/3N4/3RK3 b - - 0 1")
	MakeMove(&b,Move{From: board.E3, To: board.C2})

	moves := GenerateMovesList(&b)
	if len(moves) != 2{
		fmt.Println(moves)
		t.Errorf("expected 2 moves, got %d", len(moves))
	}

	if moves[0].To != board.F1 && moves[1].To != board.F2{
		t.Error("generating invalid moves: ", board.SquareHexToString[moves[0].To], board.F2)
	}

	// testing double attacks - black
	b = board.NewBoard("4K3/8/4Q3/8/8/4N2b/3n4/3rk3 w - - 0 1")
	MakeMove(&b,Move{From: board.E3, To: board.C2})
	b.KingLocations[1] = int8(board.E1)

	moves = GenerateMovesList(&b)
	if len(moves) != 2{
		t.Errorf("expected 2 moves, got %d", len(moves))
	}

	if moves[0].To != board.F1 && moves[1].To != board.F2{
		t.Error("generating invalid moves: ", board.SquareHexToString[moves[0].To], board.F2)
	}

	// 2r2q1k/5pp1/4p3/8/1bp4N/5P1R/6P1/2R4K w - - 0 1
	b = board.NewBoard("2r2q1k/5pp1/4p3/8/1bp4N/5P1R/6P1/2R4K w - - 0 1")
	MakeMove(&b,Move{From: board.H4, To: board.G6})
	b.KingLocations[1] = int8(board.H8)

	moves = GenerateMovesList(&b)
	if len(moves) != 1{
		t.Errorf("expected 1 move, got %d", len(moves))
	}

	if moves[0].To != board.G8 {
		t.Errorf("generated incorrect destination, %s", board.SquareHexToString[moves[0].To])
	}

	// rnbk1b1r/pp3ppp/2p5/4q3/4n3/8/PPPB1PPP/2KR1BNR b - - 0 1
	b = board.NewBoard("rnbk1b1r/pp3ppp/2p5/4q3/4n3/8/PPPB1PPP/2KR1BNR w - - 0 1")
	b.KingLocations[1] = int8(board.D8)
	MakeMove(&b,Move{From: board.D2, To: board.G5})

	moves = GenerateMovesList(&b)
	if len(moves) != 2{
		t.Errorf("expected 2 move, got %d", len(moves))
	}

	if moves[0].To != board.E8 && moves[1].To != board.C7{
		t.Error("generating invalid moves")
	}
}