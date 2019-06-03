package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

var before *TreeNode

func flatten(root *TreeNode)  {
    if root == nil {
        return
    }
    preNode := TreeNode{}
    before = &preNode
    convert(root)
    fix(root)
    return
}

func convert(now *TreeNode) {
    if now == nil {
        return
    }
    before.Left = now
    before = now
    if now.Left != nil {
        convert(now.Left)
    }
    if now.Right != nil {
        convert(now.Right)
    }
    return
}

func fix(now *TreeNode) {
    if now == nil {
        return
    }
    now.Right = now.Left
    now.Left = nil
    fix(now.Right)
}

func main() {
	fmt.Println()
}