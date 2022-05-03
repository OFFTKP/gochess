package core

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)

func testLoadFen(t *testing.T, fen string) {
	var board Board
	board.LoadFen(fen)
	res := board.GetFen()
	if res != fen {
		t.Errorf("\nExpected:%s\n     Got:%s", fen, res)
	}
}

func testMoveCount(t *testing.T, fen string, color uint8, expectedMoves int, iteration int) {
	var board Board
	board.LoadFen(fen)
	res := board.countLegalMoves(color)
	if res != expectedMoves {
		t.Errorf("\nSquare: %d\nExpected:%d\n     Got:%d", iteration, expectedMoves, res)
	}
}

func randomPiece() rune {
	var piece uint8 = 1 << rand.Intn(6)
	var color uint8 = 64 * uint8(rand.Intn(2))
	return getPieceChar(piece | color)
}

func generateFen() string {
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder
	for j := 0; j < 8; j++ {
		for i := 0; i < 8; i++ {
			sb.WriteRune(randomPiece())
			if i == 7 && j != 7 {
				sb.WriteRune('/')
			}
		}
	}
	return sb.String()
}

func generateKnightFen(knightSquare int) (string, int) {
	knightY := knightSquare / 8
	knightX := knightSquare % 8
	var sb strings.Builder
	for i := 0; i < 8; i++ {
		if i == knightY {
			r := '0' + rune(knightX)
			if r != '0' {
				sb.WriteRune(r)
			}
			sb.WriteRune('N')
			r2 := '8' - rune(knightX+1)
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
	expected := int(knightHeatTable[knightSquare])
	return sb.String(), expected
}

func TestFenStartpos(t *testing.T) {
	testLoadFen(t, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR")
}

func TestFenRandom(t *testing.T) {
	for i := 0; i < 100; i++ {
		randomFen := generateFen()
		testLoadFen(t, randomFen)
	}
}

func TestMoveCountKnight(t *testing.T) {
	for i := 0; i < 64; i++ {
		fen, expected := generateKnightFen(i)
		testMoveCount(t, fen, c_White, expected, i)
	}
}

func BenchmarkLoadFen(b *testing.B) {
	var board Board
	for i := 0; i < b.N; i++ {
		board.LoadFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR")
	}
}

func BenchmarkGetLegalMoves(b *testing.B) {
	var board Board
	board.LoadFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR")
	for i := 0; i < b.N; i++ {
		for j := uint8(0); j < 64; j++ {
			board.getLegalMoves(j)
		}
	}
}
