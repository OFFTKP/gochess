package core

import (
	"math/bits"
	"testing"
)

func TestKnightMoveCount(t *testing.T) {
	for i := 0; i < 64; i++ {
		if bits.OnesCount64(knightMovesPerSquare[i]) != int(knightHeatTable[i]) {
			t.Errorf("Knight move count failed for square %d", i)
		}
	}
}
