package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    var newHead *ListNode
    var now *ListNode
    flag := false
    before := head.Val-1
    for head != nil {
        if head.Next == nil {
            if before != head.Val {
                if !flag {
                    newHead = &ListNode{}
                    now = newHead
                    flag = true
                } else {
                    now.Next = &ListNode{}
                    now = now.Next
                }
                now.Val = head.Val
            }
        } else if head.Val != before && head.Val != head.Next.Val {
            if !flag {
                newHead = &ListNode{}
                now = newHead
                flag = true
            } else {
                now.Next = &ListNode{}
                now = now.Next
            }
            now.Val = head.Val
        }
        before = head.Val
        head = head.Next
    }
    return newHead
}

func main() {
    fmt.Println()
}
