package moves

import "github.com/logantwalker/gopher-chess/board"

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
	blackQueenSideCastingSquares []board.Square = []board.Square{board.B8, board.C8, board.D8}
)

type Move struct{
	From 		board.Square
	To 			board.Square
	Capture 	int8
	MovedPiece 	int8
	Promotion 	int8
	Castling	int8
	EnPassant	int8
}


func createMove(origin int8, dest int8) Move {
	var move Move

	if board.LegalSquare(dest){
		move = Move{From: board.Square(origin), To: board.Square(dest)}
	}
	
	return move
}
