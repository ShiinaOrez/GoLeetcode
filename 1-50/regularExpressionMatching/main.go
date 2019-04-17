package main

import "fmt"

/*func isMatch(s string, p string) bool {
    if len(s) == 1 && p == "." {
        return true
    }
    if p == "" {
        if s == "" {
            return true
        }
        return false
    }
    if s == p {
        return true
    }
    if p == ".*" {
        return true
    }
    if len(p) == 2 && p[1] == byte('*') {
        if s == "" {
            return true
        }
        ch := p[0]
        for _, c := range(s) {
            if byte(c) != ch {
                return false
            }
        }
        return true
    } else {
        var i int
        flag := false
        for i=0; i<len(s) && i<len(p); i++ {
            if i < len(p) -1 {
                if p[i+1] == byte('*') {
                    break
                }
            }
            if s[i] != p[i] {
                break
            }
        }
        if i>0 {
            s = s[i:]
            p = p[i:]
            flag = true
        }
        for i=0; len(s)-1-i>=0 && len(p)-1-i>=0; i++ {
            if s[len(s)-1-i] != p[len(p)-1-i] {
                break
            }
        }
        if i>0 {
            s = s[:len(s)-i]
            p = p[:len(p)-i]
            flag = true
        }
        if len(s) == 1 && p == "." {
            return true
        }
        if flag {
            return isMatch(s, p)
        }
        fmt.Println("Now:", s, p)
        var res1, res2 bool
//        if len(s) == 1 {
//            return isMatch(s, p)
//        } else {
            for k:=0; k<=len(s); k++ {
                for q:=0; q<len(p); q++ {
                    if k == len(s) {
                        if s == "" {
                            continue
                        }
                        res1 = isMatch(s, p[:q])
                        res2 = isMatch("", p[q:])
                        if res1 && res2 {
                            return true
                        }
                    }
                    if (k == 0 && q == 0) || p[q:][0] == byte('*') {
                        continue
                    }
                    fmt.Println(k, q)
                    res1 = isMatch(s[:k], p[:q])
                    fmt.Printf("%s - %s:", s[:k], p[:q], res1)
                    res2 = isMatch(s[k:], p[q:])
                    fmt.Println("", s[k:], "-", p[q:], ":", res2)
                    if res1 && res2 {
                        return true
                    }
                }
//            }
        }
    }
    return false
}*/

func isMatch(s string, p string) bool {
    if p == "" {
        return s == ""
    }
    firstMatch := !(s == "") && (s[0] == p[0] || p[0] == '.')
    if len(p) > 1 && p[1] == byte('*') {
        return isMatch(s, p[2:]) || (firstMatch && isMatch(s[1:], p))
    } else {
        return firstMatch && isMatch(s[1:], p[1:])
    }
}

func main() {
//    ss := "abcde"
//    for i:=0; i<len(ss); i++ {
//        fmt.Println(i, ss[:i], ss[i:])
//    }
    fmt.Println(isMatch("aaa", "ab*a*c*a"))
}