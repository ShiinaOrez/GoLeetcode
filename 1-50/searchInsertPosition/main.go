package main

import "fmt"

func searchInsert(nums []int, target int) int {
    if len(nums) == 0 {
        return 0
    }
    return find(nums, 0, len(nums)-1, target)
}

func find(nums []int, left, right, target int) int {
    if left == right {
        if nums[left] < target {
            return left + 1
        } else {
            return left
        }
    }
    
    if left == right - 1 {
        if nums[left] == target {
            return left
        } else if nums[right] == target {
            return right
        } else if nums[right] < target {
            return right + 1
        } else if nums[left] > target {
            return left
        } else {
            return left + 1
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
    fmt.Println(searchInsert([]int{0, 1, 3, 5}, 4))
}
