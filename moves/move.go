package moves

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/logantwalker/gopher-chess/board"
)

var (
	nextRank int8 = 16
	nextFile int8 = 1

	moveUp = nextRank
	moveDown = -nextRank
	moveRight = nextFile
	moveLeft = -nextFile

	moveUpandRight = moveUp + moveRight
	moveUpandLeft = moveUp + moveLeft
	moveDownandRight = moveDown + moveRight
	moveDownandLeft = moveDown + moveLeft

	knightMoves = []int8{moveUp + moveUpandRight, moveUp + moveUpandLeft, moveDown + moveDownandLeft, moveDown + moveDownandRight, moveUp + 2*moveLeft, moveDown + 2*moveLeft, moveUp + 2*moveRight, moveDown + 2*moveRight}
	rookMoves = []int8{moveUp, moveDown, moveLeft, moveRight}
	bishopMoves = []int8{moveUpandLeft, moveUpandRight, moveDownandLeft, moveDownandRight}
	queenMoves = []int8{moveUp, moveDown, moveLeft, moveRight, moveUpandLeft, moveUpandRight, moveDownandLeft, moveDownandRight}
	kingMoves = []int8{moveUp, moveDown, moveLeft, moveRight, moveUpandLeft, moveUpandRight, moveDownandLeft, moveDownandRight, 2*moveLeft, 2*moveRight}

	whitePawnMoves = []int8{moveUp, 2*moveUp, moveUpandLeft, moveUpandRight,}
	blackPawnMoves = []int8{moveDown, 2*moveDown, moveDownandLeft, moveDownandRight,}

	whitePawnStartRank int8 = 1 // rank 2
	blackPawnStartRank int8 = 6 // rank 7

	whiteKingSideCastlingSquares []board.Square = []board.Square{board.F1,board.G1}
	whiteQueenSideCastlingSquares []board.Square = []board.Square{board.B1, board.C1, board.D1}

	blackKingSideCastlingSquares []board.Square = []board.Square{board.F8,board.G8}
	blackQueenSideCastlingSquares []board.Square = []board.Square{board.B8, board.C8, board.D8}
)

// move types
var (
	moveOrdinary 	int8 = 0
	moveShortCastle int8 = 1
	moveLongCastle 	int8 = 2
	movePromote		int8 = 3
	moveEnPassant	int8 = 4
)

type Move struct{
	From 		board.Square
	To 			board.Square
	Capture 	int8
	MovedPiece 	int8
	Type 		int8
	Promotion 	int8
}

func CreateMoveFromInput(input string) (Move, error) {
	input = strings.Trim(input, " ")
	if m, _ := regexp.MatchString("^[a-h][1-8][a-h][1-8]$", input); !m {
		return Move{}, errors.New("invalid move")
	}

	from := board.SquareStringToHex[input[:2]]
	to := board.SquareStringToHex[input[2:]]

	var move Move = Move{From: from, To: to}
	return move, nil
}


func createMove(origin int8, dest int8) Move {
	var move Move

	if board.LegalSquare(dest){
		move = Move{From: board.Square(origin), To: board.Square(dest)}
	}
	
	return move
}

func PrintMoves(moves []Move) {
	for _, move := range moves {
		pieceSymbol := board.GetPieceSymbol(move.MovedPiece)
		moveString := board.SquareHexToString[move.From] + board.SquareHexToString[move.To]

		if move.Capture != board.Empty{
			captureSymbol := board.GetPieceSymbol(move.Capture)
			fmt.Println(pieceSymbol + " " + moveString + " " + captureSymbol)
		}
		fmt.Println(pieceSymbol + " " + moveString)
	}
} 

func MakeMove(b board.Board, move Move) board.Board{
	b.HalfMoveClock++

	if b.Turn == board.Black{
		b.FullMoveClock++
	}

	validMove, err := ValidateUserMove(b, move)
	if err != nil {
		fmt.Println(err.Error())
		return b
	}

	switch validMove.Type {
	case moveOrdinary:
		b.State[validMove.From] = board.Empty
		b.State[validMove.To] = validMove.MovedPiece

		switch validMove.MovedPiece {
		case board.WhiteKing:
			if validMove.From == board.WhiteKingStartSquare{
				b.WhiteCastle = board.CastleNone
			}
		}
	}
	b.Turn = -1 * b.Turn
	
	return b
}

func ValidateUserMove(b board.Board, move Move) (Move, error){
	validMoves := GenerateMovesList(b)

	for _, validMove := range validMoves{
		if validMove.To == move.To && validMove.From == move.From{
			return validMove, nil
		}
	}

	return Move{}, errors.New("invalid move")
}