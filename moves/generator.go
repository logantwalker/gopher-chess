package moves

import "github.com/logantwalker/gopher-chess/board"

type MoveGenerator struct {
	Board						*board.Board
	Moves 						[]board.Move
	LastMoveSquare				int8
	KingSquare 					int8
	IsCheck 					bool
	IsCheckByKnight				int8
	LegalSquaresWhileInCheck 	[]bool
	LegalDelta					[]int8
	IsCheckmate					bool
}


func NewGenerator(board *board.Board) *MoveGenerator {
	mg := new(MoveGenerator)
	mg.Board = board

	return mg
}


func (mg *MoveGenerator) GenerateMoves() []board.Move{
	mg.reset()

	checks := mg.findNumberofChecks()

	mg.generateKingMoves(mg.KingSquare)

	if checks > 1 {
		// need to sort the moves to speed up engine eval
		mg.Board.Status = board.StatusCheck
		return mg.Moves
	}

	if checks == 1 && mg.IsCheckByKnight != board.Empty{
		mg.Board.Status = board.StatusCheck
		mg.generateKnightCaptures(mg.IsCheckByKnight)
		// need to sort moves
		return mg.Moves
	}

	if checks == 0 {
		mg.generateCastleMoves()
	}

	// generate moves for the pieces
	for rank := int8(0); rank < rankSize; rank++ {
		for file := int8(0); file < rankSize; file++ {
			sq := board.SquareRF(rank, file)
			piece := mg.Board.State[sq] * mg.Board.Turn

			if piece > 0{
				switch piece {
				case board.Pawn:
					mg.generatePawnMoves(sq)
				case board.Rook:
					mg.generateMoves(sq, rookMoves, true)
				case board.Bishop:
					mg.generateMoves(sq, bishopMoves, true)
				case board.Queen:
					mg.generateMoves(sq, queenMoves, true)
				case board.Knight:
					mg.generateMoves(sq, knightMoves, false)
				}
			}
		}
	}

	if len(mg.Moves) == 0{
		mg.Board.Status = board.StatusCheckmate
	}
	return mg.Moves
}

func (mg *MoveGenerator) reset() {
	mg.Moves = make([]board.Move, 0, 48)

	mg.IsCheck = false
	mg.IsCheckByKnight = board.Empty

	mg.KingSquare = int8(mg.Board.KingLocations[0])
	if mg.Board.Turn == board.Black {
		mg.KingSquare = int8(mg.Board.KingLocations[1])
	}

	mg.LegalSquaresWhileInCheck = make([]bool, boardSize)
	mg.LegalDelta = make([]int8, boardSize)

	if len(mg.Board.History) > 0 {
		mg.LastMoveSquare = int8(mg.Board.History[len(mg.Board.History)-1].Move.To)
	}
}

func (mg *MoveGenerator) generateKingMoves(origin int8) {
	dest := int8(0)
	for _, delta := range kingMoves {
		dest = origin + delta
		if board.LegalSquare(dest) {
			move := createMove(origin, dest)
			if move.MovedPiece == board.Empty {
				continue
			}

			if len(mg.findThreats(dest, mg.Board.Turn, true)) == 0 {
				mg.Moves = append(mg.Moves, move)
			}
		}
	}

}

func (mg * MoveGenerator) generateCastleMoves(){
	switch mg.Board.Turn {
	case board.White:
		if mg.canCastle(mg.Board.Turn, moveShortCastle){
			mg.Moves = append(mg.Moves,board.Move{From: board.E1, To: board.G1, Capture: board.Empty, MovedPiece: board.WhiteKing, Type: moveShortCastle})
		}
		if mg.canCastle(mg.Board.Turn, moveLongCastle){
			mg.Moves = append(mg.Moves, board.Move{From: board.E1, To: board.C1, Capture: board.Empty, MovedPiece: board.WhiteKing, Type: moveLongCastle})
		}
	case board.Black:
		if mg.canCastle(mg.Board.Turn, moveShortCastle){
			mg.Moves = append(mg.Moves,board.Move{From: board.E8, To: board.G8, Capture: board.Empty, MovedPiece: board.BlackKing, Type: moveShortCastle})
		}
		if mg.canCastle(mg.Board.Turn, moveLongCastle){
			mg.Moves = append(mg.Moves, board.Move{From: board.E8, To: board.C8, Capture: board.Empty, MovedPiece: board.BlackKing, Type: moveLongCastle})
		}
	}
}

func (mg *MoveGenerator) canCastle(turn int8, move int8) bool {
	switch turn {
	case board.White:
		if mg.Board.WhiteCastle&move == move{
			if move == moveShortCastle{
				return mg.Board.IsEmpty(whiteShortCastlingSquares...) && len(mg.findThreats(int8(whiteShortCastlingSquares[0]),turn, false)) == 0 && len(mg.findThreats(int8(whiteShortCastlingSquares[1]),turn, false)) == 0 
			}

			return mg.Board.IsEmpty(whiteLongCastlingSquares...) && len(mg.findThreats(int8(whiteLongCastlingSquares[0]),turn, false)) == 0 && len(mg.findThreats(int8(whiteLongCastlingSquares[1]),turn, false)) == 0 && len(mg.findThreats(int8(whiteLongCastlingSquares[2]),turn, false)) == 0
		}

	case board.Black:
		if mg.Board.BlackCastle&move == move{
			if move == moveShortCastle{
				return mg.Board.IsEmpty(blackShortCastlingSquares...) && len(mg.findThreats(int8(blackShortCastlingSquares[0]),turn, false)) == 0 && len(mg.findThreats(int8(blackShortCastlingSquares[1]),turn, false)) == 0 
			}

			return mg.Board.IsEmpty(blackLongCastlingSquares...) && len(mg.findThreats(int8(blackLongCastlingSquares[0]),turn, false)) == 0 && len(mg.findThreats(int8(blackLongCastlingSquares[1]),turn, false)) == 0 && len(mg.findThreats(int8(blackLongCastlingSquares[2]),turn, false)) == 0
		}
	}
	return false
}

func (mg *MoveGenerator) generatePawnMoves(origin int8){
	startPos := false
	delta := whitePawnMoves[1:]

	switch mg.Board.State[origin] {
	case board.WhitePawn:
		startPos = board.Rank(origin) == whitePawnStartRank
	case board.BlackPawn:
		startPos = board.Rank(origin) == blackPawnStartRank
		delta = blackPawnMoves[1:]
	}

	for _, move := range delta{
		dest := origin + move

		if board.LegalSquare(dest) {
			if mg.LegalDelta[origin] == 0 || mg.LegalDelta[origin] == move || mg.LegalDelta[origin] == -1*move {
				mg.generatePawnMove(origin , dest, startPos)
			}
		}
	}
}

func (mg *MoveGenerator) generatePawnMove(origin int8 , dest int8, startPos bool) {
	move := board.Move{}

	move.MovedPiece = mg.Board.State[origin]
	move.Capture = mg.Board.State[dest]
	move.From = board.Square(origin)
	move.To = board.Square(dest)
	move.Promotion = board.Empty

	enPassantRemovesThreat := false

	if mg.IsCheck && mg.Board.EnPassant == board.Square(dest) {
		enPassantRemovesThreat = mg.LegalSquaresWhileInCheck[dest+(moveDown*mg.Board.Turn)]
	}

	if !mg.IsCheck || mg.LegalSquaresWhileInCheck[dest] || enPassantRemovesThreat {
		// attacks & enPassant
		if board.File(origin) != board.File(dest) {
			if move.To == mg.Board.EnPassant{
				move.Type = moveEnPassant
				move.Capture = -move.MovedPiece
			}else if mg.Board.State[origin]*mg.Board.State[dest] >= 0 {
				// must be opposite pawn
				return
			}
		} else if mg.Board.State[dest] != board.Empty {
			// moving forward requires an empty field
			return
		}

		// promotions
		if board.Rank(dest)%7 == 0 {
			move.Promotion = move.MovedPiece * board.Queen
			move.Type = movePromote
			mg.Moves = append(mg.Moves, move)

			move.Promotion = move.MovedPiece * board.Rook
			move.Type = movePromote
			mg.Moves = append(mg.Moves, move)

		} else {
			mg.Moves = append(mg.Moves, move)
		}
	}

	if startPos && board.File(origin) == board.File(dest) && mg.Board.State[dest] == board.Empty {
		move.To = board.Square(dest + dest - origin)
		if mg.Board.State[int8(move.To)] == board.Empty && (!mg.IsCheck || mg.LegalSquaresWhileInCheck[move.To]){
			mg.Moves = append(mg.Moves, move)
		}
	}
}

func (mg * MoveGenerator) generateMoves(origin int8, moves []int8, longRange bool){
	for _,delta := range moves {
		for dest := origin + delta; board.LegalSquare(dest); dest += delta{
			move := mg.createMove(origin, dest)
			
			if move.MovedPiece == board.Empty{
				break
			}

			// if there's no restriction on the direction the piece can move or if the direction matches the allowed direction to move,
			// it checks whether the move is legal considering the king's check status.
			if mg.LegalDelta[origin] == 0 || mg.LegalDelta[origin] == delta || mg.LegalDelta[origin] == -1*delta{
				if !mg.IsCheck || mg.LegalSquaresWhileInCheck[dest] {
					mg.Moves = append(mg.Moves, move)
				}
			}else{
				break
			}

			if !longRange || move.Capture != board.Empty{
				break
			}
		}
	}
}

func (mg *MoveGenerator) generateKnightCaptures(target int8) {
	attacks := mg.findThreats(target, -1*mg.Board.Turn, false)

	if len(attacks) > 0 {
		for _, atk := range attacks{
			mg.Moves = append(mg.Moves, mg.createMove(atk, target))
		}
	}
}

func (mg *MoveGenerator) createMove(origin int8, dest int8) board.Move{
	move := board.Move{From: board.Square(origin), To: board.Square(dest)}
	move.Promotion = board.Empty
	move.Capture = mg.Board.State[dest]

	if mg.Board.State[dest]*mg.Board.State[origin] > 0 {
		move.MovedPiece = board.Empty
		return move
	}

	move.MovedPiece = mg.Board.State[origin]

	return move
}

func (mg * MoveGenerator) findNumberofChecks() int {
	turn := mg.Board.Turn
	threats := 0
	depth := 0
	squareOfGuardingPiece := board.Inv


	// looking for checks by knight
	for _,delta := range knightMoves{
		dest := mg.KingSquare + delta
		if board.LegalSquare(dest) && mg.Board.State[dest] * (-1*turn) == board.Knight {
			mg.IsCheck = true
			mg.Board.IsCheck = true
			mg.IsCheckByKnight = dest
			threats++ 
			break
		}
	}

	for _, delta := range allMoves{
		depth = 0
		squareOfGuardingPiece = board.Inv

		for dest := mg.KingSquare + delta; board.LegalSquare(dest); dest += delta{
			depth++

			piece := (-1* turn)*mg.Board.State[dest]

			if piece == board.Empty{
				continue
			}

			if piece == board.Pawn || piece == board.King{
				if depth == 1{
					if mg.attackPossible(dest, -1 * delta){
						mg.IsCheck = true
						mg.Board.IsCheck = true
						threats++
						mg.LegalSquaresWhileInCheck[dest] = true
					}
				}
			}

			// piece < 0 indicates friendly piece. piece may be able to protect king
			if piece < 0 {
				if squareOfGuardingPiece == board.Inv {
					squareOfGuardingPiece = board.Square(dest)
					continue
				} else {
					break
				}
			}

			// enemy piece detected 
			if mg.attackPossible(dest, delta) {
				if squareOfGuardingPiece == board.Inv{
					threats++
					mg.IsCheck = true
					mg.Board.IsCheck = true
					mg.LegalSquaresWhileInCheck[dest] = true
					for i := mg.KingSquare + delta; i != dest; i += delta {
						mg.LegalSquaresWhileInCheck[i] = true
					}
				}else{
					mg.LegalDelta[int8(squareOfGuardingPiece)] = delta
				}
			} 
		}
	}

	return threats
}

func (mg *MoveGenerator) findThreats(square int8, turn int8, ignoreKing bool) []int8{
	threats := []int8{}

	//calculate knight attacks
	for _, delta := range knightMoves{
		dest := square + delta
		if board.LegalSquare(dest) && (-1*turn * mg.Board.State[dest]) == board.Knight{
			threats = append(threats, dest)
		}
	}

	//calculate threats from all other attacker-types
	for _, delta := range allMoves{
		depth := 0

		for dest := square + delta; board.LegalSquare(dest); dest += delta{
			depth++
			squareContent := mg.Board.State[dest] * mg.Board.Turn

			if squareContent == -board.King || squareContent == -board.Pawn{
				if depth == 1{
					if mg.attackPossible(dest, -1*delta) {
						threats = append(threats, dest)
					}
				}
			}

			// ignore all empty squares
			if squareContent == board.Empty {
				continue
			}

			if ignoreKing && squareContent == board.King {
				continue
			}

			// own piece
			if squareContent > 0 {
				break
			}

			// found opponent piece
			if mg.attackPossible(dest, delta) {
				threats = append(threats, dest)
			}

			break
		}
	}

	return threats
}

// checking attacks from the given square in the given direction
func (mg *MoveGenerator) attackPossible(attkOrigin int8, delta int8) bool {
	piece := mg.Board.State[attkOrigin]
	if piece < 0{
		piece = -1* piece
	}

	switch piece {
	case board.King:
		return mg.hasDelta(kingMoves, delta)
	case board.Queen:
		return mg.hasDelta(queenMoves, delta)
	case board.Rook:
		return mg.hasDelta(rookMoves, delta)
	case board.Bishop:
		return mg.hasDelta(bishopMoves, delta)
	case board.Pawn:
		switch mg.Board.State[attkOrigin] {
		case board.WhitePawn:
			return mg.hasDelta(whitePawnMoves[1:], delta)
		case board.BlackPawn:
			return mg.hasDelta(blackPawnMoves[1:], delta)
		}

	}

	return false
}

func (mg *MoveGenerator) hasDelta(delta []int8, direction int8) bool {
	for _, d := range delta {
		if d == direction {
			return true
		}
	}

	return false
}


