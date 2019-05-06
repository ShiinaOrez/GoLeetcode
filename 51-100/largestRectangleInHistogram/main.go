package main

import "fmt"

type Stack struct {
    Data    []int
    Size    int
}

func (stack *Stack)Init() {
    stack.Data = []int{}
    stack.Size = -1
    return
}

func (stack *Stack)IsEmpty() bool {
    return stack.Size < 0
}

func (stack *Stack)Push(x int) {
    stack.Size += 1
    stack.Data = append(stack.Data, x)
    return
}

func (stack *Stack)Top() int {
    return stack.Data[stack.Size]
}

func (stack *Stack)Pop() {
    stack.Data = stack.Data[:stack.Size]
    stack.Size -= 1
    return
}

func largestRectangleArea(heights []int) int {
    var stk Stack
    stk.Init()
    heights = append(heights, 0)
    max := 0
    n := len(heights)
    for i:=0; i<n; i++ {
        for !stk.IsEmpty() && heights[i]<heights[stk.Top()] {
            top := stk.Top()
            stk.Pop()
            els := 0
            if stk.IsEmpty() {
                els = heights[top] * i
            } else {
                els = heights[top] * (i-stk.Top()-1)
            }
            max = func(a, b int) int {
                if a >= b {
                    return a
                }
                return b
            }(max, els)
        }
        stk.Push(i)
    }
    return max
}

func main() {
    fmt.Println(largestRectangleArea([]int{2,1,5,6,2,3}))
}
