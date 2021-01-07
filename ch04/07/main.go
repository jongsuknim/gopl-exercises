package main

import (
	"fmt"
)

func reverse(arr []byte) {
	r := []rune(string(arr))

	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	copy(arr, []byte(string(r)))
}

func main() {
	s := []byte("안녕하살법 입니다.")

	reverse(s[0:9])
	fmt.Printf("%q\n", s)
	reverse(s[:])
	fmt.Printf("%q\n", s)
}
