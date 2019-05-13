package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
    if head == nil {
        return nil
    }
    
    leftHead := new(ListNode)
    leftNow := leftHead
    leftBefore := leftNow
    rightHead := new(ListNode)
    rightNow := rightHead
    rightBefore := rightNow
    
    left := 0
    right := 0
    
    for head != nil {
        if head.Val < x {
            leftNow.Next = new(ListNode)
            leftNow.Val = head.Val
            leftBefore = leftNow
            leftNow = leftNow.Next
            left += 1
        } else {
            rightNow.Next = new(ListNode)
            rightNow.Val = head.Val
            rightBefore = rightNow
            rightNow = rightNow.Next
            right += 1
        }
        head = head.Next
    }
    rightBefore.Next = nil
    if left == 0 {
        return rightHead
    } else if right != 0 {
        leftBefore.Next = rightHead
    } else {
        leftBefore.Next = nil
    }
    return leftHead
}

func main() {
    fmt.Println()
}
