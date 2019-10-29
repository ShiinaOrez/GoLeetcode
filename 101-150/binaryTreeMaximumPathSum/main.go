package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var count int

func maxPathSum(root *TreeNode) int {
	count = -99999
	getMax(root)
	return count
}

func getMax(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := max(0, getMax(root.Left))
	right := max(0, getMax(root.Right))
	count = max(count, root.Val+left+right)
	return max(left, right) + root.Val
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println()
}
