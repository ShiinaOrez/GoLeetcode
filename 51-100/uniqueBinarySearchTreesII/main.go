package main

import  "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

var mask []bool

func generateTrees(n int) []*TreeNode {
    mask = make([]bool, n+1)
    return planting(0, 1, n)
}

func planting(root, left, right int) []*TreeNode {
    forest := []*TreeNode{nil}
    newLeftBranch := []*TreeNode{}
    newRightBranch := []*TreeNode{}
    for i:=left; i<=right; i++ {
        if !mask[i] {
            mask[i] = true
            newLeftBranch = planting(i, left, i-1)
            newRightBranch = planting(i, i+1, right)
            for _, leftBranch := range newLeftBranch {
                for _, rightBranch := range newRightBranch {
                    newRoot := new(TreeNode)
                    newRoot.Val = i
                    newRoot.Left = leftBranch
                    newRoot.Right = rightBranch
                    forest = append(forest, newRoot)
                }
            }
            mask[i] = false
        }
    }
    if len(forest) > 1 {
        return forest[1:]
    }
    return forest
}

func print(node *TreeNode) {
    fmt.Println(node.Val)
    if node.Left != nil {
        fmt.Printf("left: ")
        print(node.Left)
    }
    if node.Right != nil {
        fmt.Printf("right:")
        print(node.Right)
    }
}

func main() {
    fmt.Println(generateTrees(3))
}
