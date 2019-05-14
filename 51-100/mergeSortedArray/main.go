package main

import (
    "sort"
    "fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int)  {
    copy(nums1[m:], nums2[:n])
    sort.Ints(nums1[:m+n])
}

func main() {
    a := []int{3, 4, 5, 0, 0, 0}
    b := []int{1, 2, 6}
    merge(a, 3, b, 3)
    fmt.Println(a)
}
