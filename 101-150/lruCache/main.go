package main

import (
    "fmt"
)

var timeNow int

type LRUCache struct {
    H        *Heap
    Cap      int
    ValMap   map[int]int
    NodeMap  map[int]int
}

func Constructor(capacity int) LRUCache {
    timeNow = 0
    lru := LRUCache {
        H:       new(Heap),
        Cap:     capacity,
        ValMap:  make(map[int]int),
        NodeMap: make(map[int]int),
    }
    lru.H.NodeList = []HeapNode{HeapNode{}}
    lru.H.Size = 0
    return lru
}


func (this *LRUCache) Get(key int) int {
    timeNow += 1
    ans := -1
    if v, ok := this.ValMap[key]; !ok {
        return ans
    } else {
        ans = v
        nodeIndex, _ := this.NodeMap[key]
        this.H.NodeList[nodeIndex].Time = timeNow
        this.H.Down(this, nodeIndex)
    }
    return ans
}


func (this *LRUCache) Put(key int, value int)  {
    timeNow += 1
    if _, ok := this.ValMap[key]; !ok {
        this.ValMap[key] = value
        this.H.Push(this, key, timeNow)
        if this.H.Size > this.Cap {
            k := this.H.Pop(this)
            if k != -1 {
                delete(this.ValMap, k)
                delete(this.NodeMap, k)
            }
        }
    } else {
        nodeIndex, _ := this.NodeMap[key]
        this.H.NodeList[nodeIndex].Time = timeNow
        this.H.Down(this, nodeIndex)
        this.ValMap[key] = value
    }
    return
}

type Heap struct {
    NodeList  []HeapNode
    Size      int
}

type HeapNode struct {
    Key       int
    Time      int // time
}

func (heap *Heap) Down(lru *LRUCache, nodeIndex int) {
    if nodeIndex > heap.Size {
        return
    } else {
        leftTime, rightTime := 0, 0
        if nodeIndex * 2 <= heap.Size {
            leftTime = heap.NodeList[nodeIndex*2].Time
        } else {
            leftTime = 9999999
        }
        if nodeIndex * 2 + 1 <= heap.Size {
            rightTime = heap.NodeList[nodeIndex*2+1].Time
        } else {
            rightTime = 9999999
        }
        if heap.NodeList[nodeIndex].Time > leftTime && leftTime <= rightTime { // swap now and left
            lru.NodeMap[heap.NodeList[nodeIndex].Key] = nodeIndex * 2
            lru.NodeMap[heap.NodeList[nodeIndex*2].Key] = nodeIndex
            heap.NodeList[nodeIndex], heap.NodeList[nodeIndex*2] = heap.NodeList[nodeIndex*2], heap.NodeList[nodeIndex]
            heap.Down(lru, nodeIndex*2)
        } else if heap.NodeList[nodeIndex].Time > rightTime && rightTime <= leftTime { // swap now and right
            lru.NodeMap[heap.NodeList[nodeIndex].Key] = nodeIndex * 2 + 1
            lru.NodeMap[heap.NodeList[nodeIndex*2+1].Key] = nodeIndex
            heap.NodeList[nodeIndex], heap.NodeList[nodeIndex*2+1] = heap.NodeList[nodeIndex*2+1], heap.NodeList[nodeIndex]
            heap.Down(lru, nodeIndex*2+1)
        } else { // do not swap
            return
        }
    }
    return
}

func (heap *Heap) Up(lru *LRUCache, nodeIndex int) {
    if nodeIndex > heap.Size || nodeIndex <= 1 {
        return
    } else {
        if heap.NodeList[nodeIndex].Time < heap.NodeList[nodeIndex/2].Time {
            lru.NodeMap[heap.NodeList[nodeIndex].Key] = nodeIndex / 2
            lru.NodeMap[heap.NodeList[nodeIndex/2].Key] = nodeIndex
            heap.NodeList[nodeIndex], heap.NodeList[nodeIndex/2] = heap.NodeList[nodeIndex/2], heap.NodeList[nodeIndex]
            heap.Up(lru, nodeIndex/2)
        }
    }
    return
}

func (heap *Heap) Push(lru *LRUCache, key, time int) {
    newNode := HeapNode {
        Key: key,
        Time: time,
    }
    heap.Size += 1
    if len(heap.NodeList) >= heap.Size + 1 {
        heap.NodeList[heap.Size] = newNode
    } else {
        heap.NodeList = append(heap.NodeList, newNode)
    }
    lru.NodeMap[key] = heap.Size
    heap.Up(lru, heap.Size)
    return
}

func (heap *Heap) Pop(lru *LRUCache) int {
    k := heap.NodeList[1].Key
    heap.NodeList[1] = heap.NodeList[heap.Size]
    heap.Size -= 1
    heap.Down(lru, 1)
    return k
}

func main() {
    capacity := 3
    obj := Constructor(capacity);
    obj.Put(1, 1)
    obj.Put(2, 2)
    obj.Put(3, 3)
    obj.Put(4, 4)
    fmt.Println(obj.Get(4), obj.H)
    fmt.Println(obj.Get(3), obj.H)
    fmt.Println(obj.Get(2), obj.H)
    fmt.Println(obj.Get(1), obj.H)
    obj.Put(5, 5)
    fmt.Println(obj.Get(1), obj.H)
    fmt.Println(obj.Get(2), obj.H)
    fmt.Println(obj.Get(3), obj.H)
    fmt.Println(obj.Get(4), obj.H)
    fmt.Println(obj.Get(5), obj.H)
    /* obj.Put(1, 1)
    fmt.Println(obj.H)
    obj.Put(2, 2)
    fmt.Println(obj.H)
    fmt.Println(obj.Get(1), obj.H)
    obj.Put(3, 3)
    fmt.Println(obj.H)
    fmt.Println(obj.Get(2), obj.H)
    obj.Put(4, 4)
    fmt.Println(obj.H)
    fmt.Println(obj.Get(1), obj.H)
    fmt.Println(obj.Get(3), obj.H)
    fmt.Println(obj.Get(4), obj.H) */
}