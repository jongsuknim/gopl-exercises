package main

import "fmt"

type Tree struct {
	value       int
	left, right *Tree
}

func treeSort(values []int) {

	root := &Tree{values[0], nil, nil}
	for _, v := range values[1:] {
		appendTree(root, v)
	}

	appendValues(values[:0], root)
}

func appendTree(node *Tree, v int) *Tree {
	if node == nil {
		return &Tree{value: v, left: nil, right: nil}
	}

	if v < node.value {
		node.left = appendTree(node.left, v)
	} else {
		node.right = appendTree(node.right, v)
	}

	return node
}

func appendValues(values []int, node *Tree) []int {
	if node != nil {
		values = appendValues(values, node.left)
		values = append(values, node.value)
		values = appendValues(values, node.right)
	}

	return values
}

func main() {
	arr := []int{7, 43, 7, 3, 0, 0, 7, 5, 1, 7, 5, 8, 3, 2, 6, 78}
	treeSort(arr)
	fmt.Println(arr)
}
