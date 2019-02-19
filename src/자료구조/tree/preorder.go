package main

import "fmt"

type Node struct {
	data        int
	left, right *Node
}

func preorder(root *Node) {
	if root != nil {
		fmt.Print(root.data, " ")
		preorder(root.left)
		preorder(root.right)
	}

}

func main() {

	var root Node

	root.data = 1
	root.left = &Node{data: 2, left: nil, right: nil}
	root.right = &Node{data: 3, left: nil, right: nil}

	fmt.Println(root.data)
	fmt.Println(*root.left)
	fmt.Println(*root.right)

	preorder(&root)
}
