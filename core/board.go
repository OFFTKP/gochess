package core

import (
	"strconv"
	"strings"
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

type Board struct {
	blackPawns   uint64
	blackKnights uint64
	blackBishops uint64
	blackRooks   uint64
	blackQueen   uint64
	blackKing    uint64

	whitePawns   uint64
	whiteKnights uint64
	whiteBishops uint64
	whiteRooks   uint64
	whiteQueen   uint64
	whiteKing    uint64

	// general bitboards of all bitmaps together
	whiteSquares uint64
	blackSquares uint64
	emptySquares uint64

	whiteKingsideCastle, whiteQueensideCastle,
	blackKingsideCastle, blackQueensideCastle bool

	pieceBBmap      [12]*uint64
	nextColor       int
	enPassantSquare uint8
	halfmoveClock   int
	fullmoveNumber  int
}

func (board *Board) init() {
	board.pieceBBmap = [12]*uint64{
		&board.whitePawns, &board.whiteKnights, &board.whiteBishops, &board.whiteRooks, &board.whiteQueen, &board.whiteKing,
		&board.blackPawns, &board.blackKnights, &board.blackBishops, &board.blackRooks, &board.blackQueen, &board.blackKing,
	}
	board.blackKingsideCastle = false
	board.blackQueensideCastle = false
	board.whiteKingsideCastle = false
	board.whiteQueensideCastle = false
	board.enPassantSquare = 0xFF
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
	board.whiteSquares = board.whitePawns | board.whiteKnights | board.whiteBishops | board.whiteRooks | board.whiteQueen | board.whiteKing
	board.blackSquares = board.blackPawns | board.blackKnights | board.blackBishops | board.blackRooks | board.blackQueen | board.blackKing
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
			board.blackKingsideCastle = true
		case 'q':
			board.blackQueensideCastle = true
		case 'K':
			board.whiteKingsideCastle = true
		case 'Q':
			board.whiteQueensideCastle = true
		}
	}
	if len(fenSplit[3]) == 2 {
		board.enPassantSquare = algebraicToUint8(fenSplit[3])
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
		pieceMap := board.pieceBBmap[piece+color]
		*pieceMap |= 1 << (x + y)
		x += 1
		if x > 8 {
			x = 0
			y -= 8
		}
	}
	board.recalculateGeneralMaps()
	if x != 8 && y != 0 {
		return false, "Not enough or too many pieces in FEN"
	}
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
				curBB := board.pieceBBmap[j]
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
	if board.whiteKingsideCastle {
		sb.WriteRune('K')
		castleCount++
	}
	if board.whiteQueensideCastle {
		sb.WriteRune('Q')
		castleCount++
	}
	if board.blackKingsideCastle {
		sb.WriteRune('k')
		castleCount++
	}
	if board.blackQueensideCastle {
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
