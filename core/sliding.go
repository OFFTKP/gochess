package core

func (board *Board) generateBishopMoves() uint64 {
	return 0
}

func getHorizontalSlide(square int, curRow uint8, rowIndex uint8) uint64 {
	return uint64(slidingMoves[square][curRow]) << 8 * uint64(rowIndex)
}

func generateSliding() {
	for i := 0; i < 8; i++ {
		// mask to remove our own piece
		curPieceMask := ^(1 << i)
		// 256 different occupancy possibilities per line
		for occ := 0; occ < 256; occ++ {
			var possibleMoves uint8
			possibleMoves &= uint8(curPieceMask)
			for bit := i + 2; bit < 8; bit++ {
				if (occ & (1 << bit)) != 0 {
					possibleMoves |= 1 << bit
					break
				} else {
					possibleMoves |= 1 << bit
				}
			}
			for bit := i; bit > 0; bit-- {
				if (occ & (1 << bit)) != 0 {
					possibleMoves |= 1 << bit
					break
				} else {
					possibleMoves |= 1 << bit
				}
			}
			for j := 0; j < 64; j += 8 {
				slidingMoves[i+j][occ] = possibleMoves
			}
		}
	}
}

var slidingMoves [64][256]uint8

var bishopMoveCountTable [64]uint8 = [64]uint8{
	0x7, 0x7, 0x7, 0x7, 0x7, 0x7, 0x7, 0x7,
	0x7, 0x9, 0x9, 0x9, 0x9, 0x9, 0x9, 0x7,
	0x7, 0x9, 0xB, 0xB, 0xB, 0xB, 0x9, 0x7,
	0x7, 0x9, 0xB, 0xD, 0xD, 0xB, 0x9, 0x7,
	0x7, 0x9, 0xB, 0xD, 0xD, 0xB, 0x9, 0x7,
	0x7, 0x9, 0xB, 0xB, 0xB, 0xB, 0x9, 0x7,
	0x7, 0x9, 0x9, 0x9, 0x9, 0x9, 0x9, 0x7,
	0x7, 0x7, 0x7, 0x7, 0x7, 0x7, 0x7, 0x7,
}
