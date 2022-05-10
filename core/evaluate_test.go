package core

import (
	"math"
	"sync"
	"testing"
)

func testTrickyEvaluation(t *testing.T, fen string, depth int, minimumEval int) {
	var board Board
	board.LoadFen(fen)
	eval := board.NewRoot(depth)
	if (minimumEval > 0 && eval < minimumEval) || (minimumEval <= 0 && eval > minimumEval) {
		t.Errorf("\nFailed testTrickyEvaluation at %s\nMinimum evaluation:%d\nEvaluation:%d", fen, minimumEval, eval)
	}
}

func TestObviousEvaluationW(t *testing.T) {
	var board Board
	board.LoadFen("rnb1kbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	eval := board.evaluate()
	if eval <= 0 {
		t.Errorf("TestObviousEvaluationW failed:%d", eval)
	}
}

func TestObviousEvaluationB(t *testing.T) {
	var board Board
	board.LoadFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNB1KBNR w KQkq - 0 1")
	eval := board.evaluate()
	if eval >= 0 {
		t.Errorf("TestObviousEvaluationB failed:%d", eval)
	}
}

func TestTrickyEvaluations(t *testing.T) {
	type testCase struct {
		fen         string
		depth       int
		minimumEval int
	}
	testCases := []testCase{
		// Test cases with fen, depth to test and minimum (or maximum) eval
		// -inf or inf means black or white checkmate accordingly

		// Blackburne Shilling Gambit accepted
		{"r1bqkbnr/pppp1ppp/8/4N3/2BnP3/8/PPPP1PPP/RNBQK2R b KQkq - 0 4", 5, -50},
		// Légall Mate
		{"r2qkbnr/ppp2ppp/2np4/4N3/2B1P3/2N4P/PPPP1PP1/R1BbK2R w KQkq - 0 6", 3, math.MaxInt32},
	}
	var wg sync.WaitGroup
	for _, test := range testCases {
		wg.Add(1)
		go func(test testCase) {
			defer wg.Done()
			testTrickyEvaluation(t, test.fen, test.depth, test.minimumEval)
		}(test)
	}
	wg.Wait()
}
