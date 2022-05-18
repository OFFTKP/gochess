package core

import (
	"testing"
)

func TestBBFlipVertical(t *testing.T) {
	var bb uint64 = 0x1E2222120E0A1222
	var ex uint64 = 0x22120a0e1222221e
	fbb := flipVertically(bb)
	if fbb != ex {
		t.Error("Error while flipping bitboard vertically:\n")
		DrawBitboard(t, fbb)
		DrawBitboard(t, ex)
	}
}

func TestBBFlipHorizontal(t *testing.T) {
	var bb uint64 = 0x1E2222120E0A1222
	var ex uint64 = 0x7844444870504844
	fbb := flipHorizontally(bb)
	if fbb != ex {
		t.Error("Error while flipping bitboard horizontally:\n")
		DrawBitboard(t, fbb)
		DrawBitboard(t, ex)
	}
}
