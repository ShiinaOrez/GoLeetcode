package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
    n := len(preorder)
    if n == 0 {
        return nil
    }
    rootVal := preorder[0]
    for i:=0; i<n; i++ {
        if inorder[i] == rootVal {
            leftLength := i
            rightLength := n-1-i
            var leftNode, rightNode *TreeNode
            if leftLength == 0 {
                leftNode = nil
            } else {
                leftNode = buildTree(preorder[1:leftLength+1], inorder[:i])
            }
            if rightLength == 0 {
                rightNode = nil
            } else {
                rightNode = buildTree(preorder[leftLength+1:], inorder[i+1:])
            }
            nowNode := &TreeNode{
                rootVal,
                leftNode,
                rightNode,
            }
            return nowNode
        }
    }
    return nil
}

func main() {
	fmt.Println(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
}