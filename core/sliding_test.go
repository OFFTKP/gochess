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

func TestDiagonalSliding(t *testing.T) {
	type testCase struct {
		fen      string
		expected uint64
	}
	testCases := []testCase{
		// Test cases with fen and expected bitboard outcome
		{"8/8/8/8/8/8/8/B7 w - - 0 1", 0x8040201008040200},
		{"8/8/8/8/8/8/1B6/8 w - - 0 1", 0x8040201008040001},
		{"8/8/8/8/8/2B5/8/8 w - - 0 1", 0x8040201008000201},
		{"8/8/8/8/3B4/8/8/8 w - - 0 1", 0x8040201000040201},
		{"8/8/8/4B3/8/8/8/8 w - - 0 1", 0x8040200008040201},
		{"8/8/5B2/8/8/8/8/8 w - - 0 1", 0x8040001008040201},
		{"8/6B1/8/8/8/8/8/8 w - - 0 1", 0x8000201008040201},
		{"7B/8/8/8/8/8/8/8 w - - 0 1", 0x0040201008040201},
		// Test cases blocked by own pieces and enemy
	}
	generateSliding()
	for _, test := range testCases {
		var board Board
		board.LoadFen(test.fen)
		sq := bits.TrailingZeros64(*board.PieceBBmap[p_Bishop])
		bb := slidingDiagonal[sq][^getDiagonalOccupancy(sq, ^board.emptySquares)]
		// bb &= ^board.whiteSquares
		if bb != test.expected {
			DrawBitboard(t, bb)
			DrawBitboard(t, test.expected)
		}
	}
}

func TestBishopGetDiagonalOccupancy(t *testing.T) {
	type testCase struct {
		bitboard uint64
		square   int
		expected uint8
	}
	testCases := []testCase{
		{0x8040201008040201, 0, 0b1111_1111},
		{0x0080402010080402, 1, 0b1111_1110},
		{0x0000804020100804, 2, 0b1111_1100},
		{0x0000008040201008, 3, 0b1111_1000},
		{0x0000000080402010, 4, 0b1111_0000},
		{0x0000000000804020, 5, 0b1110_0000},
		{0x0000000000008040, 6, 0b1100_0000},
		{0x0000000000000080, 7, 0b1000_0000},
		{0x4020100804020100, 8, 0b0111_1111},
		{0x2010080402010000, 16, 0b0011_1111},
		{0x1008040201000000, 24, 0b0001_1111},
		{0x0804020100000000, 32, 0b0000_1111},
		{0x0402010000000000, 40, 0b0000_0111},
		{0x0201000000000000, 48, 0b0000_0011},
		{0x0100000000000000, 56, 0b0000_0001},
	}
	for _, test := range testCases {
		cur := getDiagonalOccupancy(test.square, test.bitboard)
		if cur != test.expected {
			DrawBitboard(t, uint64(test.bitboard))
			DrawBitboard(t, uint64(cur))
		}
	}
}

func TestGetDiagonal(t *testing.T) {
	DrawBitboard(t, diagonalBitboards[diagonalTranslation[27]])
}
