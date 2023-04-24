package moves

import "github.com/logantwalker/gopher-chess/board"

type MoveGenerator struct {
	Board						*board.Board
	Moves 						[]board.Move
	LastMoveSquare				int8
	KingSquare 					int8
	IsCheck 					bool
	IsCheckByKnight				int8
	legalSquaresWhileInCheck 	[]bool
	IsCheckmate					bool
}


func NewGenerator(board *board.Board) *MoveGenerator {
	mg := new(MoveGenerator)
	mg.Board = board

	return mg
}


func (mg *MoveGenerator) GenerateMoves() []board.Move{
	mg.reset()


}

func (mg *MoveGenerator) reset() {
	mg.Moves = make([]board.Move, 0, 48)

	mg.IsCheck = false
	mg.IsCheckByKnight = board.Empty

	mg.KingSquare = int8(mg.Board.KingLocations[0])
	if mg.Board.Turn == board.Black {
		mg.KingSquare = int8(mg.Board.KingLocations[1])
	}

	mg.legalSquaresWhileInCheck = make([]bool, boardSize)
	// mg.legalDelta = make([]int8, boardSize)

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
		if mg.Board.State[dest] * (-1*turn) == board.Knight && board.LegalSquare(dest){
			mg.IsCheck = true
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
						threats++
						mg.LegalEnding[dest] = true
					}
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
		return mg.hasDelta(allMoves, delta)
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


