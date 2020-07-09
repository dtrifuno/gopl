package popcount

import (
	"testing"
)

func TestPopCouns(t *testing.T) {
	for i := uint64(0); i < 1000000; i++ {
		byByte := PopCountByByte(i)
		byByteWithLoop := PopCountByByteWithLoop(i)
		byBit := PopCountByBit(i)
		byRightmost := PopCountByRightmost(i)
		if byByte != byByteWithLoop ||
			byByteWithLoop != byBit ||
			byBit != byRightmost {
			t.Errorf("received different answers for %d: bbyte: %d, bbwloop: %d, bbit: %d, br: %d\n",
				i, byByte, byByteWithLoop, byBit, byRightmost)
		}
	}
}

func BenchmarkPopCountByByte(b *testing.B) {
	for i := uint64(0); i < 1000000; i++ {
		PopCountByByte(i)
	}
}

func BenchmarkPopCountByByteWithLoop(b *testing.B) {
	for i := uint64(0); i < 1000000; i++ {
		PopCountByByteWithLoop(i)
	}
}
func BenchmarkPopCountByBit(b *testing.B) {
	for i := uint64(0); i < 1000000; i++ {
		PopCountByBit(i)
	}
}
func BenchmarkPopCountByRightmost(b *testing.B) {
	for i := uint64(0); i < 1000000; i++ {
		PopCountByRightmost(i)
	}
}
