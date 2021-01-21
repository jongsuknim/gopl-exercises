package main

import (
	"fmt"
)

func f1() (n int) {
	defer func() {
		if p := recover(); p != nil {
			n = p.(int)
		}
	}()
	panic(1)
}

func main() {
	n := f1()
	fmt.Println(n)
}
