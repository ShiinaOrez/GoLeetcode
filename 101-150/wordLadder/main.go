package main

import "fmt"

type node struct {
    index   int
    out     []*node
}

type nodePair struct {
    node    *node
    deep    int
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
    wordList = append(wordList, beginWord)
    wordsCount := len(wordList)
    nodeList := make([]*node, wordsCount)
    var nodeStart, nodeEnd *node
    for i:=0; i<wordsCount; i++ {
        nodeList[i] = &node{
            index: i,
            out: []*node{},
        }
        if wordList[i] == endWord {
            nodeEnd = nodeList[i]
        }
    }
    if nodeEnd == nil {
        return 0
    }
    nodeStart = nodeList[wordsCount-1]
    
    // building the graph
    for i:=0; i<wordsCount-1; i++ {
        for j:=i+1; j<wordsCount; j++ {
            if jumpMatch(wordList[i], wordList[j]) {
                nodeList[i].out = append(nodeList[i].out, nodeList[j])
                nodeList[j].out = append(nodeList[j].out, nodeList[i])
            }
        }
    }
    
    length := 0
    // bfs
    mask := make([]bool, len(nodeList))
    fmt.Println(mask)
    nodeQueue := []nodePair{
        nodePair{
            node: nodeStart,
            deep: 1,
        },
    }
    headPoint := 0
    tailPoint := 0
    for headPoint <= tailPoint {
        nodeNow := nodeQueue[headPoint].node
        if nodeNow == nodeEnd {
            break
        }
        for _, to := range nodeNow.out {
            if to == nodeEnd {
                return nodeQueue[headPoint].deep+1
            }
            if !mask[to.index] {
                nodeQueue = append(nodeQueue, nodePair{
                    node: to,
                    deep: nodeQueue[headPoint].deep+1,
                })
                tailPoint += 1
            }
        }
        mask[nodeNow.index] = true
        headPoint += 1
    }
    return length
}

func jumpMatch(a, b string) bool {
    notMatch := false
    for i, _ := range a {
        if a[i] != b[i] {
            if notMatch {
                return false
            } else {
                notMatch = true
            }
        }
    }
    return notMatch
}

func main() {
    fmt.Println(ladderLength("hit", "cog", []string{
        "hot", "dot", "dog", "lot", "log", "cog",
    }))
}