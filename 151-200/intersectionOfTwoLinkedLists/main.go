package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    m := make(map[int]*ListNode)
    if headA == nil && headB == nil {
        return nil
    }
    for headA != nil {
        m[headA.Val] = headA
        headA = headA.Next
    }
    for headB != nil {
        if node, ok := m[headB.Val]; ok {
            if headB == node {
                return headB
            } else {
                headB = headB.Next
            }
        } else {
            headB = headB.Next
        }
    }
    return nil
}

func main() {
    fmt.Println()
}
