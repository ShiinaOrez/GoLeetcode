package main

import "fmt"

func countAndSay(n int) string {
	originString := "1"
	var res string
	for i := 1; i < n; i++ {
		fmt.Println(originString)
		res = ""
		for j := 0; j < len(originString); j++ {
			count := 0
			now := originString[j]
			k := j
			for originString[j] == originString[k] {
				count += 1
				k += 1
				if k == len(originString) {
					break
				}
			}
			j = k - 1
			for count > 0 {
				res = res + string(byte(count%10+'0'))
				count /= 10
			}
			res = res + string(now)
		}
		originString = res
	}
	return originString
}

func main() {
	fmt.Println(countAndSay(4))
}
