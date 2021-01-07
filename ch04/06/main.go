package main

import (
	"fmt"
	"unicode"
)

func remDupSpace(arr []byte) []byte {
	r := []rune(string(arr))

	i := 0
	var last rune
	for k, v := range r {
		if k == 0 {
			i += len([]byte(string(v))) - 1
			last = v
			continue
		}

		if last == v && unicode.IsSpace(v) {
			continue
		}

		bytes := []byte(string(v))
		for _, b := range bytes {
			i++
			arr[i] = b
			last = v
		}
	}

	return arr[:i+1]
}

func main() {
	s := []byte("안  녕세   계")

	fmt.Printf("%q\n", remDupSpace(s[:]))
	fmt.Printf("%q\n", s)
}
