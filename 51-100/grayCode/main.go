package main

import "fmt"

func grayCode(n int) []int {
    if n == 0 {
        return []int{}
    }
    if n == 1 {
        return []int{0, 1}
    }
    codes := grayCode(n-1)
    newCodes := []int{}
    m := len(codes)
    for i:=m-1; i>=0; i-- {
        newCodes = append(newCodes, codes[i] + 1 << uint(n-1))
    }
    return append(codes, newCodes...) 
}

func main() {
   fmt.Println(grayCode(11))
}
