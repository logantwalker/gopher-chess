package moves

import (
	"testing"

	"github.com/logantwalker/gopher-chess/board"
)

func TestMakeMoveUndoMove(t *testing.T) {
	b := board.NewBoard(board.StartingFen)

	MakeMove(&b, board.Move{From: board.E2,To: board.E4})
	MakeMove(&b, board.Move{From: board.E7,To: board.E5})

	MakeMove(&b, board.Move{From: board.G1,To: board.F3})
	MakeMove(&b, board.Move{From: board.G8,To: board.F6})

	MakeMove(&b, board.Move{From: board.F1,To: board.C4})
	MakeMove(&b, board.Move{From: board.F8,To: board.C5})
	
	MakeMove(&b, board.Move{From: board.E1,To: board.G1})
	MakeMove(&b, board.Move{From: board.E8,To: board.G8})

	for len(b.History) > 0 {
		UndoMove(&b)
	}

	if b.State[int8(board.E2)] == board.Empty{
		t.Errorf("failed to undo white pawn move")
	}

	if b.State[int8(board.E4)] != board.Empty{
		t.Errorf("failed to undo white pawn move")
	}

	if b.WhiteCastle != board.CastleLong + board.CastleShort{
		t.Errorf("failed to reset castle rights")
	}

	if b.BlackCastle != board.CastleLong + board.CastleShort{
		t.Errorf("failed to reset castle rights")
	}

	if b.KingLocations[0] != int8(board.E1) || b.State[int8(board.E1)] != board.WhiteKing{
		t.Errorf("failed to return king")
	}

	if b.KingLocations[1] != int8(board.E8) || b.State[int8(board.E8)] != board.BlackKing{
		t.Errorf("failed to return king")
	}
}