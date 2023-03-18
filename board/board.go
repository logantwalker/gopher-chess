package board

import (
	"errors"
	"strings"
)

const StartingFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"


type Board struct {
	// history 		[]MoveRecord
	State 			[]int8
	Turn 			int8
	HalfMoveClock 	int
	FullMoves		int
}

func ParseFen(fen string) (Board, error) {
	boardArr := make([]int8, 128)
	var board Board = Board{
		State: boardArr,
	}
	fenArray := strings.Split(fen, " ")

	i := 0
	for j := 0; j < len(fenArray[0]); j++ {
		piece := fenArray[0][j]

		switch piece {
		case 'p':
			board.State[hexBoard[i]] = BlackPawn
		case 'r':
			board.State[hexBoard[i]] = BlackRook
		case 'n':
			board.State[hexBoard[i]] = BlackKnight
		case 'b':
			board.State[hexBoard[i]] = BlackBishop
		case 'q':
			board.State[hexBoard[i]] = BlackQueen
		case 'k':
			board.State[hexBoard[i]] = BlackKing


		case 'P':
			board.State[hexBoard[i]] = WhitePawn
		case 'R':
			board.State[hexBoard[i]] = WhiteRook
		case 'N':
			board.State[hexBoard[i]] = WhiteKnight
		case 'B':
			board.State[hexBoard[i]] = WhiteBishop
		case 'Q':
			board.State[hexBoard[i]] = WhiteQueen
		case 'K':
			board.State[hexBoard[i]] = WhiteKing


		case '1':
			// do nothing
		case '2':
			i++
		case '3':
			i+= 2
		case '4':
			i+= 3
		case '5':
			i+= 4
		case '6':
			i+= 5
		case '7':
			i+= 6
		case '8':
			i+= 7

		case '/':
			i--

		default:
			return board, errors.New("invalid FEN")
		}
		i++
	}

	switch fenArray[1][0] {
	case 'w':
		board.Turn = White
	case 'b':
		board.Turn = Black
	default:
		return board, errors.New("invalid FEN")
	}
	

	return board, nil
}