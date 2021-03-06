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

func (node *TreeNode) minimum() *TreeNode {

	for node.left != nil {
		node = node.left
	}

	return node
}
func (node *TreeNode) maximum() *TreeNode {

	for node.right != nil {
		node = node.right
	}

	return node
}

func (node *TreeNode) insert(input int) {

	var y *TreeNode
	y = nil
	x := node

	fmt.Println(x)
	for x != nil {
		y = x
		if input < x.data {
			x = x.left
		} else {
			x = x.right
		}
	}

	if y == nil {
		node = &TreeNode{data: input, left: nil, right: nil}
	} else if input < y.data {
		y.left = &TreeNode{data: input, left: nil, right: nil}
	} else {
		y.right = &TreeNode{data: input, left: nil, right: nil}
	}
}

func main() {

	var root *TreeNode

	root = &TreeNode{data: 15, left: nil, right: nil}

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
	minimum := root.minimum()
	max := root.maximum()

	fmt.Println(result)
	fmt.Println(minimum)
	fmt.Println(max)

	fmt.Println(root)

	root.insert(14)

	fmt.Println(root.left.right.right.right)
}
