package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}

	rootVal := nums[n/2]
	nowNode := TreeNode{
		rootVal,
		sortedArrayToBST(nums[:n/2]),
		sortedArrayToBST(nums[n/2+1:]),
	}
	return &nowNode
}

func main() {
	array := []int{-10, -3, 3, 4, 6}
	fmt.Println(sortedArrayToBST(array))
}
