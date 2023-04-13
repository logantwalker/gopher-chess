package game

import (
	"bufio"
	"fmt"
	"os"

	"github.com/logantwalker/gopher-chess/board"
	"github.com/logantwalker/gopher-chess/moves"
)

type Game struct{
	board board.Board
}

func NewGame() Game {
	b := board.NewBoard(board.StartingFen)

	g := new(Game)

	g.board = b

	return *g
}

func (g *Game) Run(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("> ")

	for scanner.Scan() {
		input := scanner.Text()

		if input == "quit" || input == "q"{
			break
		}else if  m, err := moves.CreateMoveFromInput(input); err == nil{
			moves.MakeMove(&g.board,m)
			g.board.PrintBoard()
			fmt.Println("check: ", g.board.Check)
		}
		
		switch input {
		case "moves":
			m := moves.GenerateMovesList(&g.board)
			moves.PrintMoves(m)
		case "print":
			g.board.PrintBoard()
		}
		fmt.Printf("> ")
	}

}