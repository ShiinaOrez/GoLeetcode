package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return isSameNode(p, q)
}

func isSameNode(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	} else if p.Val != q.Val {
		return false
	} else if isSameNode(p.Left, q.Left) && isSameNode(p.Right, q.Right) {
		return true
	}
	return false
}

func main() {
	fmt.Println()
}
