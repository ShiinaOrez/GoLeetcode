package main

import "fmt"

func convertToTitle(n int) string {
  res := ""
  for n > 0 {
    n -= 1
    res = string([]byte{byte(65+n%26)}) + res
    n /= 26
  }
  return res
}

func main() {
    fmt.Println(convertToTitle(701))
}
