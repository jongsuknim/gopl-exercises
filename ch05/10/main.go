package main

import (
	"fmt"
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
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	var visitAll func(m map[string][]string)
	seen := make(map[string]bool)

	var visitTarget func(m map[string][]string, target string)
	visitTarget = func(m map[string][]string, target string) {
		if seen[target] {
			return
		}
		seen[target] = true
		if v, exist := m[target]; exist {
			for _, item := range v {
				visitTarget(m, item)
			}
		}
		order = append(order, target)
	}

	visitAll = func(m map[string][]string) {
		for k, _ := range m {
			visitTarget(m, k)
		}
	}

	visitAll(m)
	return order
}
