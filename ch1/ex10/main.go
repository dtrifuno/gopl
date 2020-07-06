package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	filename := os.Args[1]
	f, err := os.Create(filename)
	if err != nil {
		fmt.Printf("while creating %s: %v\n", filename, err)
		return
	}

	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[2:] {
		go fetch(url, ch)
	}

	for range os.Args[2:] {
		line := <-ch + "\n"
		_, err = f.WriteString(line)
		if err != nil {
			fmt.Printf("while writing to %s: %v\n", filename, err)
		}
	}
	f.WriteString(fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds()))
	f.Close()
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
