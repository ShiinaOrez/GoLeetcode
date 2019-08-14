package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
    if head == nil {
        return nil
    }
    newArray := []int{}
    for head != nil {
        newArray = append(newArray, head.Val)
        head = head.Next
    }
    return sortedArrayToBST(newArray)
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
	fmt.Println()
}