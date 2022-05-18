package core

func (board *Board) generateBishopMoves() uint64 {
	return 0
}

func (board *Board) generateRookMoves(square int) uint64 {
	bb := getHorizontalSlide(square, ^board.emptySquares)
	bb |= getVerticalSlide(square, ^board.emptySquares)
	bb &= (^(*board.ColorBBmap[board.nextColor]))
	return bb
}

func getHorizontalSlide(square int, bb uint64) uint64 {
	rowIndex := square & 0xf8
	curRow := (bb >> (uint64(rowIndex))) & 0xff
	return slidingHorizontal[square][curRow]
}

func getVerticalSlide(square int, bb uint64) uint64 {
	bbr := rotate90CW(bb)
	colIndex := 7 - (square & 0x7)
	curCol := (bbr >> (uint64(colIndex * 8))) & 0xff
	return slidingVertical[square][curCol]
}

func generateSliding() {
	for i := 0; i < 8; i++ {
		// mask to remove our own piece
		var curPieceMask uint64 = ^(1 << i)
		i_8 := i * 8
		// 256 different occupancy possibilities per line
		for occ := 0; occ < 256; occ++ {
			var possibleMoves uint64
			possibleMoves &= curPieceMask
			for bit := i + 1; bit < 8; bit++ {
				if (occ & (1 << bit)) != 0 {
					possibleMoves |= 1 << bit
					break
				} else {
					possibleMoves |= 1 << bit
				}
			}
			for bit := i - 1; bit >= 0; bit-- {
				if (occ & (1 << bit)) != 0 {
					possibleMoves |= 1 << bit
					break
				} else {
					possibleMoves |= 1 << bit
				}
			}
			for j := 0; j < 8; j += 1 {
				j_8 := j * 8
				slidingHorizontal[i+j_8][occ] = possibleMoves << j_8
				slidingVertical[i_8+(7-j)][occ] = rotate90CCW(slidingHorizontal[i+j_8][occ])
			}
		}
	}
}

var slidingHorizontal [64][256]uint64
var slidingVertical [64][256]uint64

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
