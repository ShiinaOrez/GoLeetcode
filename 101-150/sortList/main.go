package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	return mergeSort(head)
}

func mergeSort(node *ListNode) *ListNode {
	if node == nil || node.Next == nil {
		return node
	}
	fast := node
	slow := node
	breakNode := node
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		breakNode = slow
		slow = slow.Next
	}
	breakNode.Next = nil
	left := mergeSort(node)
	mid := mergeSort(slow)
	return merge(left, mid)
}

func merge(left, mid *ListNode) *ListNode {
	if left == nil {
		return mid
	}
	if mid == nil {
		return left
	}
	if left.Val <= mid.Val {
		left.Next = merge(left.Next, mid)
		return left
	}
	mid.Next = merge(mid.Next, left)
	return mid
}

func main() {
	fmt.Println()
}
