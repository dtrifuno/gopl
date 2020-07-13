package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func DifferenceCount(data1 []byte, data2 []byte) int {
	sha1 := sha256.Sum256(data1)
	sha2 := sha256.Sum256(data2)
	count := 0
	for i := 0; i < sha256.Size; i++ {
		count += int(pc[sha1[i]^sha2[i]])
	}
	return count
}

func main() {
	s1 := "The quick brown fox jumped over the lazy dog."
	s2 := "The quick brown fox jumped over the lazy cat."
	fmt.Printf(
		"The SHA256 of\n\t\"%s\"\nand\n\t\"%s\"\ndiffer in %d bits.\n",
		s1, s2, DifferenceCount([]byte(s1), []byte(s2)),
	)
}
