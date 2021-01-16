package main

import (
	"bufio"
	"fmt"
	"strings"
)

func expand(s string, f func(string) string) string {
	scaner := bufio.NewScanner(strings.NewReader(s))
	scaner.Split(bufio.ScanWords)

	m := make(map[string]string)
	for scaner.Scan() {
		text := scaner.Text()
		if text[0] == '$' {
			m[text] = f(text[1:])
		}
	}

	retStr := s
	for k, v := range m {
		retStr = strings.ReplaceAll(retStr, k, v)
	}
	return retStr
}

func main() {
	fmt.Println(expand("asdf $foo sdf $bar", strings.ToUpper))
}
