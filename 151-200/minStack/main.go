package main

import (
    "fmt"
    "sort"
)

type MinStack struct {
    Elements     []int
    MinEleCache  int
    Size         int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack {
        Elements: []int{},
        MinEleCache: 0,
        Size: 0,
    }
}


func (this *MinStack) Push(x int)  {
    if len(this.Elements) <= this.Size {
        this.Elements = append(this.Elements, x)
    } else {
        this.Elements[this.Size] = x
    }
    if x < this.MinEleCache || this.Size == 0 {
        this.MinEleCache = x
    }
    this.Size += 1
    return
}


func (this *MinStack) Pop()  {
    if this.Size >= 1 {
        if this.Elements[this.Size-1] == this.MinEleCache {
            if this.Size == 1 {
                this.MinEleCache = 0
            } else {
                cache := make([]int, this.Size-1)
                copy(cache, this.Elements)
                sort.Ints(cache)
                this.MinEleCache = cache[0]
            }
        }
        this.Size -= 1
    }
    return
}


func (this *MinStack) Top() int {
    if this.Size >= 1 {
        return this.Elements[this.Size-1]
    }
    return -1
}


func (this *MinStack) GetMin() int {
    if this.Size >= 1 {
        return this.MinEleCache
    }
    return -1
}

func main() {
    obj := Constructor()
    obj.Push(-2)
    obj.Push(0)
    obj.Push(-3)
    fmt.Println(obj.GetMin())
    obj.Pop()
    fmt.Println(obj.Top())
    fmt.Println(obj.GetMin())
}


