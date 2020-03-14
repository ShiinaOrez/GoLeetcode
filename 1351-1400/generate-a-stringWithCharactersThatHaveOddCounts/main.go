package main

import (
    "fmt"
)

func generateTheString(n int) string {
    if n == 0 {
        return ""
    }
    if n == 1 {
        return "a"
    }
    stringA := "a"
    stringB := "b"
    s := ""
    if n % 2 == 0 && (n/2) % 2 == 1 {
        for i:=0; i<n/2; i++ {
            s += (stringA + stringB)
        }
    } else if n % 2 == 0 && (n/2) % 2 == 0 {
        for i:=0; i<n/2; i++ {
            s += (stringA + stringB)
        }
        s += stringB
        s = s[1:]
    } else {
        for i:=0; i<n; i++ {
            s += stringA
        }
    }
    return s
}

func main() {
    fmt.Println(generateTheString(4))
}
