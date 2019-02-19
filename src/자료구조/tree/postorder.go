package main

import "fmt"

type Node struct {
	data        int
	left, right *Node
}

func postorder(root *Node) {
	if root != nil {
		postorder(root.left)
		postorder(root.right)
		fmt.Print(root.data, " ")
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

	postorder(&root)
}
