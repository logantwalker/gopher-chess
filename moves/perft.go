package moves

import (
	"fmt"
	"time"

	"github.com/logantwalker/gopher-chess/board"
)

var (
	Position1FEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

	Position1Table = []PerftData{
		{depth: 0, nodes: 1, captures: 0, enPassants: 0, castles: 0, promotions: 0, checks: 0, mates: 0},
		{depth: 1, nodes: 20, captures: 0, enPassants: 0, castles: 0, promotions: 0, checks: 0, mates: 0},
		{depth: 2, nodes: 400, captures: 0, enPassants: 0, castles: 0, promotions: 0, checks: 0, mates: 0},
		{depth: 3, nodes: 8902, captures: 34, enPassants: 0, castles: 0, promotions: 0, checks: 12, mates: 0},
		{depth: 4, nodes: 197281, captures: 1576, enPassants: 0, castles: 0, promotions: 0, checks: 469, mates: 8},
		{depth: 5, nodes: 4865609, captures: 82719, enPassants: 258, castles: 0, promotions: 0, checks: 27351, mates: 347},
		{depth: 6, nodes: 119060324, captures: 2812008, enPassants: 5248, castles: 0, promotions: 0, checks: 809099, mates: 10828},
	}

	Position2FEN = "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq -"

	Position2Table = []PerftData{
		{depth: 0, nodes: 1},
		{depth: 1, nodes: 48, captures: 8, enPassants: 0, castles: 2, promotions: 0, checks: 0, mates: 0},
		{depth: 2, nodes: 2039, captures: 351, enPassants: 1, castles: 91, promotions: 0, checks: 3, mates: 0},
		{depth: 3, nodes: 97862, captures: 17102, enPassants: 45, castles: 3162, promotions: 0, checks: 993, mates: 1},
		{depth: 4, nodes: 4085603, captures: 757163, enPassants: 1929, castles: 128013, promotions: 15172, checks: 25523, mates: 43},
		{depth: 5, nodes: 193690690, captures: 35043416, enPassants: 73365, castles: 4993637, promotions: 8392, checks: 3309887, mates: 30171},
	}
)

// PerftData aggregates performance test data in a structure
type PerftData struct {
	depth      int
	nodes      int64
	captures   int64
	enPassants int64
	castles    int64
	promotions int64
	checks     int64
	mates      int64
	elapsed    time.Duration
}

// Perft runs a performance test against a given FEN and expected results
func Perft(fen string, expected []PerftData) {
	b := board.NewBoard(fen)
	printPerftData(&b, expected)
}

func perft(depth int, b *board.Board) PerftData {

	data := PerftData{depth: depth}
	generator := NewGenerator(b)

	start := time.Now()

	if depth == 0 {
		data.depth = 0
		data.nodes = 1
		return data
	}

	moves := generator.GenerateMoves()

	if len(moves) == 0 {
		data.mates++
	}

	if generator.IsCheck {
		data.checks++
	}

	for _, move := range moves {
		MakeMove(b, move)

		res := perft(depth-1, b)
		data.nodes += res.nodes
		data.captures += res.captures
		data.enPassants += res.enPassants
		data.castles += res.castles
		data.promotions += res.promotions
		data.checks += res.checks
		data.mates += res.mates

		switch move.Type {
		case moveShortCastle:
			data.castles++
		case moveLongCastle:
			data.castles++
		case movePromote:
			data.promotions++
		case moveEnPassant:
			data.enPassants++
		}

		if move.Capture != board.Empty {
			data.captures++
		}

		switch b.Status {
		case board.StatusCheckmate:
			data.mates++
		case board.StatusCheck:
			data.checks++
		}

		UndoMove(b)
	}

	data.elapsed = time.Since(start)

	return data
}

func printPerftData(b *board.Board, expected []PerftData) {
	

	fmt.Printf(("D   Nodes    Capt.   E.p.   Cast.   Prom.  Checks   Mates   Time\n"))
	for i := 0; i < len(expected); i++ {

		res := perft(i, b)

		fmt.Printf("%d %7s %7s %7s %7s %7s %7s %7s  %5ss\n",
			i,
			formatNodesCount(res.nodes),
			formatNodesCount(res.captures),
			formatNodesCount(res.enPassants),
			formatNodesCount(res.castles),
			formatNodesCount(res.promotions),
			formatNodesCount(res.checks),
			formatNodesCount(res.mates),
			formatDuration(res.elapsed),
		)

		fmt.Printf("  %s %s %s %s %s %s %s\n\n",
			formatPerftEntry(res.nodes, expected[i].nodes),
			formatPerftEntry(res.captures, expected[i].captures),
			formatPerftEntry(res.enPassants, expected[i].enPassants),
			formatPerftEntry(res.castles, expected[i].castles),
			formatPerftEntry(res.promotions, expected[i].promotions),
			formatPerftEntry(res.checks, expected[i].checks),
			formatPerftEntry(res.mates, expected[i].mates))

	}

}

func formatPerftEntry(actual, expected int64) string {

	diff := actual - expected

	if diff != 0 {
		return fmt.Sprintf("%7s", formatNodesCount(diff))
	}

	return fmt.Sprint("      0")
}

func formatDuration(d time.Duration) string {
	return fmt.Sprintf("%.2f", float64(d)*float64(1e-9))
}


func formatNodesCount(nodes int64) string {
	if nodes < 1000 && nodes > -1000 {
		return fmt.Sprintf("%d", nodes)
	} else if nodes < 1000000 && nodes > -1000000 {
		return fmt.Sprintf("%.1fK", float64(nodes)/1000)
	}
	return fmt.Sprintf("%.2fM", float64(nodes)/1000000)
}