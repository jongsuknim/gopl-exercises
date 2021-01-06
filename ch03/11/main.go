package main

import (
	"fmt"
	"strings"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func commaWithPoint(s string) string {
	i := strings.Index(s, ".")
	if i == -1 {
		return comma(s)
	} else {
		return comma(s[:i]) + s[i:]
	}
}

func main() {
	fmt.Println(commaWithPoint("1234567890"))
	fmt.Println(commaWithPoint("1234567890.18238"))
	fmt.Println(commaWithPoint("123"))
	fmt.Println(commaWithPoint("123.2932"))
	fmt.Println(commaWithPoint("123456"))
}
