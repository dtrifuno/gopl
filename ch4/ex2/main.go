package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	hashPtr := flag.String("hash", "sha256", "Hash {sha256|sha384|sha512}.")
	flag.Parse()

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		bytes := input.Bytes()
		switch *hashPtr {
		case "sha256":
			fmt.Printf("%x\n", sha256.Sum256(bytes))
		case "sha384":
			fmt.Printf("%x\n", sha512.Sum384(bytes))
		case "sha512":
			fmt.Printf("%x\n", sha512.Sum512(bytes))
		}
	}
}
