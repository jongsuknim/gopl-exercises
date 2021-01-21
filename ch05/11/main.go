package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
	"linear algebra":        {"calculus"},
}

func main() {
	if r, err := topoSort(prereqs); err == nil {
		for i, course := range r {
			fmt.Printf("%d:\t%s\n", i+1, course)
		}
	} else {
		fmt.Println(err)
	}
}

func push(s *[]string, str string) {
	*s = append(*s, str)
}

func pop(s *[]string) (string, error) {
	if len(*s) == 0 {
		return "", fmt.Errorf("empty")
	}
	last := len(*s) - 1
	retval := (*s)[last]
	*s = (*s)[:last]
	return retval, nil
}

func isExist(s []string, str string) bool {
	for i := range s {
		if s[i] == str {
			return true
		}
	}
	return false
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string) error
	stack := make([]string, 0)

	visitAll = func(items []string) error {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				push(&stack, item)
				if err := visitAll(m[item]); err != nil {
					return err
				}
				pop(&stack)
				order = append(order, item)
			} else if isExist(stack, item) {
				return fmt.Errorf("colision %s", item)
			}
		}
		return nil
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	err := visitAll(keys)
	return order, err
}
