package main

import (
    "fmt"
    "strings"
)

func minWindow(s string, t string) string {
    origin := make(map[byte]int)
    for _, r := range t {
        origin[byte(r)] = strings.Count(t, string([]byte{byte(r)}))
    }
    
    n := len(s)
    ans := n+1
    res := ""
    if n == 0 {
        return ""
    }
    pattern := make(map[byte]int)
    
    start := 0
    end := 0
    status := false
    check := func(a, b map[byte]int) bool {
        for k, v := range a {
            v2, ok := b[k]
            if ok {
                if v2 < v {
                    return false
                }
            } else {
                return false
            }
        }
        return true
    }
    
    for end < n {
        fmt.Println("Now:", s[start:end+1])
        v, ok := pattern[s[end]]
        if ok {
            pattern[s[end]] = v + 1
        } else {
            pattern[s[end]] = 1
        }
        fmt.Println("Map:", pattern)
        if !status {
            now := check(origin, pattern)
            if now {
                status = now
            } else {
                end += 1
                continue
            }
        }
        fmt.Println("Status is:", status)
        if status {
            for pattern[s[start]] >= origin[s[start]] {
                fmt.Printf("\t%c | %d: %d\n", s[start], pattern[s[start]], origin[s[start]])
                if pattern[s[start]] == origin[s[start]] {
                    break
                } else {
                    pattern[s[start]] = pattern[s[start]] - 1
                    start += 1
                }
            }

            if end - start + 1 < ans {
                ans = end - start + 1
                res = s[start:end+1]
            }
        }
        end += 1
    }
    return res
}

func main() {
    fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}
