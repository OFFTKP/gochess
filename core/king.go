package core

func (board *Board) generateKingMoves(square int) uint64 {
	return kingMovesPerSquare[square] & (^(*board.ColorBBmap[board.nextColor]))
}

var kingMovesPerSquare [64]uint64 = [64]uint64{
	0x0000000000000302, 0x0000000000000705, 0x0000000000000e0a, 0x0000000000001c14, 0x0000000000003828, 0x0000000000007050, 0x000000000000e0a0, 0x000000000000c040,
	0x0000000000030203, 0x0000000000070507, 0x00000000000e0a0e, 0x00000000001c141c, 0x0000000000382838, 0x0000000000705070, 0x0000000000e0a0e0, 0x0000000000c040c0,
	0x0000000003020300, 0x0000000007050700, 0x000000000e0a0e00, 0x000000001c141c00, 0x0000000038283800, 0x0000000070507000, 0x00000000e0a0e000, 0x00000000c040c000,
	0x0000000302030000, 0x0000000705070000, 0x0000000e0a0e0000, 0x0000001c141c0000, 0x0000003828380000, 0x0000007050700000, 0x000000e0a0e00000, 0x000000c040c00000,
	0x0000030203000000, 0x0000070507000000, 0x00000e0a0e000000, 0x00001c141c000000, 0x0000382838000000, 0x0000705070000000, 0x0000e0a0e0000000, 0x0000c040c0000000,
	0x0003020300000000, 0x0007050700000000, 0x000e0a0e00000000, 0x001c141c00000000, 0x0038283800000000, 0x0070507000000000, 0x00e0a0e000000000, 0x00c040c000000000,
	0x0302030000000000, 0x0705070000000000, 0x0e0a0e0000000000, 0x1c141c0000000000, 0x3828380000000000, 0x7050700000000000, 0xe0a0e00000000000, 0xc040c00000000000,
	0x0203000000000000, 0x0507000000000000, 0x0a0e000000000000, 0x141c000000000000, 0x2838000000000000, 0x5070000000000000, 0xa0e0000000000000, 0x40c0000000000000,
}

var kingMoveCountTable [64]uint8 = [64]uint8{
	3, 5, 5, 5, 5, 5, 5, 3,
	5, 8, 8, 8, 8, 8, 8, 5,
	5, 8, 8, 8, 8, 8, 8, 5,
	5, 8, 8, 8, 8, 8, 8, 5,
	5, 8, 8, 8, 8, 8, 8, 5,
	5, 8, 8, 8, 8, 8, 8, 5,
	5, 8, 8, 8, 8, 8, 8, 5,
	3, 5, 5, 5, 5, 5, 5, 3,
}
