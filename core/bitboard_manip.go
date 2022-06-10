package core

import "math/bits"

// Contains helpful functions/constants from https://www.chessprogramming.org/General_Setwise_Operations
// Also read https://www.chessprogramming.org/Flipping_Mirroring_and_Rotating

const (
	notAFile uint64 = 0xfefefefefefefefe
	notHFile uint64 = 0x7f7f7f7f7f7f7f7f

	rank4 uint64 = 0x00000000FF000000
	rank5 uint64 = 0x000000FF00000000
)

func soutOne(b uint64) uint64 {
	return b >> 8
}

func nortOne(b uint64) uint64 {
	return b << 8
}

func eastOne(b uint64) uint64 {
	return (b << 1) & notAFile
}

func noEaOne(b uint64) uint64 {
	return (b << 9) & notAFile
}

func soEaOne(b uint64) uint64 {
	return (b >> 7) & notAFile
}

func westOne(b uint64) uint64 {
	return (b >> 1) & notHFile
}

func soWeOne(b uint64) uint64 {
	return (b >> 9) & notHFile
}

func noWeOne(b uint64) uint64 {
	return (b << 7) & notHFile
}

func flipVertically(b uint64) uint64 {
	return bits.ReverseBytes64(b)
}

func flipHorizontally(b uint64) uint64 {
	const k1 uint64 = 0x5555555555555555
	const k2 uint64 = 0x3333333333333333
	const k4 uint64 = 0x0f0f0f0f0f0f0f0f
	b = ((b >> 1) & k1) | ((b & k1) << 1)
	b = ((b >> 2) & k2) | ((b & k2) << 2)
	b = ((b >> 4) & k4) | ((b & k4) << 4)
	return b
}

func flipDiagA1H8(b uint64) uint64 {
	var t uint64
	const k1 uint64 = 0x5500550055005500
	const k2 uint64 = 0x3333000033330000
	const k4 uint64 = 0x0f0f0f0f00000000
	t = k4 & (b ^ (b << 28))
	b ^= t ^ (t >> 28)
	t = k2 & (b ^ (b << 14))
	b ^= t ^ (t >> 14)
	t = k1 & (b ^ (b << 7))
	b ^= t ^ (t >> 7)
	return b
}

func rotate90CW(b uint64) uint64 {
	return flipVertically(flipDiagA1H8(b))
}

func rotate90CCW(b uint64) uint64 {
	return flipDiagA1H8(flipVertically(b))
}

func rotate45CW(b uint64) uint64 {
	const k1 uint64 = 0xAAAAAAAAAAAAAAAA
	const k2 uint64 = 0xCCCCCCCCCCCCCCCC
	const k4 uint64 = 0xF0F0F0F0F0F0F0F0
	b ^= k1 & (b ^ bits.RotateLeft64(b, -8))
	b ^= k2 & (b ^ bits.RotateLeft64(b, -16))
	b ^= k4 & (b ^ bits.RotateLeft64(b, -32))
	return b
}

func unrotate45CW(b uint64) uint64 {
	const k1 uint64 = 0xAAAAAAAAAAAAAAAA
	const k2 uint64 = 0xCCCCCCCCCCCCCCCC
	const k4 uint64 = 0xF0F0F0F0F0F0F0F0
	b ^= k1 & (b ^ bits.RotateLeft64(b, 8))
	b ^= k2 & (b ^ bits.RotateLeft64(b, 16))
	b ^= k4 & (b ^ bits.RotateLeft64(b, 32))
	return b
}

func rotate45CCW(b uint64) uint64 {
	const k1 uint64 = 0x5555555555555555
	const k2 uint64 = 0x3333333333333333
	const k4 uint64 = 0x0f0f0f0f0f0f0f0f
	b ^= k1 & (b ^ bits.RotateLeft64(b, -8))
	b ^= k2 & (b ^ bits.RotateLeft64(b, -16))
	b ^= k4 & (b ^ bits.RotateLeft64(b, -32))
	return b
}
