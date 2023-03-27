package moves

import (
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
		case board.BlackKing:
			if b.Turn == board.Black{
				availableMoves = generateKingMoves(b, hex)
			}
		case board.WhiteQueen:
			if b.Turn == board.White{
				availableMoves = generateQueenMoves(b,hex)
			}
		case board.BlackQueen:
			if b.Turn == board.Black{
				availableMoves = generateQueenMoves(b,hex)
			}
		case board.WhiteRook:
			if b.Turn == board.White{
				availableMoves = generateRookMoves(b,hex)
			}
		case board.BlackRook:
			if b.Turn == board.Black{
				availableMoves = generateRookMoves(b,hex)
			}
		case board.WhiteBishop:
			if b.Turn == board.White{
				availableMoves = generateBishopMoves(b,hex)
			}
		case board.BlackBishop:
			if b.Turn == board.Black{
				availableMoves = generateBishopMoves(b,hex)
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
		}else{
			if delta == 2*moveRight || delta == 2*moveLeft {
				canCastle := checkCastlingAvailability(b)
				if canCastle[0] && delta == 2*moveRight{
					move := createMove(origin, dest)
					move.MovedPiece = board.WhiteKing
					move.Castling = int8(blackKingSideCastlingSquares[1])
					moves = append(moves, move)
				}
				if canCastle[1] && delta == 2*moveLeft{
					move := createMove(origin, dest)
					move.MovedPiece = board.BlackKing
					move.Castling = int8(blackQueenSideCastlingSquares[1])
					moves = append(moves, move)
				}
			}else{
				if board.LegalSquare(dest) && b.State[dest] >= 0{
					move := createMove(origin, dest)
					move.MovedPiece = board.BlackKing
					moves = append(moves, move)
				}
			}
		}
	}

	return moves
}

func generateQueenMoves(b board.Board, origin int8) []Move {
	var moves []Move
	
	// check left moves
	moves = validateLongRangeMoves(origin,-nextFile, b, moves)
	// check right moves
	moves = validateLongRangeMoves(origin, nextFile,b,moves)
	// check up moves
	moves = validateLongRangeMoves(origin, nextRank,b,moves)
	// check down moves
	moves = validateLongRangeMoves(origin,-nextFile,b,moves)
	// check upright diag
	moves = validateLongRangeMoves(origin, nextRank + nextFile,b,moves)
	// check downright diag
	moves = validateLongRangeMoves(origin,nextFile - nextRank,b,moves)
	// check upleft diag
	moves = validateLongRangeMoves(origin,nextRank - nextFile,b,moves)
	// check downleft diag
	moves = validateLongRangeMoves(origin,-nextFile - nextRank,b,moves)

	return moves
}

func generateRookMoves(b board.Board, origin int8) []Move {
	var moves []Move
	// move up
	moves = validateLongRangeMoves(origin, nextRank,b,moves)
	// move down
	moves = validateLongRangeMoves(origin,-nextRank,b,moves)
	// move left
	moves = validateLongRangeMoves(origin, -nextFile, b, moves)
	// move right
	moves = validateLongRangeMoves(origin,nextFile,b,moves)

	return moves
}

func generateBishopMoves(b board.Board, origin int8) []Move {
	var moves []Move
	// move up left
	moves = validateLongRangeMoves(origin,nextRank - nextFile,b,moves)
	// move up right
	moves = validateLongRangeMoves(origin,nextRank + nextFile,b,moves)
	// move down left
	moves = validateLongRangeMoves(origin, -nextRank - nextFile, b, moves)
	// move down right
	moves = validateLongRangeMoves(origin,nextFile - nextRank,b,moves)

	return moves
}

func validateLongRangeMoves(origin int8, delta int8, b board.Board, moves []Move) []Move {
	for i := origin; board.LegalSquare(i); i += delta{
		if b.State[i] == board.Empty{
			move := createMove(origin, i)
			move.MovedPiece = b.State[origin]
			moves = append(moves, move)
		}else if b.State[i] < board.Empty && b.Turn == board.White{ // checking if occupying piece is black and turn is white
			move := createMove(origin,i)
			move.Capture = b.State[i]
			move.MovedPiece = b.State[origin]
			moves = append(moves, move)
			break
		}else if b.State[i] > board.Empty && b.Turn == board.Black{
			move := createMove(origin,i)
			move.Capture = b.State[i]
			move.MovedPiece = b.State[origin]
			moves = append(moves, move)
			break
		}else if b.State[i] > board.Empty && b.Turn == board.White && i != origin{
			break
		}else if b.State[i] < board.Empty && b.Turn == board.Black && i != origin{
			break
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
		if (b.WhiteCastle & 1) > 0  {
			castleRights[0] = true
		}
		if (b.WhiteCastle & 2) > 0 {
			castleRights[1] = true
		}

		if castleRights[0] && b.State[whiteKingSideCastlingSquares[0]] == board.Empty && b.State[whiteKingSideCastlingSquares[1]] == board.Empty{
			castleAbility[0] = true
		}
		if castleRights[1] && b.State[whiteQueenSideCastlingSquares[0]] == board.Empty && b.State[whiteQueenSideCastlingSquares[1]] == board.Empty && b.State[whiteQueenSideCastlingSquares[2]] == board.Empty{
			castleAbility[1] = true
		}
	}else{
		if (b.WhiteCastle & 1) > 0 {
			castleRights[0] = true
		}
		if (b.BlackCastle & 2) > 0 {
			castleRights[1] = true
		}

		if castleRights[0] && b.State[blackKingSideCastlingSquares[0]] == board.Empty && b.State[blackKingSideCastlingSquares[1]] == board.Empty{
			castleAbility[0] = true
		}
		if castleRights[1] && b.State[blackQueenSideCastlingSquares[0]] == board.Empty && b.State[blackQueenSideCastlingSquares[1]] == board.Empty && b.State[blackQueenSideCastlingSquares[2]] == board.Empty{
			castleAbility[1] = true
		}
	}

	return []bool{castleRights[0] && castleAbility[0],castleRights[1] && castleAbility[1],}
}

