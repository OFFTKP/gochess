package core

import (
	"strings"
	"testing"
)

func testLoadFen(t *testing.T, fen string) {
	var board Board
	loaded, err := board.LoadFen(fen)
	if !loaded {
		t.Error(err)
		return
	}
	res := board.GetFen()
	if res != fen {
		t.Errorf("\nExpected:%s\n     Got:%s", fen, res)
	}
}

func generatePieceFen(square int, expectedTable *[64]uint8, piece rune) (string, int) {
	pieceY := square / 8
	pieceX := square % 8
	var sb strings.Builder
	for i := 0; i < 8; i++ {
		if i == pieceY {
			r := '0' + rune(pieceX)
			if r != '0' {
				sb.WriteRune(r)
			}
			sb.WriteRune(piece)
			r2 := '8' - rune(pieceX+1)
			if r2 != '0' {
				sb.WriteRune(r2)
			}
			if i < 7 {
				sb.WriteRune('/')
			}
		} else {
			sb.WriteRune('8')
			if i < 7 {
				sb.WriteRune('/')
			}
		}
	}
	expected := 0
	if expectedTable != nil {
		expected = int(expectedTable[square])
	}
	return sb.String(), expected
}

func TestFenLoadAndGet(t *testing.T) {
	testLoadFen(t, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	testLoadFen(t, "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1")
	testLoadFen(t, "r1b1k3/pp1p1prp/q1n1pbpn/2p5/4P1BB/P1QP1N2/1PP1NPP1/R5RK w - - 0 0")
}

func TestAlgebraicToUint8(t *testing.T) {
	expected := []uint8{
		algebraicToUint8("a1"), uint8(0),
		algebraicToUint8("c3"), uint8(18),
		algebraicToUint8("h8"), uint8(63),
		algebraicToUint8("a8"), uint8(56),
		algebraicToUint8("h1"), uint8(7),
	}
	for i := 0; i < len(expected); i += 2 {
		if expected[i] != expected[i+1] {
			t.Errorf("Fail, square:%d", expected[i+1])
		}
	}
}

func TestUint8ToAlgebraic(t *testing.T) {
	expected := []string{
		uint8ToAlgebraic(0), "a1",
		uint8ToAlgebraic(18), "c3",
		uint8ToAlgebraic(63), "h8",
		uint8ToAlgebraic(56), "a8",
		uint8ToAlgebraic(7), "h1",
	}
	for i := 0; i < len(expected); i += 2 {
		if expected[i] != expected[i+1] {
			t.Errorf("Fail, square:%s", expected[i+1])
		}
	}
}

func TestMakeUnmakeMove(t *testing.T) {
	var board Board
	oldFen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	board.LoadFen(oldFen)
	for i := uint8(0); i < 8; i++ {
		oldZobrist := board.zobristHash
		oldEmpty := board.emptySquares
		ret := board.MakeMove(0, 8+i, 40+i)
		board.UnmakeMove(ret, 0, 8+i, 40+i)
		newFen := board.GetFen()
		if oldZobrist != board.zobristHash {
			t.Errorf("\nBad zobrist hash\nExpected:%016x\n      Got:%016x", oldZobrist, board.zobristHash)
		}
		if oldEmpty != board.emptySquares {
			t.Errorf("\nBad empty bitboard\nExpected:%016x\n      Got:%016x", oldEmpty, board.emptySquares)
		}
		if oldFen != newFen {
			t.Errorf("\nBad zobrist hash\nExpected:%s\n     Got:%s", oldFen, newFen)
		}
	}
}

func BenchmarkMakeUnmakeMove(b *testing.B) {
	var board Board
	board.LoadFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	for i := 0; i < b.N; i++ {
		for j := uint8(0); j < 7; j++ {
			ret := board.MakeMove(0, 8, 16+j)
			board.UnmakeMove(ret, 0, 8, 16+j)
		}
	}
}
