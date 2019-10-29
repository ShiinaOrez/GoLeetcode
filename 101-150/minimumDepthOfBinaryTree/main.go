package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	Point *TreeNode
	Deep  int
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]Node, 0)
	queue = append(queue, Node{
		root, 1,
	})
	now := 0
	for !(queue[now].Point.Left == nil && queue[now].Point.Right == nil) {
		if queue[now].Point.Left != nil {
			queue = append(queue, Node{
				queue[now].Point.Left,
				queue[now].Deep + 1,
			})
		}
		if queue[now].Point.Right != nil {
			queue = append(queue, Node{
				queue[now].Point.Right,
				queue[now].Deep + 1,
			})
		}
		now += 1
	}
	return queue[now].Deep
}

func main() {
	fmt.Println()
}
