package core

import (
	"math/bits"
	"testing"
)

func TestRookSlide(t *testing.T) {
	type testCase struct {
		fen      string
		expected uint64
	}
	testCases := []testCase{
		// Test cases with fen and expected bitboard outcome
		{"8/8/8/8/1p1R2p1/8/8/8 w - - 0 1", 0x0808080876080808},
		{"8/8/4p3/4Rpp1/8/8/4p3/8 w - - 0 1", 0x0000102f10101000},
		{"8/1R6/8/8/8/8/8/8 w - - 0 1", 0x02fd020202020202},
		{"8/8/8/8/8/8/8/7R w - - 0 1", 0x808080808080807f},
		{"R7/8/8/8/8/8/8/8 w - - 0 1", 0xfe01010101010101},
		{"7R/8/8/8/8/8/8/8 w - - 0 1", 0x7f80808080808080},
		{"1n6/1n6/1n6/nRnnnnnn/1n6/1n6/1n6/1n6 w - - 0 1", 0x0000020502000000},
		// Test cases blocked by own pieces and enemy
		{"8/8/8/3P4/1p1R1Pp1/8/3P4/8 w - - 0 1", 0x0000000016080000},
	}
	generateSliding()
	for _, test := range testCases {
		var board Board
		board.LoadFen(test.fen)
		sq := bits.TrailingZeros64(*board.PieceBBmap[p_Rook])
		bb := getHorizontalSlide(sq, ^board.emptySquares)
		bb |= getVerticalSlide(sq, ^board.emptySquares)
		bb &= ^board.whiteSquares
		if bb != test.expected {
			DrawBitboard(t, bb)
			t.Errorf("0x%016x", bb)
		}
	}
}

func TestBishopSlide(t *testing.T) {
	type testCase struct {
		fen      string
		expected uint64
	}
	testCases := []testCase{
		// Test cases with fen and expected bitboard outcome
		{"8/8/8/8/3B4/8/8/8 w - - 0 1", 0x0808080876080808},
		// Test cases blocked by own pieces and enemy
		// {"8/8/8/3P4/1p1R1Pp1/8/3P4/8 w - - 0 1", 0x0000000016080000},
	}
	generateSliding()
	for _, test := range testCases {
		var board Board
		board.LoadFen(test.fen)
		sq := bits.TrailingZeros64(*board.PieceBBmap[p_Rook])
		bb := getHorizontalSlide(sq, ^board.emptySquares)
		bb |= getVerticalSlide(sq, ^board.emptySquares)
		bb &= ^board.whiteSquares
		if bb != test.expected {
			DrawBitboard(t, bb)
			t.Errorf("0x%016x", bb)
		}
	}
}
