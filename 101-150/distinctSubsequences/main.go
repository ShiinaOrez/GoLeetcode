package main

import "fmt"

func numDistinct(s string, t string) int {
    tLength := len(t)
    sLength := len(s)
    dp := make([]int, tLength+1)
    dp[0] = 1
    dict := make([]int, 129)
    for i:=0; i<=128; i++ {
        dict[i] = -1
    }
    nexts := make([]int, tLength)
    for i:=0; i<tLength; i++ {
        c := int(t[i])
        nexts[i] = dict[c]
        dict[c] = i
    }

    for i:=0; i<sLength; i++ {
        c := int(s[i])
        for j:=dict[c]; j>=0; j=nexts[j] {
            dp[j+1] += dp[j]
        }
    }
    return dp[tLength]
}

func main() {
    fmt.Println(numDistinct("rabbbit", "rabbit"))
}
