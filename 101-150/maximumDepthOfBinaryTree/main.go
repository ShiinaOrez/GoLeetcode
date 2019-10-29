package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans int

func maxDepth(root *TreeNode) int {
	ans = 0
	if root == nil {
		return 0
	}
	dfs(root, 1)
	return ans
}

func dfs(now *TreeNode, deep int) {
	if deep > ans {
		ans = deep
	}
	if now.Left != nil {
		dfs(now.Left, deep+1)
	}
	if now.Right != nil {
		dfs(now.Right, deep+1)
	}
}

func main() {
	fmt.Println()
}
