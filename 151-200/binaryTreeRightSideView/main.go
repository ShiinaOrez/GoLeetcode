package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var lefts []int

func dfs(now *TreeNode, deep int) {
	if len(lefts) < deep {
		lefts = append(lefts, now.Val)
	}
	if now.Right != nil {
		dfs(now.Right, deep+1)
	}
	if now.Left != nil {
		dfs(now.Left, deep+1)
	}
	return
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	lefts = []int{}
	dfs(root, 1)
	return lefts
}

func main() {
	fmt.Println()
}
