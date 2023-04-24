package moves

import (
	"errors"
	"fmt"
	"log"
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
	// rookMoves = []int8{moveUp, moveDown, moveLeft, moveRight}
	// bishopMoves = []int8{moveUpandLeft, moveUpandRight, moveDownandLeft, moveDownandRight}
	// queenMoves = []int8{moveUp, moveDown, moveLeft, moveRight, moveUpandLeft, moveUpandRight, moveDownandLeft, moveDownandRight}
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

func CreateMoveFromInput(b *board.Board, input string) (board.Move, error) {
	input = strings.Trim(input, " ")
	var promotion string
	if len(input) == 5 {
		promotion = string(input[len(input) - 1])
		input = strings.TrimSuffix(input, promotion)
	}
	if m, _ := regexp.MatchString("^[a-h][1-8][a-h][1-8]$", input); !m {
		return board.Move{}, errors.New("invalid move")
	}

	from := board.SquareStringToHex[input[:2]]
	to := board.SquareStringToHex[input[2:]]

	var move board.Move = board.Move{From: from, To: to}
	if len(promotion) > 0 {
		move.Type = movePromote
		switch promotion {
		case "q":
			move.Promotion = b.Turn * board.Queen
		case "r":
			move.Promotion = b.Turn * board.Rook
		case "b":
			move.Promotion = b.Turn * board.Bishop
		case "n":
			move.Promotion = b.Turn * board.Knight
		}
	}
	return move, nil
}


func createMove(origin int8, dest int8) board.Move {
	var move board.Move

	if board.LegalSquare(dest){
		move = board.Move{From: board.Square(origin), To: board.Square(dest)}
	}
	
	return move
}

func PrintMoves(moves []board.Move) {
	for _, move := range moves {
		pieceSymbol := board.GetPieceSymbol(move.MovedPiece)
		moveString := board.SquareHexToString[move.From] + board.SquareHexToString[move.To]

		if move.Capture != board.Empty{
			captureSymbol := board.GetPieceSymbol(move.Capture)
			fmt.Println(pieceSymbol + " " + moveString + " " + captureSymbol)
			return
		}
		fmt.Println(pieceSymbol + " " + moveString)
	}
} 

func MakeMove(b *board.Board, move board.Move) *board.Board{
	b.HalfMoveClock++

	if b.Turn == board.Black{
		b.FullMoveClock++
	}

	// validMove, err := ValidateUserMove(b, move)
	// if err != nil {
	// 	invalidMoveString := board.SquareHexToString[move.From] + board.SquareHexToString[move.To]
	// 	b.PrintBoard()
	// 	log.Println("legal moves: ")
	// 	m := GenerateMovesList(b)
	// 	PrintMoves(m)
	// 	fmt.Println("board state: ", b.BlackPins)
	// 	log.Fatalf("%s. move: %s, turn: %d\n", err.Error(), invalidMoveString, b.Turn)
	// 	return b
	// }

	validMove := move

	moveRecord := board.MoveRecord{
		Move: validMove,
		WhiteCastle: b.WhiteCastle,
		BlackCastle: b.BlackCastle,
		EnPassant: b.EnPassant,
		HalfMoveClock: b.HalfMoveClock,
		ZobristHash: b.ZobristHash,
	}

	b.History = append(b.History, moveRecord)

	switch validMove.Type {
	case moveOrdinary:
		b.State[validMove.From] = board.Empty
		b.State[validMove.To] = validMove.MovedPiece
		if move.Capture != board.Empty{
			b.HalfMoveClock = 0
		}
		switch validMove.MovedPiece {
		case board.WhiteKing:
			if validMove.From == board.WhiteKingStartSquare{
				b.WhiteCastle = board.CastleNone
			}
			b.KingLocations[0] = int8(validMove.To)
		case board.BlackKing:
			if validMove.From == board.BlackKingStartSquare{
				b.BlackCastle = board.CastleNone
			}
			b.KingLocations[1] = int8(validMove.To)
		case board.WhiteRook:
			if validMove.From == board.WhiteRookStartSquares[0]{
				b.WhiteCastle &= ^board.CastleLong
			}else if validMove.From == board.WhiteRookStartSquares[1]{
				b.WhiteCastle &= ^board.CastleShort
			}
		case board.BlackRook:
			if validMove.From == board.BlackRookStartSquares[0]{
				b.BlackCastle &= ^board.CastleLong
			}else if validMove.From == board.BlackRookStartSquares[1]{
				b.BlackCastle &= ^board.CastleShort
			}
		case board.WhitePawn:
			b.HalfMoveClock = 0
			delta := board.Rank(int8(move.To)) - board.Rank(int8(move.From))

			if delta > 1 && (b.State[move.To + board.Square(nextFile)] == board.BlackPawn || b.State[move.To + board.Square(-nextFile)] == board.BlackPawn){
				b.EnPassant = move.To - board.Square(nextRank)
			}
		case board.BlackPawn:
			b.HalfMoveClock = 0
			delta := board.Rank(int8(move.From)) - board.Rank(int8(move.To))

			if delta > 1 && (b.State[move.To + board.Square(nextFile)] == board.WhitePawn || b.State[move.To + board.Square(-nextFile)] == board.WhitePawn){
				b.EnPassant = move.From - board.Square(nextRank)
			}
		}
	case moveEnPassant:
		switch validMove.MovedPiece {
		case board.WhitePawn:
			b.State[validMove.From] = board.Empty
			b.State[validMove.To] = validMove.MovedPiece
			b.State[int8(b.EnPassant) - nextRank] = board.Empty

			b.HalfMoveClock = 0
		case board.BlackPawn:
			b.State[validMove.From] = board.Empty
			b.State[validMove.To] = validMove.MovedPiece
			b.State[int8(b.EnPassant) + nextRank] = board.Empty

			b.HalfMoveClock = 0
		}

		b.EnPassant = 0
	case moveShortCastle:
		b.State[validMove.From] = board.Empty
		b.State[validMove.To] = validMove.MovedPiece

		switch validMove.MovedPiece {
		case board.WhiteKing:
			b.State[board.WhiteRookStartSquares[1]] = board.Empty
			b.State[int8(validMove.To) - nextFile] = board.WhiteRook
			b.WhiteCastle &= ^board.CastleShort
			b.KingLocations[0] = int8(validMove.To)
		case board.BlackKing:
			b.State[board.BlackRookStartSquares[1]] = board.Empty
			b.State[int8(validMove.To) - nextFile] = board.BlackRook
			b.BlackCastle &= ^board.CastleShort	
			b.KingLocations[1] = int8(validMove.To)
		}
	case moveLongCastle:
		b.State[validMove.From] = board.Empty
		b.State[validMove.To] = validMove.MovedPiece
		switch validMove.MovedPiece {
		case board.WhiteKing:
			b.State[board.WhiteRookStartSquares[0]] = board.Empty
			b.State[int8(validMove.To) + nextFile] = board.WhiteRook
			b.WhiteCastle &= ^board.CastleLong
			b.KingLocations[0] = int8(validMove.To)
		case board.BlackKing:	
			b.State[board.BlackRookStartSquares[0]] = board.Empty
			b.State[int8(validMove.To) + nextFile] = board.BlackRook
			b.BlackCastle &= ^board.CastleLong
			b.KingLocations[1] = int8(validMove.To)
		}
	case movePromote:
		b.State[validMove.From] = board.Empty
		b.State[validMove.To] = validMove.Promotion

		b.HalfMoveClock = 0
	}

	resetMap := make(map[int8][]int8)
	if b.Turn == board.White{
		b.WhiteAttacks = resetMap
	}else{
		b.BlackAttacks = resetMap
	}

	if b.IsCheck && b.Status != board.StatusCheckmate{
		b.IsCheck = false
		b.Checks = []*board.Check{}
	}

	b.UpdateHash(&validMove)

	checkRepititions(b)

	generateAttacksList(b)

	b.Ply ++
	b.Turn = -1 * b.Turn

	if b.IsCheck {
		GenerateMovesList(b)
	}

	return b
}

func UndoMove(b *board.Board) {
	if len(b.History) < 1 {
		b.PrintBoard()
		log.Fatal("could not undo move")
		return
	}

	moveRecord := b.History[len(b.History) - 1]
	b.History = b.History[0 : len(b.History) - 1]

	b.WhiteCastle = moveRecord.WhiteCastle
	b.BlackCastle = moveRecord.BlackCastle
	b.EnPassant = moveRecord.EnPassant
	b.HalfMoveClock = moveRecord.HalfMoveClock
	b.ZobristHash =  moveRecord.ZobristHash

	move := moveRecord.Move

	switch move.Type {
	case moveOrdinary:
		b.State[move.From] = move.MovedPiece
		b.State[move.To] = move.Capture

		switch move.MovedPiece {
		case board.WhiteKing:
			b.KingLocations[0] = int8(move.From)
		case board.BlackKing:
			b.KingLocations[1] = int8(move.From)
		}
	case moveEnPassant:
		switch move.MovedPiece {
		case board.WhitePawn:
			b.State[move.From] = move.MovedPiece
			b.State[move.To] = board.Empty
			b.State[int8(b.EnPassant) - nextRank] = move.Capture
		case board.BlackPawn:
			b.State[move.From] = move.MovedPiece
			b.State[move.To] = board.Empty
			b.State[int8(b.EnPassant) + nextRank] = move.Capture
		}
	case moveShortCastle:
		b.State[move.From] = move.MovedPiece
		b.State[move.To] = board.Empty

		switch move.MovedPiece {
		case board.WhiteKing:
			b.State[board.WhiteRookStartSquares[1]] = board.WhiteRook
			b.State[int8(move.To) - nextFile] = board.Empty
			b.KingLocations[0] = int8(move.From)
		case board.BlackKing:
			b.State[board.BlackRookStartSquares[1]] = board.BlackRook
			b.State[int8(move.To) - nextFile] = board.Empty
			b.KingLocations[1] = int8(move.From)
		}
	case moveLongCastle:
		b.State[move.From] = move.MovedPiece
		b.State[move.To] = board.Empty
		switch move.MovedPiece {
		case board.WhiteKing:
			b.State[board.WhiteRookStartSquares[0]] = board.WhiteRook
			b.State[int8(move.To) + nextFile] = board.Empty 
			b.KingLocations[0] = int8(move.From)
		case board.BlackKing:	
			b.State[board.BlackRookStartSquares[0]] = board.BlackRook
			b.State[int8(move.To) + nextFile] = board.Empty 
			b.KingLocations[1] = int8(move.From)
		}
	case movePromote:
		b.State[move.From] = move.MovedPiece 
		b.State[move.To] = move.Capture
	}

	b.WhitePins = map[int8]board.Pin{}
	b.BlackPins = map[int8]board.Pin{}

	resetMap := make(map[int8][]int8)
	if b.Turn == board.White{
		b.WhiteAttacks = resetMap
	}else{
		b.BlackAttacks = resetMap
	}

	b.IsCheck = false
	b.Checks = nil
	b.Ply--
	b.Turn = -1 * b.Turn

	generateAttacksList(b)

	if b.Turn == board.Black {
		b.FullMoveClock--
	}

	if b.IsCheck {
		GenerateMovesList(b)
	}
}