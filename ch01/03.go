package main

import (
	"fmt"
	"strings"
	"time"
)

func join2(str []string, _sep string) string {
	s, sep := "", ""
	for _, arg := range str {
		s += sep + arg
		sep = _sep
	}
	return s
}

func test(f func([]string, string) string) float64 {
	start := time.Now()
	f(make([]string, 10000), ",")
	secs := time.Since(start).Seconds()
	return secs
}

func main() {
	fmt.Println(test(strings.Join))
	fmt.Println(test(join2))
}
