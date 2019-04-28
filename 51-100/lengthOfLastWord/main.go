package main

import (
    "fmt"
    "strings"
)

func lengthOfLastWord(s string) int {
    s = strings.TrimSpace(s)
    slice := strings.Split(s, " ")
    fmt.Println(slice)
    if len(slice) == 0 {
        return 0
    }
    return len(slice[len(slice)-1])
}

func main() {
    fmt.Println(lengthOfLastWord("a "))
}
