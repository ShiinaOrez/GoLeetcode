package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []int

func inorderTraversal(root *TreeNode) []int {
	res = []int{}
	traversal(root)
	return res
}

func traversal(root *TreeNode) {
	if root == nil {
		return
	} else {
		traversal(root.Left)
		res = append(res, root.Val)
		traversal(root.Right)
	}
	return
}

func main() {
	fmt.Println()
}
