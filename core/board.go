package core

import (
	"math"
	"math/bits"
	"math/rand"
	"time"
)

var oneTimeInitialized bool = false

const (
	p_Pawn = iota
	p_Knight
	p_Bishop
	p_Rook
	p_Queen
	p_King
)

const (
	c_White = 0
	c_Black = 6
)

var piecePower [6]int = [6]int{100, 300, 300, 500, 900, math.MaxInt32}

type Board struct {
	blackPawns   uint64
	blackKnights uint64
	blackBishops uint64
	blackRooks   uint64
	blackQueens  uint64
	blackKing    uint64

	whitePawns   uint64
	whiteKnights uint64
	whiteBishops uint64
	whiteRooks   uint64
	whiteQueens  uint64
	whiteKing    uint64

	blackPawnHashMap   [64]uint64
	blackKnightHashMap [64]uint64
	blackBishopHashMap [64]uint64
	blackRookHashMap   [64]uint64
	blackQueenHashMap  [64]uint64
	blackKingHashMap   [64]uint64

	whitePawnHashMap   [64]uint64
	whiteKnightHashMap [64]uint64
	whiteBishopHashMap [64]uint64
	whiteRookHashMap   [64]uint64
	whiteQueenHashMap  [64]uint64
	whiteKingHashMap   [64]uint64
	nextColorHashMap   [7]uint64  // 0 - white, 6 - black
	castlingHashMap    [16]uint64 // 0000 KQkq bits for speed
	enPassantHashMap   [8]uint64  // to indicate the file of the en passant square

	// maps for all piece types (0-11) to avoid branching
	PieceBBmap   [12]*uint64
	ColorBBmap   [7]*uint64
	PieceHashmap [12]*[64]uint64
	zobristHash  uint64

	// general bitboards of all bitmaps together
	whiteSquares uint64
	blackSquares uint64
	emptySquares uint64

	whiteKingsideCastle, whiteQueensideCastle,
	blackKingsideCastle, blackQueensideCastle int

	nextColor       int
	enPassantSquare uint8
	enPassantCol    uint8
	halfmoveClock   int
	fullmoveNumber  int
}

func oneTimeInit() {
	rand.Seed(time.Now().UnixNano())
	generateSliding()
}

func (board *Board) init() {
	if !oneTimeInitialized {
		oneTimeInit()
	}
	board.PieceBBmap = [12]*uint64{
		&board.whitePawns, &board.whiteKnights, &board.whiteBishops, &board.whiteRooks, &board.whiteQueens, &board.whiteKing,
		&board.blackPawns, &board.blackKnights, &board.blackBishops, &board.blackRooks, &board.blackQueens, &board.blackKing,
	}
	board.ColorBBmap = [7]*uint64{
		&board.whiteSquares, nil, nil, nil, nil, nil, &board.blackSquares,
	}
	board.blackKingsideCastle = 0
	board.blackQueensideCastle = 0
	board.whiteKingsideCastle = 0
	board.whiteQueensideCastle = 0
	board.enPassantSquare = 0xFF
	for i := 0; i < 64; i++ {
		board.whitePawnHashMap[i] = rand.Uint64()
		board.whiteKnightHashMap[i] = rand.Uint64()
		board.whiteBishopHashMap[i] = rand.Uint64()
		board.whiteRookHashMap[i] = rand.Uint64()
		board.whiteQueenHashMap[i] = rand.Uint64()
		board.whiteKingHashMap[i] = rand.Uint64()
		board.blackPawnHashMap[i] = rand.Uint64()
		board.blackKnightHashMap[i] = rand.Uint64()
		board.blackBishopHashMap[i] = rand.Uint64()
		board.blackRookHashMap[i] = rand.Uint64()
		board.blackQueenHashMap[i] = rand.Uint64()
		board.blackKingHashMap[i] = rand.Uint64()
	}
	board.nextColorHashMap[0] = rand.Uint64()
	board.nextColorHashMap[6] = rand.Uint64()
	for i := 0; i < 4; i++ {
		board.castlingHashMap[i] = rand.Uint64()
	}
	for i := 0; i < 8; i++ {
		board.enPassantHashMap[i] = rand.Uint64()
	}
	board.PieceHashmap = [12]*[64]uint64{
		&board.whitePawnHashMap, &board.whiteKnightHashMap, &board.whiteBishopHashMap, &board.whiteRookHashMap, &board.whiteQueenHashMap, &board.whiteKingHashMap,
		&board.blackPawnHashMap, &board.blackKnightHashMap, &board.blackBishopHashMap, &board.blackRookHashMap, &board.blackQueenHashMap, &board.blackKingHashMap,
	}
}

func (board *Board) Reset() {
	// TODO: implement what happens on ucinewgame?
}

func (board *Board) recalculateZobrist() {
	board.zobristHash = 0
	emptyCopy := board.emptySquares
	for emptyCopy != 0 {
		bit := bits.TrailingZeros64(emptyCopy)
		for j := 0; j < 12; j++ {
			if ((*board.PieceBBmap[j] >> bit) & 1) != 0 {
				board.zobristHash ^= (*board.PieceHashmap[j])[bit]
				break
			}
		}
		emptyCopy ^= 1 << bit
	}
	board.zobristHash ^= board.nextColorHashMap[board.nextColor]
	castling := board.whiteKingsideCastle<<3 | board.whiteQueensideCastle<<2 | board.blackKingsideCastle<<1 | board.blackQueensideCastle
	board.zobristHash ^= board.castlingHashMap[castling]
	board.zobristHash ^= board.enPassantHashMap[board.enPassantCol]
}

// function that recalculates the occupying maps
// to be run after changing the board state
func (board *Board) recalculateGeneralMaps() {
	board.whiteSquares = board.whitePawns | board.whiteKnights | board.whiteBishops | board.whiteRooks | board.whiteQueens | board.whiteKing
	board.blackSquares = board.blackPawns | board.blackKnights | board.blackBishops | board.blackRooks | board.blackQueens | board.blackKing
	board.emptySquares = (^board.whiteSquares) & (^board.blackSquares)
}

// returns -1 if no piece is captured, or > -1 if a piece is captured and the index of
// the corresponding bitboard
func (board *Board) makeMove(bitboardIndex uint8, oldSquare, newSquare uint8) int {
	var ret int = -1
	oldBitCheck := uint64(1) << oldSquare
	// remove the moving piece from hash
	board.zobristHash ^= (*board.PieceHashmap[bitboardIndex])[oldSquare]
	// remove the moving piece from bitboard
	*board.PieceBBmap[bitboardIndex] &= ^oldBitCheck
	// find if new square contains a piece thats being captured
	// 0 meaning its not empty here
	bitCheck := uint64(1) << newSquare
	if (board.emptySquares & bitCheck) == 0 {
		// find which piece it is
		for ret = 0; ; /* no test */ ret++ {
			if (*board.PieceBBmap[ret] & bitCheck) != 0 {
				break
			}
		}
		// remove the captured piece from hash
		board.zobristHash ^= (*board.PieceHashmap[ret])[newSquare]
		// remove the captured piece from bitboard
		*board.PieceBBmap[ret] &= ^bitCheck
	}
	board.zobristHash ^= (*board.PieceHashmap[bitboardIndex])[newSquare]
	*board.PieceBBmap[bitboardIndex] |= uint64(1) << newSquare
	board.recalculateGeneralMaps()
	return ret
}

func (board *Board) unmakeMove(oldCapture int, bitboardIndex uint8, oldSquare, newSquare uint8) {
	board.zobristHash ^= (*board.PieceHashmap[bitboardIndex])[newSquare]
	bitCheck := uint64(1) << newSquare
	*board.PieceBBmap[bitboardIndex] &= ^bitCheck
	if oldCapture != -1 {
		board.zobristHash ^= (*board.PieceHashmap[oldCapture])[newSquare]
		*board.PieceBBmap[oldCapture] |= bitCheck
	}
	oldBitCheck := uint64(1) << oldSquare
	board.zobristHash ^= (*board.PieceHashmap[bitboardIndex])[oldSquare]
	*board.PieceBBmap[bitboardIndex] |= oldBitCheck
	board.recalculateGeneralMaps()
}
