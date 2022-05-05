package core

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
