package core

import (
	"strings"
	"testing"
)

var pieceChars [12]byte = [12]byte{
	'P', 'N', 'B', 'R', 'Q', 'K',
	'p', 'n', 'b', 'r', 'q', 'k',
}

func DrawBitboard(t *testing.T, bb uint64) {
	var sb strings.Builder
	sb.WriteByte('\n')
	for j := uint64(0); j < 8; j++ {
		for i := uint64(0); i < 8; i++ {
			b := ((bb >> (i + (7-j)*8)) & 1) == 1
			sb.WriteByte('|')
			if b {
				sb.WriteByte('X')
			} else {
				sb.WriteByte('_')
			}
			sb.WriteByte('|')
		}
		sb.WriteByte('\n')
	}
	t.Errorf(sb.String())
}

func (board *Board) DrawChessboard(t *testing.T) {
	var boardDraw [64]byte
	for i := 0; i < 64; i++ {
		boardDraw[i] = '_'
		bitCheck := uint64(1) << i
		if (board.emptySquares & bitCheck) == 0 {
			for p := 0; p < 12; p++ {
				if (*board.PieceBBmap[p] & bitCheck) != 0 {
					boardDraw[i] = pieceChars[p]
				}
			}
		} else {
			boardDraw[i] = '_'
		}
	}
	var sb strings.Builder
	sb.WriteByte('\n')
	for j := uint64(0); j < 8; j++ {
		for i := uint64(0); i < 8; i++ {
			sb.WriteByte('|')
			ch := boardDraw[(7-j)*8+i]
			sb.WriteByte(ch)
			sb.WriteByte('|')
		}
		sb.WriteByte('\n')
	}
	t.Errorf(sb.String())
}
