package main

import (
    "fmt"
    "sort"
)

func threeSum(nums []int) [][]int {
    var res [][]int
    sort.Ints(nums)
    length := len(nums)
    for i:=0; i<length; i++ {
        if nums[i] > 0 {
            break
        }
        if i>0 && nums[i] == nums[i-1] {
            continue
        }
        j:=i+1;
        k:=length-1
        for ;j < k; {
            if nums[j] + nums[k] == -1*nums[i] {
                res = append(res, []int{nums[i], nums[j], nums[k]})
                for j < k && nums[j] == nums[j+1] {
                    j += 1
                }
                for j < k && nums[k] == nums[k-1] {
                    k -= 1
                }
                j += 1; k -= 1
            } else if nums[j] + nums[k] < -1 * nums[i] {
                j += 1
            } else {
                k -= 1
            }
        }
    }
    return res
}

func main() {
    fmt.Println(threeSum([]int{-2, 0, 1, 1, 2}))
}
