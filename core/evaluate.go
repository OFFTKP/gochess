package core

import "math/bits"

func (board *Board) evaluate() int {
	ret := 0
	// Simple piece power counting
	for i := 0; i < 6; i++ {
		power := piecePower[i]
		whitePower := power * bits.OnesCount64(*board.PieceBBmap[i])
		blackPower := -power * bits.OnesCount64(*board.PieceBBmap[c_Black+i])
		ret += whitePower + blackPower
	}
	return ret
}
