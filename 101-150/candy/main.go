package main

import "fmt"

func candy(ratings []int) int {
	min, minIndex, count := 1, 0, 0
	length := len(ratings)
	moutain := make([]int, length)
	moutain[0] = 1
	for i := 1; i < length; i++ {
		if ratings[i] > ratings[i-1] {
			moutain[i] = moutain[i-1] + 1
		} else if ratings[i] == ratings[i-1] {
			if moutain[i-1] > 0 {
				moutain[i] = 1
			} else {
				moutain[i] = moutain[i-1]
			}
		} else {
			if moutain[i-1] > 1 {
				moutain[i] = 1
			} else {
				moutain[i] = moutain[i-1] - 1
			}
		}
		if moutain[i] < min {
			min, minIndex = moutain[i], i
		}
	}

	moutain[minIndex] = 1
	for i := minIndex - 1; i >= 0; i-- {
		fmt.Println(moutain)
		if ratings[i] > ratings[i+1] {
			moutain[i] = moutain[i+1] + 1
		} else {
			if ratings[i] < ratings[i+1] {
				if moutain[i+1] == 1 {
					moutain[i] = moutain[i+1] - 1
				} else {
					moutain[i] = 1
				}
			} else {
				moutain[i] = 1
			}
		}
		if moutain[i] <= 0 {
			moutain[i] = 1
			for j := i + 1; j <= minIndex && ratings[j] > ratings[j-1]; j++ {
				if moutain[j] > moutain[j-1] {
					break
				}
				moutain[j] = moutain[j-1] + 1
			}
		}
	}
	for i := minIndex + 1; i < length; i++ {
		if ratings[i] > ratings[i-1] {
			moutain[i] = moutain[i-1] + 1
		} else {
			if ratings[i] < ratings[i-1] {
				if moutain[i-1] == 1 {
					moutain[i] = moutain[i-1] - 1
				} else {
					moutain[i] = 1
				}
			} else {
				moutain[i] = 1
			}
		}
		if moutain[i] <= 0 {
			moutain[i] = 1
			for j := i - 1; j >= minIndex && ratings[j] > ratings[j+1]; j-- {
				if moutain[j] > moutain[j+1] {
					break
				}
				moutain[j] = moutain[j+1] + 1
			}
		}
	}
	fmt.Println(moutain)
	fmt.Println(minIndex)
	for _, m := range moutain {
		count += m
	}
	return count
}
func main() {
	fmt.Println(candy([]int{
		1, 6, 10, 8, 7, 3, 2,
	}))
}
