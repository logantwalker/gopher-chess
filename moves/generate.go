package moves

import (
	"github.com/logantwalker/gopher-chess/board"
)

func GenerateMovesList(b *board.Board) []board.Move {
	var moves []board.Move

	for _, hex := range board.HexBoard {
		square := b.State[hex]
		var availableMoves []board.Move
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

	if len(moves) == 0 && !b.IsCheck{
		b.Status = board.StatusStalemate
	}

	if b.IsCheck {
		moves = generateMovesWhileInCheck(b, moves)
		if len(moves) == 0{
			b.Status = board.StatusCheckmate
		}
	}


	return moves
}

func generateMovesWhileInCheck(b *board.Board, moves []board.Move) []board.Move {
	var legalMoves []board.Move
	if len(b.Checks) == 1{
		legalMoves = findLegalMovesForCheck(b,b.Checks[0],moves)
	}

	if len(b.Checks) > 1 {
		var candidateMoves []board.Move
		moveDict := make(map[string]int) 
		for _, check := range b.Checks{
			candidateMoves = findLegalMovesForCheck(b, check, moves)
			for _, move := range candidateMoves{
				moveString := board.SquareHexToString[move.To] + board.SquareHexToString[move.From]
				if count, exists := moveDict[moveString]; !exists{
					moveDict[moveString] = 1
				}else{
					moveDict[moveString] = count + 1

					if count = moveDict[moveString]; count == len(b.Checks){
						legalMoves = append(legalMoves, move)
					}
				}
			}
		}
	}

	return legalMoves
}

func generatePawnMoves(b *board.Board, origin int8) []board.Move {
	var moves []board.Move
	if b.Turn == board.White{
		pin, isPinned := b.WhitePins[origin]
		if b.EnPassant != 0{
			// check left of black pawn
			if b.State[int8(b.EnPassant) - (nextRank + nextFile)] == board.WhitePawn && (int8(b.EnPassant) - (nextRank + nextFile) == origin){
				if isPinned && (pin.Delta == moveUpandRight || pin.Delta == -1*moveUpandRight) || !isPinned{
					move := createMove(int8(b.EnPassant) - (nextRank + nextFile), int8(b.EnPassant))
					move.MovedPiece = board.WhitePawn
					move.Type = moveEnPassant
					moves = append(moves, move)
				}
			}
			// check right of black pawn
			if b.State[int8(b.EnPassant) - (nextRank - nextFile)] == board.WhitePawn && (int8(b.EnPassant) - (nextRank - nextFile) == origin){
				if isPinned && (pin.Delta == moveUpandLeft || pin.Delta == -1*moveUpandLeft) || !isPinned{
					move := createMove(int8(b.EnPassant) - (nextRank - nextFile), int8(b.EnPassant))
					move.MovedPiece = board.WhitePawn
					move.Type = moveEnPassant
					moves = append(moves, move)
				}
			}
		}

		for _, delta := range whitePawnMoves{
			dest := origin + delta
			if isPinned && (pin.Delta == delta || pin.Delta == -1*delta) || !isPinned {
				if delta == moveUpandLeft || delta == moveUpandRight{
					validateAttacks := checkPawnAttacks(b, origin)
	
					if delta == moveUpandLeft && validateAttacks[0]{
						move := createMove(origin,dest)
						move.MovedPiece = board.WhitePawn
						move.Capture = b.State[dest]
						if board.Rank(dest) == 7{
							move.Type = movePromote
							promoMoves := generatePawnPromotions(b, move)
							moves = append(moves, promoMoves...)
						}else{
							moves = append(moves, move)
						}
						
					}
	
					if delta == moveUpandRight && validateAttacks[1]{
						move := createMove(origin,dest)
						move.MovedPiece = board.WhitePawn
						move.Capture = b.State[dest]
						if board.Rank(dest) == 7{
							move.Type = movePromote
							promoMoves := generatePawnPromotions(b, move)
							moves = append(moves, promoMoves...)
						}else{
							moves = append(moves, move)
						}
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
						if board.Rank(dest) == 7{
							move.Type = movePromote
							promoMoves := generatePawnPromotions(b, move)
							moves = append(moves, promoMoves...)
						}else{
							moves = append(moves, move)
						}
					}
				}
			}
		}
	}else{
		pin, isPinned := b.BlackPins[origin]
		if b.EnPassant != 0{
			// check left of white pawn
			if b.State[int8(b.EnPassant) + (nextRank - nextFile)] == board.BlackPawn && int8(b.EnPassant) + (nextRank - nextFile) == origin{
				if isPinned && pin.Delta == moveDownandRight || pin.Delta == -1*moveDownandRight || !isPinned{
					move := createMove(int8(b.EnPassant) + (nextRank - nextFile), int8(b.EnPassant))
					move.MovedPiece = board.BlackPawn
					move.Type = moveEnPassant
					moves = append(moves, move)
				}
			}
			// check right of white pawn
			if b.State[int8(b.EnPassant) + (nextRank + nextFile)] == board.BlackPawn && int8(b.EnPassant) + (nextRank + nextFile) == origin{
				if isPinned && (pin.Delta == moveDownandLeft || pin.Delta == -1*moveDownandLeft) || !isPinned{
					move := createMove(int8(b.EnPassant) + (nextRank + nextFile), int8(b.EnPassant))
					move.MovedPiece = board.WhitePawn
					move.Type = moveEnPassant
					moves = append(moves, move)
				}
			}
		}
		
		for _, delta := range blackPawnMoves{
			dest := origin + delta
			if isPinned && (pin.Delta == delta || pin.Delta == -1*delta) || !isPinned {
				if delta == moveDownandLeft || delta == moveDownandRight{
					validateAttacks := checkPawnAttacks(b, origin)
	
					if delta == moveDownandLeft && validateAttacks[0]{
						move := createMove(origin,dest)
						move.MovedPiece = board.BlackPawn
						move.Capture = b.State[dest]
						if board.Rank(dest) == 0{
							move.Type = movePromote
							promoMoves := generatePawnPromotions(b, move)
							moves = append(moves, promoMoves...)
						}else{
							moves = append(moves, move)
						}
						
					}
	
					if delta == moveDownandRight && validateAttacks[1]{
						move := createMove(origin,dest)
						move.MovedPiece = board.BlackPawn
						move.Capture = b.State[dest]
						if board.Rank(dest) == 0{
							move.Type = movePromote
							promoMoves := generatePawnPromotions(b, move)
							moves = append(moves, promoMoves...)
						}else{
							moves = append(moves, move)
						}
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
						if board.Rank(dest) == 0{
							move.Type = movePromote
							promoMoves := generatePawnPromotions(b, move)
							moves = append(moves, promoMoves...)
						}else{
							moves = append(moves, move)
						}
					}
				}
			}
		}
	}

	return moves
}

func generatePawnPromotions(b *board.Board, base board.Move) []board.Move {
	var moves []board.Move
	if b.Turn == board.White{
		promotions := []int8{board.WhiteQueen, board.WhiteRook, board.WhiteBishop, board.WhiteKnight}
		for _, option := range promotions{
			move := base
			move.Promotion = option
			moves = append(moves, move)
		}
	}else{
		promotions := []int8{board.BlackQueen, board.BlackRook, board.BlackBishop, board.BlackKnight}
		for _, option := range promotions{
			move := base
			move.Promotion = option
			moves = append(moves, move)
		}
	}

	return moves

}

func generateKnightMoves(b *board.Board, origin int8) []board.Move {
	var moves []board.Move

	var isPiecePinned bool
	if b.Turn == board.White{
		_, isPiecePinned = b.WhitePins[origin]
	}else{
		_, isPiecePinned = b.BlackPins[origin]
		
	}
	if isPiecePinned {
		return moves
	}

	for _, delta := range knightMoves{
		dest := origin + delta
		if board.LegalSquare(dest) && b.State[dest] <=0 && b.Turn == board.White{
			move := createMove(origin, dest)
			move.Capture = b.State[dest]
			move.MovedPiece = board.WhiteKnight
			moves = append(moves, move)
		}else if board.LegalSquare(dest) && b.State[dest] >=0 && b.Turn == board.Black{
			move := createMove(origin, dest)
			move.Capture = b.State[dest]
			move.MovedPiece = board.BlackKnight
			moves = append(moves, move)
		}		
	}

	return moves
}

func generateKingMoves(b *board.Board, origin int8) []board.Move {
	var moves []board.Move
	for _, delta := range kingMoves{
		dest := origin + delta
		if b.Turn == board.White{
			if attacks := b.BlackAttacks[dest]; len(attacks) > 0{
				continue
			}
			if delta == 2*moveRight || delta == 2*moveLeft {
				canCastle := checkCastlingAvailability(b)
				if canCastle[0] && delta == 2*moveRight{
					move := createMove(origin, dest)
					move.MovedPiece = board.WhiteKing
					move.Type = moveShortCastle
					moves = append(moves, move)
				}
				if canCastle[1] && delta == 2*moveLeft{
					move := createMove(origin, dest)
					move.MovedPiece = board.WhiteKing
					move.Type = moveLongCastle
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
			if attacks:= b.WhiteAttacks[dest]; len(attacks) > 0 {
				continue
			}
			if delta == 2*moveRight || delta == 2*moveLeft {
				canCastle := checkCastlingAvailability(b)
				if canCastle[0] && delta == 2*moveRight{
					move := createMove(origin, dest)
					move.MovedPiece = board.BlackKing
					move.Type = moveShortCastle
					moves = append(moves, move)
				}
				if canCastle[1] && delta == 2*moveLeft{
					move := createMove(origin, dest)
					move.MovedPiece = board.BlackKing
					move.Type = moveLongCastle
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

func generateQueenMoves(b *board.Board, origin int8) []board.Move {
	var moves []board.Move
	// check left moves
	moves = validateLongRangeMoves(origin,-nextFile, b, moves)
	// check right moves
	moves = validateLongRangeMoves(origin, nextFile,b,moves)
	// check up moves
	moves = validateLongRangeMoves(origin, nextRank,b,moves)
	// check down moves
	moves = validateLongRangeMoves(origin,-nextRank,b,moves)
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

func generateRookMoves(b *board.Board, origin int8) []board.Move {
	var moves []board.Move
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

func generateBishopMoves(b *board.Board, origin int8) []board.Move {
	var moves []board.Move
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

func generateAttacksList(b *board.Board){
	for _, square := range board.HexBoard{
		if b.State[square] != 0 {
			piece := b.State[square]

			if piece > 0 {
				switch piece {
				case board.WhitePawn:
					pawnPsuedoAttacks(b,square)
				case board.WhiteKnight:
					knightPsuedoAttacks(b,square)
				case board.WhiteBishop:
					bishopPsuedoAttacks(b, square)
				case board.WhiteRook:
					rookPsuedoAttacks(b,square)
				case board.WhiteQueen:
					queenPsuedoAttacks(b,square)
				case board.WhiteKing:
					kingPsuedoAttacks(b,square)
				}
			}

			if piece < 0 {
				switch piece {
				case board.BlackPawn:
					pawnPsuedoAttacks(b,square)
				case board.BlackKnight:
					knightPsuedoAttacks(b,square)
				case board.BlackBishop:
					bishopPsuedoAttacks(b, square)
				case board.BlackRook:
					rookPsuedoAttacks(b,square)
				case board.BlackQueen:
					queenPsuedoAttacks(b,square)
				case board.BlackKing:
					kingPsuedoAttacks(b,square)
				}
			}
		}
	}
}

