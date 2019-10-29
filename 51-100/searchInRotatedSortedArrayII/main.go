package main

import "fmt"

func search(nums []int, target int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}

	endpoint := 0
	for i := 0; i < n-1; i++ {
		if nums[i] > nums[i+1] {
			endpoint = i + 1
			break
		}
	}
	array := make([]int, n)
	copy(array, nums[endpoint:])
	copy(array[n-endpoint:], nums[:endpoint])
	return find(array, 0, n-1, target)
}

func find(nums []int, left, right, target int) bool {
	if left == right {
		if nums[left] != target {
			return false
		} else {
			return true
		}
	}

	if left == right-1 {
		if nums[left] == target {
			return true
		} else if nums[right] == target {
			return true
		} else {
			return false
		}
	}

	mid := (left + right) / 2
	if nums[mid] == target {
		return true
	}
	if target <= nums[mid] {
		return find(nums, left, mid, target)
	}
	return find(nums, mid+1, right, target)
}

func main() {
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 0))
}
