package main

import "fmt"

func addBinary(a string, b string) string {
    na := len(a)
    nb := len(b)
    l := make([]int, func (a, b int) int {
        if a>=b {
            return a
        }
        return b
    }(na, nb))
    n := len(l)
    for i:=1; i<=na; i++ {
        l[n-i] += int(a[na-i]) - 48
    }
    for i:=1; i<=nb; i++ {
        l[n-i] += int(b[nb-i]) - 48
    }
    next := 0
    for i:=1; i<=n; i++ {
        l[n-i] += next
        next = l[n-i] / 2
        l[n-i] %= 2
    }
    if next > 0 {
        l = append([]int{next}, l...)
    }
    bytes := []byte{}
    for _, i := range l {
        bytes = append(bytes, byte(i+48))
    }
    return string(bytes)
}

func main() {
    fmt.Println(addBinary("1", "11"))
}
