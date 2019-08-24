package main

import "fmt"

func trailingZeroes(n int) int {
  rv := 5
  five := 0
  for rv <= n {
    five += n / rv
    rv *= 5
  }
  fmt.Println(five)
  return five
}

func main() {
    fmt.Println(trailingZeroes(5))
}
