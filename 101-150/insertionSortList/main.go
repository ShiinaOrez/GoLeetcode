package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

var newHead *ListNode

func insertionSortList(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    newHead = new(ListNode)
    newHead.Next = head
    sortBeforeNode := newHead
    sortNode := head
    sortNextNode := head.Next
    for sortNode != nil {
        fmt.Printf("Now insert: %d - %d ", sortNode.Val, sortBeforeNode.Val)
        showLink(newHead.Next)
        now := newHead
        for now != sortNode && now != nil {
            if now.Next != nil && now.Next.Val >= sortNode.Val {
                (*sortBeforeNode).Next = sortNextNode
                (*sortNode).Next = now.Next
                (*now).Next = sortNode
            }
            now = now.Next
        }
        sortNode = sortNextNode
        if sortBeforeNode.Next != sortNode {
            sortBeforeNode = sortBeforeNode.Next
        }
        if sortNode == nil {
            break
        } else {
            sortNextNode = sortNode.Next
        }
    }
    return newHead.Next
}

func showLink(head *ListNode) {
    fmt.Print(head)
    for head != nil {
        fmt.Printf("%d ", head.Val)
        head = head.Next
    }
    fmt.Println()
}

func main() {
    head := &ListNode {
        Val: 3,
        Next: &ListNode {
            Val: 2,
            Next: &ListNode {
                Val: 4,
                Next: nil,
            },
        },
    }
    showLink(insertionSortList(head))
}
