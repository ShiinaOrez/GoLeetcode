package main

import "fmt"

func multiply(num1 string, num2 string) string {
    n1 := []int{}
    n2 := []int{}
    res := make([]int, 221)
    for _, i := range(num1) {
        n1 = append(n1, int(byte(i)-'0'))
    }
    for _, i := range(num2) {
        n2 = append(n2, int(byte(i)-'0'))
    }
    length1 := len(n1)
    length2 := len(n2)
    k := 0
    dig := 0
    for j:=length2-1; j>=0; j-- {
        t := k
        for i:=length1-1; i>=0; i-- {
            res[t] += n1[i] * n2[j]
            t += 1
        }
        k += 1
        dig = t
    }
    
    for i:=0; i<dig; i++ {
        for res[i] >= 10 {
            res[i] -= 10
            res[i+1] += 1
        }
    }
    
    resp := []byte{}
    for res[dig] == 0 {
        dig -= 1
        if dig == -1 {
            return "0"
        }
    }
    for i:=dig; i>=0; i-- {
        resp = append(resp, byte(res[i]+int('0')))
    }
    return string(resp)
}

func main() {
    fmt.Println(multiply("2", "3"))
}
