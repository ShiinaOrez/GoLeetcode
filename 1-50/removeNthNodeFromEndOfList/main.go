package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := 1
	h := head
	for head.Next != nil {
		head = head.Next
		length += 1
	}
	length -= n
	if length == 0 {
		return h.Next
	}
	head = h
	for i := 1; i < length; i++ {
		h = h.Next
	}
	h.Next = h.Next.Next
	return head
}

func main() {
	fmt.Println()
}
