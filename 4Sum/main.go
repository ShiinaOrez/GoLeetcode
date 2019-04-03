package main

import (
    "fmt"
    "sort"
)

func fourSum(nums []int, target int) [][]int {
    var res [][]int
    sort.Ints(nums)
    length := len(nums)
    if length == 4 {
        if nums[0] + nums[1] + nums[2] + nums[3] == target {
            res = append(res, nums)
            return res
        }
    }
    for m:=0; m<length; m++ {
        if m > 0 && nums[m] == nums[m-1] {
            continue
        }
        for i:=m+1; i<length; i++ {
            if i > m+1 && nums[i] == nums[i-1] {
                continue
            }
            j:=i+1;
            k:=length-1
            for ;j < k; {
                d := target - (nums[m] + nums[i] + nums[j] + nums[k])
                if d == 0 {
                    res = append(res, []int{nums[m], nums[i], nums[j], nums[k]})
                    for j < k && nums[j] == nums[j+1] {
                        j += 1
                    }
                    for j < k && nums[k] == nums[k-1] {
                        k -= 1
                    }
                    j += 1; k -= 1
                } else if d > 0 {
                    j += 1
                } else {
                    k -= 1
                }
//                fmt.Println(m, i, j, k)
            }
        }
    }
    return res
}

func main() {
    fmt.Println(fourSum([]int{5,5,3,5,1,-5,1,-2}, 4))
}
