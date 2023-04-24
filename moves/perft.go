package moves

import (
	"fmt"
	"time"

	"github.com/logantwalker/gopher-chess/board"
)

type PerftData struct {
	Depth      int
	Nodes      int64
	Captures   int64
	EnPassants int64
	Castles    int64
	Promotions int64
	Checks     int64
	Mates      int64
	Elapsed    time.Duration
}

var (
	Position1FEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

	Position1Table = []PerftData{
		{Depth: 0, Nodes: 1, Captures: 0, EnPassants: 0, Castles: 0, Promotions: 0, Checks: 0, Mates: 0},
		{Depth: 1, Nodes: 20, Captures: 0, EnPassants: 0, Castles: 0, Promotions: 0, Checks: 0, Mates: 0},
		{Depth: 2, Nodes: 400, Captures: 0, EnPassants: 0, Castles: 0, Promotions: 0, Checks: 0, Mates: 0},
		{Depth: 3, Nodes: 8902, Captures: 34, EnPassants: 0, Castles: 0, Promotions: 0, Checks: 12, Mates: 0},
		{Depth: 4, Nodes: 197281, Captures: 1576, EnPassants: 0, Castles: 0, Promotions: 0, Checks: 469, Mates: 8},
		{Depth: 5, Nodes: 4865609, Captures: 82719, EnPassants: 258, Castles: 0, Promotions: 0, Checks: 27351, Mates: 347},
		{Depth: 6, Nodes: 119060324, Captures: 2812008, EnPassants: 5248, Castles: 0, Promotions: 0, Checks: 809099, Mates: 10828},
	}

	Position2FEN = "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq -"

	Position2Table = []PerftData{
		{Depth: 0, Nodes: 1},
		{Depth: 1, Nodes: 48, Captures: 8, EnPassants: 0, Castles: 2, Promotions: 0, Checks: 0, Mates: 0},
		{Depth: 2, Nodes: 2039, Captures: 351, EnPassants: 1, Castles: 91, Promotions: 0, Checks: 3, Mates: 0},
		{Depth: 3, Nodes: 97862, Captures: 17102, EnPassants: 45, Castles: 3162, Promotions: 0, Checks: 993, Mates: 1},
		{Depth: 4, Nodes: 4085603, Captures: 757163, EnPassants: 1929, Castles: 128013, Promotions: 15172, Checks: 25523, Mates: 43},
		{Depth: 5, Nodes: 193690690, Captures: 35043416, EnPassants: 73365, Castles: 4993637, Promotions: 8392, Checks: 3309887, Mates: 30171},
	}
)

func Perft(fen string, expected []PerftData) {
	b := board.NewBoard(fen)
	printPerftData(&b, expected)
}

func perft(depth int, b *board.Board) PerftData {

	data := PerftData{Depth: depth}

	moves := GenerateMovesList(b)

	start := time.Now()

	if depth == 0 {
		// if len(moves) == 0 {
		// 	data.Mates++
		// }
	
		if b.IsCheck {
			data.Checks++
		}
		data.Depth = 0
		data.Nodes = 1
		return data
	}

	

	for _, move := range moves {
		switch move.Type {
		case moveShortCastle:
			data.Castles++
		case moveLongCastle:
			data.Castles++
		case movePromote:
			data.Promotions++
		case moveEnPassant:
			data.EnPassants++
		}

		if move.Capture != board.Empty {
			data.Captures++
		}

		if b.IsCheck {
			data.Checks++
		}

		MakeMove(b, move)
		res := perft(depth-1, b)
		UndoMove(b)

		data.Nodes += res.Nodes
		data.Captures += res.Captures
		data.EnPassants += res.EnPassants
		data.Castles += res.Castles
		data.Promotions += res.Promotions
		data.Checks += res.Checks
		data.Mates += res.Mates

		switch b.Status {
		case board.StatusCheckmate:
			data.Mates++
		}
	}

	data.Elapsed = time.Since(start)

	return data
}

func printPerftData(board *board.Board, expected []PerftData) {

	fmt.Printf("D   Nodes    Capt.   E.p.   Cast.   Prom.  Checks   Mates   Time\n")
	for i := 0; i < len(expected); i++ {

		res := perft(i, board)

		fmt.Printf("%d %7s %7s %7s %7s %7s %7s %7s  %5ss\n",
			i,
			formatNodesCount(res.Nodes),
			formatNodesCount(res.Captures),
			formatNodesCount(res.EnPassants),
			formatNodesCount(res.Castles),
			formatNodesCount(res.Promotions),
			formatNodesCount(res.Checks),
			formatNodesCount(res.Mates),
			formatDuration(res.Elapsed),
		)

		fmt.Printf("  %s %s %s %s %s %s %s\n\n",
			formatPerftEntry(res.Nodes, expected[i].Nodes),
			formatPerftEntry(res.Captures, expected[i].Captures),
			formatPerftEntry(res.EnPassants, expected[i].EnPassants),
			formatPerftEntry(res.Castles, expected[i].Castles),
			formatPerftEntry(res.Promotions, expected[i].Promotions),
			formatPerftEntry(res.Checks, expected[i].Checks),
			formatPerftEntry(res.Mates, expected[i].Mates))

	}

}
func formatNodesCount(nodes int64) string {
	if nodes < 1000 && nodes > -1000 {
		return fmt.Sprintf("%d", nodes)
	} else if nodes < 1000000 && nodes > -1000000 {
		return fmt.Sprintf("%.1fK", float64(nodes)/1000)
	}
	return fmt.Sprintf("%.2fM", float64(nodes)/1000000)
}

func formatPerftEntry(actual, expected int64) string {

	diff := actual - expected

	if diff != 0 {
		return formatNodesCount(diff)
	}

	return "      0"
}

func formatDuration(d time.Duration) string {
	return fmt.Sprintf("%.2f", float64(d)*float64(1e-9))
}