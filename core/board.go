package core

import (
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"
)

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
	nextColorHashMap   [2]uint64  // 0 - white, 1 - black
	castlingHashMap    [16]uint64 // 0000 KQkq bits for speed
	enPassantHashMap   [8]uint64  // to indicate the file of the en passant square

	// maps for all piece types (0-11) to avoid branching
	PieceBBmap   [12]*uint64
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

func (board *Board) init() {
	board.PieceBBmap = [12]*uint64{
		&board.whitePawns, &board.whiteKnights, &board.whiteBishops, &board.whiteRooks, &board.whiteQueens, &board.whiteKing,
		&board.blackPawns, &board.blackKnights, &board.blackBishops, &board.blackRooks, &board.blackQueens, &board.blackKing,
	}
	board.blackKingsideCastle = 0
	board.blackQueensideCastle = 0
	board.whiteKingsideCastle = 0
	board.whiteQueensideCastle = 0
	board.enPassantSquare = 0xFF
	rand.Seed(time.Now().UnixNano())
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
	board.nextColorHashMap[1] = rand.Uint64()
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

func (board *Board) recalculateZobrist() {
	board.zobristHash = 0
	for i := 0; i < 64; i++ {
		if (board.emptySquares>>i)&1 == 0 {
			for j := 0; j < 12; j++ {
				if ((*board.PieceBBmap[j] >> i) & 1) != 0 {
					board.zobristHash ^= (*board.PieceHashmap[j])[i]
					break
				}
			}
		}
	}
	if board.nextColor == c_White {
		board.zobristHash ^= board.nextColorHashMap[0]
	} else {
		board.zobristHash ^= board.nextColorHashMap[1]
	}
	castling := board.whiteKingsideCastle<<3 | board.whiteQueensideCastle<<2 | board.blackKingsideCastle<<1 | board.blackQueensideCastle
	board.zobristHash ^= board.castlingHashMap[castling]
	board.zobristHash ^= board.enPassantHashMap[board.enPassantCol]
}

func getPieceChar(piece uint8) rune {
	switch piece {
	case p_Pawn + c_Black:
		return 'p'
	case p_Knight + c_Black:
		return 'n'
	case p_Bishop + c_Black:
		return 'b'
	case p_Rook + c_Black:
		return 'r'
	case p_Queen + c_Black:
		return 'q'
	case p_King + c_Black:
		return 'k'
	case p_Pawn + c_White:
		return 'P'
	case p_Knight + c_White:
		return 'N'
	case p_Bishop + c_White:
		return 'B'
	case p_Rook + c_White:
		return 'R'
	case p_Queen + c_White:
		return 'Q'
	case p_King + c_White:
		return 'K'
	}
	return ' '
}

// function that recalculates the occupying maps
// to be run after changing the board state
func (board *Board) recalculateGeneralMaps() {
	board.whiteSquares = board.whitePawns | board.whiteKnights | board.whiteBishops | board.whiteRooks | board.whiteQueens | board.whiteKing
	board.blackSquares = board.blackPawns | board.blackKnights | board.blackBishops | board.blackRooks | board.blackQueens | board.blackKing
	board.emptySquares = (^board.whiteSquares) & (^board.blackSquares)
}

func algebraicToUint8(algebraicSquare string) uint8 {
	var ret uint8
	ch1, ch2 := algebraicSquare[0], algebraicSquare[1]
	ch1 -= 'a'
	ch2 -= '1'
	ret = ch1 + (ch2 * 8)
	return ret
}

func uint8ToAlgebraic(numSquare uint8) string {
	if numSquare == 0xFF {
		return "-"
	}
	x := numSquare & 7
	y := numSquare >> 3
	ch1, ch2 := 'a'+x, '1'+y
	var sb strings.Builder
	sb.WriteByte(ch1)
	sb.WriteByte(ch2)
	return sb.String()
}

func (board *Board) LoadFen(fen string) (bool, string) {
	board.init()
	fenSplit := strings.Split(fen, " ")
	if len(fenSplit) != 6 {
		return false, "Bad fen length: " + fen
	}
	fenPos := fenSplit[0]
	fenColor := fenSplit[1]
	fenCastling := fenSplit[2]
	if fenColor == "b" {
		board.nextColor = c_Black
	} else {
		board.nextColor = c_White
	}
	for _, ch := range fenCastling {
		switch ch {
		case 'k':
			board.blackKingsideCastle = 1
		case 'q':
			board.blackQueensideCastle = 1
		case 'K':
			board.whiteKingsideCastle = 1
		case 'Q':
			board.whiteQueensideCastle = 1
		}
	}
	if len(fenSplit[3]) == 2 {
		board.enPassantSquare = algebraicToUint8(fenSplit[3])
		board.enPassantCol = board.enPassantSquare & 0b111
	}
	{
		clk, err := strconv.Atoi(fenSplit[4])
		if err != nil {
			return false, "Could not parse halfmove clock"
		}
		board.halfmoveClock = clk
	}
	{
		clk, err := strconv.Atoi(fenSplit[5])
		if err != nil {
			return false, "Could not parse fullmove number"
		}
		board.fullmoveNumber = clk
	}
	x := uint64(0)
	y := uint64(56)
	for _, ch := range fenPos {
		if unicode.IsNumber(ch) {
			x += uint64(ch - '0')
			continue
		}
		var piece uint8
		var color uint8 = c_Black
		if !unicode.IsLower(ch) {
			color = c_White
		}
		ch_l := unicode.ToLower(ch)
		switch ch_l {
		case 'p':
			piece = p_Pawn
		case 'n':
			piece = p_Knight
		case 'b':
			piece = p_Bishop
		case 'r':
			piece = p_Rook
		case 'q':
			piece = p_Queen
		case 'k':
			piece = p_King
		case '/':
			x = 0
			y -= 8
			continue
		}
		pieceMap := board.PieceBBmap[piece+color]
		*pieceMap |= 1 << (x + y)
		x += 1
		if x > 8 {
			x = 0
			y -= 8
		}
	}
	if x != 8 && y != 0 {
		return false, "Not enough or too many pieces in FEN"
	}
	board.recalculateGeneralMaps()
	board.recalculateZobrist()
	return true, ""
}

func (board *Board) GetFen() string {
	var sb strings.Builder
	sinceLastLine := 0
	spaceCount := '0'
	for y := 7; y >= 0; y-- {
		for x := 0; x < 8; x++ {
			curIndex := x + (y * 8)
			curPiece := ' '
			for j := uint8(0); j < 12; j++ {
				curBB := board.PieceBBmap[j]
				if ((*curBB) & (1 << curIndex)) != 0 {
					curPiece = getPieceChar(j)
					break
				}
			}
			if curPiece != ' ' {
				if spaceCount > '0' {
					sb.WriteRune(spaceCount)
					spaceCount = '0'
				}
				sb.WriteRune(curPiece)
			} else {
				spaceCount++
			}
			sinceLastLine++
			if sinceLastLine > 7 {
				sinceLastLine = 0
				if spaceCount > '0' {
					sb.WriteRune(spaceCount)
					spaceCount = '0'
				}
				// final row final col
				if curIndex != 7 {
					sb.WriteRune('/')
				}
			}
		}
	}
	sb.WriteRune(' ')
	if board.nextColor == c_Black {
		sb.WriteRune('b')
	} else {
		sb.WriteRune('w')
	}
	sb.WriteRune(' ')
	castleCount := 0
	if board.whiteKingsideCastle == 1 {
		sb.WriteRune('K')
		castleCount++
	}
	if board.whiteQueensideCastle == 1 {
		sb.WriteRune('Q')
		castleCount++
	}
	if board.blackKingsideCastle == 1 {
		sb.WriteRune('k')
		castleCount++
	}
	if board.blackQueensideCastle == 1 {
		sb.WriteRune('q')
		castleCount++
	}
	if castleCount == 0 {
		sb.WriteRune('-')
	}
	sb.WriteRune(' ')
	sb.WriteString(uint8ToAlgebraic(board.enPassantSquare))
	sb.WriteRune(' ')
	sb.WriteString(strconv.Itoa(board.halfmoveClock))
	sb.WriteRune(' ')
	sb.WriteString(strconv.Itoa(board.fullmoveNumber))
	return sb.String()
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
