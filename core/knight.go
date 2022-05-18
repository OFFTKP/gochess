package core

func (board *Board) generateKnightMoves(square int) uint64 {
	return knightMovesPerSquare[square] & (^(*board.ColorBBmap[board.nextColor]))
}

var knightMovesPerSquare [64]uint64 = [64]uint64{
	0x0000000000020400, 0x0000000000050800, 0x00000000000a1100, 0x0000000000142200, 0x0000000000284400, 0x0000000000508800, 0x0000000000a01000, 0x0000000000402000,
	0x0000000002040004, 0x0000000005080008, 0x000000000a110011, 0x0000000014220022, 0x0000000028440044, 0x0000000050880088, 0x00000000a0100010, 0x0000000040200020,
	0x0000000204000402, 0x0000000508000805, 0x0000000a1100110a, 0x0000001422002214, 0x0000002844004428, 0x0000005088008850, 0x000000a0100010a0, 0x0000004020002040,
	0x0000020400040200, 0x0000050800080500, 0x00000a1100110a00, 0x0000142200221400, 0x0000284400442800, 0x0000508800885000, 0x0000a0100010a000, 0x0000402000204000,
	0x0002040004020000, 0x0005080008050000, 0x000a1100110a0000, 0x0014220022140000, 0x0028440044280000, 0x0050880088500000, 0x00a0100010a00000, 0x0040200020400000,
	0x0204000402000000, 0x0508000805000000, 0x0a1100110a000000, 0x1422002214000000, 0x2844004428000000, 0x5088008850000000, 0xa0100010a0000000, 0x4020002040000000,
	0x0400040200000000, 0x0800080500000000, 0x1100110a00000000, 0x2200221400000000, 0x4400442800000000, 0x8800885000000000, 0x100010a000000000, 0x2000204000000000,
	0x0004020000000000, 0x0008050000000000, 0x00110a0000000000, 0x0022140000000000, 0x0044280000000000, 0x0088500000000000, 0x0010a00000000000, 0x0020400000000000,
}

var knightMoveCountTable [64]uint8 = [64]uint8{
	2, 3, 4, 4, 4, 4, 3, 2,
	3, 4, 6, 6, 6, 6, 4, 3,
	4, 6, 8, 8, 8, 8, 6, 4,
	4, 6, 8, 8, 8, 8, 6, 4,
	4, 6, 8, 8, 8, 8, 6, 4,
	4, 6, 8, 8, 8, 8, 6, 4,
	3, 4, 6, 6, 6, 6, 4, 3,
	2, 3, 4, 4, 4, 4, 3, 2,
}

// Hash table of moves needed to reach a square (usage: [source][target])
var knightMoveCountToReach [64][64]uint8 = [64][64]uint8{
	{0, 3, 2, 3, 2, 3, 4, 5, 3, 4, 1, 2, 3, 4, 3, 4, 2, 1, 4, 3, 2, 3, 4, 5, 3, 2, 3, 2, 3, 4, 3, 4, 2, 3, 2, 3, 4, 3, 4, 5, 3, 4, 3, 4, 3, 4, 5, 4, 4, 3, 4, 3, 4, 5, 4, 5, 5, 4, 5, 4, 5, 4, 5, 6},
	{3, 0, 3, 2, 3, 2, 3, 4, 2, 3, 2, 1, 2, 3, 4, 3, 1, 2, 1, 4, 3, 2, 3, 4, 2, 3, 2, 3, 2, 3, 4, 3, 3, 2, 3, 2, 3, 4, 3, 4, 4, 3, 4, 3, 4, 3, 4, 5, 3, 4, 3, 4, 3, 4, 5, 4, 4, 5, 4, 5, 4, 5, 4, 5},
	{2, 3, 0, 3, 2, 3, 2, 3, 1, 2, 3, 2, 1, 2, 3, 4, 4, 1, 2, 1, 4, 3, 2, 3, 3, 2, 3, 2, 3, 2, 3, 4, 2, 3, 2, 3, 2, 3, 4, 3, 3, 4, 3, 4, 3, 4, 3, 4, 4, 3, 4, 3, 4, 3, 4, 5, 5, 4, 5, 4, 5, 4, 5, 4},
	{3, 2, 3, 0, 3, 2, 3, 2, 2, 1, 2, 3, 2, 1, 2, 3, 3, 4, 1, 2, 1, 4, 3, 2, 2, 3, 2, 3, 2, 3, 2, 3, 3, 2, 3, 2, 3, 2, 3, 4, 4, 3, 4, 3, 4, 3, 4, 3, 3, 4, 3, 4, 3, 4, 3, 4, 4, 5, 4, 5, 4, 5, 4, 5},
	{2, 3, 2, 3, 0, 3, 2, 3, 3, 2, 1, 2, 3, 2, 1, 2, 2, 3, 4, 1, 2, 1, 4, 3, 3, 2, 3, 2, 3, 2, 3, 2, 4, 3, 2, 3, 2, 3, 2, 3, 3, 4, 3, 4, 3, 4, 3, 4, 4, 3, 4, 3, 4, 3, 4, 3, 5, 4, 5, 4, 5, 4, 5, 4},
	{3, 2, 3, 2, 3, 0, 3, 2, 4, 3, 2, 1, 2, 3, 2, 1, 3, 2, 3, 4, 1, 2, 1, 4, 4, 3, 2, 3, 2, 3, 2, 3, 3, 4, 3, 2, 3, 2, 3, 2, 4, 3, 4, 3, 4, 3, 4, 3, 5, 4, 3, 4, 3, 4, 3, 4, 4, 5, 4, 5, 4, 5, 4, 5},
	{4, 3, 2, 3, 2, 3, 0, 3, 3, 4, 3, 2, 1, 2, 3, 2, 4, 3, 2, 3, 4, 1, 2, 1, 3, 4, 3, 2, 3, 2, 3, 2, 4, 3, 4, 3, 2, 3, 2, 3, 5, 4, 3, 4, 3, 4, 3, 4, 4, 5, 4, 3, 4, 3, 4, 3, 5, 4, 5, 4, 5, 4, 5, 4},
	{5, 4, 3, 2, 3, 2, 3, 0, 4, 3, 4, 3, 2, 1, 4, 3, 5, 4, 3, 2, 3, 4, 1, 2, 4, 3, 4, 3, 2, 3, 2, 3, 5, 4, 3, 4, 3, 2, 3, 2, 4, 5, 4, 3, 4, 3, 4, 3, 5, 4, 5, 4, 3, 4, 3, 4, 6, 5, 4, 5, 4, 5, 4, 5},
	{3, 2, 1, 2, 3, 4, 3, 4, 0, 3, 2, 3, 2, 3, 4, 5, 3, 2, 1, 2, 3, 4, 3, 4, 2, 1, 4, 3, 2, 3, 4, 5, 3, 2, 3, 2, 3, 4, 3, 4, 2, 3, 2, 3, 4, 3, 4, 5, 3, 4, 3, 4, 3, 4, 5, 4, 4, 3, 4, 3, 4, 5, 4, 5},
	{4, 3, 2, 1, 2, 3, 4, 3, 3, 0, 3, 2, 3, 2, 3, 4, 2, 3, 2, 1, 2, 3, 4, 3, 1, 2, 1, 4, 3, 2, 3, 4, 2, 3, 2, 3, 2, 3, 4, 3, 3, 2, 3, 2, 3, 4, 3, 4, 4, 3, 4, 3, 4, 3, 4, 5, 3, 4, 3, 4, 3, 4, 5, 4},
	{1, 2, 3, 2, 1, 2, 3, 4, 2, 3, 0, 3, 2, 3, 2, 3, 1, 2, 3, 2, 1, 2, 3, 4, 4, 1, 2, 1, 4, 3, 2, 3, 3, 2, 3, 2, 3, 2, 3, 4, 2, 3, 2, 3, 2, 3, 4, 3, 3, 4, 3, 4, 3, 4, 3, 4, 4, 3, 4, 3, 4, 3, 4, 5},
	{2, 1, 2, 3, 2, 1, 2, 3, 3, 2, 3, 0, 3, 2, 3, 2, 2, 1, 2, 3, 2, 1, 2, 3, 3, 4, 1, 2, 1, 4, 3, 2, 2, 3, 2, 3, 2, 3, 2, 3, 3, 2, 3, 2, 3, 2, 3, 4, 4, 3, 4, 3, 4, 3, 4, 3, 3, 4, 3, 4, 3, 4, 3, 4},
	{3, 2, 1, 2, 3, 2, 1, 2, 2, 3, 2, 3, 0, 3, 2, 3, 3, 2, 1, 2, 3, 2, 1, 2, 2, 3, 4, 1, 2, 1, 4, 3, 3, 2, 3, 2, 3, 2, 3, 2, 4, 3, 2, 3, 2, 3, 2, 3, 3, 4, 3, 4, 3, 4, 3, 4, 4, 3, 4, 3, 4, 3, 4, 3},
	{4, 3, 2, 1, 2, 3, 2, 1, 3, 2, 3, 2, 3, 0, 3, 2, 4, 3, 2, 1, 2, 3, 2, 1, 3, 2, 3, 4, 1, 2, 1, 4, 4, 3, 2, 3, 2, 3, 2, 3, 3, 4, 3, 2, 3, 2, 3, 2, 4, 3, 4, 3, 4, 3, 4, 3, 5, 4, 3, 4, 3, 4, 3, 4},
	{3, 4, 3, 2, 1, 2, 3, 4, 4, 3, 2, 3, 2, 3, 0, 3, 3, 4, 3, 2, 1, 2, 3, 2, 4, 3, 2, 3, 4, 1, 2, 1, 3, 4, 3, 2, 3, 2, 3, 2, 4, 3, 4, 3, 2, 3, 2, 3, 5, 4, 3, 4, 3, 4, 3, 4, 4, 5, 4, 3, 4, 3, 4, 3},
	{4, 3, 4, 3, 2, 1, 2, 3, 5, 4, 3, 2, 3, 2, 3, 0, 4, 3, 4, 3, 2, 1, 2, 3, 5, 4, 3, 2, 3, 4, 1, 2, 4, 3, 4, 3, 2, 3, 2, 3, 5, 4, 3, 4, 3, 2, 3, 2, 4, 5, 4, 3, 4, 3, 4, 3, 5, 4, 5, 4, 3, 4, 3, 4},
	{2, 1, 4, 3, 2, 3, 4, 5, 3, 2, 1, 2, 3, 4, 3, 4, 0, 3, 2, 3, 2, 3, 4, 5, 3, 2, 1, 2, 3, 4, 3, 4, 2, 1, 4, 3, 2, 3, 4, 5, 3, 2, 3, 2, 3, 4, 3, 4, 2, 3, 2, 3, 4, 3, 4, 5, 3, 4, 3, 4, 3, 4, 5, 4},
	{1, 2, 1, 4, 3, 2, 3, 4, 2, 3, 2, 1, 2, 3, 4, 3, 3, 0, 3, 2, 3, 2, 3, 4, 2, 3, 2, 1, 2, 3, 4, 3, 1, 2, 1, 4, 3, 2, 3, 4, 2, 3, 2, 3, 2, 3, 4, 3, 3, 2, 3, 2, 3, 4, 3, 4, 4, 3, 4, 3, 4, 3, 4, 5},
	{4, 1, 2, 1, 4, 3, 2, 3, 1, 2, 3, 2, 1, 2, 3, 4, 2, 3, 0, 3, 2, 3, 2, 3, 1, 2, 3, 2, 1, 2, 3, 4, 4, 1, 2, 1, 4, 3, 2, 3, 3, 2, 3, 2, 3, 2, 3, 4, 2, 3, 2, 3, 2, 3, 4, 3, 3, 4, 3, 4, 3, 4, 3, 4},
	{3, 4, 1, 2, 1, 4, 3, 2, 2, 1, 2, 3, 2, 1, 2, 3, 3, 2, 3, 0, 3, 2, 3, 2, 2, 1, 2, 3, 2, 1, 2, 3, 3, 4, 1, 2, 1, 4, 3, 2, 2, 3, 2, 3, 2, 3, 2, 3, 3, 2, 3, 2, 3, 2, 3, 4, 4, 3, 4, 3, 4, 3, 4, 3},
	{2, 3, 4, 1, 2, 1, 4, 3, 3, 2, 1, 2, 3, 2, 1, 2, 2, 3, 2, 3, 0, 3, 2, 3, 3, 2, 1, 2, 3, 2, 1, 2, 2, 3, 4, 1, 2, 1, 4, 3, 3, 2, 3, 2, 3, 2, 3, 2, 4, 3, 2, 3, 2, 3, 2, 3, 3, 4, 3, 4, 3, 4, 3, 4},
	{3, 2, 3, 4, 1, 2, 1, 4, 4, 3, 2, 1, 2, 3, 2, 1, 3, 2, 3, 2, 3, 0, 3, 2, 4, 3, 2, 1, 2, 3, 2, 1, 3, 2, 3, 4, 1, 2, 1, 4, 4, 3, 2, 3, 2, 3, 2, 3, 3, 4, 3, 2, 3, 2, 3, 2, 4, 3, 4, 3, 4, 3, 4, 3},
	{4, 3, 2, 3, 4, 1, 2, 1, 3, 4, 3, 2, 1, 2, 3, 2, 4, 3, 2, 3, 2, 3, 0, 3, 3, 4, 3, 2, 1, 2, 3, 2, 4, 3, 2, 3, 4, 1, 2, 1, 3, 4, 3, 2, 3, 2, 3, 2, 4, 3, 4, 3, 2, 3, 2, 3, 5, 4, 3, 4, 3, 4, 3, 4},
	{5, 4, 3, 2, 3, 4, 1, 2, 4, 3, 4, 3, 2, 1, 2, 3, 5, 4, 3, 2, 3, 2, 3, 0, 4, 3, 4, 3, 2, 1, 2, 3, 5, 4, 3, 2, 3, 4, 1, 2, 4, 3, 4, 3, 2, 3, 2, 3, 5, 4, 3, 4, 3, 2, 3, 2, 4, 5, 4, 3, 4, 3, 4, 3},
	{3, 2, 3, 2, 3, 4, 3, 4, 2, 1, 4, 3, 2, 3, 4, 5, 3, 2, 1, 2, 3, 4, 3, 4, 0, 3, 2, 3, 2, 3, 4, 5, 3, 2, 1, 2, 3, 4, 3, 4, 2, 1, 4, 3, 2, 3, 4, 5, 3, 2, 3, 2, 3, 4, 3, 4, 2, 3, 2, 3, 4, 3, 4, 5},
	{2, 3, 2, 3, 2, 3, 4, 3, 1, 2, 1, 4, 3, 2, 3, 4, 2, 3, 2, 1, 2, 3, 4, 3, 3, 0, 3, 2, 3, 2, 3, 4, 2, 3, 2, 1, 2, 3, 4, 3, 1, 2, 1, 4, 3, 2, 3, 4, 2, 3, 2, 3, 2, 3, 4, 3, 3, 2, 3, 2, 3, 4, 3, 4},
	{3, 2, 3, 2, 3, 2, 3, 4, 4, 1, 2, 1, 4, 3, 2, 3, 1, 2, 3, 2, 1, 2, 3, 4, 2, 3, 0, 3, 2, 3, 2, 3, 1, 2, 3, 2, 1, 2, 3, 4, 4, 1, 2, 1, 4, 3, 2, 3, 3, 2, 3, 2, 3, 2, 3, 4, 2, 3, 2, 3, 2, 3, 4, 3},
	{2, 3, 2, 3, 2, 3, 2, 3, 3, 4, 1, 2, 1, 4, 3, 2, 2, 1, 2, 3, 2, 1, 2, 3, 3, 2, 3, 0, 3, 2, 3, 2, 2, 1, 2, 3, 2, 1, 2, 3, 3, 4, 1, 2, 1, 4, 3, 2, 2, 3, 2, 3, 2, 3, 2, 3, 3, 2, 3, 2, 3, 2, 3, 4},
	{3, 2, 3, 2, 3, 2, 3, 2, 2, 3, 4, 1, 2, 1, 4, 3, 3, 2, 1, 2, 3, 2, 1, 2, 2, 3, 2, 3, 0, 3, 2, 3, 3, 2, 1, 2, 3, 2, 1, 2, 2, 3, 4, 1, 2, 1, 4, 3, 3, 2, 3, 2, 3, 2, 3, 2, 4, 3, 2, 3, 2, 3, 2, 3},
	{4, 3, 2, 3, 2, 3, 2, 3, 3, 2, 3, 4, 1, 2, 1, 4, 4, 3, 2, 1, 2, 3, 2, 1, 3, 2, 3, 2, 3, 0, 3, 2, 4, 3, 2, 1, 2, 3, 2, 1, 3, 2, 3, 4, 1, 2, 1, 4, 4, 3, 2, 3, 2, 3, 2, 3, 3, 4, 3, 2, 3, 2, 3, 2},
	{3, 4, 3, 2, 3, 2, 3, 2, 4, 3, 2, 3, 4, 1, 2, 1, 3, 4, 3, 2, 1, 2, 3, 2, 4, 3, 2, 3, 2, 3, 0, 3, 3, 4, 3, 2, 1, 2, 3, 2, 4, 3, 2, 3, 4, 1, 2, 1, 3, 4, 3, 2, 3, 2, 3, 2, 4, 3, 4, 3, 2, 3, 2, 3},
	{4, 3, 4, 3, 2, 3, 2, 3, 5, 4, 3, 2, 3, 4, 1, 2, 4, 3, 4, 3, 2, 1, 2, 3, 5, 4, 3, 2, 3, 2, 3, 0, 4, 3, 4, 3, 2, 1, 2, 3, 5, 4, 3, 2, 3, 4, 1, 2, 4, 3, 4, 3, 2, 3, 2, 3, 5, 4, 3, 4, 3, 2, 3, 2},
	{2, 3, 2, 3, 4, 3, 4, 5, 3, 2, 3, 2, 3, 4, 3, 4, 2, 1, 4, 3, 2, 3, 4, 5, 3, 2, 1, 2, 3, 4, 3, 4, 0, 3, 2, 3, 2, 3, 4, 5, 3, 2, 1, 2, 3, 4, 3, 4, 2, 1, 4, 3, 2, 3, 4, 5, 3, 2, 3, 2, 3, 4, 3, 4},
	{3, 2, 3, 2, 3, 4, 3, 4, 2, 3, 2, 3, 2, 3, 4, 3, 1, 2, 1, 4, 3, 2, 3, 4, 2, 3, 2, 1, 2, 3, 4, 3, 3, 0, 3, 2, 3, 2, 3, 4, 2, 3, 2, 1, 2, 3, 4, 3, 1, 2, 1, 4, 3, 2, 3, 4, 2, 3, 2, 3, 2, 3, 4, 3},
	{2, 3, 2, 3, 2, 3, 4, 3, 3, 2, 3, 2, 3, 2, 3, 4, 4, 1, 2, 1, 4, 3, 2, 3, 1, 2, 3, 2, 1, 2, 3, 4, 2, 3, 0, 3, 2, 3, 2, 3, 1, 2, 3, 2, 1, 2, 3, 4, 4, 1, 2, 1, 4, 3, 2, 3, 3, 2, 3, 2, 3, 2, 3, 4},
	{3, 2, 3, 2, 3, 2, 3, 4, 2, 3, 2, 3, 2, 3, 2, 3, 3, 4, 1, 2, 1, 4, 3, 2, 2, 1, 2, 3, 2, 1, 2, 3, 3, 2, 3, 0, 3, 2, 3, 2, 2, 1, 2, 3, 2, 1, 2, 3, 3, 4, 1, 2, 1, 4, 3, 2, 2, 3, 2, 3, 2, 3, 2, 3},
	{4, 3, 2, 3, 2, 3, 2, 3, 3, 2, 3, 2, 3, 2, 3, 2, 2, 3, 4, 1, 2, 1, 4, 3, 3, 2, 1, 2, 3, 2, 1, 2, 2, 3, 2, 3, 0, 3, 2, 3, 3, 2, 1, 2, 3, 2, 1, 2, 2, 3, 4, 1, 2, 1, 4, 3, 3, 2, 3, 2, 3, 2, 3, 2},
	{3, 4, 3, 2, 3, 2, 3, 2, 4, 3, 2, 3, 2, 3, 2, 3, 3, 2, 3, 4, 1, 2, 1, 4, 4, 3, 2, 1, 2, 3, 2, 1, 3, 2, 3, 2, 3, 0, 3, 2, 4, 3, 2, 1, 2, 3, 2, 1, 3, 2, 3, 4, 1, 2, 1, 4, 4, 3, 2, 3, 2, 3, 2, 3},
	{4, 3, 4, 3, 2, 3, 2, 3, 3, 4, 3, 2, 3, 2, 3, 2, 4, 3, 2, 3, 4, 1, 2, 1, 3, 4, 3, 2, 1, 2, 3, 2, 4, 3, 2, 3, 2, 3, 0, 3, 3, 4, 3, 2, 1, 2, 3, 2, 4, 3, 2, 3, 4, 1, 2, 1, 3, 4, 3, 2, 3, 2, 3, 2},
	{5, 4, 3, 4, 3, 2, 3, 2, 4, 3, 4, 3, 2, 3, 2, 3, 5, 4, 3, 2, 3, 4, 1, 2, 4, 3, 4, 3, 2, 1, 2, 3, 5, 4, 3, 2, 3, 2, 3, 0, 4, 3, 4, 3, 2, 1, 2, 3, 5, 4, 3, 2, 3, 4, 1, 2, 4, 3, 4, 3, 2, 3, 2, 3},
	{3, 4, 3, 4, 3, 4, 5, 4, 2, 3, 2, 3, 4, 3, 4, 5, 3, 2, 3, 2, 3, 4, 3, 4, 2, 1, 4, 3, 2, 3, 4, 5, 3, 2, 1, 2, 3, 4, 3, 4, 0, 3, 2, 3, 2, 3, 4, 5, 3, 2, 1, 2, 3, 4, 3, 4, 2, 1, 4, 3, 2, 3, 4, 5},
	{4, 3, 4, 3, 4, 3, 4, 5, 3, 2, 3, 2, 3, 4, 3, 4, 2, 3, 2, 3, 2, 3, 4, 3, 1, 2, 1, 4, 3, 2, 3, 4, 2, 3, 2, 1, 2, 3, 4, 3, 3, 0, 3, 2, 3, 2, 3, 4, 2, 3, 2, 1, 2, 3, 4, 3, 1, 2, 1, 4, 3, 2, 3, 4},
	{3, 4, 3, 4, 3, 4, 3, 4, 2, 3, 2, 3, 2, 3, 4, 3, 3, 2, 3, 2, 3, 2, 3, 4, 4, 1, 2, 1, 4, 3, 2, 3, 1, 2, 3, 2, 1, 2, 3, 4, 2, 3, 0, 3, 2, 3, 2, 3, 1, 2, 3, 2, 1, 2, 3, 4, 4, 1, 2, 1, 4, 3, 2, 3},
	{4, 3, 4, 3, 4, 3, 4, 3, 3, 2, 3, 2, 3, 2, 3, 4, 2, 3, 2, 3, 2, 3, 2, 3, 3, 4, 1, 2, 1, 4, 3, 2, 2, 1, 2, 3, 2, 1, 2, 3, 3, 2, 3, 0, 3, 2, 3, 2, 2, 1, 2, 3, 2, 1, 2, 3, 3, 4, 1, 2, 1, 4, 3, 2},
	{3, 4, 3, 4, 3, 4, 3, 4, 4, 3, 2, 3, 2, 3, 2, 3, 3, 2, 3, 2, 3, 2, 3, 2, 2, 3, 4, 1, 2, 1, 4, 3, 3, 2, 1, 2, 3, 2, 1, 2, 2, 3, 2, 3, 0, 3, 2, 3, 3, 2, 1, 2, 3, 2, 1, 2, 2, 3, 4, 1, 2, 1, 4, 3},
	{4, 3, 4, 3, 4, 3, 4, 3, 3, 4, 3, 2, 3, 2, 3, 2, 4, 3, 2, 3, 2, 3, 2, 3, 3, 2, 3, 4, 1, 2, 1, 4, 4, 3, 2, 1, 2, 3, 2, 1, 3, 2, 3, 2, 3, 0, 3, 2, 4, 3, 2, 1, 2, 3, 2, 1, 3, 2, 3, 4, 1, 2, 1, 4},
	{5, 4, 3, 4, 3, 4, 3, 4, 4, 3, 4, 3, 2, 3, 2, 3, 3, 4, 3, 2, 3, 2, 3, 2, 4, 3, 2, 3, 4, 1, 2, 1, 3, 4, 3, 2, 1, 2, 3, 2, 4, 3, 2, 3, 2, 3, 0, 3, 3, 4, 3, 2, 1, 2, 3, 2, 4, 3, 2, 3, 4, 1, 2, 1},
	{4, 5, 4, 3, 4, 3, 4, 3, 5, 4, 3, 4, 3, 2, 3, 2, 4, 3, 4, 3, 2, 3, 2, 3, 5, 4, 3, 2, 3, 4, 1, 2, 4, 3, 4, 3, 2, 1, 2, 3, 5, 4, 3, 2, 3, 2, 3, 0, 4, 3, 4, 3, 2, 1, 2, 3, 5, 4, 3, 2, 3, 4, 1, 2},
	{4, 3, 4, 3, 4, 5, 4, 5, 3, 4, 3, 4, 3, 4, 5, 4, 2, 3, 2, 3, 4, 3, 4, 5, 3, 2, 3, 2, 3, 4, 3, 4, 2, 1, 4, 3, 2, 3, 4, 5, 3, 2, 1, 2, 3, 4, 3, 4, 0, 3, 2, 3, 2, 3, 4, 5, 3, 2, 1, 2, 3, 4, 3, 4},
	{3, 4, 3, 4, 3, 4, 5, 4, 4, 3, 4, 3, 4, 3, 4, 5, 3, 2, 3, 2, 3, 4, 3, 4, 2, 3, 2, 3, 2, 3, 4, 3, 1, 2, 1, 4, 3, 2, 3, 4, 2, 3, 2, 1, 2, 3, 4, 3, 3, 0, 3, 2, 3, 2, 3, 4, 4, 3, 2, 1, 2, 3, 4, 3},
	{4, 3, 4, 3, 4, 3, 4, 5, 3, 4, 3, 4, 3, 4, 3, 4, 2, 3, 2, 3, 2, 3, 4, 3, 3, 2, 3, 2, 3, 2, 3, 4, 4, 1, 2, 1, 4, 3, 2, 3, 1, 2, 3, 2, 1, 2, 3, 4, 2, 3, 0, 3, 2, 3, 2, 3, 1, 2, 3, 2, 1, 2, 3, 4},
	{3, 4, 3, 4, 3, 4, 3, 4, 4, 3, 4, 3, 4, 3, 4, 3, 3, 2, 3, 2, 3, 2, 3, 4, 2, 3, 2, 3, 2, 3, 2, 3, 3, 4, 1, 2, 1, 4, 3, 2, 2, 1, 2, 3, 2, 1, 2, 3, 3, 2, 3, 0, 3, 2, 3, 2, 2, 1, 2, 3, 2, 1, 2, 3},
	{4, 3, 4, 3, 4, 3, 4, 3, 3, 4, 3, 4, 3, 4, 3, 4, 4, 3, 2, 3, 2, 3, 2, 3, 3, 2, 3, 2, 3, 2, 3, 2, 2, 3, 4, 1, 2, 1, 4, 3, 3, 2, 1, 2, 3, 2, 1, 2, 2, 3, 2, 3, 0, 3, 2, 3, 3, 2, 1, 2, 3, 2, 1, 2},
	{5, 4, 3, 4, 3, 4, 3, 4, 4, 3, 4, 3, 4, 3, 4, 3, 3, 4, 3, 2, 3, 2, 3, 2, 4, 3, 2, 3, 2, 3, 2, 3, 3, 2, 3, 4, 1, 2, 1, 4, 4, 3, 2, 1, 2, 3, 2, 1, 3, 2, 3, 2, 3, 0, 3, 2, 4, 3, 2, 1, 2, 3, 2, 1},
	{4, 5, 4, 3, 4, 3, 4, 3, 5, 4, 3, 4, 3, 4, 3, 4, 4, 3, 4, 3, 2, 3, 2, 3, 3, 4, 3, 2, 3, 2, 3, 2, 4, 3, 2, 3, 4, 1, 2, 1, 3, 4, 3, 2, 1, 2, 3, 2, 4, 3, 2, 3, 2, 3, 0, 3, 3, 4, 3, 2, 1, 2, 3, 4},
	{5, 4, 5, 4, 3, 4, 3, 4, 4, 5, 4, 3, 4, 3, 4, 3, 5, 4, 3, 4, 3, 2, 3, 2, 4, 3, 4, 3, 2, 3, 2, 3, 5, 4, 3, 2, 3, 4, 1, 2, 4, 3, 4, 3, 2, 1, 2, 3, 5, 4, 3, 2, 3, 2, 3, 0, 4, 3, 4, 3, 2, 1, 2, 3},
	{5, 4, 5, 4, 5, 4, 5, 6, 4, 3, 4, 3, 4, 5, 4, 5, 3, 4, 3, 4, 3, 4, 5, 4, 2, 3, 2, 3, 4, 3, 4, 5, 3, 2, 3, 2, 3, 4, 3, 4, 2, 1, 4, 3, 2, 3, 4, 5, 3, 4, 1, 2, 3, 4, 3, 4, 0, 3, 2, 3, 2, 3, 4, 5},
	{4, 5, 4, 5, 4, 5, 4, 5, 3, 4, 3, 4, 3, 4, 5, 4, 4, 3, 4, 3, 4, 3, 4, 5, 3, 2, 3, 2, 3, 4, 3, 4, 2, 3, 2, 3, 2, 3, 4, 3, 1, 2, 1, 4, 3, 2, 3, 4, 2, 3, 2, 1, 2, 3, 4, 3, 3, 0, 3, 2, 3, 2, 3, 4},
	{5, 4, 5, 4, 5, 4, 5, 4, 4, 3, 4, 3, 4, 3, 4, 5, 3, 4, 3, 4, 3, 4, 3, 4, 2, 3, 2, 3, 2, 3, 4, 3, 3, 2, 3, 2, 3, 2, 3, 4, 4, 1, 2, 1, 4, 3, 2, 3, 1, 2, 3, 2, 1, 2, 3, 4, 2, 3, 0, 3, 2, 3, 2, 3},
	{4, 5, 4, 5, 4, 5, 4, 5, 3, 4, 3, 4, 3, 4, 3, 4, 4, 3, 4, 3, 4, 3, 4, 3, 3, 2, 3, 2, 3, 2, 3, 4, 2, 3, 2, 3, 2, 3, 2, 3, 3, 4, 1, 2, 1, 4, 3, 2, 2, 1, 2, 3, 2, 1, 2, 3, 3, 2, 3, 0, 3, 2, 3, 2},
	{5, 4, 5, 4, 5, 4, 5, 4, 4, 3, 4, 3, 4, 3, 4, 3, 3, 4, 3, 4, 3, 4, 3, 4, 4, 3, 2, 3, 2, 3, 2, 3, 3, 2, 3, 2, 3, 2, 3, 2, 2, 3, 4, 1, 2, 1, 4, 3, 3, 2, 1, 2, 3, 2, 1, 2, 2, 3, 2, 3, 0, 3, 2, 3},
	{4, 5, 4, 5, 4, 5, 4, 5, 5, 4, 3, 4, 3, 4, 3, 4, 4, 3, 4, 3, 4, 3, 4, 3, 3, 4, 3, 2, 3, 2, 3, 2, 4, 3, 2, 3, 2, 3, 2, 3, 3, 2, 3, 4, 1, 2, 1, 4, 4, 3, 2, 1, 2, 3, 2, 1, 3, 2, 3, 2, 3, 0, 3, 2},
	{5, 4, 5, 4, 5, 4, 5, 4, 4, 5, 4, 3, 4, 3, 4, 3, 5, 4, 3, 4, 3, 4, 3, 4, 4, 3, 4, 3, 2, 3, 2, 3, 3, 4, 3, 2, 3, 2, 3, 2, 4, 3, 2, 3, 4, 1, 2, 1, 3, 4, 3, 2, 1, 2, 3, 2, 4, 3, 2, 3, 2, 3, 0, 3},
	{6, 5, 4, 5, 4, 5, 4, 5, 5, 4, 5, 4, 3, 4, 3, 4, 4, 5, 4, 3, 4, 3, 4, 3, 5, 4, 3, 4, 3, 2, 3, 2, 4, 3, 4, 3, 2, 3, 2, 3, 5, 4, 3, 2, 3, 4, 1, 2, 4, 3, 4, 3, 2, 1, 4, 3, 5, 4, 3, 2, 3, 2, 3, 0},
}
