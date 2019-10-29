// 这个题可以用倍增进行优化
package main

import "fmt"

func searchRange(nums []int, target int) []int {
	count := len(nums)
	if count == 0 {
		return []int{-1, -1}
	}
	originIndex := find(nums, 0, count-1, target)
	if originIndex == -1 {
		return []int{-1, -1}
	}

	left := originIndex
	right := originIndex
	if left > 0 {
		for nums[left-1] == nums[left] {
			left -= 1
			if left == 0 {
				break
			}
		}
	}
	if right < count-1 {
		for nums[right+1] == nums[right] {
			right += 1
			if right == count-1 {
				break
			}
		}
	}
	return []int{left, right}
}

func find(nums []int, left, right, target int) int {
	if left == right {
		if nums[left] != target {
			return -1
		} else {
			return left
		}
	}

	if left == right-1 {
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
	fmt.Println(searchRange([]int{2, 2}, 2))
}
