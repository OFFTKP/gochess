package core

import (
	"math/rand"
	"testing"
	"time"
)

func TestBBFlipVertical(t *testing.T) {
	var bb uint64 = 0x1e2222120e0a1222
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

func TestBBRotate45Unrotate(t *testing.T) {
	var testCases [15]uint64 = [15]uint64{
		0x8040201008040201,
		0x0080402010080402,
		0x0000804020100804,
		0x0000008040201008,
		0x0000000080402010,
		0x0000000000804020,
		0x0000000000008040,
		0x0000000000000080,
		0x4020100804020100,
		0x2010080402010000,
		0x1008040201000000,
		0x0804020100000000,
		0x0402010000000000,
		0x0201000000000000,
		0x0100000000000000,
	}
	for _, test := range testCases {
		bbr := rotate45CW(test)
		bbrr := unrotate45CW(bbr)
		if test != bbrr {
			DrawBitboard(t, test)
			DrawBitboard(t, bbrr)
		}
	}
}
