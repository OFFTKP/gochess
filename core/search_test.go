package core

import "testing"

func BenchmarkDepthFunc(b *testing.B) {
	var board Board
	board.NewRoot(13)
}
