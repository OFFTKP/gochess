package core

import (
	"math/bits"
	"testing"
)

func TestBishopMoveCountEmpty(t *testing.T) {
	var board Board
	board.init()
	for i := 0; i < 64; i++ {
		*board.PieceBBmap[p_Knight] = uint64(1) << i
		res := bits.OnesCount64(board.generateBishopMoves(c_White))
		if res != int(bishopHeatTable[i]) {
			t.Errorf("Bishop move count failed for square %d", i)
		}
	}
}
