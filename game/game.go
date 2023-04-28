package game

import (
	"bufio"
	"fmt"
	"os"

	"github.com/logantwalker/gopher-chess/board"
	"github.com/logantwalker/gopher-chess/moves"
)

type Game struct{
	board *board.Board
}

func NewGame() Game {
	b := board.NewBoard(board.StartingFen)

	g := new(Game)

	g.board = &b

	return *g
}

func (g *Game) Run(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("> ")

	for scanner.Scan() {
		input := scanner.Text()

		if input == "quit" || input == "q"{
			break
		}else if  m, err := moves.CreateMoveFromInput(g.board, input); err == nil{
			m, err = moves.ValidateUserMove(g.board, m)
			if err != nil{
				fmt.Println("> invalid move")
			}else{
				moves.MakeMove(g.board,m)
			}
			
			g.board.PrintBoard()
		}
		
		switch input {
		case "moves":
			mg := moves.NewGenerator(g.board)
			m := mg.GenerateMoves()
			moves.PrintMoves(m)
		case "undo":
			moves.UndoMove(g.board)
		case "print":
			g.board.PrintBoard()
		case "perft":
			moves.Perft(moves.Position1FEN,moves.Position1Table)
		}
		fmt.Printf("> ")
	}

}