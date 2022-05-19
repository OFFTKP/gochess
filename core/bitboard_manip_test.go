package core

import (
	"math/rand"
	"testing"
	"time"
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
	var bb uint64 = 0x1e2222120e0a1222
	var ex uint64 = 0x7844444870504844
	fbb := flipHorizontally(bb)
	if fbb != ex {
		t.Error("Error while flipping bitboard horizontally:\n")
		DrawBitboard(t, fbb)
		DrawBitboard(t, ex)
	}
}

func TestBBRotate90Random(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		var bb uint64 = rand.Uint64()
		rbb := rotate90CW(bb)
		rbb = rotate90CCW(rbb)
		if rbb != bb {
			t.Error("Error while rotating board by 90 degrees:\n")
			DrawBitboard(t, rbb)
			DrawBitboard(t, bb)
		}
	}
}
