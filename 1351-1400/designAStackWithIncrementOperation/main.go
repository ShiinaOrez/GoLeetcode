package main

import (
    "fmt"
)

type CustomStack struct {
    MaxSize int
    Values  []int
}


func Constructor(maxSize int) CustomStack {
    cs := CustomStack{}
    cs.MaxSize = maxSize
    return cs
}


func (this *CustomStack) Push(x int)  {
    if len(this.Values) >= this.MaxSize {
        return
    }
    this.Values = append(this.Values, x)
    return
}


func (this *CustomStack) Pop() int {
    size := len(this.Values)
    if size == 0 {
        return -1
    }
    v := this.Values[size-1]
    this.Values = this.Values[:size-1]
    return v
}


func (this *CustomStack) Increment(k int, val int)  {
    size := len(this.Values)
    if k < size {
        size = k
    }
    for i:=0; i<size; i++ {
        this.Values[i] += val
    }
    return
}

func main() {
    cs := Constructor(3)
    cs.Push(1)
    cs.Push(2)
    fmt.Println(cs.Pop())
}
