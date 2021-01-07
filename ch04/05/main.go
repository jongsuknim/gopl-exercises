package main

import "fmt"

func remDup(strings []string) []string {
	i := 0

	for _, s := range strings {
		if strings[i] == s {
			continue
		}
		i++
		strings[i] = s
	}
	return strings[:i+1]
}

func main() {
	a := []string{"a", "b", "b", "b", "c", "c", "d"}

	fmt.Println(remDup(a))
	b := []string{"a", "b", "b", "b", "c", "c"}
	fmt.Println(remDup(b))

}
