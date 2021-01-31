package main

import (
	"fmt"
	"sort"
)

func IsPalindrome(s sort.Interface) bool {
	for i := 0; i < s.Len()/2; i++ {
		j := s.Len() - i - 1
		if !(!s.Less(i, j) && !s.Less(j, i)) {
			return false
		}
	}
	return true
}

func main() {

	fmt.Println(IsPalindrome(sort.IntSlice([]int{1, 2, 3, 2, 1})))
	fmt.Println(IsPalindrome(sort.IntSlice([]int{1, 2, 3, 3, 2, 1})))
	fmt.Println(IsPalindrome(sort.IntSlice([]int{1, 2, 3, 1, 1})))
	fmt.Println(IsPalindrome(sort.IntSlice([]int{1, 2, 1, 3, 2, 1})))
}
