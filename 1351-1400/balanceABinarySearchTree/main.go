package main

import (
    "fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

var values []int

func travel(now *TreeNode) {
    if now.Left != nil {
        travel(now.Left)
    }
    values = append(values, now.Val)
    if now.Right != nil {
        travel(now.Right)
    }
}

func buildHeap(l, r int) *TreeNode {
    if l == r {
        return &TreeNode{
            Val: values[l],
        }
    }
    if l > r {
        return nil
    }
    rootIndex := (l+r) / 2
    root := &TreeNode{
        Val: values[rootIndex],
        Left: buildHeap(l, rootIndex-1),
        Right: buildHeap(rootIndex+1, r),
    }
    return root
}

func balanceBST(root *TreeNode) *TreeNode {
    values = []int{}
    if root == nil {
        return nil
    }
    travel(root)
    return  buildHeap(0, len(values)-1)
}

func main() {
    fmt.Println()
}
