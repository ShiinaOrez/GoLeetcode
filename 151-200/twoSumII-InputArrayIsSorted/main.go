package main

import (
    "fmt"
    "sort"
)

func twoSum(numbers []int, target int) []int {
    if len(numbers) == 2 {
      return []int{1, 2}
    }
    leftIndex, rightIndex := sort.SearchInts(numbers, target/2), 0
    if leftIndex + 1 < len(numbers)-1 {
        rightIndex = leftIndex + 1
    } else {
      leftIndex, rightIndex = leftIndex-1, leftIndex
    }
    for numbers[leftIndex] + numbers[rightIndex] != target {
        if numbers[leftIndex] + numbers[rightIndex] < target {
            if rightIndex == len(numbers)-1 {
                leftIndex += 1
            } else {
                rightIndex += 1
            }
        } else {
            if leftIndex == 0 {
                rightIndex -= 1
            } else {
                leftIndex -= 1
            }
        }
        if leftIndex == rightIndex {
            rightIndex = leftIndex + 1
        }
    }
    return []int{leftIndex+1, rightIndex+1}
}

func main() {
    fmt.Println()
}
