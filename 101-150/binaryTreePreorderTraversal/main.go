package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans []int

func preorderTraversal(root *TreeNode) []int {
	ans = []int{}
	if root == nil {
		return []int{}
	}
	dfs(root)
	return ans
}

func dfs(now *TreeNode) {
	ans = append(ans, now.Val)
	if now.Left != nil {
		dfs(now.Left)
	}
	if now.Right != nil {
		dfs(now.Right)
	}
	return
}

func main() {
	fmt.Println()
}
