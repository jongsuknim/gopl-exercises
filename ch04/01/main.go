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

func PopCount(b byte) int {
	return int(pc[b])
}

func diffBitCount(b1 *[32]byte, b2 *[32]byte) int {
	sum := 0
	for i := 0; i < 32; i++ {
		sum += PopCount(b1[i] ^ b2[i])
	}
	return sum
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("%d", diffBitCount(&c1, &c2))
}
