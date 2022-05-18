package core

import (
	"math/bits"
	"testing"
)

func TestKingMoveCountEmpty(t *testing.T) {
	var board Board
	board.init()
	for i := 0; i < 64; i++ {
		res := board.generateKingMoves(i)
		count := bits.OnesCount64(res)
		if count != int(kingMoveCountTable[i]) {
			t.Errorf("King move count failed for square %d", i)
			board.DrawChessboard(t)
			DrawBitboard(t, res)
		}
	}
}
