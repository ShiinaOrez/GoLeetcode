package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
    n := len(inorder)
    if n == 0 {
        return nil
    }
    rootVal := postorder[n-1]
    for i:=0; i<n; i++ {
        if inorder[i] == rootVal {
            leftLength := i
            rightLength := n-1-i
            var leftNode, rightNode *TreeNode
            if leftLength == 0 {
                leftNode = nil
            } else {
                leftNode = buildTree(inorder[:i], postorder[:leftLength])
            }
            if rightLength == 0 {
                rightNode = nil
            } else {
                rightNode = buildTree(inorder[i+1:], postorder[leftLength:n-1])
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
	fmt.Println(buildTree([]int{9,3,15,20,7}, []int{9,15,7,20,3}))
}