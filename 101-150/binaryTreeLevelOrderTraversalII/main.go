package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

type Queue struct {
    Nodes    []*TreeNode
    Deeps    []int
    Length   int
}

func levelOrderBottom(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    queue := Queue{}
    queue.Nodes = []*TreeNode{root}
    queue.Deeps = []int{0}
    queue.Length = 1
    now := 0
    for now < queue.Length {
        if queue.Nodes[now].Left != nil {
            queue.Nodes = append(queue.Nodes, queue.Nodes[now].Left)
            queue.Deeps = append(queue.Deeps, queue.Deeps[now] + 1)
            queue.Length += 1
        }
        if queue.Nodes[now].Right != nil {
            queue.Nodes = append(queue.Nodes, queue.Nodes[now].Right)
            queue.Deeps = append(queue.Deeps, queue.Deeps[now] + 1)
            queue.Length += 1
        }
        now += 1
    }
    
    res := [][]int{}
    loop := []int{}
    before := -1
    for i:=0; i<queue.Length; i++ {
        if queue.Deeps[i] > before {
            res = append(res, loop)
            loop = []int{queue.Nodes[i].Val}
            before = queue.Deeps[i]
        } else if before == queue.Deeps[i] {
            loop = append(loop, queue.Nodes[i].Val)
            before = queue.Deeps[i]
        }
    }
    res = res[1:]
    res = append(res, loop)
    ans := [][]int{}
    for i:=len(res)-1; i>=0; i-- {
        ans = append(ans, res[i])
    }
    return ans
}

func main() {
	fmt.Println()
}