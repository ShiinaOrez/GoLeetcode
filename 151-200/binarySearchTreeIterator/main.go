package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var values []int

type BSTIterator struct {
	Values []int
	Length int
	Index  int
}

func Constructor(root *TreeNode) BSTIterator {
	values = []int{}
	if root == nil {
		return BSTIterator{
			Values: []int{},
			Length: 0,
			Index:  -1,
		}
	}
	travel(root)
	return BSTIterator{
		Values: values,
		Length: len(values),
		Index:  -1,
	}
}

func (this *BSTIterator) Next() int {
	this.Index += 1
	return this.Values[this.Index]
}

func (this *BSTIterator) HasNext() bool {
	return this.Index < this.Length-1
}

func travel(now *TreeNode) {
	if now.Left != nil {
		travel(now.Left)
	}
	values = append(values, now.Val)
	if now.Right != nil {
		travel(now.Right)
	}
}

func main() {
	fmt.Println()
}
