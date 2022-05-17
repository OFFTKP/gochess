package core

import (
	"strconv"
	"strings"
	"unicode"
)

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

func AlgebraicToUint8(algebraicSquare string) uint8 {
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
		board.enPassantSquare = AlgebraicToUint8(fenSplit[3])
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
