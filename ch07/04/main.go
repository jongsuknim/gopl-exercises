package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

type sReader struct {
	str string
}

func (s *sReader) Read(p []byte) (n int, err error) {
	size := copy(p, s.str)

	return size, io.EOF
}

func NewReader(str string) io.Reader {
	return &sReader{str}
}

func main() {
	n, err := html.Parse(NewReader("<html><head></head><body>hihi</body></html>"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(n)
}
