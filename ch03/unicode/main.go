package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, 세상"

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	for i, r := range s {
		fmt.Printf("%d\t%c\n", i, r)
	}

	r := []rune(s)
	fmt.Printf("%x\n", s)
	fmt.Printf("%x\n", r)
	fmt.Println(r)
	fmt.Println(string(r))
}
