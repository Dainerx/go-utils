package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func RemoveLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return root
	}

	root.Left = RemoveLeafNodes(root.Left, target)
	root.Right = RemoveLeafNodes(root.Right, target)

	if (root.Left == nil && root.Right == nil) && root.Val == target {
		root = nil
		return root
	}

	return root
}

func rob(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	
}

func main() {
	l := &TreeNode{1, nil, nil}
	r := &TreeNode{1, nil, nil}
	root := &TreeNode{2, l, r}

	RemoveLeafNodes(root, 1)

	fmt.Println(root)
}
