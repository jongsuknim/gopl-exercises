package main

import (
	"fmt"
	"github/jongsuknim/gopl-exercises/ch02/03/popcount"
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
	fmt.Printf("%v\n", perf(func() { popcount.PopCount(0xFFFFFFFFFFFFFFFF) }))
	fmt.Printf("%v\n", perf(func() { popcount.PopCount2(0xFFFFFFFFFFFFFFFF) }))
}
