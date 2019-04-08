package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
    if head == nil {
        return nil
    }
    if k == 1 {
        return head
    }
    now := head
    newHead := &ListNode{}
    newHead.Next = head
    stop := newHead
    for true {
        fmt.Println("Stop:", stop.Val)
        now = stop.Next
        for t:=k; t>0; t-- {
            if now == nil {
                if newHead == nil {
                    return head
                } else {
                    return newHead.Next
                }
            }
            now = now.Next
        }
        stop = reverseList(stop, k)
        newHead.Print()
    }
    return newHead.Next
}

func reverseList(head *ListNode, k int) *ListNode {
    realHead := head.Next
    if k == 2 {
        head.Next = realHead.Next
        realHead.Next = head.Next.Next
        head.Next.Next = realHead
        return realHead
    }
    head.Next = realHead.Next
    tail := reverseList(realHead, k-1)
    head.Next = realHead.Next
    realHead.Next = tail.Next
    tail.Next = realHead
    return realHead
}

func (head *ListNode)Print() {
    for head != nil {
        fmt.Printf(" %d ", (*head).Val)
        head = head.Next
    }
    fmt.Println()
    return
}

func main() {
    list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, &ListNode{8, &ListNode{9, nil}}}}}}}}}
    reverseKGroup(list, 3).Print()
}