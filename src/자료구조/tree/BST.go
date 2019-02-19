package main

import "fmt"

type TreeNode struct {
	data        int
	left, right *TreeNode
}

func (node *TreeNode) search(key int) *TreeNode {

	if node == nil || key == node.data {
		return node
	}
	if key < node.data {
		return node.left.search(key)
	}
	return node.right.search(key)

}

func main() {

	var root TreeNode

	root = TreeNode{data: 15, left: nil, right: nil}

	root.left = &TreeNode{data: 6, left: nil, right: nil}
	root.right = &TreeNode{data: 18, left: nil, right: nil}

	root.left.left = &TreeNode{data: 3, left: nil, right: nil}
	root.left.right = &TreeNode{data: 7, left: nil, right: nil}

	root.left.left.left = &TreeNode{data: 2, left: nil, right: nil}
	root.left.left.right = &TreeNode{data: 4, left: nil, right: nil}

	root.left.right.right = &TreeNode{data: 13, left: nil, right: nil}
	root.left.right.right.left = &TreeNode{data: 9, left: nil, right: nil}

	root.right.left = &TreeNode{data: 17, left: nil, right: nil}
	root.right.right = &TreeNode{data: 20, left: nil, right: nil}

	result := root.search(13)

	fmt.Println(result)
}
