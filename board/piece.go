package board

const (
	White int8 = 1
	Black int8 = -1

	Empty int8 = 0

	Pawn 	int8 = 1
	Knight 	int8 = 2
	Bishop 	int8 = 3
	Rook 	int8 = 4
	Queen 	int8 = 5
	King 	int8 = 6

	WhitePawn 	int8 = White * Pawn
	WhiteKnight int8 = White * Knight
	WhiteBishop int8 = White * Bishop
	WhiteRook	int8 = White * Rook
	WhiteQueen	int8 = White * Queen
	WhiteKing 	int8 = White * King

	BlackPawn 	int8 = Black * Pawn
	BlackKnight int8 = Black * Knight
	BlackBishop int8 = Black * Bishop
	BlackRook	int8 = Black * Rook
	BlackQueen	int8 = Black * Queen
	BlackKing 	int8 = Black * King
)

var pieceChars = []string{
	".",
	"P","N","B","R","Q","K",
	"p","n","b","r","q","k",
}

var pieceSymbols = []string{
	".",
	"♟", "♞", "♝", "♜", "♛", "♚",
	"♙", "♘", "♗", "♖", "♕", "♔",
}

func GetPieceSymbol(piece int8) string {
	if piece >= 0{
		return pieceSymbols[piece]
	}

	return pieceSymbols[-1*piece + 6]
}

func GetPieceString(piece int8) string{
	if piece >= 0{
		return pieceChars[piece]
	}

	return pieceChars[-1*piece + 6]
}