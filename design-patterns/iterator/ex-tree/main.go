package main

import (
	"fmt"

	"github.com/faultyagatha/design-patterns/iterator/internal"
)

func main() {
	//   1
	//  / \
	// 2   3

	// in-order:  213
	// preorder:  123
	// postorder: 231

	root := internal.NewNode(1,
		internal.NewTerminalNode(2),
		internal.NewTerminalNode(3))
	it := internal.NewInOrderIterator(root)

	for it.MoveNext() {
		fmt.Printf("%d,", it.Current.Value)
	}
	fmt.Println("\b")

	t := internal.NewBinaryTree(root)
	for i := t.InOrder(); i.MoveNext(); {
		fmt.Printf("%d,", i.Current.Value)
	}
	fmt.Println("\b")
}
