package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
    has := make(map[ListNode]bool)
    if head == nil || head.Next == nil {
        return nil
    }
    if head.Next == head {
        return head
    }
    now := head
    for {
        if _, ok := has[*now]; ok {
            break
        }
        has[*now] = true
        now = now.Next
        if now == nil {
            return nil
        }
    }
    return now
}

func main() {
    fmt.Println()
}
