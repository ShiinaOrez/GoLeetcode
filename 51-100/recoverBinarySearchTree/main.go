package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

var order []int
var nodes []*TreeNode

func recoverTree(root *TreeNode)  {
    if root == nil {
        return
    }
    order = []int{}
    nodes = []*TreeNode{}
    traversal(root)
    n := len(order)
    ordered := make([]int, n)
    copy(ordered, order)
    sort.Ints(ordered)
    
    fmt.Println(order)
    fmt.Println(ordered)
    
    for i:=0; i<n; i++ {
        if ordered[i] != order[i] {
            for j:=i+1; j<n; j++ {
                if ordered[j] != order[j] {
                    fmt.Println(root, nodes[i], nodes[j])
                    nodes[i].Val, nodes[j].Val = nodes[j].Val, nodes[i].Val
                    fmt.Println(root, nodes[i], nodes[j])
                    return
                }
            }
        }
    }
    fmt.Println(root.Val)
    
    return
}

func traversal(root *TreeNode) {
    if root == nil {
        return
    } else {
        traversal(root.Left)
        fmt.Println(root, "->", root.Val, root.Left, root.Right)
        order = append(order, root.Val)
        nodes = append(nodes, root)
        traversal(root.Right)
    }
    return
}

func main() {
	three := new(TreeNode)
	three.Val = 3
	three.Left = new(TreeNode)
	three.Left.Val = 1
	three.Right = new(TreeNode)
	three.Right.Val = 4
	three.Right.Left = new(TreeNode)
	three.Right.Left.Val = 2
	recoverTree(three)
}