package main

import (
    "math"
    "fmt"
)

func myPow(x float64, n int) float64 {
    return math.Pow(x, float64(n))
}

func main() {
    fmt.Println(myPow(2.00000, 10))
}