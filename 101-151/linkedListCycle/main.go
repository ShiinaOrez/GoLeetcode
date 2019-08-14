package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }
    if head.Next == head {
        return true
    }
    slow, fast := head.Next, head.Next.Next
    for slow != fast {
        if slow != nil {
            slow = slow.Next
        }
        if fast != nil && fast.Next != nil {
            fast = fast.Next.Next
        }
    }
    fmt.Println(slow, fast)
    if slow == nil && fast == nil || slow.Next == nil && fast.Next == nil {
        return false
    }
    return true
}

func main() {
    fmt.Println("")
}
