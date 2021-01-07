// Charcount는 유니코드 문자의 카운트를 계산한다
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	f, err := os.Open("main.go")
	if err != nil {
		fmt.Println("file Open Error", err)
		os.Exit(1)
	}
	scaner := bufio.NewScanner(f)
	scaner.Split(bufio.ScanWords)

	for scaner.Scan() {
		word := scaner.Text()
		counts[word]++
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
}
