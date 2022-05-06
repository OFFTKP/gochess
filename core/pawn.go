package core

import "math/bits"

// https://www.chessprogramming.org/Pawn_Pushes_(Bitboards)#Generalized_Push
func (board *Board) SinglePushTargets(color int) uint64 {
	pawnMap := board.PieceBBmap[color+p_Pawn]
	// theres a potentially faster way to do this with unsafe.Pointer, converting a bool to int
	isBlack := (color >> 1) & 1 // if its black (0b110) this gets set to 1
	return bits.RotateLeft64(*pawnMap, 8-(isBlack<<4)) & board.emptySquares
}

func (board *Board) DoublePushTargets(color int) uint64 {
	// TODO: this can be branchless in the future
	if color == 0 {
		singlePushs := board.SinglePushTargets(c_White)
		return nortOne(singlePushs) & board.emptySquares & rank4
	} else {
		singlePushs := board.SinglePushTargets(c_Black)
		return soutOne(singlePushs) & board.emptySquares & rank5
	}
}

func (board *Board) PawnsAbleToPush(color int) uint64 {
	pawnMap := board.PieceBBmap[color+p_Pawn]
	if color == 0 {
		return soutOne(board.emptySquares) & *pawnMap
	} else {
		return nortOne(board.emptySquares) & *pawnMap
	}
}
