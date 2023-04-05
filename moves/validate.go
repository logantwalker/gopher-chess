package moves

import (
	"errors"
	"fmt"

	"github.com/logantwalker/gopher-chess/board"
)

func ValidateUserMove(b *board.Board, move Move) (Move, error){
	validMoves := GenerateMovesList(b)

	for _, validMove := range validMoves{
		if validMove.To == move.To && validMove.From == move.From{
			return validMove, nil
		}
	}

	return Move{}, errors.New("invalid move")
}

func validateLongRangeMoves(origin int8, delta int8, b *board.Board, moves []Move) []Move {
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
		if b.State[origin + moveDownandLeft] > int8(0) {
			attacks[0] = true
		}
		if b.State[origin + moveDownandRight] > int8(0) {
			attacks[1] = true
		}
	}

	return attacks
}

func pawnPsuedoAttacks(b *board.Board, origin int8) {
	if b.Turn == board.White{
		if board.LegalSquare(origin + nextFile + nextRank) {
			if _,exists := b.WhiteAttacks[origin + nextFile + nextRank]; !exists {
				b.WhiteAttacks[origin + nextFile + nextRank] = true
			}
		}
		if board.LegalSquare(origin - nextFile + nextRank) {
			if _,exists := b.WhiteAttacks[origin - nextFile + nextRank]; !exists {
				b.WhiteAttacks[origin - nextFile + nextRank] = true
			}
		}
	}else{
		if board.LegalSquare(origin + nextFile - nextRank) {
			if _,exists := b.BlackAttacks[origin + nextFile - nextRank]; !exists {
				b.BlackAttacks[origin + nextFile - nextRank] = true
			}
		}
		if board.LegalSquare(origin - nextFile - nextRank) {
			if _,exists := b.BlackAttacks[origin - nextFile - nextRank]; !exists {
				b.BlackAttacks[origin - nextFile - nextRank] = true
			}
		}
	}
}

func knightPsuedoAttacks(b *board.Board, origin int8) {
	for _, move := range knightMoves{
		dest := origin + move 
		if board.LegalSquare(dest){
			if b.Turn == board.White{
				if _,exists := b.WhiteAttacks[dest]; !exists{
					b.WhiteAttacks[dest] = true
				}
			}else{
				if _,exists := b.BlackAttacks[dest]; !exists{
					b.BlackAttacks[dest] = true
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

		if board.LegalSquare(dest){
			if b.Turn == board.White{
				if _,exists := b.WhiteAttacks[dest]; !exists{
					b.WhiteAttacks[dest] = true
				}
			}else{
				if _,exists := b.BlackAttacks[dest]; !exists{
					b.BlackAttacks[dest] = true
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
				if b.Turn == board.White{
					if _,exists := b.WhiteAttacks[i]; !exists{
						b.WhiteAttacks[i] = true
					}
				}else{
					if _,exists := b.BlackAttacks[i]; !exists{
						b.BlackAttacks[i] = true
					}
				}
			}else{
				if b.Turn == board.White{
					if _,exists := b.WhiteAttacks[i]; !exists{
						b.WhiteAttacks[i] = true
					}
				}else{
					if _,exists := b.BlackAttacks[i]; !exists{
						b.BlackAttacks[i] = true
					}
				}
				searchingForPin = true
				pinLocation = i
			}
		}else{
			if (b.Turn == board.White && b.State[i] != board.Empty && b.State[i] != board.BlackKing){
				break
			}
			if (b.Turn == board.Black && b.State[i] != board.Empty && b.State[i] != board.WhiteKing){
				break
			}
			if b.Turn == board.White && b.State[i] == board.BlackKing{
				fmt.Println("pin detected")
				pin := board.Pin{Delta: delta,}
				b.BlackPins[pinLocation] = pin
				break
			}
			if b.Turn == board.Black && b.State[i] == board.WhiteKing{
				fmt.Println("pin detected")
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