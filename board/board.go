package board

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

const StartingFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"


type Board struct {
	// history 		[]MoveRecord
	State 			[]int8
	Turn 			int8
	CastlingRights 	string
	HalfMoveClock 	int
	FullMoves		int
}

func NewBoard(fen string) Board {
	b, err := ParseFen(fen)

	if err != nil{
		log.Fatal(err.Error())
	}

	return b
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
			board.State[HexBoard[i]] = BlackPawn
		case 'r':
			board.State[HexBoard[i]] = BlackRook
		case 'n':
			board.State[HexBoard[i]] = BlackKnight
		case 'b':
			board.State[HexBoard[i]] = BlackBishop
		case 'q':
			board.State[HexBoard[i]] = BlackQueen
		case 'k':
			board.State[HexBoard[i]] = BlackKing


		case 'P':
			board.State[HexBoard[i]] = WhitePawn
		case 'R':
			board.State[HexBoard[i]] = WhiteRook
		case 'N':
			board.State[HexBoard[i]] = WhiteKnight
		case 'B':
			board.State[HexBoard[i]] = WhiteBishop
		case 'Q':
			board.State[HexBoard[i]] = WhiteQueen
		case 'K':
			board.State[HexBoard[i]] = WhiteKing


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

func LegalSquare(square int8) bool {
	return !(uint8(square)&0x88 != 0)
}

func (b *Board) PrintBoard(){
	for i := 0x70; i >= 0x00; i -= 0x10 {
		for j := 0; j < 8; j++ {
			square := i + j
			fmt.Printf("%v ", GetPieceSymbol(b.State[square]))
		}
		fmt.Printf("\n")
	}
}
