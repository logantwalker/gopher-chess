package moves

import (
	"fmt"

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
					moves = append(moves, move)
				}

				if delta == moveUpandRight && validateAttacks[1]{
					move := createMove(origin,dest)
					moves = append(moves, move)
				}
			}else if delta == 2*moveUp{
				if rank:= board.Rank(origin); rank == whitePawnStartRank && b.State[dest] == board.Empty && b.State[dest + moveDown] == board.Empty{
					move := createMove(origin,dest)
					moves = append(moves, move)
				}
			}else{
				if b.State[dest] == board.Empty{
					move := createMove(origin,dest)
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
					moves = append(moves, move)
				}

				if delta == moveDownandRight && validateAttacks[1]{
					move := createMove(origin,dest)
					moves = append(moves, move)
				}
			}else if delta == 2*moveDown{
				if rank:= board.Rank(origin); rank == blackPawnStartRank && b.State[dest] == board.Empty && b.State[dest + moveUp] == board.Empty{
					move := createMove(origin,dest)
					moves = append(moves, move)
				}
			}else{
				if b.State[dest] == board.Empty{
					move := createMove(origin,dest)
					moves = append(moves, move)
				}
			}
		}
	}

	return moves
}

func createMove(origin int8, dest int8) Move {
	var move Move

	if board.LegalSquare(dest){
		fmt.Println(board.SquareHexToString[board.Square(dest)])
		move = Move{From: board.Square(origin), To: board.Square(dest)}
	}
	

	return move
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

