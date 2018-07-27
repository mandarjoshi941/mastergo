package main

import (
	"fmt"
)

// IntegerTree is structure to carry tree of integer
type IntegerTree struct {
	Number      int
	Left, Right *IntegerTree
}

func populateTree(integerSlice []int) IntegerTree {
	var tree IntegerTree
	tree.Number = integerSlice[0]
	newslice := integerSlice[1:]
	for _, integer := range newslice {
		insert(&tree, integer)
	}
	return tree
}

func insert(tree *IntegerTree, val int) *IntegerTree {
	if tree == nil {
		newNode := IntegerTree{Number: val}
		return &newNode
	}
	if val < tree.Number {
		tree.Left = insert(tree.Left, val)
	} else if val > tree.Number {
		tree.Right = insert(tree.Right, val)
	}
	return tree
}

func main() {
	intSlice := []int{11, 6, 8, 19, 4, 10, 5, 17, 43, 49, 31}
	tree := populateTree(intSlice)
	printTree(tree)
}

func printTree(tree IntegerTree) {
	fmt.Print(tree.Number, ",")
	if tree.Left != nil {
		printTree(*tree.Left)
	}
	if tree.Right != nil {
		printTree(*tree.Right)
	}
}
