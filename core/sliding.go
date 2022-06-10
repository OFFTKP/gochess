package core

import "math/bits"

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

func getDiagonalOccupancy(square int, bb uint64) uint8 {
	var translation int = int(diagonalTranslation[square])
	bbd := diagonalBitboards[translation]
	var occ uint8 = uint8(bits.RotateLeft64(rotate45CW(bbd), 8*translation)) & diagonalMasks[translation]
	return occ
}

func generateSliding() {
	for i := 0; i < 8; i++ {
		// mask to remove our own piece
		i_8 := i * 8
		// 256 different occupancy possibilities per line
		for occ := 0; occ < 256; occ++ {
			var possibleMoves uint64
			//var possibleMovesDiagonal uint64
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
	// Diagonal generation
	for d := 0; d <= 7; d++ {
		for occ := 0; occ < 256; occ++ {
			slidingDiagonal[d*9][occ] = unrotate45CW(slidingHorizontal[d][occ])
		}
	}
}

var slidingHorizontal [64][256]uint64
var slidingVertical [64][256]uint64
var slidingDiagonal [64][256]uint64
var slidingAntiDiagonal [64][256]uint64

var diagonalTranslation [64]uint8 = [64]uint8{
	0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7,
	0xF, 0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6,
	0xE, 0xF, 0x0, 0x1, 0x2, 0x3, 0x4, 0x5,
	0xD, 0xE, 0xF, 0x0, 0x1, 0x2, 0x3, 0x4,
	0xC, 0xD, 0xE, 0xF, 0x0, 0x1, 0x2, 0x3,
	0xB, 0xC, 0xD, 0xE, 0xF, 0x0, 0x1, 0x2,
	0xA, 0xB, 0xC, 0xD, 0xE, 0xF, 0x0, 0x1,
	0x9, 0xA, 0xB, 0xC, 0xD, 0xE, 0xF, 0x0,
}

var diagonalMasks [16]uint8 = [16]uint8{
	0b1111_1111,
	0b1111_1110,
	0b1111_1100,
	0b1111_1000,
	0b1111_0000,
	0b1110_0000,
	0b1100_0000,
	0b1000_0000,

	0, // No '8' diagonal

	0b0000_0001,
	0b0000_0011,
	0b0000_0111,
	0b0000_1111,
	0b0001_1111,
	0b0011_1111,
	0b0111_1111,
}

var diagonalBitboards [64]uint64 = [64]uint64{
	0x8040201008040201,
	0x0080402010080402,
	0x0000804020100804,
	0x0000008040201008,
	0x0000000080402010,
	0x0000000000804020,
	0x0000000000008040,
	0x0000000000000080,

	0, // No '8' diagonal

	0x0100000000000000,
	0x0201000000000000,
	0x0402010000000000,
	0x0804020100000000,
	0x1008040201000000,
	0x2010080402010000,
	0x4020100804020100,
}

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
