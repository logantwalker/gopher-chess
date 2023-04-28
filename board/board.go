package board

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

var (
	StartingFen 	string = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	CastleNone 		int8 = 0
	CastleShort		int8 = 1
	CastleLong		int8 = 2

	WhiteKingStartSquare 	Square = E1
	WhiteRookStartSquares []Square = []Square{A1,H1}

	BlackKingStartSquare 	Square = E8
	BlackRookStartSquares []Square = []Square{A8,H8}

	StatusNormal int = 0
	StatusCheckmate int = 1
	StatusCheck int = 2
	StatusStalemate int = 3
	StatusThreeFoldRep int = 4
)

type Pin struct {
	Delta int8
}

type Check struct {
	AttackerOrigin 	int8
	AttackerDelta 	int8
}

type Move struct{
	From 		Square
	To 			Square
	Capture 	int8
	MovedPiece 	int8
	Type 		int8
	Promotion 	int8
	// Pin 		*board.Pin
}

// move types
var (
	moveOrdinary 	int8 = 0
	moveShortCastle int8 = 1
	moveLongCastle 	int8 = 2
	movePromote		int8 = 3
	moveEnPassant	int8 = 4
)

type MoveRecord struct{
	Move 			Move
	WhiteCastle		int8
	BlackCastle		int8
	EnPassant		Square
	HalfMoveClock	int
	ZobristHash		uint64
}

type Board struct {
	History 		[]MoveRecord
	State 			[]int8
	Turn 			int8
	WhiteCastle 	int8
	BlackCastle 	int8
	KingLocations	[]int8
	IsCheck 		bool
	EnPassant 		Square
	HalfMoveClock 	int
	FullMoveClock	int
	Ply 			int
	Status 			int
	ZobristTable	*ZobristTable
	ZobristHash		uint64
}

func NewBoard(fen string) Board {
	b, err := ParseFen(fen)

	b.KingLocations = append(b.KingLocations,int8(WhiteKingStartSquare))
	b.KingLocations = append(b.KingLocations,int8(BlackKingStartSquare))
	b.ZobristTable = InitZobristTable()
	b.GenerateHash()

	if err != nil{
		log.Fatal(err.Error())
	}

	return b
}

func ParseFen(fen string) (Board, error) {
	boardArr := make([]int8, 128)
	var board Board = Board{
		State: boardArr,
	}
	fenArray := strings.Split(fen, " ")

	i := 0
	for j := 0; j < len(fenArray[0]); j++ {
		piece := fenArray[0][j]

		switch piece {
		case 'p':
			board.State[HexBoard[i]] = BlackPawn
		case 'r':
			board.State[HexBoard[i]] = BlackRook
		case 'n':
			board.State[HexBoard[i]] = BlackKnight
		case 'b':
			board.State[HexBoard[i]] = BlackBishop
		case 'q':
			board.State[HexBoard[i]] = BlackQueen
		case 'k':
			board.State[HexBoard[i]] = BlackKing


		case 'P':
			board.State[HexBoard[i]] = WhitePawn
		case 'R':
			board.State[HexBoard[i]] = WhiteRook
		case 'N':
			board.State[HexBoard[i]] = WhiteKnight
		case 'B':
			board.State[HexBoard[i]] = WhiteBishop
		case 'Q':
			board.State[HexBoard[i]] = WhiteQueen
		case 'K':
			board.State[HexBoard[i]] = WhiteKing


		case '1':
			// do nothing
		case '2':
			i++
		case '3':
			i+= 2
		case '4':
			i+= 3
		case '5':
			i+= 4
		case '6':
			i+= 5
		case '7':
			i+= 6
		case '8':
			i+= 7

		case '/':
			i--

		default:
			return board, errors.New("invalid FEN")
		}
		i++
	}

	switch fenArray[1][0] {
	case 'w':
		board.Turn = White
	case 'b':
		board.Turn = Black
	default:
		return board, errors.New("invalid FEN")
	}

	castleRights := strings.Split(fenArray[2],"") 
	
	for _, char := range castleRights {
		switch char {
		case "K":
			board.WhiteCastle += CastleShort
		case "Q":
			board.WhiteCastle += CastleLong
		case "k":
			board.BlackCastle += CastleShort
		case "q":
			board.BlackCastle += CastleLong
		}
	}

	return board, nil
}

func LegalSquare(square int8) bool {
	return !(uint8(square)&0x88 != 0)
}

func (b *Board) IsEmpty(squares ...Square) bool{
	for _, sq := range squares {
		if b.State[sq] != Empty{
			return false
		}
	}

	return true
}

func (b *Board) PrintBoard(){
	for i := 0x70; i >= 0x00; i -= 0x10 {
		for j := 0; j < 8; j++ {
			square := i + j
			fmt.Printf("%v ", GetPieceSymbol(b.State[square]))
		}
		fmt.Printf("\n")
	}
	if b.Status == StatusCheckmate{
		var winner string
		if b.Turn == White{
			winner = "Black"
		}else{
			winner = "White"
		}

		fmt.Printf("Checkmate! %s wins!\n",winner)
	}

	if b.Status == StatusStalemate{
		fmt.Printf("Stalemate!\n")
	}

	if b.Status == StatusThreeFoldRep {
		fmt.Printf("Draw by Three Fold Repitition.\n")
	}
}

func (b *Board) GenerateHash() {
	key := uint64(0)

	for sq := int8(0); sq < NumSquares; sq++ {
		piece := b.State[sq]
		if piece > Empty {
			key ^= b.ZobristTable.pieceKeys[piece-1][0][sq]
		} else if piece < Empty {
			key ^= b.ZobristTable.pieceKeys[-piece-1][1][sq]
		}
	}

	key ^= b.ZobristTable.whiteCastlingKeys[b.WhiteCastle]
	key ^= b.ZobristTable.blackCastlingKeys[b.BlackCastle]

	if b.Turn == Black {
		key ^= b.ZobristTable.sideToMoveKey
	}

	b.ZobristHash = key
}

func (b* Board) UpdateHash(m *Move){
	key := b.ZobristHash
	side := 0
	if b.Turn == Black {
		side = 1
	}
	piece := m.MovedPiece
	if piece < 0{
		piece = -1*piece
	}

	piece = piece - 1

	// fmt.Println(b.ZobristTable.pieceKeys)

	key ^= b.ZobristTable.pieceKeys[piece][side][int8(m.From)]
	key ^= b.ZobristTable.pieceKeys[piece][side][int8(m.To)]

	if m.Capture != Empty {
		capture := m.Capture
		if capture < 0{
			capture = -1*capture
		}
		key ^= b.ZobristTable.pieceKeys[capture-1][side][int8(m.To)]
	}

	switch m.Type {
	case moveLongCastle:
		if side == 0 {
			key ^= b.ZobristTable.whiteCastlingKeys[CastleLong]
		} else {
			key ^= b.ZobristTable.blackCastlingKeys[CastleLong]
		}
	case moveShortCastle:
		if side == 0 {
			key ^= b.ZobristTable.whiteCastlingKeys[CastleShort]
		} else {
			key ^= b.ZobristTable.blackCastlingKeys[CastleShort]
		}
	case moveEnPassant:
		key ^= b.ZobristTable.enPassantKeys[int8(m.To)-int8(m.MovedPiece)*16]
	}

	b.ZobristHash = key
}
