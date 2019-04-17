package main

import "fmt"

func maxArea(height []int) int {
    min := func (x, y int) int {
        if x<= y {
            return x
        }
        return y
    }
    length := len(height)
    maxA := 0
    for i:=0; i<length; i++ {
        for j:=i+1; j<length; j++ {
            A := min(height[i], height[j]) * (j-i)
            if A > maxA {
                maxA = A
            }
        }
    }
    return maxA
}

func main() {
    fmt.Println(maxArea([]int{1,8,2,4,5,3,2,2}))
}
