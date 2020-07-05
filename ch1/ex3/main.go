package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	concatStart := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	concatTime := time.Since(concatStart).Nanoseconds()
	fmt.Printf("concat: %d ns\n", concatTime)

	joinStart := time.Now()
	strings.Join(os.Args[1:], s)
	joinTime := time.Since(joinStart).Nanoseconds()
	fmt.Printf("strings.Join: %d ns\n", joinTime)
}
