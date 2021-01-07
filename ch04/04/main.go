package main

import "fmt"

func rotate(s []int, k int) []int {
	i := len(s) - (k % len(s))

	return append(s[i:], s[:i]...)
}

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(rotate(a[:], 3))
	fmt.Println(rotate(a[:], 1))
	fmt.Println(rotate(a[:], 0))
	fmt.Println(rotate(a[:], 7))
	fmt.Println(rotate(a[:], 8))
}
