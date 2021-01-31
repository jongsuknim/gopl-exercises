package main

import (
	"io"
	"os"
	"strings"
)

type sReader struct {
	r     io.Reader
	limit int64
}

func (s *sReader) Read(p []byte) (n int, err error) {
	q := make([]byte, s.limit)
	s.r.Read(q)

	copy(p, q)

	return len(p), io.EOF
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &sReader{r, n}
}

func main() {
	r := LimitReader(strings.NewReader("<html><head></head><body>hihi</body></html>"), 10)

	io.Copy(os.Stdout, r)
}
