package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

var res []int

func isValidBST(root *TreeNode) bool {
    if root == nil {
        return true
    }
    res = []int{}
    traversal(root)
    n := len(res)
    if n == 1 {
        return true
    }
    for i:=1; i<n; i++ {
        if res[i] <= res[i-1] {
            return false
        }
    }
    return true
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
