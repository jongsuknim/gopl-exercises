package main

import (
	"fmt"
	"io"
	"os"
)

type ByteCounter struct {
	writer  io.Writer
	written int64
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	// c.written += int64(len(p))
	// return len(p), nil
	n, err := c.writer.Write(p)
	if err != nil {
		return n, err
	}
	c.written += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := ByteCounter{w, 0}

	return &c, &c.written
}

func main() {
	var c ByteCounter

	c = ByteCounter{os.Stdout, 0}

	d, written := CountingWriter(&c)

	d.Write([]byte("hello\n"))

	var name = "Dolly"
	fmt.Fprintf(d, "hello, %s\n", name)
	fmt.Println(*written)
}
