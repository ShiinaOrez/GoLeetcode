package main

import "fmt"

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    res := ""
    for i:=0; ; i++ {
        if i == len(strs[0]) {
            return res
        }
        c := strs[0][i]
        for j, str := range(strs) {
            if i >= len(str) {
                return res
            }
            if j == 0 {
                continue
            } else {
                if str[i] != c {
                    return res
                }
            }
        }
        res = res + string(c)
    }
    return res
}

func main() {
    fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}
