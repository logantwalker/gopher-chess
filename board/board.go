package board

import (
	"strings"
)

const StartingFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

type Board struct {
	state 			[]int8
	turn 			int8
	halfMoveClock 	int
	fullMoves		int
}

func ParseFen(fen string) []int8 {
	boardArr := make([]int8, 128)
	fenArray := strings.Split(fen, " ")

	i := 0
	for j := 0; j < len(fenArray[0]); j++ {
		piece := fenArray[0][j]

		switch piece {
		case 'p':
			boardArr[hexBoard[i]] = BlackPawn
		case 'r':
			boardArr[hexBoard[i]] = BlackRook
		case 'n':
			boardArr[hexBoard[i]] = BlackKnight
		case 'b':
			boardArr[hexBoard[i]] = BlackBishop
		case 'q':
			boardArr[hexBoard[i]] = BlackQueen
		case 'k':
			boardArr[hexBoard[i]] = BlackKing


		case 'P':
			boardArr[hexBoard[i]] = WhitePawn
		case 'R':
			boardArr[hexBoard[i]] = WhiteRook
		case 'N':
			boardArr[hexBoard[i]] = WhiteKnight
		case 'B':
			boardArr[hexBoard[i]] = WhiteBishop
		case 'Q':
			boardArr[hexBoard[i]] = WhiteQueen
		case 'K':
			boardArr[hexBoard[i]] = WhiteKing


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
			return boardArr
		}
		i++
	}

	return boardArr
}