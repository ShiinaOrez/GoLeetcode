package main

import (
    "fmt"
    "strconv"
)

func evalRPN(tokens []string) int {
    if len(tokens) == 0 {
        return 0
    }
    stack := Stack {
        Size:     0,
        Elements: []int{},
    }
    for _, token := range tokens {
        if token == "+" || token == "-" || token == "*" || token == "/" {
            b := stack.Pop()
            a := stack.Pop()
            switch token {
            case "+" :
                stack.Push(a+b)
            case "-" :
                stack.Push(a-b)
            case "*" :
                stack.Push(a*b)
            case "/" :
                stack.Push(a/b)
            }
        } else {
            num, _ := strconv.Atoi(token)
            stack.Push(num)
        }
    }
    return stack.Pop()
}

type Stack struct {
    Size     int
    Elements []int
}

func (stack *Stack) Push(ele int) {
    if len(stack.Elements) == stack.Size {
        stack.Elements = append(stack.Elements, ele)
    } else {
        stack.Elements[stack.Size] = ele
    }
    stack.Size += 1
    return
} 

func (stack *Stack) Pop() int {
    stack.Size -= 1
    return stack.Elements[stack.Size]
}

func main() {
    fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
}
