package main

import "fmt"

type node struct {
    index   int
    out     []*node
}

type nodePair struct {
    node    *node
    from    *nodePair
}

var answerArray [][]int

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
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
        return [][]string{}
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
    
    answerArray = [][]int{}
    length := -1
    // bfs
    mask := make([]bool, len(nodeList))
    fmt.Println(mask)
    nodeQueue := []nodePair{
        nodePair{
            node: nodeStart,
            from: nil,
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
                pairNow := &nodeQueue[headPoint]
                answer := []int{to.index}
                for pairNow != nil {
                    answer = append(answer, pairNow.node.index)
                    pairNow = pairNow.from
                }
                if length <= 0 {
                    length = len(answer)
                }
                if len(answer) == length {
                    answerArray = append(answerArray, answer)
                }
            }
            if !mask[to.index] {
                nodeQueue = append(nodeQueue, nodePair{
                    node: to,
                    from: &nodeQueue[headPoint],
                })
                tailPoint += 1
            }
        }
        mask[nodeNow.index] = true
        headPoint += 1
    }
    
    sss := [][]string{}
    for _, ans := range answerArray {
        ss := []string{}
        for i:=length-1; i>=0; i-- {
            ss = append(ss, wordList[ans[i]])
        }
        sss = append(sss, ss)
    }
    return sss
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
    fmt.Println(findLadders("hit", "cog", []string{
        "hot", "dot", "dog", "lot", "log", "cog",
    }))
}