package main

import "fmt"

func findOrder(numCourses int, prerequisites [][]int) []int {
	inMap := make(map[int]int)
	// outMap := make(map[int]int)
	toMap := make(map[int][]int)

	for _, pair := range prerequisites {
		// outMap[pair[1]] += 1
		inMap[pair[0]] += 1
		toMap[pair[1]] = append(toMap[pair[1]], pair[0])
	}
	ch := make(chan int, numCourses)
	ans := []int{}
	for i := 0; i < numCourses; i++ {
		if _, ok := inMap[i]; !ok {
			ch <- i
		}
	}
	if len(ch) == 0 {
		return []int{}
	}

	for item := range ch {
		ans = append(ans, item)
		if len(ans) == numCourses {
			break
		}
		toCourse, ok := toMap[item]
		if !ok {
			continue
		} else {
			for _, to := range toCourse {
				inMap[to] -= 1
				if inMap[to] == 0 {
					ch <- to
				}
			}
		}
		if len(ch) == 0 {
			return []int{}
		}
	}
	close(ch)
	return ans
}

func main() {
	fmt.Println(findOrder(4, [][]int{
		[]int{1, 0},
		[]int{0, 1},
		[]int{3, 1},
		[]int{3, 2},
	}))
}
