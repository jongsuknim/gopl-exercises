package main

import (
	"fmt"
	"github/jongsuknim/gopl-exercises/ch02/05/popcount"
	"time"
)

func perf(f func()) float64 {
	start := time.Now()
	for i := 0; i < 100000; i++ {
		f()
	}
	return time.Since(start).Seconds()
}

func main() {
	fmt.Printf("%v\n", popcount.PopCount(0xFFFFFFFFFFFFEFFF))
	fmt.Printf("%v\n", popcount.PopCount2(0xFFFFFFFFFFFEFFFF))
	fmt.Printf("%v\n", perf(func() { popcount.PopCount(0xFFFFFFFFFFFFFFFF) }))
	fmt.Printf("%v\n", perf(func() { popcount.PopCount2(0xFFFFFFFFFFFFFFFF) }))
}
