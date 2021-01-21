package main

import "fmt"

func min(val int, vals ...int) int {
	m := val

	for _, v := range vals {
		if m > v {
			m = v
		}
	}
	return m
}

func max(val int, vals ...int) int {
	m := val

	for _, v := range vals {
		if m < v {
			m = v
		}
	}
	return m
}

func main() {
	fmt.Println(min(1, 2, 3, 4, 5), max(1, 2, 3, 4, 5))
}
