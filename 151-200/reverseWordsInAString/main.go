package main

import (
    "fmt"
    "strings"
)

func reverseWords(s string) string {
    if s == "" {
        return ""
    }
    wordSlice := strings.Split(s, " ")
    for i:=len(wordSlice)-1; i>=0; i-- {
        if wordSlice[i] == "" {
            wordSlice = append(wordSlice[:i], wordSlice[i+1:]...)
        }
    }
    length := len(wordSlice)
    for i:=0; i<length/2; i++ {
        wordSlice[i], wordSlice[length-i-1] = wordSlice[length-i-1], wordSlice[i]
    }
    return strings.Join(wordSlice, " ")
}

func main() {
    fmt.Println(reverseWords("the sky is blue"))
}
