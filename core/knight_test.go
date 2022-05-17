package core

import (
	"math/bits"
	"testing"
)

func TestKnightMoveCountEmpty(t *testing.T) {
	var board Board
	board.init()
	for i := 0; i < 64; i++ {
		*board.PieceBBmap[p_Knight] = uint64(1) << i
		res := bits.OnesCount64(board.generateKnightMoves())
		if res != int(knightMoveCountTable[i]) {
			t.Errorf("Knight move count failed for square %d", i)
		}
	}
}

func TestKnightMoveCountObstacles(t *testing.T) {
	type testCase struct {
		fen           string
		expectedMoves int
	}
	testCases := []testCase{
		// Test cases with fen and expected move count
		{"8/8/8/4p3/5P2/3N4/8/8 w - - 0 1", 7},
		{"8/8/8/2P1P3/1P3P2/3N4/1P3P2/2B1B3 w - - 0 1", 0},
		{"3n4/1p3p2/2p1P3/8/8/8/8/8 b - - 0 1", 1},
	}
	for i := 0; i < len(testCases); i++ {
		var board Board
		board.LoadFen(testCases[i].fen)
		res := board.generateKnightMoves()
		count := bits.OnesCount64(res)
		if count != testCases[i].expectedMoves {
			t.Errorf("Knight move count failed for i = %d\nExpected:%d\nGot     :%d", i, testCases[i].expectedMoves, count)
			board.DrawChessboard(t)
			DrawBitboard(t, res)
		}
	}
}
