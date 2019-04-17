package main

import "fmt"

func search(nums []int, target int) int {
    count := len(nums)
    if count == 0 {
        return -1
    }
    if count == 1 {
        if nums[0] == target {
            return 0
        } else {
            return -1
        }
    }
    var i int
    for i=0; i<count-1; i++ {
        if nums[i] > nums[i+1] {
            break
        }
    }
    if i == count - 1 {
        return find(nums, 0, count-1, target)
    }
    
    if target >= nums[0] {
        return find(nums, 0, i, target)
    }
    return find(nums, i+1, count-1, target)
}

func find(nums []int, left, right, target int) int {
    fmt.Printf("Find %d in [%d, %d]\n", target, left, right)
    if left == right {
        if nums[left] != target {
            return -1
        } else {
            return left
        }
    }
    
    if left == right - 1 {
        if nums[left] == target {
            return left
        } else if nums[right] == target {
            return right
        } else {
            return -1
        }
    }
    
    mid := (left + right) / 2
    if nums[mid] == target {
        return mid
    }
    if target <= nums[mid] {
        return find(nums, left, mid, target)
    }
    return find(nums, mid+1, right, target)
}

func main() {
    fmt.Println(search([]int{1, 3}, 3))
}
