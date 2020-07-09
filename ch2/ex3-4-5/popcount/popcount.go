package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCountByByte(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountByByteWithLoop(x uint64) int {
	count := 0
	for i := 0; i < 8; i++ {
		count += int(pc[byte(x>>(i*8))])
	}
	return count
}

func PopCountByBit(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		count += int(x & 1)
		x = x >> 1
	}
	return count
}

func PopCountByRightmost(x uint64) int {
	count := 0
	for ; x != 0; x = x & (x - 1) {
		count++
	}
	return count
}
