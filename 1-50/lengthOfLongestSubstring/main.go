package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
    set := make(map[byte]int)
    maxLength := 0
    length := 0
    for i, c := range(s) {
        index, ok := set[byte(c)]
        if !ok {
            length += 1
            if length > maxLength {
                maxLength = length
            }
            set[byte(c)] = i
        } else {
            for j:=i-length; j<index; j++ {
                delete(set, s[j])
            }
            length = i - index
            set[byte(c)] = i
        }
    }
    return maxLength
}

func main() {
    fmt.Println(lengthOfLongestSubstring("gaaqfeqlqky"))
}
