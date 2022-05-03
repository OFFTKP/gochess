package core

import (
	"fmt"
	"strings"
	"unicode"
)

const (
	p_Pawn = 1 << iota
	p_Knight
	p_Bishop
	p_Rook
	p_Queen
	p_King
	p_Empty = 0
)

const (
	c_Empty   = 0
	c_Black   = 1 << 6
	c_White   = 1 << 7
	c_XorFlag = c_Black + c_White
)

const PIECE_MASK = 0b00111111
const COLOR_MASK = 0b11000000

type Board [64]uint8
type LegalMoves []uint8

var knightMovesLUT [64]uint64 = [64]uint64{
	0x00000000008a0091, 0x00000000008b9092, 0x00000000888c9193, 0x00000000898d9294, 0x000000008a8e9395, 0x000000008b8f9496, 0x000000008c009597, 0x000000008d009600,
	0x0000008200920099, 0x000000830093989a, 0x000080849094999b, 0x0000818591959a9c, 0x0000828692969b9d, 0x0000838793979c9e, 0x0000840094009d9f, 0x0000850095009e00,
	0x0081008a009a00a1, 0x8082008b009ba0a2, 0x8183888c989ca1a3, 0x8284898d999da2a4, 0x83858a8e9a9ea3a5, 0x84868b8f9b9fa4a6, 0x85878c009c00a5a7, 0x86008d009d00a600,
	0x0089009200a200a9, 0x888a009300a3a8aa, 0x898b9094a0a4a9ab, 0x8a8c9195a1a5aaac, 0x8b8d9296a2a6abad, 0x8c8e9397a3a7acae, 0x8d8f9400a400adaf, 0x8e009500a500ae00,
	0x0091009a00aa00b1, 0x9092009b00abb0b2, 0x9193989ca8acb1b3, 0x9294999da9adb2b4, 0x93959a9eaaaeb3b5, 0x94969b9fabafb4b6, 0x95979c00ac00b5b7, 0x96009d00ad00b600,
	0x009900a200b200b9, 0x989a00a300b3b8ba, 0x999ba0a4b0b4b9bb, 0x9a9ca1a5b1b5babc, 0x9b9da2a6b2b6bbbd, 0x9c9ea3a7b3b7bcbe, 0x9d9fa400b400bdbf, 0x9e00a500b500be00,
	0x00a100aa00ba0000, 0xa0a200ab00bb0000, 0xa1a3a8acb8bc0000, 0xa2a4a9adb9bd0000, 0xa3a5aaaebabe0000, 0xa4a6abafbbbf0000, 0xa5a7ac00bc000000, 0xa600ad00bd000000,
	0x00a900b200000000, 0xa8aa00b300000000, 0xa9abb0b400000000, 0xaaacb1b500000000, 0xabadb2b600000000, 0xacaeb3b700000000, 0xadafb40000000000, 0xae00b50000000000,
}

var knightHeatTable [64]uint8 = [64]uint8{
	2, 3, 4, 4, 4, 4, 3, 2,
	3, 4, 6, 6, 6, 6, 4, 3,
	4, 6, 8, 8, 8, 8, 6, 4,
	4, 6, 8, 8, 8, 8, 6, 4,
	4, 6, 8, 8, 8, 8, 6, 4,
	4, 6, 8, 8, 8, 8, 6, 4,
	3, 4, 6, 6, 6, 6, 4, 3,
	2, 3, 4, 4, 4, 4, 3, 2,
}

func getPieceChar(piece uint8) rune {
	switch piece {
	case p_Pawn | c_Black:
		return 'p'
	case p_Knight | c_Black:
		return 'n'
	case p_Bishop | c_Black:
		return 'b'
	case p_Rook | c_Black:
		return 'r'
	case p_Queen | c_Black:
		return 'q'
	case p_King | c_Black:
		return 'k'
	case p_Pawn | c_White:
		return 'P'
	case p_Knight | c_White:
		return 'N'
	case p_Bishop | c_White:
		return 'B'
	case p_Rook | c_White:
		return 'R'
	case p_Queen | c_White:
		return 'Q'
	case p_King | c_White:
		return 'K'
	}
	return ' '
}

func (board *Board) LoadFen(fen string) {
	x := 0
	y := 0
	for _, ch := range fen {
		if unicode.IsNumber(ch) {
			x += int(ch - '0')
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
			y += 8
			continue
		}
		piece |= color
		board[x+y] = piece
		x += 1
		if x > 8 {
			x = 0
			y += 8
		}
	}
}

func (board *Board) GetFen() string {
	var sb strings.Builder
	sinceLastLine := 0
	spaceCount := '0'
	for i := 0; i < 64; i++ {
		if board[i] != 0 {
			if spaceCount > '0' {
				sb.WriteRune(spaceCount)
				spaceCount = '0'
			}
			sb.WriteRune(getPieceChar(board[i]))
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
			if i != 63 {
				sb.WriteRune('/')
			}
		}
	}
	return sb.String()
}

func (board *Board) Draw() {
	fmt.Println("----------------")
	for x, y := 0, 0; y < 64; {
		fmt.Printf("|%c", getPieceChar(board[x+y]))
		x++
		if x == 8 {
			fmt.Println("|")
			y += 8
			x = 0
			fmt.Println("----------------")
		}
	}
}

func (board *Board) getLegalMoves(index uint8) []uint8 {
	piece := board[index]
	ret := make([]uint8, 0, 64)
	switch piece & PIECE_MASK {
	case p_Pawn:
		if piece&COLOR_MASK == c_White {
			// assuming pawn index always > 7
			// because otherwise it promotes
			col := board[index-8] & COLOR_MASK
			if col == c_Empty {
				ret = append(ret, index-8)
			}
			if index >= 48 {
				col := board[index-16] & COLOR_MASK
				if col == c_Empty {
					ret = append(ret, index-16)
				}
			}
		} else {
			// assuming pawn index always > 56
			// because otherwise it promotes
			col := board[index+8] & COLOR_MASK
			if col == c_Empty {
				ret = append(ret, index+8)
			}
			if index <= 15 {
				col := board[index+16] & COLOR_MASK
				if col == c_Empty {
					ret = append(ret, index+16)
				}
			}
		}
	case p_Knight:
		board.appendKnightMoves(&ret, index, piece&COLOR_MASK)
	}
	return ret
}

func (board *Board) countLegalMoves(color uint8) int {
	count := 0
	for i := uint8(0); i < 64; i++ {
		if board[i]&COLOR_MASK == color {
			moves := board.getLegalMoves(i)
			count += len(moves)
		}
	}
	return count
}

func (board *Board) appendKnightMoves(moves *[]uint8, square uint8, color uint8) {
	kmoves := knightMovesLUT[square]
	for i := 0; i < 8; i++ {
		cur := uint8(kmoves >> (i * 8))
		if ((cur & 0x80) == 0x80) && ((board[cur-0x80] & COLOR_MASK) != color) {
			*moves = append(*moves, cur&0x1F)
		}
	}
}
