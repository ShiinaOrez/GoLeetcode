package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    newList := new(ListNode)
    nowNode := newList
    up := 0
    for l1 != nil && l2 != nil {
        rest := l1.Val + l2.Val + up
        up = rest / 10
        rest = rest % 10
        nowNode.Val = rest
        l1 = l1.Next
        l2 = l2.Next
        if l1 != nil || l2 != nil {
            nowNode.Next = new(ListNode)
        }
        nowNode = nowNode.Next
    }
    var last *ListNode
    if l1 == nil && l2 != nil {
        last = l2
    } else if l2 == nil && l1 != nil {
        last = l1
    } else {
        if up != 0 {
            nowNode = newList
            for nowNode.Next != nil {
                nowNode = nowNode.Next
            }
            nowNode.Next = new(ListNode)
            nowNode.Next.Val = up
        }
        return newList
    }
    for last != nil {
        rest := last.Val + up
        up = rest / 10
        rest = rest % 10
        nowNode.Val = rest
        last = last.Next
        if last != nil {
            nowNode.Next = new(ListNode)
        } else if up != 0 {
            nowNode.Next = new(ListNode)
            nowNode.Next.Val = up
        }
        nowNode = nowNode.Next
    }
    return newList
}

func main() {
    test1 := new(ListNode)
    test2 := new(ListNode)
    test1.Val = 1
    test2.Val = 9
    test2.Next = new(ListNode)
    test2.Next.Val = 9
    ans := addTwoNumbers(test1, test2)
    for ans != nil {
        fmt.Printf("%d ", ans.Val)
        ans = ans.Next
    }
}
