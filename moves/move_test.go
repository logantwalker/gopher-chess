package moves

// import (
// 	"log"
// 	"testing"

// 	"github.com/logantwalker/gopher-chess/board"
// )

// func TestMakeMoveUndoMove(t *testing.T) {
// 	b := board.NewBoard(board.StartingFen)

// 	MakeMove(&b, board.Move{From: board.E2,To: board.E4})
// 	MakeMove(&b, board.Move{From: board.E7,To: board.E5})

// 	MakeMove(&b, board.Move{From: board.G1,To: board.F3})
// 	MakeMove(&b, board.Move{From: board.G8,To: board.F6})

// 	MakeMove(&b, board.Move{From: board.F1,To: board.C4})
// 	MakeMove(&b, board.Move{From: board.F8,To: board.C5})

// 	MakeMove(&b, board.Move{From: board.E1,To: board.G1})
// 	MakeMove(&b, board.Move{From: board.E8,To: board.G8})

// 	for len(b.History) > 0 {
// 		UndoMove(&b)
// 	}

// 	if b.State[int8(board.E2)] == board.Empty{
// 		t.Errorf("failed to undo white pawn move")
// 	}

// 	if b.State[int8(board.E4)] != board.Empty{
// 		t.Errorf("failed to undo white pawn move")
// 	}

// 	if b.State[int8(board.F1)] != board.WhiteBishop{
// 		t.Errorf("failed to undo white bishop move")
// 	}

// 	if b.WhiteCastle != board.CastleLong + board.CastleShort{
// 		t.Errorf("failed to reset castle rights")
// 	}

// 	if b.BlackCastle != board.CastleLong + board.CastleShort{
// 		t.Errorf("failed to reset castle rights")
// 	}

// 	if b.KingLocations[0] != int8(board.E1) || b.State[int8(board.E1)] != board.WhiteKing{
// 		t.Errorf("failed to return king")
// 	}

// 	if b.KingLocations[1] != int8(board.E8) || b.State[int8(board.E8)] != board.BlackKing{
// 		t.Errorf("failed to return king")
// 	}
// }

// func TestDisappearingBishop(t *testing.T){
// 	b := board.NewBoard(board.StartingFen)
// 	setupMoves := []board.Move{
// 		{From: board.D2,To: board.D3},
// 		{From: board.B8,To: board.A6},
// 		{From: board.C1,To: board.H6},
// 	}

// 	for _, move:= range setupMoves {
// 		MakeMove(&b,move)
// 	}

// 	if b.State[board.H6] != board.WhiteBishop{
// 		t.Errorf("we lost the bishop")
// 	}

// 	moves := GenerateMovesList(&b)

// 	for _, move := range moves{
// 		if b.State[board.H6] != board.WhiteBishop{
// 			log.Println("before MakeMove")
// 			log.Println("move: ", move)
// 			log.Fatal("we lost the bishop.")
// 		}

// 		MakeMove(&b, move)

// 		// if b.State[board.H6] != board.WhiteBishop && move.From != board.G8{
// 		// 	log.Println("after MakeMove, before UndoMove")
// 		// 	invalidMoveString := board.SquareHexToString[move.From] + board.SquareHexToString[move.To]
// 		// 	log.Printf("move: %s, turn: %d\n", invalidMoveString, b.Turn)
// 		// 	log.Fatal("we lost the bishop")

// 		// }

// 		UndoMove(&b)

// 		if b.State[board.H6] != board.WhiteBishop{
// 			log.Println("after UndoMove")
// 			log.Print("we lost the bishop")

// 			invalidMoveString := board.SquareHexToString[move.From] + board.SquareHexToString[move.To]
// 			offendingPiece := board.GetPieceSymbol(b.State[move.From])
// 			log.Printf("offending piece %s",offendingPiece)
// 			log.Fatalf("move: %s, turn: %d\n", invalidMoveString, b.Turn)
// 		}
// 	}
// }