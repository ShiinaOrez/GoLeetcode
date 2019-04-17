package main

import "fmt"

func intToRoman(num int) string {
    var res []byte
    first := num % 10
    second := (num / 10) % 10
    third := num - first - 10 * second
    for third >= 1000 {
        res = append(res, byte('M'))
        third -= 1000
    }
    div := third / 100
    if div >= 1 && div < 4 {
        for div > 0 {
            res = append(res, byte('C'))
            div -= 1
        }
    } else if div == 4 || div == 5 {
        if div == 4 {
            res = append(res, byte('C'))
        }
        res = append(res, byte('D'))
    } else if div > 5 && div < 9 {
        res = append(res, byte('D'))
        for div > 5 {
            res = append(res, byte('C'))
            div -= 1
        }
    } else if div == 9 {
        res = append(res, byte('C'))
        res = append(res, byte('M'))
    }
    div = second
    if div >= 1 && div < 4 {
        for div > 0 {
            res = append(res, byte('X'))
            div -= 1
        }
    } else if div == 4 || div == 5 {
        if div == 4 {
            res = append(res, byte('X'))
        }
        res = append(res, byte('L'))
    } else if div > 5 && div < 9 {
        res = append(res, byte('L'))
        for div > 5 {
            res = append(res, byte('X'))
            div -= 1
        }
    } else if div == 9 {
        res = append(res, byte('X'))
        res = append(res, byte('C'))
    }
    div = first
    if div >= 1 && div < 4 {
        for div > 0 {
            res = append(res, byte('I'))
            div -= 1
        }
    } else if div == 4 || div == 5 {
        if div == 4 {
            res = append(res, byte('I'))
        }
        res = append(res, byte('V'))
    } else if div > 5 && div < 9 {
        res = append(res, byte('V'))
        for div > 5 {
            res = append(res, byte('I'))
            div -= 1
        }
    } else if div == 9 {
        res = append(res, byte('I'))
        res = append(res, byte('X'))
    }
    return string(res)
}

func main() {
    fmt.Println(intToRoman(58))
}
