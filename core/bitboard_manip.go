package core

import "math/bits"

// Contains helpful functions/constants from https://www.chessprogramming.org/General_Setwise_Operations

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