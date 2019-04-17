package main

import "fmt"

func trap(height []int) int {
    if len(height) == 0 {
        return 0
    }
    length := len(height)
    maxLeft := 0
    wholeArea := 0
    count := 0
    for ind, i := range height {
        count += i
        if i > maxLeft {
            h := maxLeft
            maxLeft = i
            for j:=length-1; j>=0; j-- {
                sub := i - h
                if height[j] >= i {
                    if j >= ind {
                        wholeArea += sub * (j-ind+1)
                        fmt.Println("New:", sub * (j-ind+1))
                    }
                    break
                }
                if height[j] > h {
                    if j >= ind {
                        wholeArea += (height[j]-h) * (j-ind+1)
                        fmt.Println("New:", (height[j]-h) * (j-ind+1))
                        h = height[j]
                    }
                }
            }
        }
    }
    return wholeArea - count
}

func main() {
    fmt.Println(trap([]int{4, 2, 3}))
}