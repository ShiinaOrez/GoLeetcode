package main

import "fmt"

var res []string

func generateParenthesis(n int) []string {
    l := n
    r := n
    res = []string{}
    generate("", l, r)
    return res
}

func generate(mid string, l, r int) {
    if l < 0 || r < 0 || r < l {
        return 
    }
    if l == 0 && r == 0 {
        res = append(res, mid)
    }
    generate(mid+"(", l-1, r)
    generate(mid+")", l, r-1)
}

func main() {
    fmt.Println(generateParenthesis(4))
}
