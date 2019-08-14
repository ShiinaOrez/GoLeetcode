package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

var ans []int

func postorderTraversal(root *TreeNode) []int {
    ans = []int{}
    if root == nil {
        return []int{}
    }
    dfs(root)
    return ans
}

func dfs(now *TreeNode) {
    if now.Left != nil {
        dfs(now.Left)
    }
    if now.Right != nil {
        dfs(now.Right)
    }
    ans = append(ans, now.Val)
    return
}

func main() {
    fmt.Println()
}
