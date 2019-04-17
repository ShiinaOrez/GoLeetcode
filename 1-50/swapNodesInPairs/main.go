package main

import (
    "fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
    now := head
    if now == nil || now.Next == nil {
        return head
    } else {
        next := now.Next
        now.Next = next.Next
        next.Next = now
        head = next
    }
    for now != nil {
        if now.Next != nil && now.Next.Next == nil {
            break
        }
        next := now.Next
        if next != nil {
            now.Next = next.Next
            if next.Next != nil {
                next.Next = next.Next.Next
                now.Next.Next = next
            }
        } else {
            break
        }
        now = next
    }
    return head
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
    list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
    list.Print()
    fmt.Println(list)
    newList := swapPairs(list)
    newList.Print()
    fmt.Println(newList)
}
