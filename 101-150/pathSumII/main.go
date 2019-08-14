package main

import (
	"sync"
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

var res [][]int
var lock *sync.Mutex

func pathSum(root *TreeNode, sum int) [][]int {
    if root == nil {
        return [][]int{}
    }
    res = [][]int{}
    lock = new(sync.Mutex)
    search(root, []int{}, 0, sum)
    return res
}

func next(node *TreeNode, line []int, score, target int) {
    search(node, line, score, target)
}

func search(now *TreeNode, line []int, score, target int) {
    //fmt.Println(line)
    if now == nil {
        if score == target {
            lock.Lock()
            defer lock.Unlock()
            
            res = append(res, line)
        }
        return
    }
    
    newLine := make([]int, len(line) + 1)
    copy(newLine, append(line, now.Val))
    
    if now.Left == nil && now.Right == nil {
        if target - score - now.Val == 0 {
            lock.Lock()
            defer lock.Unlock()
            
            res = append(res, newLine)
        }
        return
    }
    
    if now.Left != nil {
        next(now.Left, newLine, score + now.Val, target)
    }
    
    if now.Right != nil {
        next(now.Right, newLine, score + now.Val, target)
    }
    return
}

func main() {
	fmt.Println()
}