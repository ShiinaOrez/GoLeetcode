package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }
    c := deep(root.Left) - deep(root.Right)
    if -1 <= c &&  c <= 1 {
        if isBalanced(root.Left) && isBalanced(root.Right) {
            return true
        }
    }
    return false
}

func deep(now *TreeNode) int {
    if now == nil {
        return 0
    }
    return 1 + max(deep(now.Left), deep(now.Right))
} 

func max(x, y int) int {
    if x >= y {
        return x
    }
    return y
}

func main() {
	fmt.Println()
}