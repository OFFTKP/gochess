package core

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strconv"
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
	res, moves := board.countLegalMoves(color)
	if res != expectedMoves {
		t.Errorf("\nSquare: %d\nExpected:%d\n     Got:%d\n   Table:\n%s", iteration, expectedMoves, res, board.Draw(&moves))
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
	expected := int(expectedTable[square])
	return sb.String(), expected
}

func TestMoveCountFile(t *testing.T) {
	file, err := os.Open("data/simple_positions.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	line := 0
	for scanner.Scan() {
		var board Board
		words := strings.Split(scanner.Text(), " ")
		board.LoadFen(words[0])
		var color uint8
		if words[1][0] == 'w' {
			color = c_White
		} else {
			color = c_Black
		}
		count, moves := board.countLegalMoves(color)
		expected, err := strconv.Atoi(words[6])
		if err != nil {
			panic("Could not parse expected move count")
		}
		if expected != count {
			t.Errorf("\nExpected:%d\n     Got:%d\n  Squares:\n%s", expected, count, board.Draw(&moves))
		}
		line++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
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
		fen, expected := generatePieceFen(i, &knightHeatTable, 'N')
		testMoveCount(t, fen, c_White, expected, i)
	}
}

func TestMoveCountBishop(t *testing.T) {
	for i := 0; i < 64; i++ {
		fen, expected := generatePieceFen(i, &bishopHeatTable, 'B')
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
