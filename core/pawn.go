package core

import "math/bits"

// https://www.chessprogramming.org/Pawn_Pushes_(Bitboards)#Generalized_Push
func (board *Board) singlePushTargets(color int) uint64 {
	pawnMap := board.PieceBBmap[color+p_Pawn]
	// theres a potentially faster way to do this with unsafe.Pointer, converting a bool to int
	isBlack := (color >> 1) & 1 // if its black (0b110) this gets set to 1
	return bits.RotateLeft64(*pawnMap, 8-(isBlack<<4)) & board.emptySquares
}
