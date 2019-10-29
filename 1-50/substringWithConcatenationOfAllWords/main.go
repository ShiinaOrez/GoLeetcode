package main

import "fmt"

func findSubstring(s string, words []string) []int {
	wordsLength := len(words)
	if wordsLength == 0 {
		return []int{}
	}
	stepLength := len(words[0])
	sliceLength := wordsLength * stepLength
	var res []int
	flag := true
	length := len(s)
	for start := 0; length >= start+sliceLength; start++ {
		wordsMap := make(map[string]int)
		for _, str := range words {
			wordsMap[str] += 1
		}
		flag = true
		slice := s[start : start+sliceLength]

		fmt.Printf("Slice: %s ", slice)

		for i := 0; i < wordsLength; i++ {
			sliceStep := slice[i*stepLength : (i+1)*stepLength]
			fmt.Printf("[%s]:", sliceStep)
			if ok, _ := wordsMap[sliceStep]; ok > 0 {
				wordsMap[sliceStep] = ok - 1
				fmt.Printf("%d Yes\n", ok)
			} else {
				fmt.Printf("No\n")
				flag = false
				break
			}
		}
		if flag {
			res = append(res, start)
		}
	}
	return res
}

func main() {
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}))
}
