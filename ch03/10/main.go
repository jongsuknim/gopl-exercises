package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer

	m := len(s) % 3
	d := len(s) / 3

	fmt.Fprintf(&buf, "%s", s[:m])

	for i := 0; i < d; i++ {
		if m == 0 && i == 0 {
			fmt.Fprintf(&buf, "%s", s[m+i*3:m+(i+1)*3])
		} else {
			fmt.Fprintf(&buf, ",%s", s[m+i*3:m+(i+1)*3])
		}
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("1234567890"))
	fmt.Println(comma("123"))
	fmt.Println(comma("12"))
	fmt.Println(comma("123456"))
}
