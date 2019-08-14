package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

var ans int

func sumNumbers(root *TreeNode) int {
    if root == nil {
        return 0
    }
    ans = 0
    dfs(root, 0)
    return ans
}

func dfs(now *TreeNode, tot int) {
    newTot := tot * 10 + now.Val
    if now.Left == nil && now.Right == nil {
        ans += newTot
    } else {
        if now.Left != nil {
            dfs(now.Left, newTot)
        }
        if now.Right != nil {
            dfs(now.Right, newTot)
        }
    }
    return
}

func main() {
	fmt.Println()
}