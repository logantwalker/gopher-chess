package moves

import (
	"strings"

	"github.com/logantwalker/gopher-chess/board"
)

func GenerateMovesList(b board.Board) []Move {
	var moves []Move

	for _, hex := range board.HexBoard {
		square := b.State[hex]
		var availableMoves []Move
		switch square {
		case board.WhitePawn:
			if b.Turn == board.White{
				availableMoves = generatePawnMoves(b, hex)
			}
		case board.BlackPawn:
			if b.Turn == board.Black{
				availableMoves = generatePawnMoves(b, hex)
			}
		case board.WhiteKnight:
			if b.Turn == board.White{
				availableMoves = generateKnightMoves(b, hex)
			}
		case board.BlackKnight:
			if b.Turn == board.Black{
				availableMoves = generateKnightMoves(b, hex)
			}
		case board.WhiteKing:
			if b.Turn == board.White{
				availableMoves = generateKingMoves(b, hex)
			}
		}
		moves = append(moves, availableMoves...)
	}
	return moves
}

func generatePawnMoves(b board.Board, origin int8) []Move {
	var moves []Move
	if b.Turn == board.White{
		for _, delta := range whitePawnMoves{
			dest := origin + delta
			if delta == moveUpandLeft || delta == moveUpandRight{
				validateAttacks := checkPawnAttacks(b, origin)

				if delta == moveUpandLeft && validateAttacks[0]{
					move := createMove(origin,dest)
					move.MovedPiece = board.WhitePawn
					moves = append(moves, move)
				}

				if delta == moveUpandRight && validateAttacks[1]{
					move := createMove(origin,dest)
					move.MovedPiece = board.WhitePawn
					moves = append(moves, move)
				}
			}else if delta == 2*moveUp{
				if rank:= board.Rank(origin); rank == whitePawnStartRank && b.State[dest] == board.Empty && b.State[dest + moveDown] == board.Empty{
					move := createMove(origin,dest)
					move.MovedPiece = board.WhitePawn
					moves = append(moves, move)
				}
			}else{
				if b.State[dest] == board.Empty{
					move := createMove(origin,dest)
					move.MovedPiece = board.WhitePawn
					moves = append(moves, move)
				}
			}
		}
	}else{
		for _, delta := range blackPawnMoves{
			dest := origin + delta
			if delta == moveDownandLeft || delta == moveDownandRight{
				validateAttacks := checkPawnAttacks(b, origin)

				if delta == moveDownandLeft && validateAttacks[0]{
					move := createMove(origin,dest)
					move.MovedPiece = board.BlackPawn
					moves = append(moves, move)
				}

				if delta == moveDownandRight && validateAttacks[1]{
					move := createMove(origin,dest)
					move.MovedPiece = board.BlackPawn
					moves = append(moves, move)
				}
			}else if delta == 2*moveDown{
				if rank:= board.Rank(origin); rank == blackPawnStartRank && b.State[dest] == board.Empty && b.State[dest + moveUp] == board.Empty{
					move := createMove(origin,dest)
					move.MovedPiece = board.BlackPawn
					moves = append(moves, move)
				}
			}else{
				if b.State[dest] == board.Empty{
					move := createMove(origin,dest)
					move.MovedPiece = board.BlackPawn
					moves = append(moves, move)
				}
			}
		}
	}

	return moves
}

func generateKnightMoves(b board.Board, origin int8) []Move {
	var moves []Move

	for _, delta := range knightMoves{
		dest := origin + delta
		if board.LegalSquare(dest) && b.State[dest] <=0 && b.Turn == board.White{
			move := createMove(origin, dest)
			move.MovedPiece = board.WhiteKnight
			moves = append(moves, move)
		}else if board.LegalSquare(dest) && b.State[dest] >=0 && b.Turn == board.Black{
			move := createMove(origin, dest)
			move.MovedPiece = board.BlackKnight
			moves = append(moves, move)
		}		
	}

	return moves
}

//regarding king moves, I need to check king safety in the future before alowing a move
func generateKingMoves(b board.Board, origin int8) []Move {
	var moves []Move
	for _, delta := range kingMoves{
		dest := origin + delta
		if b.Turn == board.White{
			if delta == 2*moveRight || delta == 2*moveLeft {
				canCastle := checkCastlingAvailability(b)
				if canCastle[0] && delta == 2*moveRight{
					move := createMove(origin, dest)
					move.MovedPiece = board.WhiteKing
					move.Castling = int8(whiteKingSideCastlingSquares[1])
					moves = append(moves, move)
				}
				if canCastle[1] && delta == 2*moveLeft{
					move := createMove(origin, dest)
					move.MovedPiece = board.WhiteKing
					move.Castling = int8(whiteQueenSideCastlingSquares[1])
					moves = append(moves, move)
				}
			}else{
				if board.LegalSquare(dest) && b.State[dest] <= 0{
					move := createMove(origin, dest)
					move.MovedPiece = board.WhiteKing
					moves = append(moves, move)
				}
			}
		}
	}

	return moves
}

func checkPawnAttacks(b board.Board, origin int8) []bool{
	var attacks = []bool{false, false}
	if b.Turn == board.White{
		if b.State[origin + moveUpandLeft] < int8(0) {
			attacks[0] = true
		}
		if b.State[origin + moveUpandRight] < int8(0) {
			attacks[1] = true
		}
	}else{
		if b.State[origin + moveDownandLeft] > int8(0) {
			attacks[0] = true
		}
		if b.State[origin + moveDownandRight] > int8(0) {
			attacks[1] = true
		}
	}

	return attacks
}

func checkCastlingAvailability(b board.Board) []bool {
	castleAbility := []bool{false, false}
	castleRights := []bool{false, false}

	if b.Turn == board.White{
		if strings.Contains(b.CastlingRights,"K") {
			castleRights[0] = true
		}
		if strings.Contains(b.CastlingRights, "Q"){
			castleRights[1] = true
		}

		if castleRights[0] && b.State[whiteKingSideCastlingSquares[0]] == board.Empty && b.State[whiteKingSideCastlingSquares[1]] == board.Empty{
			castleAbility[0] = true
		}
		if castleRights[1] && b.State[whiteQueenSideCastlingSquares[0]] == board.Empty && b.State[whiteQueenSideCastlingSquares[1]] == board.Empty && b.State[whiteQueenSideCastlingSquares[2]] == board.Empty{
			castleAbility[1] = true
		}
	}else{

	}

	return []bool{castleRights[0] && castleAbility[0],castleRights[1] && castleAbility[1],}
}

