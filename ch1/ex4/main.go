package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	appearsIn := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, appearsIn)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, appearsIn)
			f.Close()
		}
	}
	for line, files := range appearsIn {
		if len(files) > 1 {
			fmt.Printf("%s\t%v\n", line, appearsIn[line])
		}
	}
}

func countLines(f *os.File, appearsIn map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		s := input.Text()
		appearsIn[s] = append(appearsIn[s], f.Name())
	}
}
