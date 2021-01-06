package main

import (
	"fmt"
)

func isAnagrams(s1, s2 string) bool {
	l1, l2 := len(s1), len(s2)
	if l1 != l2 {
		return false
	}

	if l1 == 0 {
		return true
	}

	if s1[l1-1] != s2[0] {
		return false
	}

	return isAnagrams(s1[:l1-1], s2[1:])
}

func main() {
	fmt.Println(isAnagrams("123", "321"))
	fmt.Println(isAnagrams("123", "231"))
	fmt.Println(isAnagrams("123", "123"))
}
