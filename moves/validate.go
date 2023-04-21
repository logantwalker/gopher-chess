package moves

import (
	"errors"

	"github.com/logantwalker/gopher-chess/board"
)

func ValidateUserMove(b *board.Board, move board.Move) (board.Move, error){
	validMoves := GenerateMovesList(b)

	for _, validMove := range validMoves{
		if move.Type == movePromote{
			if validMove.To == move.To && validMove.From == move.From && validMove.Promotion == move.Promotion{
				return validMove, nil
			}
		}else{
			if validMove.To == move.To && validMove.From == move.From{
				return validMove, nil
			}
		}	
	}

	return board.Move{}, errors.New("invalid move")
}

func validateLongRangeMoves(origin int8, delta int8, b *board.Board, moves []board.Move) []board.Move {
	var isPiecePinned bool
	var pin board.Pin
	if b.Turn == board.White{
		pin, isPiecePinned = b.WhitePins[origin]
	}else{
		pin, isPiecePinned = b.BlackPins[origin]
	}

	if isPiecePinned && !(pin.Delta == delta || pin.Delta == -1 *delta) {
		return moves
	}

	for i := origin + delta; board.LegalSquare(i); i += delta{
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
		}else if b.State[i] > board.Empty && b.Turn == board.White{
			break
		}else if b.State[i] < board.Empty && b.Turn == board.Black{
			break
		}
	}

	return moves
}

func checkPawnAttacks(b *board.Board, origin int8) []bool{
	var attacks = []bool{false, false}
	if b.Turn == board.White{
		if b.State[origin + moveUpandLeft] < int8(0) {
			attacks[0] = true
		}
		if b.State[origin + moveUpandRight] < int8(0) {
			attacks[1] = true
		}
	}else{
		if board.LegalSquare(origin + moveDownandLeft) {
			if b.State[origin + moveDownandLeft] > int8(0) {
				attacks[0] = true
			}
			
		}
		if board.LegalSquare(origin + moveDownandRight) {
			if b.State[origin + moveDownandRight] > int8(0) {
				attacks[1] = true
			}
		}
	}

	return attacks
}

func pawnPsuedoAttacks(b *board.Board, origin int8) {
	if b.State[origin] > 0 {
		if board.LegalSquare(origin + nextFile + nextRank) {
			if _,exists := b.WhiteAttacks[origin + nextFile + nextRank]; !exists {
				b.WhiteAttacks[origin + nextFile + nextRank] = []int8{origin}
				if b.State[origin + nextFile + nextRank] == board.BlackKing {
					b.IsCheck = true
					b.Checks = append(b.Checks, &board.Check{AttackerOrigin: origin, AttackerDelta: nextFile + nextRank})
				}
			}else{
				b.WhiteAttacks[origin + nextFile + nextRank] = append(b.WhiteAttacks[origin+nextFile+nextRank], origin)
				if b.State[origin + nextFile + nextRank] == board.BlackKing {
					b.IsCheck = true
					b.Checks = append(b.Checks, &board.Check{AttackerOrigin: origin, AttackerDelta: nextFile + nextRank})
				}
			}
		}
		if board.LegalSquare(origin - nextFile + nextRank) {
			if _,exists := b.WhiteAttacks[origin - nextFile + nextRank]; !exists {
				b.WhiteAttacks[origin - nextFile + nextRank] = []int8{origin}
				if b.State[origin - nextFile + nextRank] == board.BlackKing{
					b.IsCheck = true
					b.Checks = append(b.Checks, &board.Check{AttackerOrigin: origin, AttackerDelta: nextRank - nextFile})
				}
			}else{
				b.WhiteAttacks[origin - nextFile + nextRank] = append(b.WhiteAttacks[origin-nextFile+nextRank], origin)
				if b.State[origin - nextFile + nextRank] == board.BlackKing{
					b.IsCheck = true
					b.Checks = append(b.Checks, &board.Check{AttackerOrigin: origin, AttackerDelta: nextRank - nextFile})
				}
			}
		}
	}else{
		if board.LegalSquare(origin + nextFile - nextRank) {
			if _,exists := b.BlackAttacks[origin + nextFile - nextRank]; !exists {
				b.BlackAttacks[origin + nextFile - nextRank] = []int8{origin}
				if b.State[origin + nextFile - nextRank] == board.WhiteKing{
					b.IsCheck = true
					b.Checks = append(b.Checks, &board.Check{AttackerOrigin: origin, AttackerDelta: nextFile - nextRank})
				}
			}else{
				b.BlackAttacks[origin + nextFile - nextRank] = append(b.BlackAttacks[origin+nextFile-nextRank], origin)
				if b.State[origin + nextFile - nextRank] == board.WhiteKing{
					b.IsCheck = true
					b.Checks = append(b.Checks, &board.Check{AttackerOrigin: origin, AttackerDelta: nextFile - nextRank})
				}
			}
		}
		if board.LegalSquare(origin - nextFile - nextRank) {
			if _,exists := b.BlackAttacks[origin - nextFile - nextRank]; !exists {
				b.BlackAttacks[origin - nextFile - nextRank] = []int8{origin}
				if b.State[origin - nextFile - nextRank] == board.WhiteKing{
					b.IsCheck = true
					b.Checks = append(b.Checks, &board.Check{AttackerOrigin: origin, AttackerDelta: 0 - nextFile - nextRank})
				}
			}else{
				b.BlackAttacks[origin - nextFile - nextRank] = append(b.BlackAttacks[origin-nextFile-nextRank], origin)
				if b.State[origin - nextFile - nextRank] == board.WhiteKing{
					b.IsCheck = true
					b.Checks = append(b.Checks, &board.Check{AttackerOrigin: origin, AttackerDelta: 0 - nextFile - nextRank})
				}
			}
		}
	}
}

func knightPsuedoAttacks(b *board.Board, origin int8) {
	for _, move := range knightMoves{
		dest := origin + move 
		if board.LegalSquare(dest){
			if b.State[origin] > 0 {
				if _,exists := b.WhiteAttacks[dest]; !exists{
					b.WhiteAttacks[dest] = []int8{origin}
					if b.State[dest] == board.BlackKing{
						b.IsCheck = true
						b.Checks = append(b.Checks, &board.Check{AttackerOrigin: origin, AttackerDelta: move})
					}
				}else{
					b.WhiteAttacks[dest] = append(b.WhiteAttacks[dest], origin)
					if b.State[dest] == board.BlackKing{
						b.IsCheck = true
						b.Checks = append(b.Checks, &board.Check{AttackerOrigin: origin, AttackerDelta: move})
					}
				}
			}else{
				if _,exists := b.BlackAttacks[dest]; !exists{
					b.BlackAttacks[dest] = []int8{origin}
					if b.State[dest] == board.WhiteKing{
						b.IsCheck = true
						b.Checks = append(b.Checks, &board.Check{AttackerOrigin: origin, AttackerDelta: move})
					}
				}else{
					b.BlackAttacks[dest] = append(b.BlackAttacks[dest], origin)
					if b.State[dest] == board.WhiteKing{
						b.IsCheck = true
						b.Checks = append(b.Checks, &board.Check{AttackerOrigin: origin, AttackerDelta: move})
					}
				}
			}
		}
	}
}

func bishopPsuedoAttacks(b *board.Board, origin int8){
	// move up left
	psuedoLongRangeAttacks(b,origin,nextRank - nextFile)
	// move up right
	psuedoLongRangeAttacks(b,origin,nextRank + nextFile)
	// move down left
	psuedoLongRangeAttacks(b,origin, -nextRank - nextFile)
	// move down right
	psuedoLongRangeAttacks(b,origin,nextFile - nextRank)
}

func rookPsuedoAttacks(b *board.Board, origin int8){
	// move up 
	psuedoLongRangeAttacks(b,origin,nextRank)
	// move right
	psuedoLongRangeAttacks(b,origin,nextFile)
	// move down
	psuedoLongRangeAttacks(b,origin,-nextRank)
	// move left
	psuedoLongRangeAttacks(b,origin,-nextFile)
}

func queenPsuedoAttacks(b *board.Board, origin int8){
	// move up 
	psuedoLongRangeAttacks(b,origin,nextRank)
	// move right
	psuedoLongRangeAttacks(b,origin,nextFile)
	// move down
	psuedoLongRangeAttacks(b,origin,-nextRank)
	// move left
	psuedoLongRangeAttacks(b,origin,-nextFile)
	// move up left
	psuedoLongRangeAttacks(b,origin,nextRank - nextFile)
	// move up right
	psuedoLongRangeAttacks(b,origin,nextRank + nextFile)
	// move down left
	psuedoLongRangeAttacks(b,origin, -nextRank - nextFile)
	// move down right
	psuedoLongRangeAttacks(b,origin,nextFile - nextRank)
}

func kingPsuedoAttacks(b *board.Board, origin int8){
	for _, move := range kingMoves{
		dest := origin + move
		piece := b.State[origin]
		if board.LegalSquare(dest){
			if piece == board.WhiteKing{
				if _,exists := b.WhiteAttacks[dest]; !exists{
					b.WhiteAttacks[dest] = []int8{origin}
				}else{
					b.WhiteAttacks[dest] = append(b.WhiteAttacks[dest], origin)
				}
			}else{
				if _,exists := b.BlackAttacks[dest]; !exists{
					b.BlackAttacks[dest] = []int8{origin}
				}else{
					b.BlackAttacks[dest] = append(b.BlackAttacks[dest], origin)
				}
			}
		}
	}
}

func psuedoLongRangeAttacks(b *board.Board, origin int8, delta int8){
	searchingForPin := false
	var pinLocation int8
	for i := origin + delta; board.LegalSquare(i); i += delta{
		if !searchingForPin{
			if b.State[i] == board.Empty{
				if b.State[origin] > 0{
					if _,exists := b.WhiteAttacks[i]; !exists{
						b.WhiteAttacks[i] = []int8{origin}
					}else{
						b.WhiteAttacks[i] = append(b.WhiteAttacks[i], origin)
					}
				}else{
					if _,exists := b.BlackAttacks[i]; !exists{
						b.BlackAttacks[i] = []int8{origin}
					}else{
						b.BlackAttacks[i] = append(b.BlackAttacks[i], origin)
					}
				}
			}else{
				if b.State[origin] > 0 {
					if _,exists := b.WhiteAttacks[i]; !exists{
						b.WhiteAttacks[i] = []int8{origin}
						if b.State[i] == board.BlackKing{
							b.IsCheck = true
							b.Checks = append(b.Checks, &board.Check{AttackerOrigin: origin, AttackerDelta: delta})
						}
					}else{
						b.WhiteAttacks[i] = append(b.WhiteAttacks[i], origin)
						if b.State[i] == board.BlackKing{
							b.IsCheck = true
							b.Checks = append(b.Checks, &board.Check{AttackerOrigin: origin, AttackerDelta: delta})
						}
					}
				}else{
					if _,exists := b.BlackAttacks[i]; !exists{
						b.BlackAttacks[i] = []int8{origin}
						if b.State[i] == board.WhiteKing{
							b.IsCheck = true
							b.Checks = append(b.Checks, &board.Check{AttackerOrigin: origin, AttackerDelta: delta})
						}
					}else{
						b.BlackAttacks[i] = append(b.BlackAttacks[i], origin)
						if b.State[i] == board.WhiteKing{
							b.IsCheck = true
							b.Checks = append(b.Checks, &board.Check{AttackerOrigin: origin, AttackerDelta: delta})
						}
					}
				}
				searchingForPin = true
				pinLocation = i
			}
		}else{
			if (b.State[origin] > 0 && b.State[i] != board.Empty && b.State[i] != board.BlackKing){
				break
			}
			if (b.State[origin] < 0 && b.State[i] != board.Empty && b.State[i] != board.WhiteKing){
				break
			}
			if b.State[origin] > 0 && b.State[i] == board.BlackKing{
				pin := board.Pin{Delta: delta,}
				b.BlackPins[pinLocation] = pin
				break
			}
			if b.State[origin] < 0 && b.State[i] == board.WhiteKing{
				pin := board.Pin{Delta: delta,}
				b.WhitePins[pinLocation] = pin
				break
			}
		}
	}
}

func checkCastlingAvailability(b *board.Board) []bool {
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
			_,sq1Attacked := b.BlackAttacks[int8(whiteKingSideCastlingSquares[0])]
			_,sq2Attacked := b.BlackAttacks[int8(whiteKingSideCastlingSquares[1])]

			if !sq1Attacked && !sq2Attacked{
				castleAbility[0] = true
			}
		}
		if castleRights[1] && b.State[whiteQueenSideCastlingSquares[0]] == board.Empty && b.State[whiteQueenSideCastlingSquares[1]] == board.Empty && b.State[whiteQueenSideCastlingSquares[2]] == board.Empty{
			_,sq1Attacked := b.BlackAttacks[int8(whiteQueenSideCastlingSquares[0])]
			_,sq2Attacked := b.BlackAttacks[int8(whiteQueenSideCastlingSquares[1])]
			_,sq3Attacked := b.BlackAttacks[int8(whiteQueenSideCastlingSquares[2])]

			if !sq1Attacked && !sq2Attacked && !sq3Attacked{
				castleAbility[1] = true
			}
		}
	}else{
		if (b.BlackCastle & 1) > 0 {
			castleRights[0] = true
		}
		if (b.BlackCastle & 2) > 0 {
			castleRights[1] = true
		}

		if castleRights[0] && b.State[blackKingSideCastlingSquares[0]] == board.Empty && b.State[blackKingSideCastlingSquares[1]] == board.Empty{
			_,sq1Attacked := b.WhiteAttacks[int8(blackKingSideCastlingSquares[0])]
			_,sq2Attacked := b.WhiteAttacks[int8(blackKingSideCastlingSquares[1])]

			if !sq1Attacked && !sq2Attacked{
				castleAbility[0] = true
			}
		}
		if castleRights[1] && b.State[blackQueenSideCastlingSquares[0]] == board.Empty && b.State[blackQueenSideCastlingSquares[1]] == board.Empty && b.State[blackQueenSideCastlingSquares[2]] == board.Empty{
			_,sq1Attacked := b.WhiteAttacks[int8(blackQueenSideCastlingSquares[0])]
			_,sq2Attacked := b.WhiteAttacks[int8(blackQueenSideCastlingSquares[1])]
			_,sq3Attacked := b.WhiteAttacks[int8(blackQueenSideCastlingSquares[2])]

			if !sq1Attacked && !sq2Attacked && !sq3Attacked{
				castleAbility[1] = true
			}
		}
	}

	return []bool{castleRights[0] && castleAbility[0],castleRights[1] && castleAbility[1],}
}

func findLegalMovesForCheck(b *board.Board, Check *board.Check, moves []board.Move) []board.Move{
	attacker := b.State[Check.AttackerOrigin]
	attackDelta := Check.AttackerDelta

	var legalMoves []board.Move
	if b.Turn == board.White {
		if attacker == board.BlackRook || attacker == board.BlackQueen || attacker == board.BlackBishop {
			blockSquares := findBlockingSquares(Check.AttackerOrigin, b.KingLocations[0], attackDelta)

			for _, sq := range blockSquares {
				for _, move := range moves{
					if int8(move.To) == sq{
						legalMoves = append(legalMoves, move)
					}
				}
			}
		}
	}else{
		if attacker == board.WhiteRook || attacker == board.WhiteQueen || attacker == board.WhiteBishop {
			blockSquares := findBlockingSquares(Check.AttackerOrigin, b.KingLocations[1], attackDelta)
			for _, sq := range blockSquares {
				for _, move := range moves{
					if int8(move.To) == sq{
						legalMoves = append(legalMoves, move)
					}
				}
			}
		}
	}

	//find captures
	for _, move := range moves{
		if int8(move.To) == Check.AttackerOrigin && move.Capture == attacker{
			legalMoves = append(legalMoves, move)
		}
	}

	//find legal king moves
	var kingMoves []board.Move
	if b.Turn == board.White{
		kingMoves = generateKingMoves(b, b.KingLocations[0])
	}else{
		kingMoves = generateKingMoves(b, b.KingLocations[1])
	}

	legalMoves = append(legalMoves, kingMoves...)
	return legalMoves
}

func findBlockingSquares(origin int8, dest int8, delta int8) []int8 {
	var blockingSquares []int8
	for sq := origin + delta; board.LegalSquare(sq); sq += delta {
		if sq == dest{
			break
		}
		blockingSquares = append(blockingSquares, sq)
	}

	return blockingSquares
}

func checkRepititions(b *board.Board) {
	r := 0
	first := len(b.History) - b.HalfMoveClock
	if first >= 0 {
		for i := first; i < len(b.History)-1; i++ {
			if b.History[i].ZobristHash == b.ZobristHash {
				r++
			}
		}
	}

	if r >= 3 {
		b.Status = board.StatusThreeFoldRep
	}
}