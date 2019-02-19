package main

import (
	"fmt"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

func inorder(root *Node) {
	if root != nil {
		inorder(root.left)
		fmt.Print(root.data, " ")
		inorder(root.right)
	}
}

func main() {

	var root Node

	root.data = 5
	root.left = &Node{data: 3, left: nil, right: nil}
	root.right = &Node{data: 7, left: nil, right: nil}

	root.left.left = &Node{data: 2, left: nil, right: nil}
	root.left.right = &Node{data: 5, left: nil, right: nil}

	root.right.right = &Node{data: 8, left: nil, right: nil}

	fmt.Println(root.data)
	fmt.Println(*root.left)
	fmt.Println(*root.right)

	inorder(&root)

}
