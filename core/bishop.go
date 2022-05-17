package core

func (board *Board) generateBishopMoves(color int) uint64 {
	return 0
}

var bishopHeatTable [64]uint8 = [64]uint8{
	0x7, 0x7, 0x7, 0x7, 0x7, 0x7, 0x7, 0x7,
	0x7, 0x9, 0x9, 0x9, 0x9, 0x9, 0x9, 0x7,
	0x7, 0x9, 0xB, 0xB, 0xB, 0xB, 0x9, 0x7,
	0x7, 0x9, 0xB, 0xD, 0xD, 0xB, 0x9, 0x7,
	0x7, 0x9, 0xB, 0xD, 0xD, 0xB, 0x9, 0x7,
	0x7, 0x9, 0xB, 0xB, 0xB, 0xB, 0x9, 0x7,
	0x7, 0x9, 0x9, 0x9, 0x9, 0x9, 0x9, 0x7,
	0x7, 0x7, 0x7, 0x7, 0x7, 0x7, 0x7, 0x7,
}
