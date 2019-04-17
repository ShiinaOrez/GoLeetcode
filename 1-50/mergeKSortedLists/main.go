package main

import (
    "fmt"
    "sort"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
    res := []int{}
    for _, list := range(lists) {
        for list != nil {
            res = append(res, list.Val)
            list = list.Next
        }
    }
    sort.Ints(res)
    var list *ListNode
    if len(res) > 0 {
        list = &ListNode{}
        now := list
        for ind, i := range(res) {
            now.Val = i
            if ind == len(res) - 1 {
                break
            }
            now.Next = &ListNode{}
            now = now.Next
        }
        return list
    } else {
        return nil
    }
}

func main() {
    fmt.Println()
}
