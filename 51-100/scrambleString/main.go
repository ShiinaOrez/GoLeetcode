package main

import (
    "strings"
    "fmt"
)

func isScramble(s1 string, s2 string) bool {
    if s1 == s2 {
        return true
    }
    
    n := len(s1)
    m := len(s2)
    if n != m {
        return false
    }
    
    map1 := make(map[rune]int)
    map2 := make(map[rune]int)
    for _, r := range s1 {
        if _, ok := map1[r]; !ok {
            map1[r] = strings.Count(s1, string([]byte{byte(r)}))
        }
    } 
    for _, r := range s2 {
        if _, ok := map2[r]; !ok {
            map2[r] = strings.Count(s2, string([]byte{byte(r)}))
        }
    }
    
    for k, v := range map1 {
        if map2[k] != v {
            return false
        }
    }
    
    for i:=1; i<n; i++ {
        if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) {
            return true
        }
        if isScramble(s1[:i], s2[n-i:]) && isScramble(s1[i:], s2[:n-i]) {
            return true
        }
    }
    return false
}

func main() {
    fmt.Println(isScramble("great", "rgeat"))
}
