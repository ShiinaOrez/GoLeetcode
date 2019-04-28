package main

import (
    "fmt"
)

func getPermutation(n int, k int) string {
	facs := []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}
    nums := []int{}
    for i:=1; i<=n; i++ {
        nums = append(nums, i)
    }
    res := []int{}
    k -= 1
    for j:=n; j>=1; j-- {
        r := k % facs[j-1];
        t := k / facs[j-1];
        k = r;
        fmt.Println(t, nums)
        res = append(res, nums[t])
        nums = append(nums[:t], nums[t+1:]...)
    }
 //   res = append(res, nums[0])
    bytes := []byte{}
    for _, v := range res {
        bytes = append(bytes, byte(48 + v))
    }
    
    return string(bytes)
}

func main() {
    fmt.Println(getPermutation(4, 9))
}
