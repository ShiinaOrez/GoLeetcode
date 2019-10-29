package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	var Head, Now *ListNode
	Head = new(ListNode)
	Now = Head
	for l1 != nil || l2 != nil {
		Now.Next = new(ListNode)
		if l1 == nil {
			Now.Val = l2.Val
			l2 = l2.Next
		} else if l2 == nil {
			Now.Val = l1.Val
			l1 = l1.Next
		} else {
			if l1.Val <= l2.Val {
				Now.Val = l1.Val
				l1 = l1.Next
			} else {
				Now.Val = l2.Val
				l2 = l2.Next
			}
		}
		if l1 == nil && l2 == nil {
			Now.Next = nil
		}
		Now = Now.Next
	}
	return Head
}

func main() {
	fmt.Println()
}
