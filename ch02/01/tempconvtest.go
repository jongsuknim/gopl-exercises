package main

import (
	"github/jongsuknim/gopl-exercises/ch02/01/tempconv"
	"fmt"
)

func main() {
	fmt.Printf("Brrrr! %v\n", tempconv.ZeroK)
	fmt.Printf("Brrrr! %v\n", tempconv.KToC(tempconv.ZeroK))
}