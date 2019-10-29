package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	Elements []*TreeNode
	Length   int
}

type Layers struct {
	Map    map[int]Stack
	Length int
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	initMap := make(map[int]Stack)
	initMap[0] = Stack{[]*TreeNode{root}, 1}
	layers := Layers{initMap, 1}
	now := 0
	for now < layers.Length {
		layers.Map[now+1] = Stack{[]*TreeNode{}, 0}
		nodes := layers.Map[now]
		for i := nodes.Length - 1; i >= 0; i-- {
			if now%2 == 0 {
				if nodes.Elements[i].Right != nil {
					layers.Map[now+1] = Stack{append(layers.Map[now+1].Elements, nodes.Elements[i].Right), layers.Map[now+1].Length + 1}
				}
				if nodes.Elements[i].Left != nil {
					layers.Map[now+1] = Stack{append(layers.Map[now+1].Elements, nodes.Elements[i].Left), layers.Map[now+1].Length + 1}
				}
			} else {
				if nodes.Elements[i].Left != nil {
					layers.Map[now+1] = Stack{append(layers.Map[now+1].Elements, nodes.Elements[i].Left), layers.Map[now+1].Length + 1}
				}
				if nodes.Elements[i].Right != nil {
					layers.Map[now+1] = Stack{append(layers.Map[now+1].Elements, nodes.Elements[i].Right), layers.Map[now+1].Length + 1}
				}
			}
		}
		if layers.Map[now+1].Length > 0 {
			layers.Length += 1
		}
		now += 1
	}
	res := [][]int{}
	for i := 0; i < layers.Length; i++ {
		loop := []int{}
		for j := 0; j < layers.Map[i].Length; j++ {
			loop = append(loop, layers.Map[i].Elements[j].Val)
		}
		res = append(res, loop)
	}
	return res
}

func main() {
	fmt.Println()
}
