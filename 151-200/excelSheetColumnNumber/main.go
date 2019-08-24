package main

import "fmt"

func titleToNumber(s string) int {
  length := len(s)
  if length == 0 {
    return 0
  }
  res := 0
  rv := 1
  for i:=length-1; i>=0; i-- {
    res += rv * (int(s[i]) - 64)
    rv *= 26
  }
  return res
}

func main() {
    fmt.Println(titleToNumber("ABBA"))
}
