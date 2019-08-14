package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func reorderList(head *ListNode)  {
    if head == nil || head.Next == nil {
        return
    }
    slow := head
    fast := head.Next

    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    // head -> ... -> slow -> ... -> fast

    reverse := func(head *ListNode) *ListNode {
        var p1, p2, p3 *ListNode
        p1, p2, p3 = nil, head, head
        for p2 != nil {
            p3 = p2.Next
            p2.Next = p1
            p1 = p2
            p2 = p3
        }
        return p1
    }

    reverser := slow.Next
    slow.Next = nil
    reverser = reverse(reverser)
    // head -> ... -> slow slow.Next <- ... <- fast

    cur := head
    printList(cur)
    printList(reverser)
    for cur != nil || reverser != nil {
        curNext := reverser
        reverser = reverser.Next
        nextCur := cur.Next
        curNext.Next = cur.Next
        cur.Next = curNext

        if nextCur == nil && reverser != nil {
            curNext.Next = reverser
            break
        }
        cur = nextCur
    }
    printList(cur)
    printList(reverser)
    // head -> fast -> head.Next -> fast.Next -> ... -> slow

    return
}

func printList(head *ListNode) {
    for head != nil {
        fmt.Printf("%d ", head.Val)
        head = head.Next
    }
    fmt.Println()
}

func main() {
    n5 := ListNode{5, nil}
    n4 := ListNode{4, &n5}
    n3 := ListNode{3, &n4}
    n2 := ListNode{2, &n3}
    n1 := ListNode{1, &n2}
    reorderList(&n1)
    printList(&n1)
    return
}