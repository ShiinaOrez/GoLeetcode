package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {
        return nil
    }
    now := &ListNode{}
    now = head
    n := 1
    for now.Next != nil {
        now = now.Next
        n += 1
    }
    now.Next = head
    k %= n
    for i:=0; i<n-k; i++ {
        now = now.Next
    }
    newHead := now.Next
    now.Next = nil
    return newHead
}

func main() {
    fmt.Println()
}
