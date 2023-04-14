package moves

import (
	"testing"

	"github.com/logantwalker/gopher-chess/board"
)

func TestPawnMoves(t *testing.T){
	b := board.NewBoard(board.StartingFen)

	//testing starting row - white
	for sq := board.A2; board.LegalSquare(int8(sq)); sq += board.Square(nextFile){
		moves := generatePawnMoves(&b,int8(sq))
		if len(moves) != 2{
			t.Errorf("Should be generating 2 moves, got %d", len(moves))
		}

		if moves[0].To != sq + board.Square(nextRank){
			t.Errorf("expected %s, got %s",board.SquareHexToString[sq + board.Square(nextRank)],board.SquareHexToString[moves[0].To])
		}

		if moves[1].To != sq + board.Square(2*nextRank){
			t.Errorf("expected %s, got %s",board.SquareHexToString[sq + board.Square(2*nextRank)],board.SquareHexToString[moves[0].To])
		}
	}

	b.Turn = board.Black
	//testing starting row - black
	for sq := board.A7; board.LegalSquare(int8(sq)); sq += board.Square(nextFile){
		moves := generatePawnMoves(&b,int8(sq))
		if len(moves) != 2{
			t.Errorf("Should be generating 2 moves, got %d", len(moves))
		}

		if moves[0].To != sq - board.Square(nextRank){
			t.Errorf("expected %s, got %s",board.SquareHexToString[sq + board.Square(nextRank)],board.SquareHexToString[moves[0].To])
		}

		if moves[1].To != sq - board.Square(2*nextRank){
			t.Errorf("expected %s, got %s",board.SquareHexToString[sq + board.Square(2*nextRank)],board.SquareHexToString[moves[0].To])
		}
	}

	//testing captures & pawns blocked by pieces - white
	b = board.NewBoard("rnbqkbnr/8/8/pppppppp/PPPPPPPP/8/8/RNBQKBNR w KQkq - 0 1")
	
	for sq := board.A4; board.LegalSquare(int8(sq)); sq += board.Square(nextFile){
		moves := generatePawnMoves(&b,int8(sq))
		if sq == board.A4 || sq == board.H4{
			if len(moves) != 1{
				t.Errorf("Should be generating 1 moves, got %d", len(moves))
			}

			if sq == board.A4{
				if moves[0].To != board.B5{
					t.Errorf("expected %s, got %s",board.SquareHexToString[board.B5],board.SquareHexToString[moves[0].To])
				}

				if moves[0].Capture != board.BlackPawn{
					t.Errorf("expected to capture %d, got %d", board.BlackPawn, moves[0].Capture)
				}
			}

			if sq == board.H4{
				if moves[0].To != board.G5{
					t.Errorf("expected %s, got %s",board.SquareHexToString[board.G5],board.SquareHexToString[moves[0].To])
				}

				if moves[0].Capture != board.BlackPawn{
					t.Errorf("expected to capture %d, got %d", board.BlackPawn, moves[0].Capture)
				}
			}
		}else{
			if len(moves) != 2{
				t.Errorf("Should be generating 2 moves, got %d", len(moves))
			}

			if moves[0].To != sq + board.Square(nextRank - nextFile){
				t.Errorf("expected %s, got %s",board.SquareHexToString[sq + board.Square(nextRank - nextFile)],board.SquareHexToString[moves[0].To])
			}	

			if moves[1].To != sq + board.Square(nextRank + nextFile){
				t.Errorf("expected %s, got %s",board.SquareHexToString[sq + board.Square(nextRank + nextFile)],board.SquareHexToString[moves[0].To])
			}	

			if moves[0].Capture != board.BlackPawn{
				t.Errorf("expected to capture %d, got %d", board.BlackPawn, moves[0].Capture)
			}

			if moves[1].Capture != board.BlackPawn{
				t.Errorf("expected to capture %d, got %d", board.BlackPawn, moves[1].Capture)
			}
		}
	}


	//testing captures & pawns blocked by pieces - black
	b.Turn = board.Black

	for sq := board.A5; board.LegalSquare(int8(sq)); sq += board.Square(nextFile){
		moves := generatePawnMoves(&b,int8(sq))
		if sq == board.A5 || sq == board.H5{
			if len(moves) != 1{
				t.Errorf("Should be generating 1 moves, got %d", len(moves))
			}

			if sq == board.A5{
				if moves[0].To != board.B4{
					t.Errorf("expected %s, got %s",board.SquareHexToString[board.B4],board.SquareHexToString[moves[0].To])
				}

				if moves[0].Capture != board.WhitePawn{
					t.Errorf("expected to capture %d, got %d", board.WhitePawn, moves[0].Capture)
				}
			}

			if sq == board.H5{
				if moves[0].To != board.G4{
					t.Errorf("expected %s, got %s",board.SquareHexToString[board.G4],board.SquareHexToString[moves[0].To])
				}

				if moves[0].Capture != board.WhitePawn{
					t.Errorf("expected to capture %d, got %d", board.WhitePawn, moves[0].Capture)
				}
			}
		}else{
			if len(moves) != 2{
				t.Errorf("Should be generating 2 moves, got %d", len(moves))
			}

			if moves[0].To != sq - board.Square(nextRank + nextFile){
				t.Errorf("expected %s, got %s",board.SquareHexToString[sq - board.Square(nextRank - nextFile)],board.SquareHexToString[moves[0].To])
			}	

			if moves[1].To != sq - board.Square(nextRank - nextFile){
				t.Errorf("expected %s, got %s",board.SquareHexToString[sq - board.Square(nextRank + nextFile)],board.SquareHexToString[moves[0].To])
			}	

			if moves[0].Capture != board.WhitePawn{
				t.Errorf("expected to capture %d, got %d", board.WhitePawn, moves[0].Capture)
			}

			if moves[1].Capture != board.WhitePawn{
				t.Errorf("expected to capture %d, got %d", board.WhitePawn, moves[1].Capture)
			}
		}
	}

	// testing En Passant - white
	b = board.NewBoard("rnbqkbnr/pppppppp/8/3P4/8/8/PPP1PPPP/RNBQKBNR b KQkq - 0 1")
	setupMove := Move{From: board.E7, To: board.E5}
	MakeMove(&b, setupMove)

	moves := generatePawnMoves(&b, int8(board.D5))

	if len(moves) != 2 {
		t.Errorf("expected 2 moves, generated %d", len(moves))
	}

	if moves[0].Type != moveEnPassant {
		t.Errorf("expected En Passant available")
	}

	if moves[0].To != board.E6 {
		t.Errorf("expected en passant square e6, got %s", board.SquareHexToString[moves[0].To])
	}

	if b.EnPassant != board.E6 {
		t.Errorf("board did not record correct en passant square, got %s", board.SquareHexToString[b.EnPassant])
	}

	// testing En Passant - black
	b = board.NewBoard("rnbqkbnr/ppp1pppp/8/8/3p4/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	setupMove = Move{From: board.E2, To: board.E4}
	MakeMove(&b, setupMove)

	moves = generatePawnMoves(&b, int8(board.D4))

	if len(moves) != 2 {
		t.Errorf("expected 2 moves, generated %d", len(moves))
	}

	if moves[0].Type != moveEnPassant {
		t.Errorf("expected En Passant available")
	}

	if moves[0].To != board.E3 {
		t.Errorf("expected en passant square e3, got %s", board.SquareHexToString[moves[0].To])
	}

	if b.EnPassant != board.E3 {
		t.Errorf("board did not record correct en passant square, got %s", board.SquareHexToString[b.EnPassant])
	}


	// testing pinned pawns - white 
	b = board.NewBoard("rnbqkbnr/pppp1ppp/8/4p3/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 1")
	setupMove = Move{From: board.D8, To: board.H4}
	MakeMove(&b, setupMove)

	moves = generatePawnMoves(&b, int8(board.F2))

	if len(moves) != 0{
		t.Errorf("expected 0 moves, generated %d", len(moves))
	}

	b = board.NewBoard("rnbqkbnr/ppp2ppp/8/3p4/4P3/8/PPPP1PPP/RNBQKBNR b KQkq - 0 1")
	setupMove = Move{From: board.D8, To: board.E7}
	MakeMove(&b, setupMove)

	moves = generatePawnMoves(&b, int8(board.E4))
	if len(moves) != 1{
		t.Errorf("expected 1 move, generated %d", len(moves))
	}

	if moves[0].To != board.E5 {
		t.Errorf("expected move to e5, got %s", board.SquareHexToString[moves[0].To])
	}

	// testing pinned pawns - black 
	b = board.NewBoard("rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w KQkq - 0 1")
	setupMove = Move{From: board.D1, To: board.H5}
	MakeMove(&b, setupMove)

	moves = generatePawnMoves(&b, int8(board.F7))

	if len(moves) != 0{
		t.Errorf("expected 0 moves, generated %d", len(moves))
	}

	b = board.NewBoard("rnbqkbnr/pppp1ppp/8/4p3/3P4/8/PPPP1PPP/RNBQKBNR w KQkq - 0 1")
	setupMove = Move{From: board.D1, To: board.E2}
	MakeMove(&b, setupMove)

	moves = generatePawnMoves(&b, int8(board.E5))
	if len(moves) != 1{
		t.Errorf("expected 1 move, generated %d", len(moves))
	}

	if moves[0].To != board.E4 {
		t.Errorf("expected move to e4, got %s", board.SquareHexToString[moves[0].To])
	}

	// testing promotions - white
	b = board.NewBoard("8/PPPPPPPP/8/8/8/3K4/8/3k4 w - - 0 1")

	for sq := board.A7; board.LegalSquare(int8(sq)); sq += board.Square(nextFile){
		moves = generatePawnMoves(&b, int8(sq))

		if len(moves) != 4{
			t.Errorf("expected 4 moves, generated %d", len(moves))
		}

		for _, move := range moves {
			if move.Type != movePromote{
				t.Errorf("expected promotion move type, got %d", move.Type)
			}
		}

		if moves[0].Promotion != b.Turn * board.Queen{
			t.Errorf("expected queen promotion, got %d", moves[0].Promotion)
		}
		if moves[1].Promotion != b.Turn * board.Rook{
			t.Errorf("expected rook promotion, got %d", moves[0].Promotion)
		}
		if moves[2].Promotion != b.Turn * board.Bishop{
			t.Errorf("expected bishop promotion, got %d", moves[0].Promotion)
		}
		if moves[3].Promotion != b.Turn * board.Knight{
			t.Errorf("expected knight promotion, got %d", moves[0].Promotion)
		}
	}

	// testing promotions - black
	b = board.NewBoard("8/8/8/3K4/8/3k4/pppppppp/8 b - - 0 1")

	for sq := board.A2; board.LegalSquare(int8(sq)); sq += board.Square(nextFile){
		moves = generatePawnMoves(&b, int8(sq))

		if len(moves) != 4{
			t.Errorf("expected 4 moves, generated %d", len(moves))
		}

		for _, move := range moves {
			if move.Type != movePromote{
				t.Errorf("expected promotion move type, got %d", move.Type)
			}
		}

		if moves[0].Promotion != b.Turn * board.Queen{
			t.Errorf("expected queen promotion, got %d", moves[0].Promotion)
		}
		if moves[1].Promotion != b.Turn * board.Rook{
			t.Errorf("expected rook promotion, got %d", moves[0].Promotion)
		}
		if moves[2].Promotion != b.Turn * board.Bishop{
			t.Errorf("expected bishop promotion, got %d", moves[0].Promotion)
		}
		if moves[3].Promotion != b.Turn * board.Knight{
			t.Errorf("expected knight promotion, got %d", moves[0].Promotion)
		}
	}

}