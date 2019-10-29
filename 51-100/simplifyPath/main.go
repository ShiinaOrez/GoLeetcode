package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	path = strings.Replace(path, "/./", "/", -1)
	path = strings.TrimRight(path, "/")
	dirs := strings.Split(path, "/")
	n := len(dirs)
	for i := 1; i <= n; i++ {
		if dirs[n-i] == "." {
			dirs[n-i] = ""
		}
		if dirs[n-i] == ".." {
			dirs[n-i] = ""
			for j := i; j <= n; j++ {
				if dirs[n-j] == "." {
					dirs[n-j] = ""
				}
				if dirs[n-j] != "" && dirs[n-j] != ".." {
					dirs[n-j] = ""
					break
				}
			}
		}
	}
	newPath := ""
	for _, dir := range dirs {
		if dir != "" {
			newPath = newPath + "/" + dir
		}
	}
	if newPath == "" {
		return "/"
	}
	return newPath
}

func main() {
	fmt.Println(simplifyPath("/a//b////c/d//././/.."))
}
