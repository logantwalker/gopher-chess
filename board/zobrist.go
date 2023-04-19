package board

import (
	"crypto/rand"
	"encoding/binary"
)

const (
	NumPieces = 6
	NumSides = 2
	NumSquares int8 = 120
)

type ZobristTable struct {
	pieceKeys 			[NumPieces][NumSides][NumSquares]uint64
	enPassantKeys 		[NumSquares]uint64
	whiteCastlingKeys 	[4]uint64
	blackCastlingKeys 	[4]uint64
	sideToMoveKey 		uint64
}

func randomUint64() uint64 {
	var buf [8]byte
	_, err := rand.Read(buf[:])
	if err != nil {
		panic(err)
	}
	return binary.LittleEndian.Uint64(buf[:])
}

func InitZobristTable() *ZobristTable {
	z := ZobristTable{}

	// pieces
	for piece := 0; piece < NumPieces; piece++ {
		for side := 0; side < NumSides; side++ {
			for sq := int8(0); sq < NumSquares; sq++ {
				z.pieceKeys[piece][side][sq] = randomUint64()
			}
		}
	}
	// en passant
	for sq := int8(0); sq < NumPieces; sq++ {
		z.enPassantKeys[sq] = randomUint64()
	}
	// castling options
	for i := int8(0); i < 4; i++ {
		z.whiteCastlingKeys[i] = randomUint64()
		z.blackCastlingKeys[i] = randomUint64()
	}

	// side
	z.sideToMoveKey = randomUint64()

	return &z
}